/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarProjectListV2ResponseData
type StarProjectListV2ResponseData struct {
	PageInfo *StarProjectListV2ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	Projects []*StarProjectListV2ResponseDataProjectsInner `json:"projects,omitempty"`
}
