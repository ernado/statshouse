// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
)

var _ = basictl.NatWrite

type EnginePid struct {
}

func (EnginePid) TLName() string { return "engine.pid" }
func (EnginePid) TLTag() uint32  { return 0x559d6e36 }

func (item *EnginePid) Reset()                         {}
func (item *EnginePid) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EnginePid) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EnginePid) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0x559d6e36)
}
func (item *EnginePid) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0x559d6e36), nil
}

func (item *EnginePid) ReadResult(w []byte, ret *NetPid) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EnginePid) WriteResult(w []byte, ret NetPid) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EnginePid) ReadResultJSON(j interface{}, ret *NetPid) error {
	if err := NetPid__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *EnginePid) WriteResultJSON(w []byte, ret NetPid) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EnginePid) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret NetPid
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EnginePid) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.pid", err.Error())
	}
	var ret NetPid
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EnginePid) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EnginePid__ReadJSON(item *EnginePid, j interface{}) error { return item.readJSON(j) }
func (item *EnginePid) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.pid", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.pid", k)
	}
	return nil
}

func (item *EnginePid) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EnginePid) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EnginePid) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.pid", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.pid", err.Error())
	}
	return nil
}
