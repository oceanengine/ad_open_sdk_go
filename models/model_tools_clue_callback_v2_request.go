/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueCallbackV2Request struct for ToolsClueCallbackV2Request
type ToolsClueCallbackV2Request struct {
	//
	AdvertiserIds []int64 `json:"advertiser_ids"`
	//
	ClueConvertState int64 `json:"clue_convert_state"`
	//
	ClueId string `json:"clue_id"`
	//
	Source int64 `json:"source"`
}
