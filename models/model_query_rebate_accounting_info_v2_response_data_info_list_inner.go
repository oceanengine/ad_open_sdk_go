/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryRebateAccountingInfoV2ResponseDataInfoListInner struct for QueryRebateAccountingInfoV2ResponseDataInfoListInner
type QueryRebateAccountingInfoV2ResponseDataInfoListInner struct {
	// 代理商ID
	AgentId *int64 `json:"agent_id,omitempty"`
	// 代理商name
	AgentName *string `json:"agent_name,omitempty"`
	// 返点金额 单位 元 保留2位小数
	Amount *float64 `json:"amount,omitempty"`
	// 合同编号
	ContractSerial *string `json:"contract_serial,omitempty"`
	// 媒体签约主体
	ContractSubjectName *string `json:"contract_subject_name,omitempty"`
	// 创建时间 2000-01-01 00:00:00
	CreateTime *string `json:"create_time,omitempty"`
	// 当前审批节点名称 （待代理商审批/待平台复核终审/审批完成）
	CurrentApprovalStatusName *string `json:"current_approval_status_name,omitempty"`
	// 是否已创建返点 1:是 2:否
	IsCreateRebate *int32 `json:"is_create_rebate,omitempty"`
	// 是否已创建返点name (\"是\"、\"否\")
	IsCreateRebateName *string `json:"is_create_rebate_name,omitempty"`
	// 结算季度/月 （1:1月，2:2月，3:3月，4:4月，5:5月，6:6月，7:7月，8:8月，9:9月，10:10月，11:11月，12:12月，13:Q1，14:Q2，15:Q3，16:Q4，17:上半年，18:下半年，19:全年）
	MonthQuarter *int32 `json:"month_quarter,omitempty"`
	// 结算季度/月name
	MonthQuarterName *string `json:"month_quarter_name,omitempty"`
	// 业绩基数 单位 元 保留2位小数
	PerformanceAmount *float64 `json:"performance_amount,omitempty"`
	// 业务线/平台
	Platform *int64 `json:"platform,omitempty"`
	// 业务线/平台name
	PlatformName *string `json:"platform_name,omitempty"`
	// 返点核算信息ID
	RebateAccountingInfoId *int64 `json:"rebate_accounting_info_id,omitempty"`
	// 返点核算信息编号（与平台披露编号一致，建议使用）
	RebateAccountingInfoSerial *string `json:"rebate_accounting_info_serial,omitempty"`
	// 返点流水ID
	RebateBalanceId *int64 `json:"rebate_balance_id,omitempty"`
	// 返点流水编号（与平台披露编号一致，建议使用）
	RebateBalanceSerial *string `json:"rebate_balance_serial,omitempty"`
	// 审批状态code
	Status *int32 `json:"status,omitempty"`
	// 审批状态name
	StatusName *string `json:"status_name,omitempty"`
	// 返点结算类型
	Type *int32 `json:"type,omitempty"`
	// 返点结算类型name
	TypeName *string `json:"type_name,omitempty"`
	// 使用类型name
	UseTypeNames []string `json:"use_type_names,omitempty"`
	// 使用类型list
	UseTypes []int32 `json:"use_types,omitempty"`
	// 结算年份 例：2024
	Year *int32 `json:"year,omitempty"`
}
