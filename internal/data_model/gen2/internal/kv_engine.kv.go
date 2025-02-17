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

type KvEngineKv struct {
	Key   int64
	Value int64
}

func (KvEngineKv) TLName() string { return "kv_engine.kv" }
func (KvEngineKv) TLTag() uint32  { return 0x18f34950 }

func (item *KvEngineKv) Reset() {
	item.Key = 0
	item.Value = 0
}

func (item *KvEngineKv) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.LongRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.Value)
}

func (item *KvEngineKv) Write(w []byte) (_ []byte, err error) {
	w = basictl.LongWrite(w, item.Key)
	return basictl.LongWrite(w, item.Value), nil
}

func (item *KvEngineKv) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x18f34950); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *KvEngineKv) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x18f34950)
	return item.Write(w)
}

func (item KvEngineKv) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func KvEngineKv__ReadJSON(item *KvEngineKv, j interface{}) error { return item.readJSON(j) }
func (item *KvEngineKv) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("kv_engine.kv", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadInt64(_jKey, &item.Key); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadInt64(_jValue, &item.Value); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("kv_engine.kv", k)
	}
	return nil
}

func (item *KvEngineKv) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Key != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteInt64(w, item.Key)
	}
	if item.Value != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteInt64(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *KvEngineKv) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *KvEngineKv) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("kv_engine.kv", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("kv_engine.kv", err.Error())
	}
	return nil
}

func VectorKvEngineKvBoxed0Read(w []byte, vec *[]KvEngineKv) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]KvEngineKv, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].ReadBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorKvEngineKvBoxed0Write(w []byte, vec []KvEngineKv) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.WriteBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorKvEngineKvBoxed0ReadJSON(j interface{}, vec *[]KvEngineKv) error {
	l, _arr, err := JsonReadArray("[]KvEngineKv", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]KvEngineKv, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := KvEngineKv__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorKvEngineKvBoxed0WriteJSON(w []byte, vec []KvEngineKv) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
