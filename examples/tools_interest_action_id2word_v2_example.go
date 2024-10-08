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

type ApiOpenApi2ToolsInterestActionId2wordGetRequestExample struct {
	ActionDays    ToolsInterestActionId2wordV2ActionDays    `json:"action_days,omitempty"`
	AdvertiserId  int64                                     `json:"advertiser_id,omitempty"`
	Ids           []int64                                   `json:"ids,omitempty"`
	TagType       ToolsInterestActionId2wordV2TagType       `json:"tag_type,omitempty"`
	TargetingType ToolsInterestActionId2wordV2TargetingType `json:"targeting_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/interest_action/id2word/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsInterestActionId2wordGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsInterestActionId2wordV2Api().
		Get(ctx).
		AccessToken(accessToken).
		ActionDays(request.ActionDays).AdvertiserId(request.AdvertiserId).Ids(request.Ids).TagType(request.TagType).TargetingType(request.TargetingType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
