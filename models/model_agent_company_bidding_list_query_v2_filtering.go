/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentCompanyBiddingListQueryV2Filtering
type AgentCompanyBiddingListQueryV2Filtering struct {
	// 广告公司 id 列表，若选填该字段，则最少应上传1个广告主公司id，最多支持同时查询100个广告主公司。
	CompanyIds []int64 `json:"company_ids,omitempty"`
	// 广告主所属公司名称，若选填该字段，限制最小长度为1，最大长度为223。支持模糊查询。
	CompanyName *string `json:"company_name,omitempty"`
}
