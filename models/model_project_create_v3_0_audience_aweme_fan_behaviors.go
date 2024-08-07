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

// ProjectCreateV30AudienceAwemeFanBehaviors
type ProjectCreateV30AudienceAwemeFanBehaviors string

// List of project_create_v3.0_audience_aweme_fan_behaviors
const (
	COMMENTED_USER_ProjectCreateV30AudienceAwemeFanBehaviors       ProjectCreateV30AudienceAwemeFanBehaviors = "COMMENTED_USER"
	FOLLOWED_USER_ProjectCreateV30AudienceAwemeFanBehaviors        ProjectCreateV30AudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_ProjectCreateV30AudienceAwemeFanBehaviors    ProjectCreateV30AudienceAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_ProjectCreateV30AudienceAwemeFanBehaviors    ProjectCreateV30AudienceAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_ProjectCreateV30AudienceAwemeFanBehaviors           ProjectCreateV30AudienceAwemeFanBehaviors = "LIKED_USER"
	LIVE_COMMENT_ProjectCreateV30AudienceAwemeFanBehaviors         ProjectCreateV30AudienceAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_ProjectCreateV30AudienceAwemeFanBehaviors ProjectCreateV30AudienceAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_ProjectCreateV30AudienceAwemeFanBehaviors     ProjectCreateV30AudienceAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_ProjectCreateV30AudienceAwemeFanBehaviors     ProjectCreateV30AudienceAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_ProjectCreateV30AudienceAwemeFanBehaviors     ProjectCreateV30AudienceAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_ProjectCreateV30AudienceAwemeFanBehaviors           ProjectCreateV30AudienceAwemeFanBehaviors = "LIVE_WATCH"
	SHARED_USER_ProjectCreateV30AudienceAwemeFanBehaviors          ProjectCreateV30AudienceAwemeFanBehaviors = "SHARED_USER"
)

// All allowed values of ProjectCreateV30AudienceAwemeFanBehaviors enum
var AllowedProjectCreateV30AudienceAwemeFanBehaviorsEnumValues = []ProjectCreateV30AudienceAwemeFanBehaviors{
	"COMMENTED_USER",
	"FOLLOWED_USER",
	"GOODS_CARTS_CLICK",
	"GOODS_CARTS_ORDER",
	"LIKED_USER",
	"LIVE_COMMENT",
	"LIVE_EFFECTIVE_WATCH",
	"LIVE_EXCEPTIONAL",
	"LIVE_GOODS_CLICK",
	"LIVE_GOODS_ORDER",
	"LIVE_WATCH",
	"SHARED_USER",
}

// NewProjectCreateV30AudienceAwemeFanBehaviorsFromValue returns a pointer to a valid ProjectCreateV30AudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceAwemeFanBehaviorsFromValue(v string) (*ProjectCreateV30AudienceAwemeFanBehaviors, error) {
	ev := ProjectCreateV30AudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceAwemeFanBehaviors: valid values are %v", v, AllowedProjectCreateV30AudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_aweme_fan_behaviors value
func (v ProjectCreateV30AudienceAwemeFanBehaviors) Ptr() *ProjectCreateV30AudienceAwemeFanBehaviors {
	return &v
}
