/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
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

type ApiOpenApiV10QianchuanSuggestBidGetRequestExample struct {
	AdvertiserId   int64                                `json:"advertiser_id"`
	MarketingGoal  QianchuanSuggestBidV10MarketingGoal  `json:"marketing_goal"`
	ExternalAction QianchuanSuggestBidV10ExternalAction `json:"external_action"`
	AwemeId        int64                                `json:"aweme_id,omitempty"`
	ProductId      int64                                `json:"product_id,omitempty"`
	CampaignScene  QianchuanSuggestBidV10CampaignScene  `json:"campaign_scene,omitempty"`
	EcomGuestType  QianchuanSuggestBidV10EcomGuestType  `json:"ecom_guest_type,omitempty"`
	ShopId         int64                                `json:"shop_id,omitempty"`
	BrandId        int64                                `json:"brand_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/suggest_bid/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanSuggestBidGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanSuggestBidV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).MarketingGoal(request.MarketingGoal).ExternalAction(request.ExternalAction).AwemeId(request.AwemeId).ProductId(request.ProductId).CampaignScene(request.CampaignScene).EcomGuestType(request.EcomGuestType).ShopId(request.ShopId).BrandId(request.BrandId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
