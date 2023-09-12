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

type ApiOpenApi2ReportAdvertiserGetGetRequestExample struct {
	AdvertiserId    int64                                `json:"advertiser_id,omitempty"`
	EndDate         *string                              `json:"end_date,omitempty"`
	Fields          []string                             `json:"fields,omitempty"`
	Filtering       ReportAdvertiserGetV2Filtering       `json:"filtering,omitempty"`
	GroupBy         []*ReportAdvertiserGetV2GroupBy      `json:"group_by,omitempty"`
	OrderField      ReportAdvertiserGetV2OrderField      `json:"order_field,omitempty"`
	OrderType       ReportAdvertiserGetV2OrderType       `json:"order_type,omitempty"`
	Page            int64                                `json:"page,omitempty"`
	PageSize        int64                                `json:"page_size,omitempty"`
	StartDate       *string                              `json:"start_date,omitempty"`
	TimeGranularity ReportAdvertiserGetV2TimeGranularity `json:"time_granularity,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/report/advertiser/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ReportAdvertiserGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportAdvertiserGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).EndDate(request.EndDate).Fields(request.Fields).Filtering(request.Filtering).GroupBy(request.GroupBy).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).StartDate(request.StartDate).TimeGranularity(request.TimeGranularity).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
