/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAnchorListV30DataAnchorListAnchorStatus
type BrandAnchorListV30DataAnchorListAnchorStatus string

// List of brand_anchor_list_v3.0_data_anchor_list_anchor_status
const (
	AUDITING_BrandAnchorListV30DataAnchorListAnchorStatus BrandAnchorListV30DataAnchorListAnchorStatus = "AUDITING"
	DELETE_BrandAnchorListV30DataAnchorListAnchorStatus   BrandAnchorListV30DataAnchorListAnchorStatus = "DELETE"
	PASS_BrandAnchorListV30DataAnchorListAnchorStatus     BrandAnchorListV30DataAnchorListAnchorStatus = "PASS"
	REJECT_BrandAnchorListV30DataAnchorListAnchorStatus   BrandAnchorListV30DataAnchorListAnchorStatus = "REJECT"
)

// Ptr returns reference to brand_anchor_list_v3.0_data_anchor_list_anchor_status value
func (v BrandAnchorListV30DataAnchorListAnchorStatus) Ptr() *BrandAnchorListV30DataAnchorListAnchorStatus {
	return &v
}
