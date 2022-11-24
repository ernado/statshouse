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

type StatshouseStringTopElement struct {
	Key   string
	Value float32
}

func (StatshouseStringTopElement) TLName() string { return "statshouse.string_top_element" }
func (StatshouseStringTopElement) TLTag() uint32  { return 0xec2e097d }

func (item *StatshouseStringTopElement) Reset() {
	item.Key = ""
	item.Value = 0
}

func (item *StatshouseStringTopElement) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.FloatRead(w, &item.Value)
}

func (item *StatshouseStringTopElement) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWrite(w, item.Key); err != nil {
		return w, err
	}
	return basictl.FloatWrite(w, item.Value), nil
}

func (item *StatshouseStringTopElement) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xec2e097d); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseStringTopElement) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xec2e097d)
	return item.Write(w)
}

func (item StatshouseStringTopElement) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseStringTopElement__ReadJSON(item *StatshouseStringTopElement, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseStringTopElement) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.string_top_element", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadString(_jKey, &item.Key); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadFloat32(_jValue, &item.Value); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.string_top_element", k)
	}
	return nil
}

func (item *StatshouseStringTopElement) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Key) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteString(w, item.Key)
	}
	if item.Value != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteFloat32(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *StatshouseStringTopElement) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseStringTopElement) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.string_top_element", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.string_top_element", err.Error())
	}
	return nil
}

type StatshouseStringTopElementBytes struct {
	Key   []byte
	Value float32
}

func (StatshouseStringTopElementBytes) TLName() string { return "statshouse.string_top_element" }
func (StatshouseStringTopElementBytes) TLTag() uint32  { return 0xec2e097d }

func (item *StatshouseStringTopElementBytes) Reset() {
	item.Key = item.Key[:0]
	item.Value = 0
}

func (item *StatshouseStringTopElementBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringReadBytes(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.FloatRead(w, &item.Value)
}

func (item *StatshouseStringTopElementBytes) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWriteBytes(w, item.Key); err != nil {
		return w, err
	}
	return basictl.FloatWrite(w, item.Value), nil
}

func (item *StatshouseStringTopElementBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xec2e097d); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseStringTopElementBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xec2e097d)
	return item.Write(w)
}

func (item StatshouseStringTopElementBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseStringTopElementBytes__ReadJSON(item *StatshouseStringTopElementBytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseStringTopElementBytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.string_top_element", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadStringBytes(_jKey, &item.Key); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadFloat32(_jValue, &item.Value); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.string_top_element", k)
	}
	return nil
}

func (item *StatshouseStringTopElementBytes) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Key) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteStringBytes(w, item.Key)
	}
	if item.Value != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteFloat32(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *StatshouseStringTopElementBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseStringTopElementBytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.string_top_element", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.string_top_element", err.Error())
	}
	return nil
}

func VectorStatshouseStringTopElement0Read(w []byte, vec *[]StatshouseStringTopElement) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseStringTopElement, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseStringTopElement0Write(w []byte, vec []StatshouseStringTopElement) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseStringTopElement0ReadJSON(j interface{}, vec *[]StatshouseStringTopElement) error {
	l, _arr, err := JsonReadArray("[]StatshouseStringTopElement", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseStringTopElement, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseStringTopElement__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshouseStringTopElement0WriteJSON(w []byte, vec []StatshouseStringTopElement) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

func VectorStatshouseStringTopElement0BytesRead(w []byte, vec *[]StatshouseStringTopElementBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseStringTopElementBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseStringTopElement0BytesWrite(w []byte, vec []StatshouseStringTopElementBytes) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseStringTopElement0BytesReadJSON(j interface{}, vec *[]StatshouseStringTopElementBytes) error {
	l, _arr, err := JsonReadArray("[]StatshouseStringTopElementBytes", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseStringTopElementBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseStringTopElementBytes__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshouseStringTopElement0BytesWriteJSON(w []byte, vec []StatshouseStringTopElementBytes) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}