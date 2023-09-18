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

type ApiOpenApi2ClueCouponCodeGetGetRequestExample struct {
	ActivityId   int64       `json:"activity_id,omitempty"`
	AdvertiserId int64       `json:"advertiser_id,omitempty"`
	CouponId     int64       `json:"coupon_id,omitempty"`
	EndTime      *string     `json:"end_time,omitempty"`
	Page         int64       `json:"page,omitempty"`
	PageSize     int64       `json:"page_size,omitempty"`
	ResourceId   int64       `json:"resource_id,omitempty"`
	StartTime    *string     `json:"start_time,omitempty"`
	Status       interface{} `json:"status,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/clue/coupon/code/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ClueCouponCodeGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ClueCouponCodeGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		ActivityId(request.ActivityId).AdvertiserId(request.AdvertiserId).CouponId(request.CouponId).EndTime(request.EndTime).Page(request.Page).PageSize(request.PageSize).ResourceId(request.ResourceId).StartTime(request.StartTime).Status(request.Status).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}