/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentCompanyBiddingListQueryV2Response struct for AgentCompanyBiddingListQueryV2Response
type AgentCompanyBiddingListQueryV2Response struct {
	//
	Code *int64                                      `json:"code,omitempty"`
	Data *AgentCompanyBiddingListQueryV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
