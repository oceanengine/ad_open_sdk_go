/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsBidSuggestV2ApiService ToolsBidSuggestV2Api service
type ToolsBidSuggestV2ApiService service

type ApiOpenApi2ToolsBidSuggestGetRequest struct {
	ctx                       context.Context
	ApiService                *ToolsBidSuggestV2ApiService
	ac                        *[]string
	actionCategories          *[]int64
	actionDays                *ToolsBidSuggestV2ActionDays
	actionScene               *[]string
	actionWords               *[]int64
	activateType              *[]string
	adId                      *int64
	adTag                     *[]int64
	advertiserId              *int64
	age                       *[]string
	androidOsv                *interface{}
	appBehaviorTarget         *interface{}
	appCategory               *[]int64
	appIds                    *[]int64
	articleCategory           *[]string
	audiencePackageId         *int64
	autoExtendTargets         *[]string
	awemeFanAccounts          *[]int64
	awemeFanBehaviors         *[]string
	awemeFanCategories        *[]int64
	awemeFanTimeScope         *interface{}
	awemeFansNumbers          *[]int64
	bidMode                   *interface{}
	budget                    *int64
	budgetMode                *interface{}
	businessIds               *[]int64
	campaignId                *int64
	career                    *[]string
	carrier                   *[]string
	city                      *[]int64
	convertId                 *int64
	convertedTimeDuration     *interface{}
	deviceBrand               *[]string
	deviceType                *[]string
	district                  *interface{}
	dpaLocalAudience          *ToolsBidSuggestV2DpaLocalAudience
	dpaRtaRecommendType       *interface{}
	dpaRtaSwitch              *interface{}
	excludeCustomActions      *[]*ToolsBidSuggestV2ExcludeCustomActionsInner
	excludeFlowPackage        *[]int64
	filterAwemeAbnormalActive *ToolsBidSuggestV2FilterAwemeAbnormalActive
	filterAwemeFansCount      *int64
	filterOwnAwemeFans        *ToolsBidSuggestV2FilterOwnAwemeFans
	flowControlMode           *interface{}
	flowPackage               *[]int64
	gender                    *interface{}
	geolocation               *[]*AudiencePackageUpdateV2RequestGeolocationInner
	hideIfConverted           *interface{}
	hideIfExists              *ToolsBidSuggestV2HideIfExists
	includeCustomActions      *[]*ToolsBidSuggestV2ExcludeCustomActionsInner
	interestActionMode        *interface{}
	interestCategories        *[]int64
	interestTags              *[]int64
	interestWords             *[]int64
	iosOsv                    *interface{}
	launchPrice               *[]int64
	locationType              *interface{}
	platform                  *[]string
	pricing                   *interface{}
	regionVersion             *string
	retargetingTags           *[]int64
	retargetingTagsExclude    *[]int64
	retargetingTagsInclude    *[]int64
	retargetingType           *interface{}
	scheduleType              *interface{}
	superiorPopularityType    *interface{}
	userType                  *[]int64
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Ac(ac []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.ac = &ac
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ActionCategories(actionCategories []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.actionCategories = &actionCategories
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ActionDays(actionDays ToolsBidSuggestV2ActionDays) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.actionDays = &actionDays
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ActionScene(actionScene []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.actionScene = &actionScene
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ActionWords(actionWords []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.actionWords = &actionWords
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ActivateType(activateType []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.activateType = &activateType
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AdId(adId int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AdTag(adTag []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.adTag = &adTag
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Age(age []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.age = &age
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AndroidOsv(androidOsv interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.androidOsv = &androidOsv
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AppBehaviorTarget(appBehaviorTarget interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.appBehaviorTarget = &appBehaviorTarget
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AppCategory(appCategory []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.appCategory = &appCategory
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AppIds(appIds []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.appIds = &appIds
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ArticleCategory(articleCategory []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.articleCategory = &articleCategory
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AudiencePackageId(audiencePackageId int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.audiencePackageId = &audiencePackageId
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AutoExtendTargets(autoExtendTargets []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.autoExtendTargets = &autoExtendTargets
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AwemeFanAccounts(awemeFanAccounts []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.awemeFanAccounts = &awemeFanAccounts
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AwemeFanBehaviors(awemeFanBehaviors []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.awemeFanBehaviors = &awemeFanBehaviors
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AwemeFanCategories(awemeFanCategories []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.awemeFanCategories = &awemeFanCategories
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AwemeFanTimeScope(awemeFanTimeScope interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.awemeFanTimeScope = &awemeFanTimeScope
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AwemeFansNumbers(awemeFansNumbers []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.awemeFansNumbers = &awemeFansNumbers
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) BidMode(bidMode interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.bidMode = &bidMode
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Budget(budget int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.budget = &budget
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) BudgetMode(budgetMode interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.budgetMode = &budgetMode
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) BusinessIds(businessIds []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.businessIds = &businessIds
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) CampaignId(campaignId int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.campaignId = &campaignId
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Career(career []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.career = &career
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Carrier(carrier []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.carrier = &carrier
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) City(city []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.city = &city
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ConvertId(convertId int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.convertId = &convertId
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ConvertedTimeDuration(convertedTimeDuration interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.convertedTimeDuration = &convertedTimeDuration
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) DeviceBrand(deviceBrand []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.deviceBrand = &deviceBrand
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) DeviceType(deviceType []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.deviceType = &deviceType
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) District(district interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.district = &district
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) DpaLocalAudience(dpaLocalAudience ToolsBidSuggestV2DpaLocalAudience) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.dpaLocalAudience = &dpaLocalAudience
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) DpaRtaRecommendType(dpaRtaRecommendType interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.dpaRtaRecommendType = &dpaRtaRecommendType
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) DpaRtaSwitch(dpaRtaSwitch interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.dpaRtaSwitch = &dpaRtaSwitch
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ExcludeCustomActions(excludeCustomActions []*ToolsBidSuggestV2ExcludeCustomActionsInner) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.excludeCustomActions = &excludeCustomActions
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ExcludeFlowPackage(excludeFlowPackage []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.excludeFlowPackage = &excludeFlowPackage
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) FilterAwemeAbnormalActive(filterAwemeAbnormalActive ToolsBidSuggestV2FilterAwemeAbnormalActive) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.filterAwemeAbnormalActive = &filterAwemeAbnormalActive
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) FilterAwemeFansCount(filterAwemeFansCount int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.filterAwemeFansCount = &filterAwemeFansCount
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) FilterOwnAwemeFans(filterOwnAwemeFans ToolsBidSuggestV2FilterOwnAwemeFans) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.filterOwnAwemeFans = &filterOwnAwemeFans
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) FlowControlMode(flowControlMode interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.flowControlMode = &flowControlMode
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) FlowPackage(flowPackage []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.flowPackage = &flowPackage
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Gender(gender interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.gender = &gender
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Geolocation(geolocation []*AudiencePackageUpdateV2RequestGeolocationInner) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.geolocation = &geolocation
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) HideIfConverted(hideIfConverted interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.hideIfConverted = &hideIfConverted
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) HideIfExists(hideIfExists ToolsBidSuggestV2HideIfExists) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.hideIfExists = &hideIfExists
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) IncludeCustomActions(includeCustomActions []*ToolsBidSuggestV2ExcludeCustomActionsInner) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.includeCustomActions = &includeCustomActions
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) InterestActionMode(interestActionMode interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.interestActionMode = &interestActionMode
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) InterestCategories(interestCategories []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.interestCategories = &interestCategories
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) InterestTags(interestTags []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.interestTags = &interestTags
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) InterestWords(interestWords []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.interestWords = &interestWords
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) IosOsv(iosOsv interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.iosOsv = &iosOsv
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) LaunchPrice(launchPrice []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.launchPrice = &launchPrice
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) LocationType(locationType interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.locationType = &locationType
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Platform(platform []string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.platform = &platform
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Pricing(pricing interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.pricing = &pricing
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) RegionVersion(regionVersion string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.regionVersion = &regionVersion
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) RetargetingTags(retargetingTags []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.retargetingTags = &retargetingTags
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) RetargetingTagsExclude(retargetingTagsExclude []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.retargetingTagsExclude = &retargetingTagsExclude
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) RetargetingTagsInclude(retargetingTagsInclude []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.retargetingTagsInclude = &retargetingTagsInclude
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) RetargetingType(retargetingType interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.retargetingType = &retargetingType
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) ScheduleType(scheduleType interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.scheduleType = &scheduleType
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) SuperiorPopularityType(superiorPopularityType interface{}) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.superiorPopularityType = &superiorPopularityType
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) UserType(userType []int64) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.userType = &userType
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) Execute() (*ToolsBidSuggestV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsBidSuggestGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsBidSuggestGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsBidSuggestGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsBidSuggestGet Method for OpenApi2ToolsBidSuggestGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsBidSuggestGetRequest
*/
func (a *ToolsBidSuggestV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsBidSuggestGetRequest {
	return &ApiOpenApi2ToolsBidSuggestGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsBidSuggestV2Response
func (a *ToolsBidSuggestV2ApiService) getExecute(r *ApiOpenApi2ToolsBidSuggestGetRequest) (*ToolsBidSuggestV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsBidSuggestV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/bid/suggest/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ac", r.ac)
	}
	if r.actionCategories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_categories", r.actionCategories)
	}
	if r.actionDays != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_days", r.actionDays)
	}
	if r.actionScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_scene", r.actionScene)
	}
	if r.actionWords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_words", r.actionWords)
	}
	if r.activateType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activate_type", r.activateType)
	}
	if r.adId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	}
	if r.adTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_tag", r.adTag)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.age != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "age", r.age)
	}
	if r.androidOsv != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "android_osv", r.androidOsv)
	}
	if r.appBehaviorTarget != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_behavior_target", r.appBehaviorTarget)
	}
	if r.appCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_category", r.appCategory)
	}
	if r.appIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_ids", r.appIds)
	}
	if r.articleCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "article_category", r.articleCategory)
	}
	if r.audiencePackageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audience_package_id", r.audiencePackageId)
	}
	if r.autoExtendTargets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auto_extend_targets", r.autoExtendTargets)
	}
	if r.awemeFanAccounts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_fan_accounts", r.awemeFanAccounts)
	}
	if r.awemeFanBehaviors != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_fan_behaviors", r.awemeFanBehaviors)
	}
	if r.awemeFanCategories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_fan_categories", r.awemeFanCategories)
	}
	if r.awemeFanTimeScope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_fan_time_scope", r.awemeFanTimeScope)
	}
	if r.awemeFansNumbers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_fans_numbers", r.awemeFansNumbers)
	}
	if r.bidMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bid_mode", r.bidMode)
	}
	if r.budget != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "budget", r.budget)
	}
	if r.budgetMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "budget_mode", r.budgetMode)
	}
	if r.businessIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "business_ids", r.businessIds)
	}
	if r.campaignId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_id", r.campaignId)
	}
	if r.career != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "career", r.career)
	}
	if r.carrier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "carrier", r.carrier)
	}
	if r.city != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "city", r.city)
	}
	if r.convertId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "convert_id", r.convertId)
	}
	if r.convertedTimeDuration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "converted_time_duration", r.convertedTimeDuration)
	}
	if r.deviceBrand != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "device_brand", r.deviceBrand)
	}
	if r.deviceType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "device_type", r.deviceType)
	}
	if r.district != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "district", r.district)
	}
	if r.dpaLocalAudience != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dpa_local_audience", r.dpaLocalAudience)
	}
	if r.dpaRtaRecommendType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dpa_rta_recommend_type", r.dpaRtaRecommendType)
	}
	if r.dpaRtaSwitch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dpa_rta_switch", r.dpaRtaSwitch)
	}
	if r.excludeCustomActions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_custom_actions", r.excludeCustomActions)
	}
	if r.excludeFlowPackage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_flow_package", r.excludeFlowPackage)
	}
	if r.filterAwemeAbnormalActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_aweme_abnormal_active", r.filterAwemeAbnormalActive)
	}
	if r.filterAwemeFansCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_aweme_fans_count", r.filterAwemeFansCount)
	}
	if r.filterOwnAwemeFans != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_own_aweme_fans", r.filterOwnAwemeFans)
	}
	if r.flowControlMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flow_control_mode", r.flowControlMode)
	}
	if r.flowPackage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flow_package", r.flowPackage)
	}
	if r.gender != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gender", r.gender)
	}
	if r.geolocation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "geolocation", r.geolocation)
	}
	if r.hideIfConverted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hide_if_converted", r.hideIfConverted)
	}
	if r.hideIfExists != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hide_if_exists", r.hideIfExists)
	}
	if r.includeCustomActions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_custom_actions", r.includeCustomActions)
	}
	if r.interestActionMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interest_action_mode", r.interestActionMode)
	}
	if r.interestCategories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interest_categories", r.interestCategories)
	}
	if r.interestTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interest_tags", r.interestTags)
	}
	if r.interestWords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interest_words", r.interestWords)
	}
	if r.iosOsv != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ios_osv", r.iosOsv)
	}
	if r.launchPrice != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "launch_price", r.launchPrice)
	}
	if r.locationType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "location_type", r.locationType)
	}
	if r.platform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform", r.platform)
	}
	if r.pricing != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pricing", r.pricing)
	}
	if r.regionVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region_version", r.regionVersion)
	}
	if r.retargetingTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retargeting_tags", r.retargetingTags)
	}
	if r.retargetingTagsExclude != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retargeting_tags_exclude", r.retargetingTagsExclude)
	}
	if r.retargetingTagsInclude != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retargeting_tags_include", r.retargetingTagsInclude)
	}
	if r.retargetingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retargeting_type", r.retargetingType)
	}
	if r.scheduleType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "schedule_type", r.scheduleType)
	}
	if r.superiorPopularityType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "superior_popularity_type", r.superiorPopularityType)
	}
	if r.userType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_type", r.userType)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
