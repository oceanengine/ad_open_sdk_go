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

type ApiOpenApi2ReportRtaGetGetRequestExample struct {
	AdvertiserId    int64                         `json:"advertiser_id,omitempty"`
	EndDate         string                        `json:"end_date,omitempty"`
	OrderField      ReportRtaGetV2OrderField      `json:"order_field,omitempty"`
	OrderType       ReportRtaGetV2OrderType       `json:"order_type,omitempty"`
	Page            int64                         `json:"page,omitempty"`
	PageSize        int64                         `json:"page_size,omitempty"`
	ReportType      ReportRtaGetV2ReportType      `json:"report_type,omitempty"`
	StartDate       string                        `json:"start_date,omitempty"`
	TimeGranularity ReportRtaGetV2TimeGranularity `json:"time_granularity,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/report/rta/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ReportRtaGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportRtaGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).EndDate(request.EndDate).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).ReportType(request.ReportType).StartDate(request.StartDate).TimeGranularity(request.TimeGranularity).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
