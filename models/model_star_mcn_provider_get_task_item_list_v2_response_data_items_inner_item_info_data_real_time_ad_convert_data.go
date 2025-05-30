/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnProviderGetTaskItemListV2ResponseDataItemsInnerItemInfoDataRealTimeAdConvertData 广告流量实时转化数据
type StarMcnProviderGetTaskItemListV2ResponseDataItemsInnerItemInfoDataRealTimeAdConvertData struct {
	// 激活数
	ActiveCnt *int64 `json:"active_cnt,omitempty"`
	// 激活成本 cpa（单位 厘）
	CostPerActive *int64 `json:"cost_per_active,omitempty"`
	// cpfp 首次付费成本（单位 厘）
	CostPerFirstPay *int64 `json:"cost_per_first_pay,omitempty"`
	// 次日留存成本（单位 厘）
	CostPerNextDayOpen *int64 `json:"cost_per_next_day_open,omitempty"`
	// 付费成本（单位 厘）
	CostPerPay *int64 `json:"cost_per_pay,omitempty"`
	// 注册成本 cpr（单位 厘）
	CostPerRegister *int64 `json:"cost_per_register,omitempty"`
	// 首次付费次数
	FirstPayCnt *int64 `json:"first_pay_cnt,omitempty"`
	// 次日留存数
	NextDayOpenCnt *int64 `json:"next_day_open_cnt,omitempty"`
	// 1: optional i64 ctr //组件点击率 (点击量/展示量) （单位：万分比 小数扩大了1W倍） 2: optional i64 click_cnt //点击量 付费次数
	PayCnt *int64 `json:"pay_cnt,omitempty"`
	// 注册数
	RegisterCnt *int64 `json:"register_cnt,omitempty"`
}
