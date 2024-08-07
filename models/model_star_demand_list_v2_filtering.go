/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandListV2Filtering
type StarDemandListV2Filtering struct {
	ComponentType *StarDemandListV2FilteringComponentType `json:"component_type,omitempty"`
	//
	Name                    *string                                           `json:"name,omitempty"`
	UniversalOrderStatus    *StarDemandListV2FilteringUniversalOrderStatus    `json:"universal_order_status,omitempty"`
	UniversalSettlementType *StarDemandListV2FilteringUniversalSettlementType `json:"universal_settlement_type,omitempty"`
}
