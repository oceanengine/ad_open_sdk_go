/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueGetV2ResponseDataListInner struct for ToolsClueGetV2ResponseDataListInner
type ToolsClueGetV2ResponseDataListInner struct {
	//
	AdId *string `json:"ad_id,omitempty"`
	//
	AdName *string `json:"ad_name,omitempty"`
	//
	Address *string `json:"address,omitempty"`
	//
	AdvertiserId *string `json:"advertiser_id,omitempty"`
	//
	AdvertiserName *string `json:"advertiser_name,omitempty"`
	//
	Age *int64 `json:"age,omitempty"`
	//
	AllocationStatus *int64 `json:"allocation_status,omitempty"`
	//
	AppName *string `json:"app_name,omitempty"`
	//
	CityName *string `json:"city_name,omitempty"`
	//
	ClueId *string `json:"clue_id,omitempty"`
	//
	ClueOwnerName *string `json:"clue_owner_name,omitempty"`
	//
	ClueReturnStatus *string `json:"clue_return_status,omitempty"`
	//
	ClueSource *int64 `json:"clue_source,omitempty"`
	//
	ClueState *int64 `json:"clue_state,omitempty"`
	//
	ClueStateName *string `json:"clue_state_name,omitempty"`
	//
	ClueType *int64 `json:"clue_type,omitempty"`
	//
	ConvertStatus *string `json:"convert_status,omitempty"`
	//
	CountryName *string `json:"country_name,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	CreateTimeDetail *string `json:"create_time_detail,omitempty"`
	//
	CreativeId *string `json:"creative_id,omitempty"`
	//
	Date *string `json:"date,omitempty"`
	//
	Email *string `json:"email,omitempty"`
	//
	ExtInfo *string `json:"ext_info,omitempty"`
	//
	ExternalUrl *string `json:"external_url,omitempty"`
	//
	ExtraInfo map[string]interface{} `json:"extra_info,omitempty"`
	//
	FollowStateName *string `json:"follow_state_name,omitempty"`
	//
	FormRemark *string `json:"form_remark,omitempty"`
	//
	Gender *int64 `json:"gender,omitempty"`
	//
	IntentionEstimation *string `json:"intention_estimation,omitempty"`
	//
	Location *string `json:"location,omitempty"`
	//
	MidInfo *string `json:"mid_info,omitempty"`
	//
	ModuleId *string `json:"module_id,omitempty"`
	//
	ModuleName *string `json:"module_name,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	ProjectId *string `json:"project_id,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
	//
	PromotionName *string `json:"promotion_name,omitempty"`
	//
	ProvinceName *string `json:"province_name,omitempty"`
	//
	Qq *string `json:"qq,omitempty"`
	//
	Remark *string `json:"remark,omitempty"`
	//
	RemarkDict map[string]interface{} `json:"remark_dict,omitempty"`
	//
	ReqId *string `json:"req_id,omitempty"`
	//
	SiteId *string `json:"site_id,omitempty"`
	//
	SiteName *string                                   `json:"site_name,omitempty"`
	Store    *ToolsClueGetV2ResponseDataListInnerStore `json:"store,omitempty"`
	//
	SystemTags []string `json:"system_tags,omitempty"`
	//
	Tags []string `json:"tags,omitempty"`
	//
	Telephone *string `json:"telephone,omitempty"`
	//
	TelephoneAddress *string `json:"telephone_address,omitempty"`
	//
	Weixin *string `json:"weixin,omitempty"`
}
