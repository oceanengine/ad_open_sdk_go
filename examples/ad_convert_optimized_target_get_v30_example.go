/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

type ApiOpenApiV30AdConvertOptimizedTargetGetGetRequestExample struct {
	AdvertiserId       int64                                          `json:"advertiser_id"`
	LandingType        AdConvertOptimizedTargetGetV30LandingType      `json:"landing_type"`
	MarketingPurpose   AdConvertOptimizedTargetGetV30MarketingPurpose `json:"marketing_purpose"`
	AppType            AdConvertOptimizedTargetGetV30AppType          `json:"app_type,omitempty"`
	PromotionContent   AdConvertOptimizedTargetGetV30PromotionContent `json:"promotion_content,omitempty"`
	ConvertTracking    AdConvertOptimizedTargetGetV30ConvertTracking  `json:"convert_tracking,omitempty"`
	AdType             AdConvertOptimizedTargetGetV30AdType           `json:"ad_type,omitempty"`
	DeliveryRange      AdConvertOptimizedTargetGetV30DeliveryRange    `json:"delivery_range,omitempty"`
	ExternalUrl        string                                         `json:"external_url,omitempty"`
	AndroidPackageName string                                         `json:"android_package_name,omitempty"`
	ItunesUrl          string                                         `json:"itunes_url,omitempty"`
	ConvertId          int64                                          `json:"convert_id,omitempty"`
	Page               int64                                          `json:"page,omitempty"`
	PageSize           int64                                          `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/ad_convert/optimized_target/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30AdConvertOptimizedTargetGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.AdConvertOptimizedTargetGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).LandingType(request.LandingType).MarketingPurpose(request.MarketingPurpose).AppType(request.AppType).PromotionContent(request.PromotionContent).ConvertTracking(request.ConvertTracking).AdType(request.AdType).DeliveryRange(request.DeliveryRange).ExternalUrl(request.ExternalUrl).AndroidPackageName(request.AndroidPackageName).ItunesUrl(request.ItunesUrl).ConvertId(request.ConvertId).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
