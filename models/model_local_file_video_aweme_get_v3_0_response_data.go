/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalFileVideoAwemeGetV30ResponseData
type LocalFileVideoAwemeGetV30ResponseData struct {
	PageInfo *LocalFileVideoAwemeGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	VideoList []*LocalFileVideoAwemeGetV30ResponseDataVideoListInner `json:"video_list"`
}