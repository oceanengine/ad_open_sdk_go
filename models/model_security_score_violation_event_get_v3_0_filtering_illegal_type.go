/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SecurityScoreViolationEventGetV30FilteringIllegalType
type SecurityScoreViolationEventGetV30FilteringIllegalType string

// List of security_score_violation_event_get_v3.0_filtering_illegal_type
const (
	CRITICAL_SecurityScoreViolationEventGetV30FilteringIllegalType      SecurityScoreViolationEventGetV30FilteringIllegalType = "CRITICAL"
	GENERAL_SecurityScoreViolationEventGetV30FilteringIllegalType       SecurityScoreViolationEventGetV30FilteringIllegalType = "GENERAL"
	MINOR_SecurityScoreViolationEventGetV30FilteringIllegalType         SecurityScoreViolationEventGetV30FilteringIllegalType = "MINOR"
	ONECLASS_SecurityScoreViolationEventGetV30FilteringIllegalType      SecurityScoreViolationEventGetV30FilteringIllegalType = "ONECLASS"
	SERIOUS_SecurityScoreViolationEventGetV30FilteringIllegalType       SecurityScoreViolationEventGetV30FilteringIllegalType = "SERIOUS"
	TWOTHREECLASS_SecurityScoreViolationEventGetV30FilteringIllegalType SecurityScoreViolationEventGetV30FilteringIllegalType = "TWOTHREECLASS"
)

// Ptr returns reference to security_score_violation_event_get_v3.0_filtering_illegal_type value
func (v SecurityScoreViolationEventGetV30FilteringIllegalType) Ptr() *SecurityScoreViolationEventGetV30FilteringIllegalType {
	return &v
}
