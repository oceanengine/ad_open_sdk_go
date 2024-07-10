/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetV2ListV2ResponseDataAssetListInner struct for DpaAssetV2ListV2ResponseDataAssetListInner
type DpaAssetV2ListV2ResponseDataAssetListInner struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 投放条件创建时间，格式: yyyy-MM-dd HH:mm:ss
	AssetCreateTime string `json:"asset_create_time"`
	// 投放条件id
	AssetId int64 `json:"asset_id"`
	// 投放条件最近一次修改时间，格式: yyyy-MM-dd HH:mm:ss
	AssetModifyTime string `json:"asset_modify_time"`
	// 投放条件名称
	AssetName   string                                   `json:"asset_name"`
	AssetType   DpaAssetV2ListV2DataAssetListAssetType   `json:"asset_type"`
	AuditStatus DpaAssetV2ListV2DataAssetListAuditStatus `json:"audit_status"`
	Source      DpaAssetV2ListV2DataAssetListSource      `json:"source"`
	Status      DpaAssetV2ListV2DataAssetListStatus      `json:"status"`
	// 线索商品id
	UniqueProductId int64 `json:"unique_product_id"`
}
