// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rpc

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"os"
	"sync"
	"time"

	"github.com/vkcom/statshouse/internal/vkgo/basictl"
)

const (
	packetOverhead       = 4 * 4
	maxPacketLen         = 16*1024*1024 - 1
	maxNonceHandshakeLen = 1024 // this limit is not exact, but it still prevents attack to force allocating lots of memory
	blockSize            = 16
	padVal               = 4
	startSeqNum          = -2

	memcachedStatsReqRN  = "stats\r\n"
	memcachedStatsReqN   = "stats\n"
	memcachedGetStatsReq = "get stats\r\n"
	memcachedVersionReq  = "version\r\n"
)

var (
	memcachedCommands = []string{
		memcachedStatsReqRN,
		memcachedStatsReqN,
		memcachedGetStatsReq,
		memcachedVersionReq,
	}

	castagnoliTable = crc32.MakeTable(crc32.Castagnoli)

	ErrInvalidPacketLength = fmt.Errorf("invalid packet length")
)

// transport stream, encrypted using standard VK rpc scheme
type PacketConn struct {
	conn       net.Conn
	tcpconn_fd *os.File

	remoteAddr      string
	localAddr       string
	timeoutAccuracy time.Duration

	flagCancelReq bool

	readMu        sync.Mutex
	r             *cryptoReader
	rDeadline     time.Time
	headerReadBuf [packetOverhead - 4]byte
	readSeqNum    int64

	writeMu        sync.Mutex
	w              *cryptoWriter
	wDeadline      time.Time
	headerWriteBuf []byte // contains either crc from previous packet or nothing after flush
	writeSeqNum    int64

	keyID [4]byte // to identify clients for Server with more than 1 crypto key
	table *crc32.Table

	closeOnce sync.Once
	closeErr  error
}

func NewPacketConn(c net.Conn, readBufSize int, writeBufSize int, timeoutAccuracy time.Duration) *PacketConn {
	pc := &PacketConn{
		conn:            c,
		remoteAddr:      c.RemoteAddr().String(),
		localAddr:       c.LocalAddr().String(),
		timeoutAccuracy: timeoutAccuracy,
		r:               newCryptoReader(c, readBufSize),
		w:               newCryptoWriter(c, writeBufSize),
		readSeqNum:      startSeqNum,
		writeSeqNum:     startSeqNum,
		table:           crc32.IEEETable,
	}

	if tcpconn, ok := c.(*net.TCPConn); ok { // pay cast and dup() const only on start
		if fd, err := tcpconn.File(); err == nil { // ok, will work as not a tcp connection
			pc.tcpconn_fd = fd
		}
	}

	return pc
}

// Negotiated during handshake, does not change after handshake
func (pc *PacketConn) FlagCancelReq() bool {
	return pc.flagCancelReq
}

func (pc *PacketConn) LocalAddr() string {
	return pc.localAddr
}

func (pc *PacketConn) RemoteAddr() string {
	return pc.remoteAddr
}

func (pc *PacketConn) Close() error {
	pc.closeOnce.Do(func() {
		if pc.tcpconn_fd != nil {
			_ = pc.tcpconn_fd.Close()
		}
		pc.closeErr = pc.conn.Close()
	})
	return pc.closeErr
}

func (pc *PacketConn) setCRC32C() {
	pc.table = castagnoliTable
}

func (pc *PacketConn) setReadTimeoutUnlocked(timeout time.Duration) error {
	if timeout == 0 {
		if pc.rDeadline == (time.Time{}) {
			return nil
		}
		err := pc.conn.SetReadDeadline(time.Time{})
		if err != nil {
			return err
		}
		pc.rDeadline = time.Time{}
		return nil
	}
	deadlineDiff := time.Until(pc.rDeadline) - timeout
	if deadlineDiff < -pc.timeoutAccuracy || deadlineDiff > pc.timeoutAccuracy {
		deadline := time.Now().Add(timeout)
		err := pc.conn.SetReadDeadline(deadline)
		if err != nil {
			return err
		}
		pc.rDeadline = deadline
	}
	return nil
}

