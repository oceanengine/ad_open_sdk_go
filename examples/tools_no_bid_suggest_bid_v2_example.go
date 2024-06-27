/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

type ApiOpenApi2ToolsNoBidSuggestBidGetRequestExample struct {
	AdvertiserId   int64                                `json:"advertiser_id,omitempty"`
	Budget         int64                                `json:"budget,omitempty"`
	BudgetMode     ToolsNoBidSuggestBidV2BudgetMode     `json:"budget_mode,omitempty"`
	ExternalAction ToolsNoBidSuggestBidV2ExternalAction `json:"external_action,omitempty"`
	Filtering      ToolsNoBidSuggestBidV2Filtering      `json:"filtering,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/no_bid/suggest_bid/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsNoBidSuggestBidGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsNoBidSuggestBidV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Budget(request.Budget).BudgetMode(request.BudgetMode).ExternalAction(request.ExternalAction).Filtering(request.Filtering).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
