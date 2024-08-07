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

// StardeliveryTaskListV30DataListStarTaskSource
type StardeliveryTaskListV30DataListStarTaskSource string

// List of stardelivery_task_list_v3.0_data_list_star_task_source
const (
	MY_CREATIONS_StardeliveryTaskListV30DataListStarTaskSource  StardeliveryTaskListV30DataListStarTaskSource = "MY_CREATIONS"
	SHARING_StardeliveryTaskListV30DataListStarTaskSource       StardeliveryTaskListV30DataListStarTaskSource = "SHARING"
	SHATE_EXPIRED_StardeliveryTaskListV30DataListStarTaskSource StardeliveryTaskListV30DataListStarTaskSource = "SHATE_EXPIRED"
)

// All allowed values of StardeliveryTaskListV30DataListStarTaskSource enum
var AllowedStardeliveryTaskListV30DataListStarTaskSourceEnumValues = []StardeliveryTaskListV30DataListStarTaskSource{
	"MY_CREATIONS",
	"SHARING",
	"SHATE_EXPIRED",
}

// NewStardeliveryTaskListV30DataListStarTaskSourceFromValue returns a pointer to a valid StardeliveryTaskListV30DataListStarTaskSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskListV30DataListStarTaskSourceFromValue(v string) (*StardeliveryTaskListV30DataListStarTaskSource, error) {
	ev := StardeliveryTaskListV30DataListStarTaskSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskListV30DataListStarTaskSource: valid values are %v", v, AllowedStardeliveryTaskListV30DataListStarTaskSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskListV30DataListStarTaskSource) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskListV30DataListStarTaskSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_list_v3.0_data_list_star_task_source value
func (v StardeliveryTaskListV30DataListStarTaskSource) Ptr() *StardeliveryTaskListV30DataListStarTaskSource {
	return &v
}
