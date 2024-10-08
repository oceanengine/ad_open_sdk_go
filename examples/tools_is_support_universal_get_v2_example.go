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

type ApiOpenApi2ToolsIsSupportUniversalGetGetRequestExample struct {
	AdvertiserId       int64                                          `json:"advertiser_id"`
	LandingType        ToolsIsSupportUniversalGetV2LandingType        `json:"landing_type"`
	ExternalAction     ToolsIsSupportUniversalGetV2ExternalAction     `json:"external_action"`
	DeepExternalAction ToolsIsSupportUniversalGetV2DeepExternalAction `json:"deep_external_action,omitempty"`
	DeepBidType        ToolsIsSupportUniversalGetV2DeepBidType        `json:"deep_bid_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/is_support_universal/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsIsSupportUniversalGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsIsSupportUniversalGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).LandingType(request.LandingType).ExternalAction(request.ExternalAction).DeepExternalAction(request.DeepExternalAction).DeepBidType(request.DeepBidType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
