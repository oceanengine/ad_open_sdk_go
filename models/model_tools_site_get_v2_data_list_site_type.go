/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteGetV2DataListSiteType
type ToolsSiteGetV2DataListSiteType string

// List of tools_site_get_v2_data_list_siteType
const (
	ADL_AB_TEST_ToolsSiteGetV2DataListSiteType               ToolsSiteGetV2DataListSiteType = "ADL_AB_TEST"
	ADL_BIZ_ToolsSiteGetV2DataListSiteType                   ToolsSiteGetV2DataListSiteType = "ADL_BIZ"
	ADL_DCAR_ToolsSiteGetV2DataListSiteType                  ToolsSiteGetV2DataListSiteType = "ADL_DCAR"
	ADL_EDU_ToolsSiteGetV2DataListSiteType                   ToolsSiteGetV2DataListSiteType = "ADL_EDU"
	ADL_FORM_GROUP_ToolsSiteGetV2DataListSiteType            ToolsSiteGetV2DataListSiteType = "ADL_FORM_GROUP"
	ADL_GAME_ToolsSiteGetV2DataListSiteType                  ToolsSiteGetV2DataListSiteType = "ADL_GAME"
	ADL_HOME_ToolsSiteGetV2DataListSiteType                  ToolsSiteGetV2DataListSiteType = "ADL_HOME"
	ADL_LIFE_ToolsSiteGetV2DataListSiteType                  ToolsSiteGetV2DataListSiteType = "ADL_LIFE"
	ADL_NOVEL_ToolsSiteGetV2DataListSiteType                 ToolsSiteGetV2DataListSiteType = "ADL_NOVEL"
	AD_APP_ToolsSiteGetV2DataListSiteType                    ToolsSiteGetV2DataListSiteType = "AD_APP"
	AUTOGEN_CAR_ToolsSiteGetV2DataListSiteType               ToolsSiteGetV2DataListSiteType = "AUTOGEN_CAR"
	AUTOGEN_GAME_ToolsSiteGetV2DataListSiteType              ToolsSiteGetV2DataListSiteType = "AUTOGEN_GAME"
	AUTOGEN_NOVEL_ToolsSiteGetV2DataListSiteType             ToolsSiteGetV2DataListSiteType = "AUTOGEN_NOVEL"
	BUSINESS_FORM_COMMENT_ToolsSiteGetV2DataListSiteType     ToolsSiteGetV2DataListSiteType = "BUSINESS_FORM_COMMENT"
	BUSINESS_FORM_INTERACTIVE_ToolsSiteGetV2DataListSiteType ToolsSiteGetV2DataListSiteType = "BUSINESS_FORM_INTERACTIVE"
	BUSINESS_FORM_SEARCH_ToolsSiteGetV2DataListSiteType      ToolsSiteGetV2DataListSiteType = "BUSINESS_FORM_SEARCH"
	BUSINESS_FORM_V4_ToolsSiteGetV2DataListSiteType          ToolsSiteGetV2DataListSiteType = "BUSINESS_FORM_V4"
	CONSULT_PAGE_ToolsSiteGetV2DataListSiteType              ToolsSiteGetV2DataListSiteType = "CONSULT_PAGE"
	CREATIVE_COMP_ToolsSiteGetV2DataListSiteType             ToolsSiteGetV2DataListSiteType = "CREATIVE_COMP"
	CREATIVE_FORM_ToolsSiteGetV2DataListSiteType             ToolsSiteGetV2DataListSiteType = "CREATIVE_FORM"
	CREATIVE_FORM_OLD_ToolsSiteGetV2DataListSiteType         ToolsSiteGetV2DataListSiteType = "CREATIVE_FORM_OLD"
	DPA_ToolsSiteGetV2DataListSiteType                       ToolsSiteGetV2DataListSiteType = "DPA"
	ENTERPRISE_ToolsSiteGetV2DataListSiteType                ToolsSiteGetV2DataListSiteType = "ENTERPRISE"
	ENTERPRISE_CLOUD_CAR_SHOP_ToolsSiteGetV2DataListSiteType ToolsSiteGetV2DataListSiteType = "ENTERPRISE_CLOUD_CAR_SHOP"
	ENTERPRISE_CLUE_ToolsSiteGetV2DataListSiteType           ToolsSiteGetV2DataListSiteType = "ENTERPRISE_CLUE"
	ENTERPRISE_CONTACT_ToolsSiteGetV2DataListSiteType        ToolsSiteGetV2DataListSiteType = "ENTERPRISE_CONTACT"
	ENTERPRISE_GROUP_BUY_ToolsSiteGetV2DataListSiteType      ToolsSiteGetV2DataListSiteType = "ENTERPRISE_GROUP_BUY"
	ENTERPRISE_HOMEPAGE_ToolsSiteGetV2DataListSiteType       ToolsSiteGetV2DataListSiteType = "ENTERPRISE_HOMEPAGE"
	ENTERPRISE_RESERVE_ToolsSiteGetV2DataListSiteType        ToolsSiteGetV2DataListSiteType = "ENTERPRISE_RESERVE"
	ENTERPRISE_TOOLS_ToolsSiteGetV2DataListSiteType          ToolsSiteGetV2DataListSiteType = "ENTERPRISE_TOOLS"
	FORM_ToolsSiteGetV2DataListSiteType                      ToolsSiteGetV2DataListSiteType = "FORM"
	HEALTH_ToolsSiteGetV2DataListSiteType                    ToolsSiteGetV2DataListSiteType = "HEALTH"
	INNER_EXP_ToolsSiteGetV2DataListSiteType                 ToolsSiteGetV2DataListSiteType = "INNER_EXP"
	LOCAL_INDUSTRY_ToolsSiteGetV2DataListSiteType            ToolsSiteGetV2DataListSiteType = "LOCAL_INDUSTRY"
	MICRO_APP_ToolsSiteGetV2DataListSiteType                 ToolsSiteGetV2DataListSiteType = "MICRO_APP"
	MICRO_APP_SHELL_ToolsSiteGetV2DataListSiteType           ToolsSiteGetV2DataListSiteType = "MICRO_APP_SHELL"
	MICRO_GAME_ToolsSiteGetV2DataListSiteType                ToolsSiteGetV2DataListSiteType = "MICRO_GAME"
	MINIAPP_ToolsSiteGetV2DataListSiteType                   ToolsSiteGetV2DataListSiteType = "MINIAPP"
	MULTI_GOODS_ToolsSiteGetV2DataListSiteType               ToolsSiteGetV2DataListSiteType = "MULTI_GOODS"
	NATIVE_ToolsSiteGetV2DataListSiteType                    ToolsSiteGetV2DataListSiteType = "NATIVE"
	NATIVE_FORM_ToolsSiteGetV2DataListSiteType               ToolsSiteGetV2DataListSiteType = "NATIVE_FORM"
	NO_PUBLISH_SITE_ToolsSiteGetV2DataListSiteType           ToolsSiteGetV2DataListSiteType = "NO_PUBLISH_SITE"
	OCEAN_ENGINE_APP_ToolsSiteGetV2DataListSiteType          ToolsSiteGetV2DataListSiteType = "OCEAN_ENGINE_APP"
	POLL_ToolsSiteGetV2DataListSiteType                      ToolsSiteGetV2DataListSiteType = "POLL"
	PRIVACY_POLICY_ToolsSiteGetV2DataListSiteType            ToolsSiteGetV2DataListSiteType = "PRIVACY_POLICY"
	PROGRAM_ToolsSiteGetV2DataListSiteType                   ToolsSiteGetV2DataListSiteType = "PROGRAM"
	SELF_ToolsSiteGetV2DataListSiteType                      ToolsSiteGetV2DataListSiteType = "SELF"
	SELF_CREATIVE_ToolsSiteGetV2DataListSiteType             ToolsSiteGetV2DataListSiteType = "SELF_CREATIVE"
	SLIDE_ToolsSiteGetV2DataListSiteType                     ToolsSiteGetV2DataListSiteType = "SLIDE"
	STAR_ToolsSiteGetV2DataListSiteType                      ToolsSiteGetV2DataListSiteType = "STAR"
	STAR_API_ToolsSiteGetV2DataListSiteType                  ToolsSiteGetV2DataListSiteType = "STAR_API"
	STORE_SITE_ToolsSiteGetV2DataListSiteType                ToolsSiteGetV2DataListSiteType = "STORE_SITE"
	STREAMING_ToolsSiteGetV2DataListSiteType                 ToolsSiteGetV2DataListSiteType = "STREAMING"
	STRUCTURED_LANDING_PAGE_ToolsSiteGetV2DataListSiteType   ToolsSiteGetV2DataListSiteType = "STRUCTURED_LANDING_PAGE"
	SUBCHAIN_ToolsSiteGetV2DataListSiteType                  ToolsSiteGetV2DataListSiteType = "SUBCHAIN"
	THIRD_EXP_ToolsSiteGetV2DataListSiteType                 ToolsSiteGetV2DataListSiteType = "THIRD_EXP"
	THIRD_SITE_ToolsSiteGetV2DataListSiteType                ToolsSiteGetV2DataListSiteType = "THIRD_SITE"
	ULTRA_ToolsSiteGetV2DataListSiteType                     ToolsSiteGetV2DataListSiteType = "ULTRA"
	UNIVERSAL_ToolsSiteGetV2DataListSiteType                 ToolsSiteGetV2DataListSiteType = "UNIVERSAL"
	WECHAT_APPLET_ToolsSiteGetV2DataListSiteType             ToolsSiteGetV2DataListSiteType = "WECHAT_APPLET"
	WECHAT_GAME_ToolsSiteGetV2DataListSiteType               ToolsSiteGetV2DataListSiteType = "WECHAT_GAME"
)

