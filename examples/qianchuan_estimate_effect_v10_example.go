/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
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

type ApiOpenApiV10QianchuanEstimateEffectGetRequestExample struct {
	AdvertiserId       int64                                        `json:"advertiser_id"`
	AwemeId            int64                                        `json:"aweme_id"`
	ExternalAction     QianchuanEstimateEffectV10ExternalAction     `json:"external_action"`
	BudgetMode         QianchuanEstimateEffectV10BudgetMode         `json:"budget_mode"`
	Budget             float64                                      `json:"budget"`
	LiveScheduleType   QianchuanEstimateEffectV10LiveScheduleType   `json:"live_schedule_type"`
	DeepExternalAction QianchuanEstimateEffectV10DeepExternalAction `json:"deep_external_action,omitempty"`
	DeepBidType        QianchuanEstimateEffectV10DeepBidType        `json:"deep_bid_type,omitempty"`
	StartTime          string                                       `json:"start_time,omitempty"`
	EndTime            string                                       `json:"end_time,omitempty"`
	ScheduleTime       string                                       `json:"schedule_time,omitempty"`
	ScheduleFixedRange int64                                        `json:"schedule_fixed_range,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/estimate/effect/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanEstimateEffectGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanEstimateEffectV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AwemeId(request.AwemeId).ExternalAction(request.ExternalAction).BudgetMode(request.BudgetMode).Budget(request.Budget).LiveScheduleType(request.LiveScheduleType).DeepExternalAction(request.DeepExternalAction).DeepBidType(request.DeepBidType).StartTime(request.StartTime).EndTime(request.EndTime).ScheduleTime(request.ScheduleTime).ScheduleFixedRange(request.ScheduleFixedRange).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
