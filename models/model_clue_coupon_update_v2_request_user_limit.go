/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponUpdateV2RequestUserLimit
type ClueCouponUpdateV2RequestUserLimit struct {
	//
	DayLimit *int64 `json:"day_limit,omitempty"`
	//
	TotalLimit *int64 `json:"total_limit,omitempty"`
}