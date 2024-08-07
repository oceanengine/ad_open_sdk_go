/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource
type QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource string

// List of qianchuan_today_live_room_flow_performance_get_v1.0_flow_source
const (
	ALL_ECOM_FLOW_SOURCE_QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource = "AllEcomFlowSource"
	ALL_ECOM_GMV_SOURCE_QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource  QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource = "AllEcomGmvSource"
)

// All allowed values of QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource enum
var AllowedQianchuanTodayLiveRoomFlowPerformanceGetV10FlowSourceEnumValues = []QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource{
	"AllEcomFlowSource",
	"AllEcomGmvSource",
}

// NewQianchuanTodayLiveRoomFlowPerformanceGetV10FlowSourceFromValue returns a pointer to a valid QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomFlowPerformanceGetV10FlowSourceFromValue(v string) (*QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource, error) {
	ev := QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource: valid values are %v", v, AllowedQianchuanTodayLiveRoomFlowPerformanceGetV10FlowSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomFlowPerformanceGetV10FlowSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_flow_performance_get_v1.0_flow_source value
func (v QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource) Ptr() *QianchuanTodayLiveRoomFlowPerformanceGetV10FlowSource {
	return &v
}