func (pc *PacketConn) setWriteTimeoutUnlocked(timeout time.Duration) error {
	if timeout == 0 {
		if pc.wDeadline == (time.Time{}) {
			return nil
		}
		err := pc.conn.SetWriteDeadline(time.Time{})
		if err != nil {
			return err
		}
		pc.wDeadline = time.Time{}
		return nil
	}
	deadlineDiff := time.Until(pc.wDeadline) - timeout
	if deadlineDiff < -pc.timeoutAccuracy || deadlineDiff > pc.timeoutAccuracy {
		deadline := time.Now().Add(timeout)
		err := pc.conn.SetWriteDeadline(deadline)
		if err != nil {
			return err
		}
		pc.wDeadline = deadline
	}
	return nil
}

// ReadPacket will resize/reuse body to size of packet
func (pc *PacketConn) ReadPacket(body []byte, timeout time.Duration) (tip uint32, _ []byte, err error) {
	tip, _, body, err = pc.readPacketWithMagic(body, timeout)
	return tip, body, err
}

// supports sending ascii command via terminal instead of first TL RPC packet, returns command in
func (pc *PacketConn) readPacketWithMagic(body []byte, timeout time.Duration) (tip uint32, magic []byte, _ []byte, err error) {
	pc.readMu.Lock()
	defer pc.readMu.Unlock()

	var header packetHeader
	magicHead, err := pc.readPacketHeaderUnlocked(&header, timeout)
	if err != nil {
		return 0, magicHead, body, err
	}

	body, err = pc.readPacketBodyUnlocked(&header, body, false, 0)
	return header.tip, nil, body, err
}

func (pc *PacketConn) readPacketHeaderUnlocked(header *packetHeader, timeout time.Duration) (magicHead []byte, err error) {
	if err = pc.setReadTimeoutUnlocked(timeout); err != nil {
		return nil, err
	}

	// special-case first packet: it can't contain padding, but can be a memcached command
	if pc.readSeqNum == startSeqNum {
		n, err := readFullOrMagic(pc.r, pc.headerReadBuf[:12], memcachedCommands)
		if err != nil {
			return pc.headerReadBuf[:n], err
		}
		header.length = binary.LittleEndian.Uint32(pc.headerReadBuf[:4])
		header.seqNum = binary.LittleEndian.Uint32(pc.headerReadBuf[4:8])
		header.tip = binary.LittleEndian.Uint32(pc.headerReadBuf[8:12])
	} else {
		header.length = padVal
		// it is important to return (0, eof) when FIN is read on the message boundary
		for i := 0; header.length == padVal; i++ {
			if i >= blockSize/4 {
				return nil, fmt.Errorf("excessive (%d) padding", i)
			}
			n, err := io.ReadFull(pc.r, pc.headerReadBuf[:4])
			if err != nil {
				return pc.headerReadBuf[:n], err
			}
			header.length = binary.LittleEndian.Uint32(pc.headerReadBuf[:4])
		}

		n, err := io.ReadFull(pc.r, pc.headerReadBuf[4:12])
		if err != nil {
			return pc.headerReadBuf[:4+n], err
		}
		header.seqNum = binary.LittleEndian.Uint32(pc.headerReadBuf[4:8])
		header.tip = binary.LittleEndian.Uint32(pc.headerReadBuf[8:12])
	}

	if header.length < packetOverhead || header.length > maxPacketLen {
		return pc.headerReadBuf[:12], fmt.Errorf("packet size %v outside  [%v, %v]", header.length, packetOverhead, maxPacketLen)
	}
	if pc.readSeqNum < 0 && header.length > maxNonceHandshakeLen {
		return pc.headerReadBuf[:12], fmt.Errorf("nonce/handshake packet size %v outside  [%v, %v]", header.length, packetOverhead, maxNonceHandshakeLen)
	}
	if header.seqNum != uint32(pc.readSeqNum) {
		return pc.headerReadBuf[:12], fmt.Errorf("seqnum mismatch: read %v, expected %v", header.seqNum, pc.readSeqNum)
	}
	if pc.readSeqNum == startSeqNum && header.tip != packetTypeRPCNonce { // this check is in nonceExchangeServer, but repeated here to detect non-RPC protocol earlier for connection hijack
		return pc.headerReadBuf[:12], fmt.Errorf("nonce packet type 0x%x instead of 0x%x", header.tip, packetTypeRPCNonce)
	}
	pc.readSeqNum++

	return nil, nil
}

