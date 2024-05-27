/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

type ApiOpenApi2ToolsAdConvertQueryGetRequestExample struct {
	AdvancedCreativeType ToolsAdConvertQueryV2AdvancedCreativeType `json:"advanced_creative_type,omitempty"`
	AdvertiserId         int64                                     `json:"advertiser_id,omitempty"`
	AppSchema            string                                    `json:"app_schema,omitempty"`
	AppType              ToolsAdConvertQueryV2AppType              `json:"app_type,omitempty"`
	DedicateType         ToolsAdConvertQueryV2DedicateType         `json:"dedicate_type,omitempty"`
	DeliveryRange        ToolsAdConvertQueryV2DeliveryRange        `json:"delivery_range,omitempty"`
	ExternalUrl          string                                    `json:"external_url,omitempty"`
	ItunesUrl            string                                    `json:"itunes_url,omitempty"`
	LandingType          ToolsAdConvertQueryV2LandingType          `json:"landing_type,omitempty"`
	MarketingScene       ToolsAdConvertQueryV2MarketingScene       `json:"marketing_scene,omitempty"`
	PackageName          string                                    `json:"package_name,omitempty"`
	PromotionContent     ToolsAdConvertQueryV2PromotionContent     `json:"promotion_content,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/ad_convert/query/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAdConvertQueryGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAdConvertQueryV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvancedCreativeType(request.AdvancedCreativeType).AdvertiserId(request.AdvertiserId).AppSchema(request.AppSchema).AppType(request.AppType).DedicateType(request.DedicateType).DeliveryRange(request.DeliveryRange).ExternalUrl(request.ExternalUrl).ItunesUrl(request.ItunesUrl).LandingType(request.LandingType).MarketingScene(request.MarketingScene).PackageName(request.PackageName).PromotionContent(request.PromotionContent).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
