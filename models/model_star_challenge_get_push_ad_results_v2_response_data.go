/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeGetPushAdResultsV2ResponseData
type StarChallengeGetPushAdResultsV2ResponseData struct {
	// 广告推送结果
	PushResults []*StarChallengeGetPushAdResultsV2ResponseDataPushResultsInner `json:"push_results,omitempty"`
}