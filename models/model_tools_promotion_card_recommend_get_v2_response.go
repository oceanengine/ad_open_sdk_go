/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionCardRecommendGetV2Response struct for ToolsPromotionCardRecommendGetV2Response
type ToolsPromotionCardRecommendGetV2Response struct {
	//
	Code *int64                                        `json:"code,omitempty"`
	Data *ToolsPromotionCardRecommendGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
