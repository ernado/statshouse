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

type StatshouseApiSeriesMeta struct {
	FieldsMask uint32
	TimeShift  int64
	Tags       map[string]string
	What       StatshouseApiFunction // Conditional: item.FieldsMask.1
	Name       string                // Conditional: nat_query_fields_mask.4
	Color      string                // Conditional: nat_query_fields_mask.5
	Total      int32                 // Conditional: nat_query_fields_mask.6
	MaxHosts   []string              // Conditional: nat_query_fields_mask.7
}

func (StatshouseApiSeriesMeta) TLName() string { return "statshouseApi.seriesMeta" }
func (StatshouseApiSeriesMeta) TLTag() uint32  { return 0x5c2bf286 }

func (item *StatshouseApiSeriesMeta) SetWhat(v StatshouseApiFunction) {
	item.What = v
	item.FieldsMask |= 1 << 1
}
func (item *StatshouseApiSeriesMeta) ClearWhat() {
	item.What.Reset()
	item.FieldsMask &^= 1 << 1
}
func (item StatshouseApiSeriesMeta) IsSetWhat() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *StatshouseApiSeriesMeta) SetName(v string, nat_query_fields_mask *uint32) {
	item.Name = v
	if nat_query_fields_mask != nil {
		*nat_query_fields_mask |= 1 << 4
	}
}
func (item *StatshouseApiSeriesMeta) ClearName(nat_query_fields_mask *uint32) {
	item.Name = ""
	if nat_query_fields_mask != nil {
		*nat_query_fields_mask &^= 1 << 4
	}
}
func (item StatshouseApiSeriesMeta) IsSetName(nat_query_fields_mask uint32) bool {
	return nat_query_fields_mask&(1<<4) != 0
}

func (item *StatshouseApiSeriesMeta) SetColor(v string, nat_query_fields_mask *uint32) {
	item.Color = v
	if nat_query_fields_mask != nil {
		*nat_query_fields_mask |= 1 << 5
	}
}
func (item *StatshouseApiSeriesMeta) ClearColor(nat_query_fields_mask *uint32) {
	item.Color = ""
	if nat_query_fields_mask != nil {
		*nat_query_fields_mask &^= 1 << 5
	}
}
func (item StatshouseApiSeriesMeta) IsSetColor(nat_query_fields_mask uint32) bool {
	return nat_query_fields_mask&(1<<5) != 0
}

func (item *StatshouseApiSeriesMeta) SetTotal(v int32, nat_query_fields_mask *uint32) {
	item.Total = v
	if nat_query_fields_mask != nil {
		*nat_query_fields_mask |= 1 << 6
	}
}
func (item *StatshouseApiSeriesMeta) ClearTotal(nat_query_fields_mask *uint32) {
	item.Total = 0
	if nat_query_fields_mask != nil {
		*nat_query_fields_mask &^= 1 << 6
	}
}
func (item StatshouseApiSeriesMeta) IsSetTotal(nat_query_fields_mask uint32) bool {
	return nat_query_fields_mask&(1<<6) != 0
}

func (item *StatshouseApiSeriesMeta) SetMaxHosts(v []string, nat_query_fields_mask *uint32) {
	item.MaxHosts = v
	if nat_query_fields_mask != nil {
		*nat_query_fields_mask |= 1 << 7
	}
}
func (item *StatshouseApiSeriesMeta) ClearMaxHosts(nat_query_fields_mask *uint32) {
	item.MaxHosts = item.MaxHosts[:0]
	if nat_query_fields_mask != nil {
		*nat_query_fields_mask &^= 1 << 7
	}
}
func (item StatshouseApiSeriesMeta) IsSetMaxHosts(nat_query_fields_mask uint32) bool {
	return nat_query_fields_mask&(1<<7) != 0
}

func (item *StatshouseApiSeriesMeta) Reset() {
	item.FieldsMask = 0
	item.TimeShift = 0
	VectorDictionaryFieldString0Reset(item.Tags)
	item.What.Reset()
	item.Name = ""
	item.Color = ""
	item.Total = 0
	item.MaxHosts = item.MaxHosts[:0]
}

