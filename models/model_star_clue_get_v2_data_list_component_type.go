/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
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
	ANCHOR_MICRO_APP_POI_StarClueGetV2DataListComponentType     StarClueGetV2DataListComponentType = "ANCHOR_MICRO_APP_POI"
	POI_StarClueGetV2DataListComponentType                      StarClueGetV2DataListComponentType = "POI"
	ANCHOR_MOVIE_StarClueGetV2DataListComponentType             StarClueGetV2DataListComponentType = "ANCHOR_MOVIE"
	GAME_ANCHOR_StarClueGetV2DataListComponentType              StarClueGetV2DataListComponentType = "GAME_ANCHOR"
	ANCHOR_ECOM_StarClueGetV2DataListComponentType              StarClueGetV2DataListComponentType = "ANCHOR_ECOM"
	CART_StarClueGetV2DataListComponentType                     StarClueGetV2DataListComponentType = "CART"
	BRAND_ANCHOR_StarClueGetV2DataListComponentType             StarClueGetV2DataListComponentType = "BRAND_ANCHOR"
	LINK_StarClueGetV2DataListComponentType                     StarClueGetV2DataListComponentType = "LINK"
	ENTERPRISE_SALON_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ENTERPRISE_SALON"
	ANCHOR_CAR_StarClueGetV2DataListComponentType               StarClueGetV2DataListComponentType = "ANCHOR_CAR"
	ANCHOR_INSURANCE_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ANCHOR_INSURANCE"
	ANCHOR_HOME_StarClueGetV2DataListComponentType              StarClueGetV2DataListComponentType = "ANCHOR_HOME"
	ANCHOR_TELECOM_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "ANCHOR_TELECOM"
	ENTERPRISE_WEDDING_PHOTO_StarClueGetV2DataListComponentType StarClueGetV2DataListComponentType = "ENTERPRISE_WEDDING_PHOTO"
	ANCHOR_INVESTMENT_StarClueGetV2DataListComponentType        StarClueGetV2DataListComponentType = "ANCHOR_INVESTMENT"
	ENTERPRISE_NOVEL_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ENTERPRISE_NOVEL"
	ANCHOR_E_GAME_StarClueGetV2DataListComponentType            StarClueGetV2DataListComponentType = "ANCHOR_E_GAME"
	ALL_StarClueGetV2DataListComponentType                      StarClueGetV2DataListComponentType = "ALL"
	ANCHOR_ESTATE_SERVICE_StarClueGetV2DataListComponentType    StarClueGetV2DataListComponentType = "ANCHOR_ESTATE_SERVICE"
	MICROAPP_ANCHOR_StarClueGetV2DataListComponentType          StarClueGetV2DataListComponentType = "MICROAPP_ANCHOR"
	ENTERPRISE_DOWNLOAD_StarClueGetV2DataListComponentType      StarClueGetV2DataListComponentType = "ENTERPRISE_DOWNLOAD"
	ANCHOR_EDUCATION_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ANCHOR_EDUCATION"
	ENTERPRISE_COUPON_StarClueGetV2DataListComponentType        StarClueGetV2DataListComponentType = "ENTERPRISE_COUPON"
	ENTERPRISE_MICRO_APP_StarClueGetV2DataListComponentType     StarClueGetV2DataListComponentType = "ENTERPRISE_MICRO_APP"
	ANCHOR_XIGUA_StarClueGetV2DataListComponentType             StarClueGetV2DataListComponentType = "ANCHOR_XIGUA"
	ENTERPRISE_DOWNLOAD_APP_StarClueGetV2DataListComponentType  StarClueGetV2DataListComponentType = "ENTERPRISE_DOWNLOAD_APP"
	ENTERPRISE_ORDER_SERVICE_StarClueGetV2DataListComponentType StarClueGetV2DataListComponentType = "ENTERPRISE_ORDER_SERVICE"
	ENTERPRISE_CAR_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "ENTERPRISE_CAR"
	ENTERPRISE_ECOM_StarClueGetV2DataListComponentType          StarClueGetV2DataListComponentType = "ENTERPRISE_ECOM"
	VARIETY_ANCHOR_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "VARIETY_ANCHOR"
	ANCHOR_DOWNLOAD_StarClueGetV2DataListComponentType          StarClueGetV2DataListComponentType = "ANCHOR_DOWNLOAD"
	LIVE_ORDER_COMPONENT_StarClueGetV2DataListComponentType     StarClueGetV2DataListComponentType = "LIVE_ORDER_COMPONENT"
	ANCHOR_TOURISM_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "ANCHOR_TOURISM"
)

// All allowed values of StarClueGetV2DataListComponentType enum
var AllowedStarClueGetV2DataListComponentTypeEnumValues = []StarClueGetV2DataListComponentType{
	"ANCHOR_MICRO_APP_POI",
	"POI",
	"ANCHOR_MOVIE",
	"GAME_ANCHOR",
	"ANCHOR_ECOM",
	"CART",
	"BRAND_ANCHOR",
	"LINK",
	"ENTERPRISE_SALON",
	"ANCHOR_CAR",
	"ANCHOR_INSURANCE",
	"ANCHOR_HOME",
	"ANCHOR_TELECOM",
	"ENTERPRISE_WEDDING_PHOTO",
	"ANCHOR_INVESTMENT",
	"ENTERPRISE_NOVEL",
	"ANCHOR_E_GAME",
	"ALL",
	"ANCHOR_ESTATE_SERVICE",
	"MICROAPP_ANCHOR",
	"ENTERPRISE_DOWNLOAD",
	"ANCHOR_EDUCATION",
	"ENTERPRISE_COUPON",
	"ENTERPRISE_MICRO_APP",
	"ANCHOR_XIGUA",
	"ENTERPRISE_DOWNLOAD_APP",
	"ENTERPRISE_ORDER_SERVICE",
	"ENTERPRISE_CAR",
	"ENTERPRISE_ECOM",
	"VARIETY_ANCHOR",
	"ANCHOR_DOWNLOAD",
	"LIVE_ORDER_COMPONENT",
	"ANCHOR_TOURISM",
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
