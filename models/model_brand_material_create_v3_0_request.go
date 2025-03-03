/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandMaterialCreateV30Request struct for BrandMaterialCreateV30Request
type BrandMaterialCreateV30Request struct {
	//
	AdvertiserId        *int64                                     `json:"advertiser_id,omitempty"`
	CategoryInfo        BrandMaterialCreateV30RequestCategoryInfo  `json:"category_info"`
	CreativeDisplayMode *BrandMaterialCreateV30CreativeDisplayMode `json:"creative_display_mode,omitempty"`
	//
	MaterialList []*BrandMaterialCreateV30RequestMaterialListInner `json:"material_list"`
	//
	OrderId      int64                                      `json:"order_id"`
	TrackUrlInfo *BrandMaterialCreateV30RequestTrackUrlInfo `json:"track_url_info,omitempty"`
}
