/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOperationLogQueryV30ObjectType
type BrandOperationLogQueryV30ObjectType string

// List of brand_operation_log_query_v3.0_object_type
const (
	CAMPAIGN_BrandOperationLogQueryV30ObjectType BrandOperationLogQueryV30ObjectType = "Campaign"
	ORDER_BrandOperationLogQueryV30ObjectType    BrandOperationLogQueryV30ObjectType = "Order"
)

// Ptr returns reference to brand_operation_log_query_v3.0_object_type value
func (v BrandOperationLogQueryV30ObjectType) Ptr() *BrandOperationLogQueryV30ObjectType {
	return &v
}
