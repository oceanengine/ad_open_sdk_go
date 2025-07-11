/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductUpdateV2ProductInfoSpuSource
type DpaProductUpdateV2ProductInfoSpuSource string

// List of dpa_product_update_v2_product_info_spu_source
const (
	YUNTU_SPU_DpaProductUpdateV2ProductInfoSpuSource          DpaProductUpdateV2ProductInfoSpuSource = "YUNTU_SPU"
	OUTER_PLATFORM_SPU_DpaProductUpdateV2ProductInfoSpuSource DpaProductUpdateV2ProductInfoSpuSource = "OUTER_PLATFORM_SPU"
	WITHOUT_SPU_DpaProductUpdateV2ProductInfoSpuSource        DpaProductUpdateV2ProductInfoSpuSource = "WITHOUT_SPU"
)

// Ptr returns reference to dpa_product_update_v2_product_info_spu_source value
func (v DpaProductUpdateV2ProductInfoSpuSource) Ptr() *DpaProductUpdateV2ProductInfoSpuSource {
	return &v
}
