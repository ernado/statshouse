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

type EngineRegisterDynamicLib struct {
	LibId string
}

func (EngineRegisterDynamicLib) TLName() string { return "engine.registerDynamicLib" }
func (EngineRegisterDynamicLib) TLTag() uint32  { return 0x2f86f276 }

func (item *EngineRegisterDynamicLib) Reset() {
	item.LibId = ""
}

func (item *EngineRegisterDynamicLib) Read(w []byte) (_ []byte, err error) {
	return basictl.StringRead(w, &item.LibId)
}

func (item *EngineRegisterDynamicLib) Write(w []byte) (_ []byte, err error) {
	return basictl.StringWrite(w, item.LibId)
}

func (item *EngineRegisterDynamicLib) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x2f86f276); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineRegisterDynamicLib) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x2f86f276)
	return item.Write(w)
}

func (item *EngineRegisterDynamicLib) ReadResult(w []byte, ret *BoolStat) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineRegisterDynamicLib) WriteResult(w []byte, ret BoolStat) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineRegisterDynamicLib) ReadResultJSON(j interface{}, ret *BoolStat) error {
	if err := BoolStat__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *EngineRegisterDynamicLib) WriteResultJSON(w []byte, ret BoolStat) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineRegisterDynamicLib) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret BoolStat
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineRegisterDynamicLib) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.registerDynamicLib", err.Error())
	}
	var ret BoolStat
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineRegisterDynamicLib) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineRegisterDynamicLib__ReadJSON(item *EngineRegisterDynamicLib, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineRegisterDynamicLib) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.registerDynamicLib", "expected json object")
	}
	_jLibId := _jm["lib_id"]
	delete(_jm, "lib_id")
	if err := JsonReadString(_jLibId, &item.LibId); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.registerDynamicLib", k)
	}
	return nil
}

func (item *EngineRegisterDynamicLib) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.LibId) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"lib_id":`...)
		w = basictl.JSONWriteString(w, item.LibId)
	}
	return append(w, '}'), nil
}

func (item *EngineRegisterDynamicLib) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineRegisterDynamicLib) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.registerDynamicLib", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.registerDynamicLib", err.Error())
	}
	return nil
}