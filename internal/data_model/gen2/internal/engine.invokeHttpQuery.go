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

type EngineInvokeHttpQuery struct {
	Query EngineHttpQuery
}

func (EngineInvokeHttpQuery) TLName() string { return "engine.invokeHttpQuery" }
func (EngineInvokeHttpQuery) TLTag() uint32  { return 0xf4c73c0b }

func (item *EngineInvokeHttpQuery) Reset() {
	item.Query.Reset()
}

func (item *EngineInvokeHttpQuery) Read(w []byte) (_ []byte, err error) {
	return item.Query.Read(w)
}

func (item *EngineInvokeHttpQuery) Write(w []byte) (_ []byte, err error) {
	return item.Query.Write(w)
}

func (item *EngineInvokeHttpQuery) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf4c73c0b); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineInvokeHttpQuery) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xf4c73c0b)
	return item.Write(w)
}

func (item *EngineInvokeHttpQuery) ReadResult(w []byte, ret *EngineHttpQueryResponse) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineInvokeHttpQuery) WriteResult(w []byte, ret EngineHttpQueryResponse) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineInvokeHttpQuery) ReadResultJSON(j interface{}, ret *EngineHttpQueryResponse) error {
	if err := EngineHttpQueryResponse__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *EngineInvokeHttpQuery) WriteResultJSON(w []byte, ret EngineHttpQueryResponse) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineInvokeHttpQuery) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret EngineHttpQueryResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineInvokeHttpQuery) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.invokeHttpQuery", err.Error())
	}
	var ret EngineHttpQueryResponse
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineInvokeHttpQuery) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineInvokeHttpQuery__ReadJSON(item *EngineInvokeHttpQuery, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineInvokeHttpQuery) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.invokeHttpQuery", "expected json object")
	}
	_jQuery := _jm["query"]
	delete(_jm, "query")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.invokeHttpQuery", k)
	}
	if err := EngineHttpQuery__ReadJSON(&item.Query, _jQuery); err != nil {
		return err
	}
	return nil
}

func (item *EngineInvokeHttpQuery) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"query":`...)
	if w, err = item.Query.WriteJSON(w); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *EngineInvokeHttpQuery) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineInvokeHttpQuery) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.invokeHttpQuery", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.invokeHttpQuery", err.Error())
	}
	return nil
}
