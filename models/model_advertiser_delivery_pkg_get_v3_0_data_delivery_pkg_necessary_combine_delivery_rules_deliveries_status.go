/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus
type AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus string

// List of advertiser_delivery_pkg_get_v3.0_data_delivery_pkg_necessary_combine_delivery_rules_deliveries_status
const (
	STATUS_CONFIRM_AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus         AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus    AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus      AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus    AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus = "STATUS_WAIT_CONFIRM"
)

// Ptr returns reference to advertiser_delivery_pkg_get_v3.0_data_delivery_pkg_necessary_combine_delivery_rules_deliveries_status value
func (v AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus) Ptr() *AdvertiserDeliveryPkgGetV30DataDeliveryPkgNecessaryCombineDeliveryRulesDeliveriesStatus {
	return &v
}
