/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus
type DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus string

// List of dmp_data_source_read_v2_data_data_list_default_audience_delivery_status
const (
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH_DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus    DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE_DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus  DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE_DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus    DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH_DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH"
)

// Ptr returns reference to dmp_data_source_read_v2_data_data_list_default_audience_delivery_status value
func (v DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus) Ptr() *DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus {
	return &v
}
