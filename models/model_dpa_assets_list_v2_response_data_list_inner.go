/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetsListV2ResponseDataListInner struct for DpaAssetsListV2ResponseDataListInner
type DpaAssetsListV2ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AssetCreateTime **string `json:"asset_create_time,omitempty"`
	//
	AssetId *int64 `json:"asset_id,omitempty"`
	//
	AssetModifyTime **string                            `json:"asset_modify_time,omitempty"`
	AssetType       *DpaAssetsListV2DataListAssetType   `json:"asset_type,omitempty"`
	AuditStatus     *DpaAssetsListV2DataListAuditStatus `json:"audit_status,omitempty"`
	//
	PlatformId *int64 `json:"platform_id,omitempty"`
	//
	ProductId *int64                         `json:"product_id,omitempty"`
	Source    *DpaAssetsListV2DataListSource `json:"source,omitempty"`
	Status    *DpaAssetsListV2DataListStatus `json:"status,omitempty"`
}
