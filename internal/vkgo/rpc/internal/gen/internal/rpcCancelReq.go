// Copyright 2022 V Kontakte LLC
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

type RpcCancelReq struct {
	QueryId int64
}

func (RpcCancelReq) TLName() string { return "rpcCancelReq" }
func (RpcCancelReq) TLTag() uint32  { return 0x193f1b22 }

func (item *RpcCancelReq) Reset() {
	item.QueryId = 0
}

func (item *RpcCancelReq) Read(w []byte) (_ []byte, err error) {
	return basictl.LongRead(w, &item.QueryId)
}

func (item *RpcCancelReq) Write(w []byte) (_ []byte, err error) {
	return basictl.LongWrite(w, item.QueryId), nil
}

func (item *RpcCancelReq) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x193f1b22); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *RpcCancelReq) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x193f1b22)
	return item.Write(w)
}

func (item RpcCancelReq) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func RpcCancelReq__ReadJSON(item *RpcCancelReq, j interface{}) error { return item.readJSON(j) }
func (item *RpcCancelReq) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("rpcCancelReq", "expected json object")
	}
	_jQueryId := _jm["query_id"]
	delete(_jm, "query_id")
	if err := JsonReadInt64(_jQueryId, &item.QueryId); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("rpcCancelReq", k)
	}
	return nil
}

func (item *RpcCancelReq) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.QueryId != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"query_id":`...)
		w = basictl.JSONWriteInt64(w, item.QueryId)
	}
	return append(w, '}'), nil
}

func (item *RpcCancelReq) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *RpcCancelReq) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("rpcCancelReq", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("rpcCancelReq", err.Error())
	}
	return nil
}
