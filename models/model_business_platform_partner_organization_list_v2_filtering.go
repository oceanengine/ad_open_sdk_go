/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BusinessPlatformPartnerOrganizationListV2Filtering
type BusinessPlatformPartnerOrganizationListV2Filtering struct {
	// 合作的巨量纵横组织id，最大不超过200
	PartnerOrganizationIds []int64 `json:"partner_organization_ids,omitempty"`
	// 合作状态，默认查询绑定成功的状态<br>枚举值：`BOUND`（已绑定）、`BINDING`（绑定中）、`INVALID`（失效）、`DELETED`（删除）
	Status []*BusinessPlatformPartnerOrganizationListV2FilteringStatus `json:"status,omitempty"`
}
