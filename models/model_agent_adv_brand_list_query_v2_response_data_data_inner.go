/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvBrandListQueryV2ResponseDataDataInner struct for AgentAdvBrandListQueryV2ResponseDataDataInner
type AgentAdvBrandListQueryV2ResponseDataDataInner struct {
	AccountStatus *AgentAdvBrandListQueryV2DataDataAccountStatus `json:"account_status,omitempty"`
	// 广告主账户ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 广告主账户名称
	AdvertiserName *string `json:"advertiser_name,omitempty"`
	// 代理商账户ID
	AgentId *int64 `json:"agent_id,omitempty"`
	// 代理商账户名称
	AgentName *string `json:"agent_name,omitempty"`
	// 品牌余额
	BrandBalance *int64 `json:"brand_balance,omitempty"`
	// 品牌可用余额
	BrandBalanceValid *int64 `json:"brand_balance_valid,omitempty"`
	// 品牌赠款余额
	BrandGrantBalance *int64 `json:"brand_grant_balance,omitempty"`
	// 品牌非赠款可用余额
	BrandNonGrantBalanceValid *int64 `json:"brand_non_grant_balance_valid,omitempty"`
	// 非赠款消耗
	CashCost *int64 `json:"cash_cost,omitempty"`
	// 广告主公司ID
	CompanyId *int64 `json:"company_id,omitempty"`
	// 广告主公司名称
	CompanyName *string `json:"company_name,omitempty"`
	// 广告主账户一级行业名称
	FirstIndustry *string `json:"first_industry,omitempty"`
	// 广告主账户一级行业ID
	FirstIndustryId *int64 `json:"first_industry_id,omitempty"`
	// 赠款消耗
	GrantCost *int64 `json:"grant_cost,omitempty"`
	// 广告主账户二级行业名称
	SecondIndustry *string `json:"second_industry,omitempty"`
	// 广告主账户二级行业ID
	SecondIndustryId *int64 `json:"second_industry_id,omitempty"`
	// 总消耗
	StatCost *int64 `json:"stat_cost,omitempty"`
}
