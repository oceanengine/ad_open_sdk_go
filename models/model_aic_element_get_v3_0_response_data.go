/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicElementGetV30ResponseData
type AicElementGetV30ResponseData struct {
	// 返回结果
	Elements []*AicElementGetV30ResponseDataElementsInner `json:"elements,omitempty"`
	PageInfo *AicElementGetV30ResponseDataPageInfo        `json:"page_info,omitempty"`
}
