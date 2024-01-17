/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeSimilarAuthorSearchV2ResponseDataListInner struct for ToolsAwemeSimilarAuthorSearchV2ResponseDataListInner
type ToolsAwemeSimilarAuthorSearchV2ResponseDataListInner struct {
	//
	AuthorName *string `json:"author_name,omitempty"`
	//
	Avatar *string `json:"avatar,omitempty"`
	//
	AwemeId *string `json:"aweme_id,omitempty"`
	//
	CategoryName *string `json:"category_name,omitempty"`
	//
	CoverNumStr *string `json:"cover_num_str,omitempty"`
	//
	LabelId *int64 `json:"label_id,omitempty"`
	//
	TotalFansNumStr *string `json:"total_fans_num_str,omitempty"`
}