func (pc *PacketConn) readPacketBodyUnlocked(header *packetHeader, body []byte, setTimeout bool, timeout time.Duration) (_ []byte, err error) {
	if setTimeout {
		if err = pc.setReadTimeoutUnlocked(timeout); err != nil {
			return body, err
		}
	}
	if header.length < packetOverhead || header.length > maxPacketLen {
		panic(fmt.Sprintf("packet size %v outside [%v, %v], was checked in readPacketHeaderUnlocked", header.length, packetOverhead, maxPacketLen))
	}
	sz := int(header.length) - packetOverhead + 4
	if cap(body) < sz {
		body = make([]byte, sz)
	} else {
		body = body[:sz]
	}
	n := 0
	for n < len(body) && err == nil {
		var nn int
		nn, err = pc.r.Read(body[n:])
		n += nn
	}
	if n < len(body) { // implies err != nil
		return body, err
	}
	// we forget error here, if err != nil, expecting to receive it again on reading next packet
	header.crc = binary.LittleEndian.Uint32(body[sz-4:])
	body = body[:sz-4]

	crc := pc.updateCRC(0, pc.headerReadBuf[:12])
	crc = pc.updateCRC(crc, body)

	if header.crc != crc {
		return body, fmt.Errorf("CRC mismatch: read 0x%x, expected 0x%x", header.crc, crc)
	}

	return body, nil
}

func (pc *PacketConn) WritePacket(packetType uint32, body []byte, timeout time.Duration) error {
	pc.writeMu.Lock()
	defer pc.writeMu.Unlock()

	crc, err := pc.writePacketHeaderUnlocked(packetType, len(body), timeout)
	if err != nil {
		return err
	}
	crc, err = pc.writePacketBodyUnlocked(crc, body)
	if err != nil {
		return err
	}
	pc.writePacketTrailerUnlocked(crc)
	return pc.FlushUnlocked()
}

func (pc *PacketConn) WritePacketNoFlush(packetType uint32, body []byte, timeout time.Duration) error {
	pc.writeMu.Lock()
	defer pc.writeMu.Unlock()

	return pc.WritePacketNoFlushUnlocked(packetType, body, timeout)
}

// If all writing is performed from the same goroutine, you can call Unlocked version of Write and Flush
func (pc *PacketConn) WritePacketNoFlushUnlocked(packetType uint32, body []byte, timeout time.Duration) error {
	crc, err := pc.writePacketHeaderUnlocked(packetType, len(body), timeout)
	if err != nil {
		return err
	}
	crc, err = pc.writePacketBodyUnlocked(crc, body)
	if err != nil {
		return err
	}
	pc.writePacketTrailerUnlocked(crc)
	return nil
}

func (pc *PacketConn) Flush() error {
	pc.writeMu.Lock()
	defer pc.writeMu.Unlock()

	return pc.FlushUnlocked()
}

// If all writing is performed from the same goroutine, you can call Unlocked version of Write and Flush
func (pc *PacketConn) FlushUnlocked() error {
	prevBytes := len(pc.headerWriteBuf)
	toWrite := prevBytes + pc.w.Padding(prevBytes)
	if toWrite != 0 {
		if toWrite&3 != 0 { // if p < 0 || p > 12, will panic in [] below
			panic(fmt.Sprintf("invalid crypto padding toWrite=%d", toWrite))
		}
		const padding = "\x04\x00\x00\x00\x04\x00\x00\x00\x04\x00\x00\x00" // []byte{padVal, 0, 0, 0, padVal, 0, 0, 0, padVal, 0, 0, 0}
		pc.headerWriteBuf = append(pc.headerWriteBuf, padding...)

		if _, err := pc.w.Write(pc.headerWriteBuf[:toWrite]); err != nil {
			return err
		}
		pc.headerWriteBuf = pc.headerWriteBuf[:0]
	}
	return pc.w.Flush()
}

type closeWriter interface {
	CloseWrite() error
}

// Motivation - you call ShutdownWrite, and your blocking ReadPacket* will stop after receiveing FIN with compatible sockets
// if you receive error for this method, you should call Close()
func (pc *PacketConn) ShutdownWrite() error {
	if err := pc.Flush(); err != nil { // Rare, so no problem to make excess locked call
		return err
	}
	cw, ok := pc.conn.(closeWriter) // UnixConn, TCPConn, and any other
	if !ok {
		return io.ErrShortWrite // TODO - better error
	}
	return cw.CloseWrite()
}

