/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
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

type ApiOpenApi2ToolsPromotionCardRecommendGetGetRequestExample struct {
	AdId                 int64                                             `json:"ad_id,omitempty"`
	AdvancedCreativeType string                                            `json:"advanced_creative_type,omitempty"`
	AdvertiserId         int64                                             `json:"advertiser_id,omitempty"`
	DownloadType         string                                            `json:"download_type,omitempty"`
	RecommendType        string                                            `json:"recommend_type,omitempty"`
	TitleList            []*ToolsPromotionCardRecommendGetV2TitleListInner `json:"title_list,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/promotion_card/recommend/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsPromotionCardRecommendGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsPromotionCardRecommendGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdId(request.AdId).AdvancedCreativeType(request.AdvancedCreativeType).AdvertiserId(request.AdvertiserId).DownloadType(request.DownloadType).RecommendType(request.RecommendType).TitleList(request.TitleList).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}