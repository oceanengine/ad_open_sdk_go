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

// ProjectCreateV30MicroPromotionType
type ProjectCreateV30MicroPromotionType string

// List of project_create_v3.0_micro_promotion_type
const (
	BYTE_APP_ProjectCreateV30MicroPromotionType    ProjectCreateV30MicroPromotionType = "BYTE_APP"
	BYTE_GAME_ProjectCreateV30MicroPromotionType   ProjectCreateV30MicroPromotionType = "BYTE_GAME"
	WECHAT_APP_ProjectCreateV30MicroPromotionType  ProjectCreateV30MicroPromotionType = "WECHAT_APP"
	WECHAT_GAME_ProjectCreateV30MicroPromotionType ProjectCreateV30MicroPromotionType = "WECHAT_GAME"
)

// All allowed values of ProjectCreateV30MicroPromotionType enum
var AllowedProjectCreateV30MicroPromotionTypeEnumValues = []ProjectCreateV30MicroPromotionType{
	"BYTE_APP",
	"BYTE_GAME",
	"WECHAT_APP",
	"WECHAT_GAME",
}

// NewProjectCreateV30MicroPromotionTypeFromValue returns a pointer to a valid ProjectCreateV30MicroPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30MicroPromotionTypeFromValue(v string) (*ProjectCreateV30MicroPromotionType, error) {
	ev := ProjectCreateV30MicroPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30MicroPromotionType: valid values are %v", v, AllowedProjectCreateV30MicroPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30MicroPromotionType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30MicroPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_micro_promotion_type value
func (v ProjectCreateV30MicroPromotionType) Ptr() *ProjectCreateV30MicroPromotionType {
	return &v
}
