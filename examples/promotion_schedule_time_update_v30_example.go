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

type ApiOpenApiV30PromotionScheduleTimeUpdatePostRequestExample struct {
	PromotionScheduleTimeUpdateV30Request PromotionScheduleTimeUpdateV30Request `json:"PromotionScheduleTimeUpdateV30Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/promotion/schedule_time/update/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30PromotionScheduleTimeUpdatePostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.PromotionScheduleTimeUpdateV30Api().
		Post(ctx).
		AccessToken(accessToken).
		PromotionScheduleTimeUpdateV30Request(request.PromotionScheduleTimeUpdateV30Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
