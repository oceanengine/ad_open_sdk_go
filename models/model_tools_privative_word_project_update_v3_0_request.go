/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordProjectUpdateV30Request struct for ToolsPrivativeWordProjectUpdateV30Request
type ToolsPrivativeWordProjectUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ProjectList []*ToolsPrivativeWordProjectUpdateV30RequestProjectListInner `json:"project_list"`
}
