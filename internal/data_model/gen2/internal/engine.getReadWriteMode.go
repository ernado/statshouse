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

type EngineGetReadWriteMode struct {
	FieldsMask uint32
}

func (EngineGetReadWriteMode) TLName() string { return "engine.getReadWriteMode" }
func (EngineGetReadWriteMode) TLTag() uint32  { return 0x61b3f593 }

func (item *EngineGetReadWriteMode) Reset() {
	item.FieldsMask = 0
}

func (item *EngineGetReadWriteMode) Read(w []byte) (_ []byte, err error) {
	return basictl.NatRead(w, &item.FieldsMask)
}

func (item *EngineGetReadWriteMode) Write(w []byte) (_ []byte, err error) {
	return basictl.NatWrite(w, item.FieldsMask), nil
}

func (item *EngineGetReadWriteMode) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x61b3f593); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineGetReadWriteMode) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x61b3f593)
	return item.Write(w)
}

func (item *EngineGetReadWriteMode) ReadResult(w []byte, ret *EngineReadWriteMode) (_ []byte, err error) {
	return ret.ReadBoxed(w, item.FieldsMask)
}

func (item *EngineGetReadWriteMode) WriteResult(w []byte, ret EngineReadWriteMode) (_ []byte, err error) {
	return ret.WriteBoxed(w, item.FieldsMask)
}

func (item *EngineGetReadWriteMode) ReadResultJSON(j interface{}, ret *EngineReadWriteMode) error {
	if err := EngineReadWriteMode__ReadJSON(ret, j, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *EngineGetReadWriteMode) WriteResultJSON(w []byte, ret EngineReadWriteMode) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w, item.FieldsMask); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineGetReadWriteMode) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret EngineReadWriteMode
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineGetReadWriteMode) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.getReadWriteMode", err.Error())
	}
	var ret EngineReadWriteMode
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineGetReadWriteMode) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineGetReadWriteMode__ReadJSON(item *EngineGetReadWriteMode, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineGetReadWriteMode) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.getReadWriteMode", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.getReadWriteMode", k)
	}
	return nil
}

func (item *EngineGetReadWriteMode) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	return append(w, '}'), nil
}

func (item *EngineGetReadWriteMode) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineGetReadWriteMode) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.getReadWriteMode", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.getReadWriteMode", err.Error())
	}
	return nil
}