func (item *StatshouseApiSeriesMeta) Read(w []byte, nat_query_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.TimeShift); err != nil {
		return w, err
	}
	if w, err = VectorDictionaryFieldString0Read(w, &item.Tags); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = item.What.ReadBoxed(w); err != nil {
			return w, err
		}
	} else {
		item.What.Reset()
	}
	if nat_query_fields_mask&(1<<4) != 0 {
		if w, err = basictl.StringRead(w, &item.Name); err != nil {
			return w, err
		}
	} else {
		item.Name = ""
	}
	if nat_query_fields_mask&(1<<5) != 0 {
		if w, err = basictl.StringRead(w, &item.Color); err != nil {
			return w, err
		}
	} else {
		item.Color = ""
	}
	if nat_query_fields_mask&(1<<6) != 0 {
		if w, err = basictl.IntRead(w, &item.Total); err != nil {
			return w, err
		}
	} else {
		item.Total = 0
	}
	if nat_query_fields_mask&(1<<7) != 0 {
		if w, err = VectorString0Read(w, &item.MaxHosts); err != nil {
			return w, err
		}
	} else {
		item.MaxHosts = item.MaxHosts[:0]
	}
	return w, nil
}

func (item *StatshouseApiSeriesMeta) Write(w []byte, nat_query_fields_mask uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.LongWrite(w, item.TimeShift)
	if w, err = VectorDictionaryFieldString0Write(w, item.Tags); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = item.What.WriteBoxed(w); err != nil {
			return w, err
		}
	}
	if nat_query_fields_mask&(1<<4) != 0 {
		if w, err = basictl.StringWrite(w, item.Name); err != nil {
			return w, err
		}
	}
	if nat_query_fields_mask&(1<<5) != 0 {
		if w, err = basictl.StringWrite(w, item.Color); err != nil {
			return w, err
		}
	}
	if nat_query_fields_mask&(1<<6) != 0 {
		w = basictl.IntWrite(w, item.Total)
	}
	if nat_query_fields_mask&(1<<7) != 0 {
		if w, err = VectorString0Write(w, item.MaxHosts); err != nil {
			return w, err
		}
	}
	return w, nil
}

func (item *StatshouseApiSeriesMeta) ReadBoxed(w []byte, nat_query_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x5c2bf286); err != nil {
		return w, err
	}
	return item.Read(w, nat_query_fields_mask)
}

func (item *StatshouseApiSeriesMeta) WriteBoxed(w []byte, nat_query_fields_mask uint32) ([]byte, error) {
	w = basictl.NatWrite(w, 0x5c2bf286)
	return item.Write(w, nat_query_fields_mask)
}

