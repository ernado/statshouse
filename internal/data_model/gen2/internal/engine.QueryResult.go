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

func (item EngineQueryResult) AsUnion() EngineQueryResultUnion {
	var ret EngineQueryResultUnion
	ret.SetQueryResult(item)
	return ret
}

// AsUnion will be here
type EngineQueryResult struct {
	Size int32
}

func (EngineQueryResult) TLName() string { return "engine.queryResult" }
func (EngineQueryResult) TLTag() uint32  { return 0xac4d6fe9 }

func (item *EngineQueryResult) Reset() {
	item.Size = 0
}

func (item *EngineQueryResult) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.Size)
}

func (item *EngineQueryResult) Write(w []byte) (_ []byte, err error) {
	return basictl.IntWrite(w, item.Size), nil
}

func (item *EngineQueryResult) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xac4d6fe9); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineQueryResult) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xac4d6fe9)
	return item.Write(w)
}

func (item EngineQueryResult) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineQueryResult__ReadJSON(item *EngineQueryResult, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineQueryResult) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.queryResult", "expected json object")
	}
	_jSize := _jm["size"]
	delete(_jm, "size")
	if err := JsonReadInt32(_jSize, &item.Size); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.queryResult", k)
	}
	return nil
}

func (item *EngineQueryResult) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Size != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"size":`...)
		w = basictl.JSONWriteInt32(w, item.Size)
	}
	return append(w, '}'), nil
}

func (item *EngineQueryResult) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineQueryResult) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.queryResult", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.queryResult", err.Error())
	}
	return nil
}

func (item EngineQueryResultAio) AsUnion() EngineQueryResultUnion {
	var ret EngineQueryResultUnion
	ret.SetAio()
	return ret
}

// AsUnion will be here
type EngineQueryResultAio struct {
}

func (EngineQueryResultAio) TLName() string { return "engine.queryResultAio" }
func (EngineQueryResultAio) TLTag() uint32  { return 0xee2879b0 }

func (item *EngineQueryResultAio) Reset()                         {}
func (item *EngineQueryResultAio) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EngineQueryResultAio) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EngineQueryResultAio) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0xee2879b0)
}
func (item *EngineQueryResultAio) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0xee2879b0), nil
}

func (item EngineQueryResultAio) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineQueryResultAio__ReadJSON(item *EngineQueryResultAio, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineQueryResultAio) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.queryResultAio", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.queryResultAio", k)
	}
	return nil
}

func (item *EngineQueryResultAio) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineQueryResultAio) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineQueryResultAio) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.queryResultAio", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.queryResultAio", err.Error())
	}
	return nil
}

func (item EngineQueryResultError) AsUnion() EngineQueryResultUnion {
	var ret EngineQueryResultUnion
	ret.SetError(item)
	return ret
}

// AsUnion will be here
type EngineQueryResultError struct {
	ErrorCode   int32
	ErrorString string
}

func (EngineQueryResultError) TLName() string { return "engine.queryResultError" }
func (EngineQueryResultError) TLTag() uint32  { return 0x2b4dd0ba }

func (item *EngineQueryResultError) Reset() {
	item.ErrorCode = 0
	item.ErrorString = ""
}

func (item *EngineQueryResultError) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.ErrorCode); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.ErrorString)
}

func (item *EngineQueryResultError) Write(w []byte) (_ []byte, err error) {
	w = basictl.IntWrite(w, item.ErrorCode)
	return basictl.StringWrite(w, item.ErrorString)
}

func (item *EngineQueryResultError) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x2b4dd0ba); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineQueryResultError) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x2b4dd0ba)
	return item.Write(w)
}

func (item EngineQueryResultError) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineQueryResultError__ReadJSON(item *EngineQueryResultError, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineQueryResultError) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.queryResultError", "expected json object")
	}
	_jErrorCode := _jm["error_code"]
	delete(_jm, "error_code")
	if err := JsonReadInt32(_jErrorCode, &item.ErrorCode); err != nil {
		return err
	}
	_jErrorString := _jm["error_string"]
	delete(_jm, "error_string")
	if err := JsonReadString(_jErrorString, &item.ErrorString); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.queryResultError", k)
	}
	return nil
}

func (item *EngineQueryResultError) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.ErrorCode != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"error_code":`...)
		w = basictl.JSONWriteInt32(w, item.ErrorCode)
	}
	if len(item.ErrorString) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"error_string":`...)
		w = basictl.JSONWriteString(w, item.ErrorString)
	}
	return append(w, '}'), nil
}

func (item *EngineQueryResultError) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineQueryResultError) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.queryResultError", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.queryResultError", err.Error())
	}
	return nil
}

var _EngineQueryResultUnion = [3]UnionElement{
	{TLTag: 0xac4d6fe9, TLName: "engine.queryResult", TLString: "engine.queryResult#ac4d6fe9"},
	{TLTag: 0x2b4dd0ba, TLName: "engine.queryResultError", TLString: "engine.queryResultError#2b4dd0ba"},
	{TLTag: 0xee2879b0, TLName: "engine.queryResultAio", TLString: "engine.queryResultAio#ee2879b0"},
}

type EngineQueryResultUnion struct {
	valueQueryResult EngineQueryResult
	valueError       EngineQueryResultError
	index            int
}

func (item EngineQueryResultUnion) TLName() string { return _EngineQueryResultUnion[item.index].TLName }
func (item EngineQueryResultUnion) TLTag() uint32  { return _EngineQueryResultUnion[item.index].TLTag }

func (item *EngineQueryResultUnion) Reset() { item.ResetToQueryResult() }

func (item *EngineQueryResultUnion) IsQueryResult() bool { return item.index == 0 }

func (item *EngineQueryResultUnion) AsQueryResult() (*EngineQueryResult, bool) {
	if item.index != 0 {
		return nil, false
	}
	return &item.valueQueryResult, true
}
func (item *EngineQueryResultUnion) ResetToQueryResult() *EngineQueryResult {
	item.index = 0
	item.valueQueryResult.Reset()
	return &item.valueQueryResult
}
func (item *EngineQueryResultUnion) SetQueryResult(value EngineQueryResult) {
	item.index = 0
	item.valueQueryResult = value
}

func (item *EngineQueryResultUnion) IsError() bool { return item.index == 1 }

func (item *EngineQueryResultUnion) AsError() (*EngineQueryResultError, bool) {
	if item.index != 1 {
		return nil, false
	}
	return &item.valueError, true
}
func (item *EngineQueryResultUnion) ResetToError() *EngineQueryResultError {
	item.index = 1
	item.valueError.Reset()
	return &item.valueError
}
func (item *EngineQueryResultUnion) SetError(value EngineQueryResultError) {
	item.index = 1
	item.valueError = value
}

func (item *EngineQueryResultUnion) IsAio() bool { return item.index == 2 }

func (item *EngineQueryResultUnion) AsAio() (EngineQueryResultAio, bool) {
	var value EngineQueryResultAio
	return value, item.index == 2
}
func (item *EngineQueryResultUnion) ResetToAio() { item.index = 2 }
func (item *EngineQueryResultUnion) SetAio()     { item.index = 2 }

func (item *EngineQueryResultUnion) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0xac4d6fe9:
		item.index = 0
		return item.valueQueryResult.Read(w)
	case 0x2b4dd0ba:
		item.index = 1
		return item.valueError.Read(w)
	case 0xee2879b0:
		item.index = 2
		return w, nil
	default:
		return w, ErrorInvalidUnionTag("engine.QueryResult", tag)
	}
}

func (item *EngineQueryResultUnion) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, _EngineQueryResultUnion[item.index].TLTag)
	switch item.index {
	case 0:
		return item.valueQueryResult.Write(w)
	case 1:
		return item.valueError.Write(w)
	case 2:
		return w, nil
	default: // Impossible due to panic above
		return w, nil
	}
}

func EngineQueryResultUnion__ReadJSON(item *EngineQueryResultUnion, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineQueryResultUnion) readJSON(j interface{}) error {
	_jm, _tag, err := JsonReadUnionType("engine.QueryResult", j)
	if err != nil {
		return err
	}
	jvalue := _jm["value"]
	switch _tag {
	case "engine.queryResult#ac4d6fe9", "engine.queryResult", "#ac4d6fe9":
		item.index = 0
		if err := EngineQueryResult__ReadJSON(&item.valueQueryResult, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	case "engine.queryResultError#2b4dd0ba", "engine.queryResultError", "#2b4dd0ba":
		item.index = 1
		if err := EngineQueryResultError__ReadJSON(&item.valueError, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	case "engine.queryResultAio#ee2879b0", "engine.queryResultAio", "#ee2879b0":
		item.index = 2
	default:
		return ErrorInvalidUnionTagJSON("engine.QueryResult", _tag)
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.QueryResult", k)
	}
	return nil
}

func (item *EngineQueryResultUnion) WriteJSON(w []byte) (_ []byte, err error) {
	switch item.index {
	case 0:
		w = append(w, `{"type":"engine.queryResult#ac4d6fe9","value":`...)
		if w, err = item.valueQueryResult.WriteJSON(w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	case 1:
		w = append(w, `{"type":"engine.queryResultError#2b4dd0ba","value":`...)
		if w, err = item.valueError.WriteJSON(w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	case 2:
		return append(w, `{"type":"engine.queryResultAio#ee2879b0"}`...), nil
	default: // Impossible due to panic above
		return w, nil
	}
}

func (item EngineQueryResultUnion) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}
