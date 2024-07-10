/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StardeliveryTaskDetailV30DataStarTaskSource
type StardeliveryTaskDetailV30DataStarTaskSource string

// List of stardelivery_task_detail_v3.0_data_star_task_source
const (
	MY_CREATIONS_StardeliveryTaskDetailV30DataStarTaskSource StardeliveryTaskDetailV30DataStarTaskSource = "MY_CREATIONS"
	SHARING_StardeliveryTaskDetailV30DataStarTaskSource      StardeliveryTaskDetailV30DataStarTaskSource = "SHARING"
)

// All allowed values of StardeliveryTaskDetailV30DataStarTaskSource enum
var AllowedStardeliveryTaskDetailV30DataStarTaskSourceEnumValues = []StardeliveryTaskDetailV30DataStarTaskSource{
	"MY_CREATIONS",
	"SHARING",
}

// NewStardeliveryTaskDetailV30DataStarTaskSourceFromValue returns a pointer to a valid StardeliveryTaskDetailV30DataStarTaskSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskDetailV30DataStarTaskSourceFromValue(v string) (*StardeliveryTaskDetailV30DataStarTaskSource, error) {
	ev := StardeliveryTaskDetailV30DataStarTaskSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskDetailV30DataStarTaskSource: valid values are %v", v, AllowedStardeliveryTaskDetailV30DataStarTaskSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskDetailV30DataStarTaskSource) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskDetailV30DataStarTaskSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_detail_v3.0_data_star_task_source value
func (v StardeliveryTaskDetailV30DataStarTaskSource) Ptr() *StardeliveryTaskDetailV30DataStarTaskSource {
	return &v
}
