/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SecurityScoreViolationEventGetV30DataAdvScoreEventStatus
type SecurityScoreViolationEventGetV30DataAdvScoreEventStatus string

// List of security_score_violation_event_get_v3.0_data_adv_score_event_status
const (
	APPEAL_SecurityScoreViolationEventGetV30DataAdvScoreEventStatus     SecurityScoreViolationEventGetV30DataAdvScoreEventStatus = "APPEAL"
	FAILAPPEAL_SecurityScoreViolationEventGetV30DataAdvScoreEventStatus SecurityScoreViolationEventGetV30DataAdvScoreEventStatus = "FAILAPPEAL"
	ONAPPEAL_SecurityScoreViolationEventGetV30DataAdvScoreEventStatus   SecurityScoreViolationEventGetV30DataAdvScoreEventStatus = "ONAPPEAL"
	VALID_SecurityScoreViolationEventGetV30DataAdvScoreEventStatus      SecurityScoreViolationEventGetV30DataAdvScoreEventStatus = "VALID"
)

// Ptr returns reference to security_score_violation_event_get_v3.0_data_adv_score_event_status value
func (v SecurityScoreViolationEventGetV30DataAdvScoreEventStatus) Ptr() *SecurityScoreViolationEventGetV30DataAdvScoreEventStatus {
	return &v
}