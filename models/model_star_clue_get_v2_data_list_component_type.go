/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarClueGetV2DataListComponentType
type StarClueGetV2DataListComponentType string

// List of star_clue_get_v2_data_list_component_type
const (
	ANCHOR_INVESTMENT_StarClueGetV2DataListComponentType        StarClueGetV2DataListComponentType = "ANCHOR_INVESTMENT"
	GAME_ANCHOR_StarClueGetV2DataListComponentType              StarClueGetV2DataListComponentType = "GAME_ANCHOR"
	ANCHOR_INSURANCE_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ANCHOR_INSURANCE"
	LINK_StarClueGetV2DataListComponentType                     StarClueGetV2DataListComponentType = "LINK"
	ALL_StarClueGetV2DataListComponentType                      StarClueGetV2DataListComponentType = "ALL"
	ANCHOR_TELECOM_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "ANCHOR_TELECOM"
	BRAND_ANCHOR_StarClueGetV2DataListComponentType             StarClueGetV2DataListComponentType = "BRAND_ANCHOR"
	ANCHOR_ESTATE_SERVICE_StarClueGetV2DataListComponentType    StarClueGetV2DataListComponentType = "ANCHOR_ESTATE_SERVICE"
	POI_StarClueGetV2DataListComponentType                      StarClueGetV2DataListComponentType = "POI"
	ENTERPRISE_DOWNLOAD_StarClueGetV2DataListComponentType      StarClueGetV2DataListComponentType = "ENTERPRISE_DOWNLOAD"
	ANCHOR_HOME_StarClueGetV2DataListComponentType              StarClueGetV2DataListComponentType = "ANCHOR_HOME"
	ANCHOR_DOWNLOAD_StarClueGetV2DataListComponentType          StarClueGetV2DataListComponentType = "ANCHOR_DOWNLOAD"
	ENTERPRISE_WEDDING_PHOTO_StarClueGetV2DataListComponentType StarClueGetV2DataListComponentType = "ENTERPRISE_WEDDING_PHOTO"
	ANCHOR_ECOM_StarClueGetV2DataListComponentType              StarClueGetV2DataListComponentType = "ANCHOR_ECOM"
	ANCHOR_EDUCATION_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ANCHOR_EDUCATION"
	ANCHOR_MOVIE_StarClueGetV2DataListComponentType             StarClueGetV2DataListComponentType = "ANCHOR_MOVIE"
	ENTERPRISE_SALON_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ENTERPRISE_SALON"
	ENTERPRISE_ECOM_StarClueGetV2DataListComponentType          StarClueGetV2DataListComponentType = "ENTERPRISE_ECOM"
	ANCHOR_XIGUA_StarClueGetV2DataListComponentType             StarClueGetV2DataListComponentType = "ANCHOR_XIGUA"
	MICROAPP_ANCHOR_StarClueGetV2DataListComponentType          StarClueGetV2DataListComponentType = "MICROAPP_ANCHOR"
	ANCHOR_CAR_StarClueGetV2DataListComponentType               StarClueGetV2DataListComponentType = "ANCHOR_CAR"
	LIVE_ORDER_COMPONENT_StarClueGetV2DataListComponentType     StarClueGetV2DataListComponentType = "LIVE_ORDER_COMPONENT"
	ENTERPRISE_NOVEL_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ENTERPRISE_NOVEL"
	CART_StarClueGetV2DataListComponentType                     StarClueGetV2DataListComponentType = "CART"
	ENTERPRISE_MICRO_APP_StarClueGetV2DataListComponentType     StarClueGetV2DataListComponentType = "ENTERPRISE_MICRO_APP"
	ENTERPRISE_COUPON_StarClueGetV2DataListComponentType        StarClueGetV2DataListComponentType = "ENTERPRISE_COUPON"
	ANCHOR_E_GAME_StarClueGetV2DataListComponentType            StarClueGetV2DataListComponentType = "ANCHOR_E_GAME"
	ANCHOR_MICRO_APP_POI_StarClueGetV2DataListComponentType     StarClueGetV2DataListComponentType = "ANCHOR_MICRO_APP_POI"
	ANCHOR_TOURISM_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "ANCHOR_TOURISM"
	ENTERPRISE_ORDER_SERVICE_StarClueGetV2DataListComponentType StarClueGetV2DataListComponentType = "ENTERPRISE_ORDER_SERVICE"
	ENTERPRISE_CAR_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "ENTERPRISE_CAR"
	ENTERPRISE_DOWNLOAD_APP_StarClueGetV2DataListComponentType  StarClueGetV2DataListComponentType = "ENTERPRISE_DOWNLOAD_APP"
	VARIETY_ANCHOR_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "VARIETY_ANCHOR"
)

// All allowed values of StarClueGetV2DataListComponentType enum
var AllowedStarClueGetV2DataListComponentTypeEnumValues = []StarClueGetV2DataListComponentType{
	"ANCHOR_INVESTMENT",
	"GAME_ANCHOR",
	"ANCHOR_INSURANCE",
	"LINK",
	"ALL",
	"ANCHOR_TELECOM",
	"BRAND_ANCHOR",
	"ANCHOR_ESTATE_SERVICE",
	"POI",
	"ENTERPRISE_DOWNLOAD",
	"ANCHOR_HOME",
	"ANCHOR_DOWNLOAD",
	"ENTERPRISE_WEDDING_PHOTO",
	"ANCHOR_ECOM",
	"ANCHOR_EDUCATION",
	"ANCHOR_MOVIE",
	"ENTERPRISE_SALON",
	"ENTERPRISE_ECOM",
	"ANCHOR_XIGUA",
	"MICROAPP_ANCHOR",
	"ANCHOR_CAR",
	"LIVE_ORDER_COMPONENT",
	"ENTERPRISE_NOVEL",
	"CART",
	"ENTERPRISE_MICRO_APP",
	"ENTERPRISE_COUPON",
	"ANCHOR_E_GAME",
	"ANCHOR_MICRO_APP_POI",
	"ANCHOR_TOURISM",
	"ENTERPRISE_ORDER_SERVICE",
	"ENTERPRISE_CAR",
	"ENTERPRISE_DOWNLOAD_APP",
	"VARIETY_ANCHOR",
}

// NewStarClueGetV2DataListComponentTypeFromValue returns a pointer to a valid StarClueGetV2DataListComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarClueGetV2DataListComponentTypeFromValue(v string) (*StarClueGetV2DataListComponentType, error) {
	ev := StarClueGetV2DataListComponentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarClueGetV2DataListComponentType: valid values are %v", v, AllowedStarClueGetV2DataListComponentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarClueGetV2DataListComponentType) IsValid() bool {
	for _, existing := range AllowedStarClueGetV2DataListComponentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_clue_get_v2_data_list_component_type value
func (v StarClueGetV2DataListComponentType) Ptr() *StarClueGetV2DataListComponentType {
	return &v
}
