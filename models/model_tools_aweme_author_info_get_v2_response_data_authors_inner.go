/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthorInfoGetV2ResponseDataAuthorsInner struct for ToolsAwemeAuthorInfoGetV2ResponseDataAuthorsInner
type ToolsAwemeAuthorInfoGetV2ResponseDataAuthorsInner struct {
	//
	AuthorName *string `json:"author_name,omitempty"`
	//
	Avatar *string `json:"avatar,omitempty"`
	//
	AwemeId *string `json:"aweme_id,omitempty"`
	//
	CoverNumStr *string `json:"cover_num_str,omitempty"`
	//
	LabelId *int64 `json:"label_id,omitempty"`
	//
	TotalFansNumStr *string `json:"total_fans_num_str,omitempty"`
}
