/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType
type StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType int64

// List of star_demand_om_expand_challenge_v2_om_participate_author_range_operation_type
const (
	Enum_0_StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType = 0
	Enum_1_StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType = 1
	Enum_2_StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType = 2
)

// All allowed values of StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType enum
var AllowedStarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationTypeEnumValues = []StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType{
	0,
	1,
	2,
}

// NewStarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationTypeFromValue returns a pointer to a valid StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationTypeFromValue(v int64) (*StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType, error) {
	ev := StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType: valid values are %v", v, AllowedStarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType) IsValid() bool {
	for _, existing := range AllowedStarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_om_expand_challenge_v2_om_participate_author_range_operation_type value
func (v StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType) Ptr() *StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType {
	return &v
}
