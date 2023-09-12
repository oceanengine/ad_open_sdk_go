/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "code.byted.org/ad/ad_open_sdk_go"
	"code.byted.org/ad/ad_open_sdk_go/config"
	. "code.byted.org/ad/ad_open_sdk_go/models"
)

type ApiOpenApi2AgentAdvertiserSelectGetRequestExample struct {
	AdvertiserId int64   `json:"advertiser_id,omitempty"`
	CompanyIds   []int64 `json:"company_ids,omitempty"`
	Count        int64   `json:"count,omitempty"`
	Cursor       int64   `json:"cursor,omitempty"`
	Page         int64   `json:"page,omitempty"`
	PageSize     int64   `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/agent/advertiser/select/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2AgentAdvertiserSelectGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.AgentAdvertiserSelectV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).CompanyIds(request.CompanyIds).Count(request.Count).Cursor(request.Cursor).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
