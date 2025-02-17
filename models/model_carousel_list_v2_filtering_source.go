/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselListV2FilteringSource
type CarouselListV2FilteringSource string

// List of carousel_list_v2_filtering_source
const (
	ACCOUNT_PUSH_CarouselListV2FilteringSource       CarouselListV2FilteringSource = "ACCOUNT_PUSH"
	AD_SITE_CarouselListV2FilteringSource            CarouselListV2FilteringSource = "AD_SITE"
	AIC_CarouselListV2FilteringSource                CarouselListV2FilteringSource = "AIC"
	BP_CarouselListV2FilteringSource                 CarouselListV2FilteringSource = "BP"
	CEWEBRITY_CAROUSEL_CarouselListV2FilteringSource CarouselListV2FilteringSource = "CEWEBRITY_CAROUSEL"
	OPEN_API_CarouselListV2FilteringSource           CarouselListV2FilteringSource = "OPEN_API"
)

// Ptr returns reference to carousel_list_v2_filtering_source value
func (v CarouselListV2FilteringSource) Ptr() *CarouselListV2FilteringSource {
	return &v
}
