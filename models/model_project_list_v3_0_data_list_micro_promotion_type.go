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

// ProjectListV30DataListMicroPromotionType
type ProjectListV30DataListMicroPromotionType string

// List of project_list_v3.0_data_list_micro_promotion_type
const (
	BYTE_APP_ProjectListV30DataListMicroPromotionType    ProjectListV30DataListMicroPromotionType = "BYTE_APP"
	BYTE_GAME_ProjectListV30DataListMicroPromotionType   ProjectListV30DataListMicroPromotionType = "BYTE_GAME"
	WECHAT_APP_ProjectListV30DataListMicroPromotionType  ProjectListV30DataListMicroPromotionType = "WECHAT_APP"
	WECHAT_GAME_ProjectListV30DataListMicroPromotionType ProjectListV30DataListMicroPromotionType = "WECHAT_GAME"
)

// All allowed values of ProjectListV30DataListMicroPromotionType enum
var AllowedProjectListV30DataListMicroPromotionTypeEnumValues = []ProjectListV30DataListMicroPromotionType{
	"BYTE_APP",
	"BYTE_GAME",
	"WECHAT_APP",
	"WECHAT_GAME",
}

// NewProjectListV30DataListMicroPromotionTypeFromValue returns a pointer to a valid ProjectListV30DataListMicroPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListMicroPromotionTypeFromValue(v string) (*ProjectListV30DataListMicroPromotionType, error) {
	ev := ProjectListV30DataListMicroPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListMicroPromotionType: valid values are %v", v, AllowedProjectListV30DataListMicroPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListMicroPromotionType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListMicroPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_micro_promotion_type value
func (v ProjectListV30DataListMicroPromotionType) Ptr() *ProjectListV30DataListMicroPromotionType {
	return &v
}
