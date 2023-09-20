/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerDeepBidTypeGetV30DataDeepBidType
type EventManagerDeepBidTypeGetV30DataDeepBidType string

// List of event_manager_deep_bid_type_get_v3.0_data_deep_bid_type
const (
	ALI_FL_EventManagerDeepBidTypeGetV30DataDeepBidType                       EventManagerDeepBidTypeGetV30DataDeepBidType = "ALI_FL"
	AUTO_MIN_SECOND_STAGE_EventManagerDeepBidTypeGetV30DataDeepBidType        EventManagerDeepBidTypeGetV30DataDeepBidType = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_EventManagerDeepBidTypeGetV30DataDeepBidType               EventManagerDeepBidTypeGetV30DataDeepBidType = "BID_PER_ACTION"
	DEEP_BID_DEFAULT_EventManagerDeepBidTypeGetV30DataDeepBidType             EventManagerDeepBidTypeGetV30DataDeepBidType = "DEEP_BID_DEFAULT"
	DEEP_BID_MIN_EventManagerDeepBidTypeGetV30DataDeepBidType                 EventManagerDeepBidTypeGetV30DataDeepBidType = "DEEP_BID_MIN"
	DEEP_BID_PACING_EventManagerDeepBidTypeGetV30DataDeepBidType              EventManagerDeepBidTypeGetV30DataDeepBidType = "DEEP_BID_PACING"
	DEEP_BID_TYPE_RETENTION_DAYS_EventManagerDeepBidTypeGetV30DataDeepBidType EventManagerDeepBidTypeGetV30DataDeepBidType = "DEEP_BID_TYPE_RETENTION_DAYS"
	FIRST_AND_SEVEN_PAY_ROI_EventManagerDeepBidTypeGetV30DataDeepBidType      EventManagerDeepBidTypeGetV30DataDeepBidType = "FIRST_AND_SEVEN_PAY_ROI"
	FORM_BID_EventManagerDeepBidTypeGetV30DataDeepBidType                     EventManagerDeepBidTypeGetV30DataDeepBidType = "FORM_BID"
	MIN_SECOND_STAGE_EventManagerDeepBidTypeGetV30DataDeepBidType             EventManagerDeepBidTypeGetV30DataDeepBidType = "MIN_SECOND_STAGE"
	PACING_SECOND_STAGE_EventManagerDeepBidTypeGetV30DataDeepBidType          EventManagerDeepBidTypeGetV30DataDeepBidType = "PACING_SECOND_STAGE"
	PER_AND_SEVEN_PAY_ROI_EventManagerDeepBidTypeGetV30DataDeepBidType        EventManagerDeepBidTypeGetV30DataDeepBidType = "PER_AND_SEVEN_PAY_ROI"
	PHONE_CONNECT_BID_EventManagerDeepBidTypeGetV30DataDeepBidType            EventManagerDeepBidTypeGetV30DataDeepBidType = "PHONE_CONNECT_BID"
	ROI_COEFFICIENT_EventManagerDeepBidTypeGetV30DataDeepBidType              EventManagerDeepBidTypeGetV30DataDeepBidType = "ROI_COEFFICIENT"
	ROI_DIRECT_MAIL_EventManagerDeepBidTypeGetV30DataDeepBidType              EventManagerDeepBidTypeGetV30DataDeepBidType = "ROI_DIRECT_MAIL"
	ROI_PACING_EventManagerDeepBidTypeGetV30DataDeepBidType                   EventManagerDeepBidTypeGetV30DataDeepBidType = "ROI_PACING"
	SMARTBID_EventManagerDeepBidTypeGetV30DataDeepBidType                     EventManagerDeepBidTypeGetV30DataDeepBidType = "SMARTBID"
	SOCIAL_ROI_EventManagerDeepBidTypeGetV30DataDeepBidType                   EventManagerDeepBidTypeGetV30DataDeepBidType = "SOCIAL_ROI"
)

// All allowed values of EventManagerDeepBidTypeGetV30DataDeepBidType enum
var AllowedEventManagerDeepBidTypeGetV30DataDeepBidTypeEnumValues = []EventManagerDeepBidTypeGetV30DataDeepBidType{
	"ALI_FL",
	"AUTO_MIN_SECOND_STAGE",
	"BID_PER_ACTION",
	"DEEP_BID_DEFAULT",
	"DEEP_BID_MIN",
	"DEEP_BID_PACING",
	"DEEP_BID_TYPE_RETENTION_DAYS",
	"FIRST_AND_SEVEN_PAY_ROI",
	"FORM_BID",
	"MIN_SECOND_STAGE",
	"PACING_SECOND_STAGE",
	"PER_AND_SEVEN_PAY_ROI",
	"PHONE_CONNECT_BID",
	"ROI_COEFFICIENT",
	"ROI_DIRECT_MAIL",
	"ROI_PACING",
	"SMARTBID",
	"SOCIAL_ROI",
}

// NewEventManagerDeepBidTypeGetV30DataDeepBidTypeFromValue returns a pointer to a valid EventManagerDeepBidTypeGetV30DataDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerDeepBidTypeGetV30DataDeepBidTypeFromValue(v string) (*EventManagerDeepBidTypeGetV30DataDeepBidType, error) {
	ev := EventManagerDeepBidTypeGetV30DataDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerDeepBidTypeGetV30DataDeepBidType: valid values are %v", v, AllowedEventManagerDeepBidTypeGetV30DataDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerDeepBidTypeGetV30DataDeepBidType) IsValid() bool {
	for _, existing := range AllowedEventManagerDeepBidTypeGetV30DataDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_deep_bid_type_get_v3.0_data_deep_bid_type value
func (v EventManagerDeepBidTypeGetV30DataDeepBidType) Ptr() *EventManagerDeepBidTypeGetV30DataDeepBidType {
	return &v
}
