/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2DeliveryRange
type AudiencePackageCreateV2DeliveryRange string

// List of audience_package_create_v2_delivery_range
const (
	DEFAULT_AudiencePackageCreateV2DeliveryRange   AudiencePackageCreateV2DeliveryRange = "DEFAULT"
	UNIVERSAL_AudiencePackageCreateV2DeliveryRange AudiencePackageCreateV2DeliveryRange = "UNIVERSAL"
	UNION_AudiencePackageCreateV2DeliveryRange     AudiencePackageCreateV2DeliveryRange = "UNION"
)

// Ptr returns reference to audience_package_create_v2_delivery_range value
func (v AudiencePackageCreateV2DeliveryRange) Ptr() *AudiencePackageCreateV2DeliveryRange {
	return &v
}
