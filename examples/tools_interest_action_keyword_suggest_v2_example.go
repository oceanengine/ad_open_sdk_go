/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


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

type ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequestExample struct {
	AdvertiserId  int64                                            `json:"advertiser_id"`
	Id            int64                                            `json:"id"`
	TagType       ToolsInterestActionKeywordSuggestV2TagType       `json:"tag_type"`
	TargetingType ToolsInterestActionKeywordSuggestV2TargetingType `json:"targeting_type"`
	ActionDays    ToolsInterestActionKeywordSuggestV2ActionDays    `json:"action_days,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/interest_action/keyword/suggest/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsInterestActionKeywordSuggestV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Id(request.Id).TagType(request.TagType).TargetingType(request.TargetingType).ActionDays(request.ActionDays).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
