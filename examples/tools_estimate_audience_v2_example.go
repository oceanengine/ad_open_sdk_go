/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

type ApiOpenApi2ToolsEstimateAudienceGetRequestExample struct {
	Ac                        []string                                          `json:"ac,omitempty"`
	ActionCategories          []int64                                           `json:"action_categories,omitempty"`
	ActionDays                ToolsEstimateAudienceV2ActionDays                 `json:"action_days,omitempty"`
	ActionWords               []int64                                           `json:"action_words,omitempty"`
	ActivateType              []string                                          `json:"activate_type,omitempty"`
	AdTag                     []int64                                           `json:"ad_tag,omitempty"`
	AdvertiserId              int64                                             `json:"advertiser_id,omitempty"`
	Age                       []string                                          `json:"age,omitempty"`
	AndroidOsv                interface{}                                       `json:"android_osv,omitempty"`
	AppBehaviorTarget         interface{}                                       `json:"app_behavior_target,omitempty"`
	AppCategory               []int64                                           `json:"app_category,omitempty"`
	AppIds                    []int64                                           `json:"app_ids,omitempty"`
	ArticleCategory           []string                                          `json:"article_category,omitempty"`
	AudiencePackageId         int64                                             `json:"audience_package_id,omitempty"`
	AutoExtendEnabled         ToolsEstimateAudienceV2AutoExtendEnabled          `json:"auto_extend_enabled,omitempty"`
	AutoExtendTargets         []string                                          `json:"auto_extend_targets,omitempty"`
	AwemeFansNumbers          []int64                                           `json:"aweme_fans_numbers,omitempty"`
	Carrier                   []string                                          `json:"carrier,omitempty"`
	City                      []int64                                           `json:"city,omitempty"`
	DeviceBrand               []string                                          `json:"device_brand,omitempty"`
	DeviceType                []string                                          `json:"device_type,omitempty"`
	District                  interface{}                                       `json:"district,omitempty"`
	DpaLocalAudience          ToolsEstimateAudienceV2DpaLocalAudience           `json:"dpa_local_audience,omitempty"`
	ExcludeCustomActions      []*ToolsBidSuggestV2ExcludeCustomActionsInner     `json:"exclude_custom_actions,omitempty"`
	ExcludeFlowPackage        []int64                                           `json:"exclude_flow_package,omitempty"`
	FilterAwemeAbnormalActive ToolsEstimateAudienceV2FilterAwemeAbnormalActive  `json:"filter_aweme_abnormal_active,omitempty"`
	FilterAwemeFansCount      int64                                             `json:"filter_aweme_fans_count,omitempty"`
	FilterOwnAwemeFans        ToolsEstimateAudienceV2FilterOwnAwemeFans         `json:"filter_own_aweme_fans,omitempty"`
	FlowPackage               []int64                                           `json:"flow_package,omitempty"`
	Gender                    interface{}                                       `json:"gender,omitempty"`
	Geolocation               []*AudiencePackageUpdateV2RequestGeolocationInner `json:"geolocation,omitempty"`
	IncludeCustomActions      []*ToolsBidSuggestV2ExcludeCustomActionsInner     `json:"include_custom_actions,omitempty"`
	InterestActionMode        interface{}                                       `json:"interest_action_mode,omitempty"`
	InterestCategories        []int64                                           `json:"interest_categories,omitempty"`
	InterestTags              []int64                                           `json:"interest_tags,omitempty"`
	InterestWords             []int64                                           `json:"interest_words,omitempty"`
	IosOsv                    interface{}                                       `json:"ios_osv,omitempty"`
	LaunchPrice               []int64                                           `json:"launch_price,omitempty"`
	LocationType              interface{}                                       `json:"location_type,omitempty"`
	Platform                  []string                                          `json:"platform,omitempty"`
	RetargetingTags           []int64                                           `json:"retargeting_tags,omitempty"`
	RetargetingTagsExclude    []int64                                           `json:"retargeting_tags_exclude,omitempty"`
	RetargetingTagsInclude    []int64                                           `json:"retargeting_tags_include,omitempty"`
	RetargetingType           interface{}                                       `json:"retargeting_type,omitempty"`
	SuperiorPopularityType    interface{}                                       `json:"superior_popularity_type,omitempty"`
	UserType                  []int64                                           `json:"user_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/estimate_audience/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsEstimateAudienceGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsEstimateAudienceV2Api().
		Get(ctx).
		AccessToken(accessToken).
		Ac(request.Ac).ActionCategories(request.ActionCategories).ActionDays(request.ActionDays).ActionWords(request.ActionWords).ActivateType(request.ActivateType).AdTag(request.AdTag).AdvertiserId(request.AdvertiserId).Age(request.Age).AndroidOsv(request.AndroidOsv).AppBehaviorTarget(request.AppBehaviorTarget).AppCategory(request.AppCategory).AppIds(request.AppIds).ArticleCategory(request.ArticleCategory).AudiencePackageId(request.AudiencePackageId).AutoExtendEnabled(request.AutoExtendEnabled).AutoExtendTargets(request.AutoExtendTargets).AwemeFansNumbers(request.AwemeFansNumbers).Carrier(request.Carrier).City(request.City).DeviceBrand(request.DeviceBrand).DeviceType(request.DeviceType).District(request.District).DpaLocalAudience(request.DpaLocalAudience).ExcludeCustomActions(request.ExcludeCustomActions).ExcludeFlowPackage(request.ExcludeFlowPackage).FilterAwemeAbnormalActive(request.FilterAwemeAbnormalActive).FilterAwemeFansCount(request.FilterAwemeFansCount).FilterOwnAwemeFans(request.FilterOwnAwemeFans).FlowPackage(request.FlowPackage).Gender(request.Gender).Geolocation(request.Geolocation).IncludeCustomActions(request.IncludeCustomActions).InterestActionMode(request.InterestActionMode).InterestCategories(request.InterestCategories).InterestTags(request.InterestTags).InterestWords(request.InterestWords).IosOsv(request.IosOsv).LaunchPrice(request.LaunchPrice).LocationType(request.LocationType).Platform(request.Platform).RetargetingTags(request.RetargetingTags).RetargetingTagsExclude(request.RetargetingTagsExclude).RetargetingTagsInclude(request.RetargetingTagsInclude).RetargetingType(request.RetargetingType).SuperiorPopularityType(request.SuperiorPopularityType).UserType(request.UserType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
