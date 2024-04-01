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

type ApiOpenApi2StarOrderListByCampaignGetRequestExample struct {
	StarId      int64   `json:"star_id"`
	CampaignIds []int64 `json:"campaign_ids"`
	Page        int32   `json:"page,omitempty"`
	Limit       int32   `json:"limit,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/order/list_by_campaign/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarOrderListByCampaignGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarOrderListByCampaignV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).CampaignIds(request.CampaignIds).Page(request.Page).Limit(request.Limit).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