// All allowed values of ToolsSiteGetV2DataListSiteType enum
var AllowedToolsSiteGetV2DataListSiteTypeEnumValues = []ToolsSiteGetV2DataListSiteType{
	"ADL_AB_TEST",
	"ADL_BIZ",
	"ADL_DCAR",
	"ADL_EDU",
	"ADL_FORM_GROUP",
	"ADL_GAME",
	"ADL_HOME",
	"ADL_LIFE",
	"ADL_NOVEL",
	"AD_APP",
	"AUTOGEN_CAR",
	"AUTOGEN_GAME",
	"AUTOGEN_NOVEL",
	"BUSINESS_FORM_COMMENT",
	"BUSINESS_FORM_INTERACTIVE",
	"BUSINESS_FORM_SEARCH",
	"BUSINESS_FORM_V4",
	"CONSULT_PAGE",
	"CREATIVE_COMP",
	"CREATIVE_FORM",
	"CREATIVE_FORM_OLD",
	"DPA",
	"ENTERPRISE",
	"ENTERPRISE_CLOUD_CAR_SHOP",
	"ENTERPRISE_CLUE",
	"ENTERPRISE_CONTACT",
	"ENTERPRISE_GROUP_BUY",
	"ENTERPRISE_HOMEPAGE",
	"ENTERPRISE_RESERVE",
	"ENTERPRISE_TOOLS",
	"FORM",
	"HEALTH",
	"INNER_EXP",
	"LOCAL_INDUSTRY",
	"MICRO_APP",
	"MICRO_APP_SHELL",
	"MICRO_GAME",
	"MINIAPP",
	"MULTI_GOODS",
	"NATIVE",
	"NATIVE_FORM",
	"NO_PUBLISH_SITE",
	"OCEAN_ENGINE_APP",
	"POLL",
	"PRIVACY_POLICY",
	"PROGRAM",
	"SELF",
	"SELF_CREATIVE",
	"SLIDE",
	"STAR",
	"STAR_API",
	"STORE_SITE",
	"STREAMING",
	"STRUCTURED_LANDING_PAGE",
	"SUBCHAIN",
	"THIRD_EXP",
	"THIRD_SITE",
	"ULTRA",
	"UNIVERSAL",
	"WECHAT_APPLET",
	"WECHAT_GAME",
}

// NewToolsSiteGetV2DataListSiteTypeFromValue returns a pointer to a valid ToolsSiteGetV2DataListSiteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteGetV2DataListSiteTypeFromValue(v string) (*ToolsSiteGetV2DataListSiteType, error) {
	ev := ToolsSiteGetV2DataListSiteType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteGetV2DataListSiteType: valid values are %v", v, AllowedToolsSiteGetV2DataListSiteTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteGetV2DataListSiteType) IsValid() bool {
	for _, existing := range AllowedToolsSiteGetV2DataListSiteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_get_v2_data_list_siteType value
func (v ToolsSiteGetV2DataListSiteType) Ptr() *ToolsSiteGetV2DataListSiteType {
	return &v
}
