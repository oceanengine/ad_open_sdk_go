/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType
type SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType string

// List of security_score_violation_event_get_v3.0_data_adv_score_event_illegal_type
const (
	GENERAL_SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType       SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType = "GENERAL"
	ONECLASS_SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType      SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType = "ONECLASS"
	SERIOUS_SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType       SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType = "SERIOUS"
	TWOTHREECLASS_SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType = "TWOTHREECLASS"
)

// Ptr returns reference to security_score_violation_event_get_v3.0_data_adv_score_event_illegal_type value
func (v SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType) Ptr() *SecurityScoreViolationEventGetV30DataAdvScoreEventIllegalType {
	return &v
}
