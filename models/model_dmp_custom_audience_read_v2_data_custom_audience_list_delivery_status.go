/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus
type DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus string

// List of dmp_custom_audience_read_v2_data_custom_audience_list_delivery_status
const (
	CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE_DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus    DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH_DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH_DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus    DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE_DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus  DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE"
)

// Ptr returns reference to dmp_custom_audience_read_v2_data_custom_audience_list_delivery_status value
func (v DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus) Ptr() *DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus {
	return &v
}
