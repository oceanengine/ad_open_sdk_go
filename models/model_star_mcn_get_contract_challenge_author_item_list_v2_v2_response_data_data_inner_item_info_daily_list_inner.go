/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInnerItemInfoDailyListInner struct for StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInnerItemInfoDailyListInner
type StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInnerItemInfoDailyListInner struct {
	// 达人已发放收益
	AuthorAmount *int64 `json:"author_amount,omitempty"`
	//
	AuthorAmountIaa *int64 `json:"author_amount_iaa,omitempty"`
	//
	AuthorAmountIap *int64 `json:"author_amount_iap,omitempty"`
	// 达人预估广告消耗收益
	AuthorCostAdReward *int64 `json:"author_cost_ad_reward,omitempty"`
	// 达人预估cpa收益
	AuthorEstAmountCpa *int64 `json:"author_est_amount_cpa,omitempty"`
	// 达人预估cps收益
	AuthorEstAmountCps *int64 `json:"author_est_amount_cps,omitempty"`
	// 达人 预估 除广告消耗收益之外的收益
	AuthorEstAmountNonAd *int64 `json:"author_est_amount_non_ad,omitempty"`
	// 达人已结算广告消耗收益
	AuthorSettleAdReward *int64 `json:"author_settle_ad_reward,omitempty"`
	// 达人 已结算 除广告消耗收益之外的收益
	AuthorSettleAmountNonAd *int64 `json:"author_settle_amount_non_ad,omitempty"`
	// item_id 产生的分天广告收益，单位为厘（元*1000）
	IaaIncomeDaily *string `json:"iaa_income_daily,omitempty"`
	//
	PDate *string `json:"p_date,omitempty"`
	// 服务商预估广告消耗收益
	ProviderCostAdReward *int64 `json:"provider_cost_ad_reward,omitempty"`
	// 服务商预估收益
	ProviderEstAmount *int64 `json:"provider_est_amount,omitempty"`
	// 服务商预估cpa收益
	ProviderEstAmountCpa *int64 `json:"provider_est_amount_cpa,omitempty"`
	// 服务商预估cps收益
	ProviderEstAmountCps *int64 `json:"provider_est_amount_cps,omitempty"`
	// 服务商预估收益iaa
	ProviderEstAmountIaa *int64 `json:"provider_est_amount_iaa,omitempty"`
	// 服务商预估收益iap
	ProviderEstAmountIap *int64 `json:"provider_est_amount_iap,omitempty"`
	// 服务商 预估 除广告消耗收益之外的收益
	ProviderEstAmountNonAd *int64 `json:"provider_est_amount_non_ad,omitempty"`
	// 服务商已结算广告消耗收益
	ProviderSettleAdReward *int64 `json:"provider_settle_ad_reward,omitempty"`
	// 服务商 已结算 除广告消耗收益之外的收益
	ProviderSettleAmountNonAd *int64 `json:"provider_settle_amount_non_ad,omitempty"`
	// 撮合中介已发放佣金(和mcn_amount类似)，单位为厘（元*1000）
	ServiceProviderAmount *int64 `json:"service_provider_amount,omitempty"`
	//
	ServiceProviderAmountIaa *int64 `json:"service_provider_amount_iaa,omitempty"`
	//
	ServiceProviderAmountIap *int64 `json:"service_provider_amount_iap,omitempty"`
	//
	SharePrice *int64 `json:"share_price,omitempty"`
	//
	SharePriceIaa *int64 `json:"share_price_iaa,omitempty"`
	//
	SharePriceIap *int64 `json:"share_price_iap,omitempty"`
	//
	SharePriceTotal *int64 `json:"share_price_total,omitempty"`
	//
	SharePriceTotalIaa *int64 `json:"share_price_total_iaa,omitempty"`
	//
	SharePriceTotalIap *int64 `json:"share_price_total_iap,omitempty"`
	// 达人+服务商 预估 广告消耗收益
	TotalCostAdReward *int64 `json:"total_cost_ad_reward,omitempty"`
	// 达人+服务商 已结算 除广告消耗收益之外的收益
	TotalEstAmountNonAd *int64 `json:"total_est_amount_non_ad,omitempty"`
}
