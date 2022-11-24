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

type StatshouseSourceBucket struct {
	Metrics           []StatshouseItem
	SampleFactors     []StatshouseSampleFactor
	IngestionStatusOk []StatshouseIngestionStatus
	MissedSeconds     uint32
	AgentEnv          int32
}

func (StatshouseSourceBucket) TLName() string { return "statshouse.sourceBucket" }
func (StatshouseSourceBucket) TLTag() uint32  { return 0x3af6c822 }

func (item *StatshouseSourceBucket) Reset() {
	item.Metrics = item.Metrics[:0]
	item.SampleFactors = item.SampleFactors[:0]
	item.IngestionStatusOk = item.IngestionStatusOk[:0]
	item.MissedSeconds = 0
	item.AgentEnv = 0
}

func (item *StatshouseSourceBucket) Read(w []byte) (_ []byte, err error) {
	if w, err = VectorStatshouseItem0Read(w, &item.Metrics); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Read(w, &item.SampleFactors); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseIngestionStatus0Read(w, &item.IngestionStatusOk); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.MissedSeconds); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.AgentEnv)
}

func (item *StatshouseSourceBucket) Write(w []byte) (_ []byte, err error) {
	if w, err = VectorStatshouseItem0Write(w, item.Metrics); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Write(w, item.SampleFactors); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseIngestionStatus0Write(w, item.IngestionStatusOk); err != nil {
		return w, err
	}
	w = basictl.NatWrite(w, item.MissedSeconds)
	return basictl.IntWrite(w, item.AgentEnv), nil
}

func (item *StatshouseSourceBucket) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3af6c822); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseSourceBucket) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x3af6c822)
	return item.Write(w)
}

