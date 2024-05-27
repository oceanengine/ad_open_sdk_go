/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StardeliveryTaskListV30FilteringStarTaskSource
type StardeliveryTaskListV30FilteringStarTaskSource string

// List of stardelivery_task_list_v3.0_filtering_star_task_source
const (
	MY_CREATIONS_StardeliveryTaskListV30FilteringStarTaskSource  StardeliveryTaskListV30FilteringStarTaskSource = "MY_CREATIONS"
	SHARING_StardeliveryTaskListV30FilteringStarTaskSource       StardeliveryTaskListV30FilteringStarTaskSource = "SHARING"
	SHATE_EXPIRED_StardeliveryTaskListV30FilteringStarTaskSource StardeliveryTaskListV30FilteringStarTaskSource = "SHATE_EXPIRED"
)

// All allowed values of StardeliveryTaskListV30FilteringStarTaskSource enum
var AllowedStardeliveryTaskListV30FilteringStarTaskSourceEnumValues = []StardeliveryTaskListV30FilteringStarTaskSource{
	"MY_CREATIONS",
	"SHARING",
	"SHATE_EXPIRED",
}

// NewStardeliveryTaskListV30FilteringStarTaskSourceFromValue returns a pointer to a valid StardeliveryTaskListV30FilteringStarTaskSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskListV30FilteringStarTaskSourceFromValue(v string) (*StardeliveryTaskListV30FilteringStarTaskSource, error) {
	ev := StardeliveryTaskListV30FilteringStarTaskSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskListV30FilteringStarTaskSource: valid values are %v", v, AllowedStardeliveryTaskListV30FilteringStarTaskSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskListV30FilteringStarTaskSource) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskListV30FilteringStarTaskSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_list_v3.0_filtering_star_task_source value
func (v StardeliveryTaskListV30FilteringStarTaskSource) Ptr() *StardeliveryTaskListV30FilteringStarTaskSource {
	return &v
}
