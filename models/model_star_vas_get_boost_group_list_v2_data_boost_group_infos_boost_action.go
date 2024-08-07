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

// StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction
type StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction string

// List of star_vas_get_boost_group_list_v2_data_boost_group_infos_boost_action
const (
	LINK_ACTION_StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction   StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction = "LINK_ACTION"
	NATIVE_ACTION_StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction = "NATIVE_ACTION"
)

// All allowed values of StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction enum
var AllowedStarVasGetBoostGroupListV2DataBoostGroupInfosBoostActionEnumValues = []StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction{
	"LINK_ACTION",
	"NATIVE_ACTION",
}

// NewStarVasGetBoostGroupListV2DataBoostGroupInfosBoostActionFromValue returns a pointer to a valid StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarVasGetBoostGroupListV2DataBoostGroupInfosBoostActionFromValue(v string) (*StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction, error) {
	ev := StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction: valid values are %v", v, AllowedStarVasGetBoostGroupListV2DataBoostGroupInfosBoostActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction) IsValid() bool {
	for _, existing := range AllowedStarVasGetBoostGroupListV2DataBoostGroupInfosBoostActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_vas_get_boost_group_list_v2_data_boost_group_infos_boost_action value
func (v StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction) Ptr() *StarVasGetBoostGroupListV2DataBoostGroupInfosBoostAction {
	return &v
}
