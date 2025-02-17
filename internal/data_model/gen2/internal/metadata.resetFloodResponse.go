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

type MetadataResetFloodResponse struct {
}

func (MetadataResetFloodResponse) TLName() string { return "metadata.resetFloodResponse" }
func (MetadataResetFloodResponse) TLTag() uint32  { return 0x9286abee }

func (item *MetadataResetFloodResponse) Reset()                         {}
func (item *MetadataResetFloodResponse) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *MetadataResetFloodResponse) Write(w []byte) ([]byte, error) { return w, nil }
func (item *MetadataResetFloodResponse) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0x9286abee)
}
func (item *MetadataResetFloodResponse) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0x9286abee), nil
}

func (item MetadataResetFloodResponse) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func MetadataResetFloodResponse__ReadJSON(item *MetadataResetFloodResponse, j interface{}) error {
	return item.readJSON(j)
}
func (item *MetadataResetFloodResponse) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.resetFloodResponse", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.resetFloodResponse", k)
	}
	return nil
}

func (item *MetadataResetFloodResponse) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *MetadataResetFloodResponse) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataResetFloodResponse) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("metadata.resetFloodResponse", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("metadata.resetFloodResponse", err.Error())
	}
	return nil
}
