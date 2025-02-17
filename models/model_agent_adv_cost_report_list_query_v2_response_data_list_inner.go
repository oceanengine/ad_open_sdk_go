/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvCostReportListQueryV2ResponseDataListInner struct for AgentAdvCostReportListQueryV2ResponseDataListInner
type AgentAdvCostReportListQueryV2ResponseDataListInner struct {
	// 账户ID
	AdvId *int64 `json:"adv_id,omitempty"`
	// 账户名称
	AdvName *string `json:"adv_name,omitempty"`
	// 代理商客户ID
	AgentId *int64 `json:"agent_id,omitempty"`
	// 代理商客户名称
	AgentName *string `json:"agent_name,omitempty"`
	// 代理商业绩消耗(单位：千分之一分)
	AgentPerformanceCost *int64 `json:"agent_performance_cost,omitempty"`
	// 赠款消耗（计返点）(单位：千分之一分)
	AgentPerformanceGrantCost *int64                                        `json:"agent_performance_grant_cost,omitempty"`
	AppName                   *AgentAdvCostReportListQueryV2DataListAppName `json:"app_name,omitempty"`
	// 时间(格式：YYYY-MM-DD)
	AttributionDate *string `json:"attribution_date,omitempty"`
	// 品牌二级类别code
	BrandProductSecondType *int64 `json:"brand_product_second_type,omitempty"`
	// 品牌二级类别名称
	BrandProductSecondTypeName *string                                            `json:"brand_product_second_type_name,omitempty"`
	BusinessType               *AgentAdvCostReportListQueryV2DataListBusinessType `json:"business_type,omitempty"`
	// 业务来源名称
	BusinessTypeName *string                                            `json:"business_type_name,omitempty"`
	CashbackType     *AgentAdvCostReportListQueryV2DataListCashbackType `json:"cashback_type,omitempty"`
	// 资源包类型名称
	CashbackTypeName *string `json:"cashback_type_name,omitempty"`
	// 总消耗(单位：千分之一分)
	Cost       *int64                                           `json:"cost,omitempty"`
	CostSource *AgentAdvCostReportListQueryV2DataListCostSource `json:"cost_source,omitempty"`
	// 消耗来源名称
	CostSourceName *string                                   `json:"cost_source_name,omitempty"`
	Cpt            *AgentAdvCostReportListQueryV2DataListCpt `json:"cpt,omitempty"`
	// 是否CPT描述
	CptName *string `json:"cpt_name,omitempty"`
	// 授信消耗(单位：千分之一分)
	CreditCost *int64 `json:"credit_cost,omitempty"`
	// 广告主客户ID
	CustomerId *int64 `json:"customer_id,omitempty"`
	// 广告主客户名称
	CustomerName *string `json:"customer_name,omitempty"`
	// 赠款消耗（计返货）(单位：千分之一分)
	CustomerPerformanceGrantCost *int64                                              `json:"customer_performance_grant_cost,omitempty"`
	EcommerceType                *AgentAdvCostReportListQueryV2DataListEcommerceType `json:"ecommerce_type,omitempty"`
	Feedslive                    *AgentAdvCostReportListQueryV2DataListFeedslive     `json:"feedslive,omitempty"`
	// 是否FeedsLive描述
	FeedsliveName *string `json:"feedslive_name,omitempty"`
	// 一级代理商账户ID
	FirstAgentId *int64 `json:"first_agent_id,omitempty"`
	// 一级代理商账户名称
	FirstAgentName *string `json:"first_agent_name,omitempty"`
	// 客户一级行业
	FirstIndustry *string `json:"first_industry,omitempty"`
	// 赠款消耗(单位：千分之一分)
	GrantCost          *int64                                                   `json:"grant_cost,omitempty"`
	IsMatchingProvince *AgentAdvCostReportListQueryV2DataListIsMatchingProvince `json:"is_matching_province,omitempty"`
	// 非赠款消耗(单位：千分之一分)
	NoGrantCost *int64 `json:"no_grant_cost,omitempty"`
	// 预付消耗(单位：千分之一分)
	PrepayCost      *int64                                                `json:"prepay_cost,omitempty"`
	PricingCategory *AgentAdvCostReportListQueryV2DataListPricingCategory `json:"pricing_category,omitempty"`
	// 广告类型名称
	PricingCategoryName *string `json:"pricing_category_name,omitempty"`
	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`
	// 项目编号
	ProjectSerial *string `json:"project_serial,omitempty"`
	// 推广形式(1: 直播计划, 2: 非直播计划)
	PromotionType *int64 `json:"promotion_type,omitempty"`
	// 注册时间(格式：YYYY-MM-DD HH:MM:SS)
	RegisterTime *string `json:"register_time,omitempty"`
	// 销售ID
	SaleId *int64 `json:"sale_id,omitempty"`
	// 销售名称
	SaleName *string `json:"sale_name,omitempty"`
	// 代理商子账户ID
	SecondAgentId *int64 `json:"second_agent_id,omitempty"`
	// 代理商子账户名称
	SecondAgentName *string `json:"second_agent_name,omitempty"`
	// 客户二级行业
	SecondIndustry *string `json:"second_industry,omitempty"`
	// 当月结算一级行业
	SettlementFirstIndustryName *string `json:"settlement_first_industry_name,omitempty"`
	// 结算行业内部行业分类
	SettlementInternalIndustryCategory *string `json:"settlement_internal_industry_category,omitempty"`
	// 当月结算二级行业
	SettlementSecondIndustryName *string                                                   `json:"settlement_second_industry_name,omitempty"`
	SettlementStatsType          *AgentAdvCostReportListQueryV2DataListSettlementStatsType `json:"settlement_stats_type,omitempty"`
	// 商品ID
	SpuId        *int64                                             `json:"spu_id,omitempty"`
	SpuLabelName *AgentAdvCostReportListQueryV2DataListSpuLabelName `json:"spu_label_name,omitempty"`
	// 商品名称
	SpuName *string `json:"spu_name,omitempty"`
	// 共享钱包消耗(单位：千分之一分)
	SubSharedWalletCost *int64 `json:"sub_shared_wallet_cost,omitempty"`
	// 共享授信消耗(单位：千分之一分)
	SubSharedWalletCreditCost *int64 `json:"sub_shared_wallet_credit_cost,omitempty"`
	// 共享子钱包ID
	SubSharedWalletId *int64 `json:"sub_shared_wallet_id,omitempty"`
	// 共享子钱包名称
	SubSharedWalletName *string `json:"sub_shared_wallet_name,omitempty"`
	// 共享预付消耗(单位：千分之一分)
	SubSharedWalletPrepayCost *int64                                             `json:"sub_shared_wallet_prepay_cost,omitempty"`
	SystemOrigin              *AgentAdvCostReportListQueryV2DataListSystemOrigin `json:"system_origin,omitempty"`
}