// how to use:
// first call writePacketHeaderUnlocked with sum of all body chunk lengths you are going to write on the next step
// then call writePacketBodyUnlocked 0 or more times
// then call writePacketTrailerUnlocked
func (pc *PacketConn) writePacketHeaderUnlocked(packetType uint32, packetBodyLen int, timeout time.Duration) (uint32, error) {
	if err := validBodyLen(packetBodyLen); err != nil {
		return 0, err
	}
	if err := pc.setWriteTimeoutUnlocked(timeout); err != nil {
		return 0, err
	}
	prevBytes := len(pc.headerWriteBuf)
	buf := pc.headerWriteBuf

	buf = basictl.NatWrite(buf, uint32(packetBodyLen+packetOverhead))
	buf = basictl.NatWrite(buf, uint32(pc.writeSeqNum))
	buf = basictl.NatWrite(buf, packetType)
	pc.headerWriteBuf = buf[:0] // reuse, prepare to accept crc32
	pc.writeSeqNum++

	if _, err := pc.w.Write(buf); err != nil { // with prevBytes
		return 0, err
	}
	return pc.updateCRC(0, buf[prevBytes:]), nil // without prevBytes
}

func (pc *PacketConn) writePacketBodyUnlocked(crc uint32, body []byte) (uint32, error) {
	if _, err := pc.w.Write(body); err != nil {
		return 0, err
	}
	return pc.updateCRC(crc, body), nil
}

func (pc *PacketConn) writePacketTrailerUnlocked(crc uint32) {
	pc.headerWriteBuf = binary.LittleEndian.AppendUint32(pc.headerWriteBuf, crc)
}

func (pc *PacketConn) updateCRC(crc uint32, data []byte) uint32 {
	return crc32.Update(crc, pc.table, data)
}

func (pc *PacketConn) encrypt(readKey []byte, readIV []byte, writeKey []byte, writeIV []byte) error {
	pc.readMu.Lock()
	defer pc.readMu.Unlock()

	pc.writeMu.Lock()
	defer pc.writeMu.Unlock()

	rc, err := aes.NewCipher(readKey)
	if err != nil {
		return fmt.Errorf("read AES init failed: %w", err)
	}

	wc, err := aes.NewCipher(writeKey)
	if err != nil {
		return fmt.Errorf("write AES init failed: %w", err)
	}

	if len(readIV) != blockSize {
		return fmt.Errorf("read IV size must be %v, not %v", blockSize, len(readIV))
	}
	if len(writeIV) != blockSize {
		return fmt.Errorf("write IV size must be %v, not %v", blockSize, len(writeIV))
	}

	rcbc := cipher.NewCBCDecrypter(rc, readIV)
	wcbc := cipher.NewCBCEncrypter(wc, writeIV)
	if rcbc.BlockSize() != blockSize {
		return fmt.Errorf("CBC read decrypter BlockSize must be %v, not %v", blockSize, rcbc.BlockSize())
	}
	if wcbc.BlockSize() != blockSize {
		return fmt.Errorf("CBC write decrypter BlockSize must be %v, not %v", blockSize, wcbc.BlockSize())
	}

	pc.r.encrypt(rcbc)
	pc.w.encrypt(wcbc)

	return nil
}

func validBodyLen(n int) error { // Motivation - high byte was used for some flags, we must not use it
	if n > maxPacketLen-packetOverhead {
		return fmt.Errorf("packet size (metadata+extra+request) %v exceeds maximum %v: %w", n, maxPacketLen, ErrInvalidPacketLength)
	}
	if n%4 != 0 {
		return fmt.Errorf("packet size %v must be a multiple of 4: %w", n, ErrInvalidPacketLength)
	}
	return nil
}

// Merge of io.ReadFull and io.ReadAtLeast, but with a twist:
// we return early with io.EOF if we have read a magic sequence of bytes.
func readFullOrMagic(r io.Reader, buf []byte, magics []string) (n int, err error) {
	m := len(buf)
	for n < m && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += nn

		// a twist:
		for _, magic := range magics {
			// cmd/compile does not allocate for this string conversion
			if string(buf[:n]) == magic {
				if err == nil {
					err = io.EOF
				}
				return
			}
		}
	}
	if n >= m { // actually never >
		err = nil
	} else if n > 0 && err == io.EOF {
		err = io.ErrUnexpectedEOF
	}
	return
}
