/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAgentGetV2V2ResponseDataListInner struct for ReportAgentGetV2V2ResponseDataListInner
type ReportAgentGetV2V2ResponseDataListInner struct {
	AccountSource *ReportAgentGetV2V2DataListAccountSource `json:"account_source,omitempty"`
	AccountStatus *ReportAgentGetV2V2DataListAccountStatus `json:"account_status,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AdvertiserName *string `json:"advertiser_name,omitempty"`
	//
	AgentId *int64 `json:"agent_id,omitempty"`
	//
	AgentName *string `json:"agent_name,omitempty"`
	//
	Arpu3cd *float64 `json:"arpu_3cd,omitempty"`
	//
	AuditPassTime *string `json:"audit_pass_time,omitempty"`
	//
	BidCost *float64 `json:"bid_cost,omitempty"`
	//
	BrandCost *float64 `json:"brand_cost,omitempty"`
	//
	CashBalance *float64 `json:"cash_balance,omitempty"`
	//
	CashCost *float64 `json:"cash_cost,omitempty"`
	//
	Click *int64 `json:"click,omitempty"`
	//
	CompanyId *int64 `json:"company_id,omitempty"`
	//
	CompanyName *string `json:"company_name,omitempty"`
	//
	Cost *float64 `json:"cost,omitempty"`
	//
	CpaCost *float64 `json:"cpa_cost,omitempty"`
	//
	Cpc *float64 `json:"cpc,omitempty"`
	//
	CpcCost *float64 `json:"cpc_cost,omitempty"`
	//
	Cpm *float64 `json:"cpm,omitempty"`
	//
	CpmCost *float64 `json:"cpm_cost,omitempty"`
	//
	CptCost *float64 `json:"cpt_cost,omitempty"`
	//
	CpvCost *float64 `json:"cpv_cost,omitempty"`
	//
	Ctr          *float64                                `json:"ctr,omitempty"`
	CustomerType *ReportAgentGetV2V2DataListCustomerType `json:"customer_type,omitempty"`
	//
	FirstCostTime *string `json:"first_cost_time,omitempty"`
	//
	FirstIndustry *string `json:"first_industry,omitempty"`
	//
	FirstRechargeAmount *float64 `json:"first_recharge_amount,omitempty"`
	//
	FirstRechargeTime *string `json:"first_recharge_time,omitempty"`
	//
	GdCost *float64 `json:"gd_cost,omitempty"`
	//
	GrantBalance *float64 `json:"grant_balance,omitempty"`
	//
	GrantsCost *float64 `json:"grants_cost,omitempty"`
	//
	HistoryCashCost *float64 `json:"history_cash_cost,omitempty"`
	//
	HistoryCost *float64 `json:"history_cost,omitempty"`
	//
	LastRechargeTime *string `json:"last_recharge_time,omitempty"`
	//
	LastRenewAmount *float64 `json:"last_renew_amount,omitempty"`
	//
	LastRenewTime *string `json:"last_renew_time,omitempty"`
	//
	OcpcCost *float64 `json:"ocpc_cost,omitempty"`
	//
	OcpmCost *float64 `json:"ocpm_cost,omitempty"`
	//
	QueryEndTime *int64 `json:"query_end_time,omitempty"`
	//
	RegisterDays *float64 `json:"register_days,omitempty"`
	//
	RegisterTime *string `json:"register_time,omitempty"`
	//
	SecondIndustry *string `json:"second_industry,omitempty"`
	//
	Show *int64 `json:"show,omitempty"`
	//
	TodayCost *float64 `json:"today_cost,omitempty"`
	//
	TodayXHourClick *int64 `json:"today_x_hour_click,omitempty"`
	//
	TodayXHourConvert *int64 `json:"today_x_hour_convert,omitempty"`
	//
	TodayXHourCost *float64 `json:"today_x_hour_cost,omitempty"`
	//
	TodayXHourCpa *float64 `json:"today_x_hour_cpa,omitempty"`
	//
	TodayXHourCpc *float64 `json:"today_x_hour_cpc,omitempty"`
	//
	TotalBalance *float64 `json:"total_balance,omitempty"`
	//
	TotalRenewAmount *float64 `json:"total_renew_amount,omitempty"`
	//
	TotalRenewNumber *float64 `json:"total_renew_number,omitempty"`
	//
	TransferAmount *float64 `json:"transfer_amount,omitempty"`
	//
	TransferCount *int64 `json:"transfer_count,omitempty"`
	//
	ValidBalance *float64 `json:"valid_balance,omitempty"`
	//
	XHourClickDiff *int64 `json:"x_hour_click_diff,omitempty"`
	//
	XHourConvertDiff *int64 `json:"x_hour_convert_diff,omitempty"`
	//
	XHourCostDiff *float64 `json:"x_hour_cost_diff,omitempty"`
	//
	XHourCpaDiff *float64 `json:"x_hour_cpa_diff,omitempty"`
	//
	XHourCpcDiff *float64 `json:"x_hour_cpc_diff,omitempty"`
	//
	YesterdayAdNumber *float64 `json:"yesterday_ad_number,omitempty"`
	//
	YesterdayBidCashCost *float64 `json:"yesterday_bid_cash_cost,omitempty"`
	//
	YesterdayBrandCashCost *float64 `json:"yesterday_brand_cash_cost,omitempty"`
	//
	YesterdayCashCost *float64 `json:"yesterday_cash_cost,omitempty"`
	//
	YesterdayCost *float64 `json:"yesterday_cost,omitempty"`
	//
	YesterdayCostDiff *float64 `json:"yesterday_cost_diff,omitempty"`
	//
	YesterdayGrantsCost *float64 `json:"yesterday_grants_cost,omitempty"`
	//
	YesterdayMomCost *float64 `json:"yesterday_mom_cost,omitempty"`
	//
	YesterdayXHourClick *int64 `json:"yesterday_x_hour_click,omitempty"`
	//
	YesterdayXHourConvert *int64 `json:"yesterday_x_hour_convert,omitempty"`
	//
	YesterdayXHourCost *float64 `json:"yesterday_x_hour_cost,omitempty"`
	//
	YesterdayXHourCpa *float64 `json:"yesterday_x_hour_cpa,omitempty"`
	//
	YesterdayXHourCpc *float64 `json:"yesterday_x_hour_cpc,omitempty"`
}