func (item StatshouseSourceBucket) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseSourceBucket__ReadJSON(item *StatshouseSourceBucket, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseSourceBucket) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.sourceBucket", "expected json object")
	}
	_jMetrics := _jm["metrics"]
	delete(_jm, "metrics")
	_jSampleFactors := _jm["sample_factors"]
	delete(_jm, "sample_factors")
	_jIngestionStatusOk := _jm["ingestion_status_ok"]
	delete(_jm, "ingestion_status_ok")
	_jMissedSeconds := _jm["missed_seconds"]
	delete(_jm, "missed_seconds")
	if err := JsonReadUint32(_jMissedSeconds, &item.MissedSeconds); err != nil {
		return err
	}
	_jAgentEnv := _jm["agent_env"]
	delete(_jm, "agent_env")
	if err := JsonReadInt32(_jAgentEnv, &item.AgentEnv); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.sourceBucket", k)
	}
	if err := VectorStatshouseItem0ReadJSON(_jMetrics, &item.Metrics); err != nil {
		return err
	}
	if err := VectorStatshouseSampleFactor0ReadJSON(_jSampleFactors, &item.SampleFactors); err != nil {
		return err
	}
	if err := VectorStatshouseIngestionStatus0ReadJSON(_jIngestionStatusOk, &item.IngestionStatusOk); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseSourceBucket) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Metrics) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"metrics":`...)
		if w, err = VectorStatshouseItem0WriteJSON(w, item.Metrics); err != nil {
			return w, err
		}
	}
	if len(item.SampleFactors) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"sample_factors":`...)
		if w, err = VectorStatshouseSampleFactor0WriteJSON(w, item.SampleFactors); err != nil {
			return w, err
		}
	}
	if len(item.IngestionStatusOk) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"ingestion_status_ok":`...)
		if w, err = VectorStatshouseIngestionStatus0WriteJSON(w, item.IngestionStatusOk); err != nil {
			return w, err
		}
	}
	if item.MissedSeconds != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"missed_seconds":`...)
		w = basictl.JSONWriteUint32(w, item.MissedSeconds)
	}
	if item.AgentEnv != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"agent_env":`...)
		w = basictl.JSONWriteInt32(w, item.AgentEnv)
	}
	return append(w, '}'), nil
}

func (item *StatshouseSourceBucket) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseSourceBucket) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.sourceBucket", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.sourceBucket", err.Error())
	}
	return nil
}

type StatshouseSourceBucketBytes struct {
	Metrics           []StatshouseItemBytes
	SampleFactors     []StatshouseSampleFactor
	IngestionStatusOk []StatshouseIngestionStatus
	MissedSeconds     uint32
	AgentEnv          int32
}

func (StatshouseSourceBucketBytes) TLName() string { return "statshouse.sourceBucket" }
func (StatshouseSourceBucketBytes) TLTag() uint32  { return 0x3af6c822 }

func (item *StatshouseSourceBucketBytes) Reset() {
	item.Metrics = item.Metrics[:0]
	item.SampleFactors = item.SampleFactors[:0]
	item.IngestionStatusOk = item.IngestionStatusOk[:0]
	item.MissedSeconds = 0
	item.AgentEnv = 0
}

func (item *StatshouseSourceBucketBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = VectorStatshouseItem0BytesRead(w, &item.Metrics); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Read(w, &item.SampleFactors); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseIngestionStatus0Read(w, &item.IngestionStatusOk); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.MissedSeconds); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.AgentEnv)
}

func (item *StatshouseSourceBucketBytes) Write(w []byte) (_ []byte, err error) {
	if w, err = VectorStatshouseItem0BytesWrite(w, item.Metrics); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Write(w, item.SampleFactors); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseIngestionStatus0Write(w, item.IngestionStatusOk); err != nil {
		return w, err
	}
	w = basictl.NatWrite(w, item.MissedSeconds)
	return basictl.IntWrite(w, item.AgentEnv), nil
}

func (item *StatshouseSourceBucketBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3af6c822); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseSourceBucketBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x3af6c822)
	return item.Write(w)
}

func (item StatshouseSourceBucketBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseSourceBucketBytes__ReadJSON(item *StatshouseSourceBucketBytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseSourceBucketBytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.sourceBucket", "expected json object")
	}
	_jMetrics := _jm["metrics"]
	delete(_jm, "metrics")
	_jSampleFactors := _jm["sample_factors"]
	delete(_jm, "sample_factors")
	_jIngestionStatusOk := _jm["ingestion_status_ok"]
	delete(_jm, "ingestion_status_ok")
	_jMissedSeconds := _jm["missed_seconds"]
	delete(_jm, "missed_seconds")
	if err := JsonReadUint32(_jMissedSeconds, &item.MissedSeconds); err != nil {
		return err
	}
	_jAgentEnv := _jm["agent_env"]
	delete(_jm, "agent_env")
	if err := JsonReadInt32(_jAgentEnv, &item.AgentEnv); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.sourceBucket", k)
	}
	if err := VectorStatshouseItem0BytesReadJSON(_jMetrics, &item.Metrics); err != nil {
		return err
	}
	if err := VectorStatshouseSampleFactor0ReadJSON(_jSampleFactors, &item.SampleFactors); err != nil {
		return err
	}
	if err := VectorStatshouseIngestionStatus0ReadJSON(_jIngestionStatusOk, &item.IngestionStatusOk); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseSourceBucketBytes) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Metrics) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"metrics":`...)
		if w, err = VectorStatshouseItem0BytesWriteJSON(w, item.Metrics); err != nil {
			return w, err
		}
	}
	if len(item.SampleFactors) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"sample_factors":`...)
		if w, err = VectorStatshouseSampleFactor0WriteJSON(w, item.SampleFactors); err != nil {
			return w, err
		}
	}
	if len(item.IngestionStatusOk) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"ingestion_status_ok":`...)
		if w, err = VectorStatshouseIngestionStatus0WriteJSON(w, item.IngestionStatusOk); err != nil {
			return w, err
		}
	}
	if item.MissedSeconds != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"missed_seconds":`...)
		w = basictl.JSONWriteUint32(w, item.MissedSeconds)
	}
	if item.AgentEnv != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"agent_env":`...)
		w = basictl.JSONWriteInt32(w, item.AgentEnv)
	}
	return append(w, '}'), nil
}

func (item *StatshouseSourceBucketBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseSourceBucketBytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.sourceBucket", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.sourceBucket", err.Error())
	}
	return nil
}