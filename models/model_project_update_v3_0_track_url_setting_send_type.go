/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30TrackUrlSettingSendType
type ProjectUpdateV30TrackUrlSettingSendType string

// List of project_update_v3.0_track_url_setting_send_type
const (
	CLIENT_SEND_ProjectUpdateV30TrackUrlSettingSendType ProjectUpdateV30TrackUrlSettingSendType = "CLIENT_SEND"
	SERVER_SEND_ProjectUpdateV30TrackUrlSettingSendType ProjectUpdateV30TrackUrlSettingSendType = "SERVER_SEND"
)

// Ptr returns reference to project_update_v3.0_track_url_setting_send_type value
func (v ProjectUpdateV30TrackUrlSettingSendType) Ptr() *ProjectUpdateV30TrackUrlSettingSendType {
	return &v
}
