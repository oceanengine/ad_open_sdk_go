/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

type ApiOpenApiV30ToolsBidsSuggestGetRequestExample struct {
	AdvertiserId       int64                                 `json:"advertiser_id"`
	Pricing            ToolsBidsSuggestV30Pricing            `json:"pricing"`
	ExternalAction     ToolsBidsSuggestV30ExternalAction     `json:"external_action"`
	ProjectId          int64                                 `json:"project_id,omitempty"`
	DeepExternalAction ToolsBidsSuggestV30DeepExternalAction `json:"deep_external_action,omitempty"`
	DeepBidType        ToolsBidsSuggestV30DeepBidType        `json:"deep_bid_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/tools/bids/suggest/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ToolsBidsSuggestGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsBidsSuggestV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Pricing(request.Pricing).ExternalAction(request.ExternalAction).ProjectId(request.ProjectId).DeepExternalAction(request.DeepExternalAction).DeepBidType(request.DeepBidType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