func StatshouseApiSeriesMeta__ReadJSON(item *StatshouseApiSeriesMeta, j interface{}, nat_query_fields_mask uint32) error {
	return item.readJSON(j, nat_query_fields_mask)
}
func (item *StatshouseApiSeriesMeta) readJSON(j interface{}, nat_query_fields_mask uint32) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouseApi.seriesMeta", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jTimeShift := _jm["time_shift"]
	delete(_jm, "time_shift")
	if err := JsonReadInt64(_jTimeShift, &item.TimeShift); err != nil {
		return err
	}
	_jTags := _jm["tags"]
	delete(_jm, "tags")
	_jWhat := _jm["what"]
	delete(_jm, "what")
	_jName := _jm["name"]
	delete(_jm, "name")
	_jColor := _jm["color"]
	delete(_jm, "color")
	_jTotal := _jm["total"]
	delete(_jm, "total")
	_jMaxHosts := _jm["max_hosts"]
	delete(_jm, "max_hosts")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouseApi.seriesMeta", k)
	}
	if _jWhat != nil {
		item.FieldsMask |= 1 << 1
	}
	if nat_query_fields_mask&(1<<4) == 0 && _jName != nil {
		return ErrorInvalidJSON("statshouseApi.seriesMeta", "field 'name' is defined, while corresponding implicit fieldmask bit is 0")
	}
	if nat_query_fields_mask&(1<<5) == 0 && _jColor != nil {
		return ErrorInvalidJSON("statshouseApi.seriesMeta", "field 'color' is defined, while corresponding implicit fieldmask bit is 0")
	}
	if nat_query_fields_mask&(1<<6) == 0 && _jTotal != nil {
		return ErrorInvalidJSON("statshouseApi.seriesMeta", "field 'total' is defined, while corresponding implicit fieldmask bit is 0")
	}
	if nat_query_fields_mask&(1<<7) == 0 && _jMaxHosts != nil {
		return ErrorInvalidJSON("statshouseApi.seriesMeta", "field 'max_hosts' is defined, while corresponding implicit fieldmask bit is 0")
	}
	if err := VectorDictionaryFieldString0ReadJSON(_jTags, &item.Tags); err != nil {
		return err
	}
	if _jWhat != nil {
		if err := StatshouseApiFunction__ReadJSON(&item.What, _jWhat); err != nil {
			return err
		}
	} else {
		item.What.Reset()
	}
	if nat_query_fields_mask&(1<<4) != 0 {
		if err := JsonReadString(_jName, &item.Name); err != nil {
			return err
		}
	} else {
		item.Name = ""
	}
	if nat_query_fields_mask&(1<<5) != 0 {
		if err := JsonReadString(_jColor, &item.Color); err != nil {
			return err
		}
	} else {
		item.Color = ""
	}
	if nat_query_fields_mask&(1<<6) != 0 {
		if err := JsonReadInt32(_jTotal, &item.Total); err != nil {
			return err
		}
	} else {
		item.Total = 0
	}
	if nat_query_fields_mask&(1<<7) != 0 {
		if err := VectorString0ReadJSON(_jMaxHosts, &item.MaxHosts); err != nil {
			return err
		}
	} else {
		item.MaxHosts = item.MaxHosts[:0]
	}
	return nil
}

func (item *StatshouseApiSeriesMeta) WriteJSON(w []byte, nat_query_fields_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if item.TimeShift != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"time_shift":`...)
		w = basictl.JSONWriteInt64(w, item.TimeShift)
	}
	if len(item.Tags) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"tags":`...)
		if w, err = VectorDictionaryFieldString0WriteJSON(w, item.Tags); err != nil {
			return w, err
		}
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"what":`...)
		if w, err = item.What.WriteJSON(w); err != nil {
			return w, err
		}
	}
	if nat_query_fields_mask&(1<<4) != 0 {
		if len(item.Name) != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"name":`...)
			w = basictl.JSONWriteString(w, item.Name)
		}
	}
	if nat_query_fields_mask&(1<<5) != 0 {
		if len(item.Color) != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"color":`...)
			w = basictl.JSONWriteString(w, item.Color)
		}
	}
	if nat_query_fields_mask&(1<<6) != 0 {
		if item.Total != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"total":`...)
			w = basictl.JSONWriteInt32(w, item.Total)
		}
	}
	if nat_query_fields_mask&(1<<7) != 0 {
		if len(item.MaxHosts) != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"max_hosts":`...)
			if w, err = VectorString0WriteJSON(w, item.MaxHosts); err != nil {
				return w, err
			}
		}
	}
	return append(w, '}'), nil
}

func VectorStatshouseApiSeriesMeta0Read(w []byte, vec *[]StatshouseApiSeriesMeta, nat_t uint32) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseApiSeriesMeta, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w, nat_t); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseApiSeriesMeta0Write(w []byte, vec []StatshouseApiSeriesMeta, nat_t uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w, nat_t); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseApiSeriesMeta0ReadJSON(j interface{}, vec *[]StatshouseApiSeriesMeta, nat_t uint32) error {
	l, _arr, err := JsonReadArray("[]StatshouseApiSeriesMeta", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseApiSeriesMeta, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseApiSeriesMeta__ReadJSON(&(*vec)[i], _arr[i], nat_t); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshouseApiSeriesMeta0WriteJSON(w []byte, vec []StatshouseApiSeriesMeta, nat_t uint32) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w, nat_t); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
