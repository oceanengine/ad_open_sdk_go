/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnGetContractChallengeAuthorItemListV2V2Response struct for StarMcnGetContractChallengeAuthorItemListV2V2Response
type StarMcnGetContractChallengeAuthorItemListV2V2Response struct {
	//
	Code *int64                                                     `json:"code,omitempty"`
	Data *StarMcnGetContractChallengeAuthorItemListV2V2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
