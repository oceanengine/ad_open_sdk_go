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

type ApiOpenApiV30ReportBrandAdGetGetRequestExample struct {
	AdvertiserId int64                          `json:"advertiser_id"`
	Page         int64                          `json:"page"`
	Size         int64                          `json:"size"`
	LandingType  ReportBrandAdGetV30LandingType `json:"landing_type,omitempty"`
	PricingType  ReportBrandAdGetV30PricingType `json:"pricing_type,omitempty"`
	AdIds        []string                       `json:"ad_ids,omitempty"`
	StartTime    string                         `json:"start_time,omitempty"`
	EndTime      string                         `json:"end_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/report/brand/ad/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ReportBrandAdGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportBrandAdGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Page(request.Page).Size(request.Size).LandingType(request.LandingType).PricingType(request.PricingType).AdIds(request.AdIds).StartTime(request.StartTime).EndTime(request.EndTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
