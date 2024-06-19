/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
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

type ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequestExample struct {
	LandingType        ToolsEventConvertOptimizedGoalGetV30LandingType      `json:"landing_type"`
	MarketingPurpose   ToolsEventConvertOptimizedGoalGetV30MarketingPurpose `json:"marketing_purpose"`
	AssetType          ToolsEventConvertOptimizedGoalGetV30AssetType        `json:"asset_type"`
	AdvertiserId       int64                                                `json:"advertiser_id,omitempty"`
	SiteId             int64                                                `json:"site_id,omitempty"`
	MiniProgramId      string                                               `json:"mini_program_id,omitempty"`
	AssetId            int64                                                `json:"asset_id,omitempty"`
	CampaignType       ToolsEventConvertOptimizedGoalGetV30CampaignType     `json:"campaign_type,omitempty"`
	MicroAppInstanceId int64                                                `json:"micro_app_instance_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/tools/event_convert/optimized_goal/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ToolsEventConvertOptimizedGoalGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsEventConvertOptimizedGoalGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		LandingType(request.LandingType).MarketingPurpose(request.MarketingPurpose).AssetType(request.AssetType).AdvertiserId(request.AdvertiserId).SiteId(request.SiteId).MiniProgramId(request.MiniProgramId).AssetId(request.AssetId).CampaignType(request.CampaignType).MicroAppInstanceId(request.MicroAppInstanceId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
