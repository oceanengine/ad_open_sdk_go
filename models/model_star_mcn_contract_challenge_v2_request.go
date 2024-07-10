/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnContractChallengeV2Request struct for StarMcnContractChallengeV2Request
type StarMcnContractChallengeV2Request struct {
	//
	DemandId int64 `json:"demand_id"`
	//
	DeveloperId *int64 `json:"developer_id,omitempty"`
	//
	McnProfitRate int32 `json:"mcn_profit_rate"`
	//
	ProfitRateChannel *map[string]int32 `json:"profit_rate_channel,omitempty"`
	//
	StarId int64 `json:"star_id"`
}
