/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeAuthorListV2ResponseData
type StarChallengeAuthorListV2ResponseData struct {
	//
	AuthorList []*StarChallengeAuthorListV2ResponseDataAuthorListInner `json:"author_list,omitempty"`
	PageInfo   *StarChallengeAuthorListV2ResponseDataPageInfo          `json:"page_info,omitempty"`
}
