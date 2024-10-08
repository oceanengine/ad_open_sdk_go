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

type ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequestExample struct {
	AdvertiserId    int64                                                `json:"advertiser_id"`
	AdId            int64                                                `json:"ad_id"`
	AdRaiseVersion  int64                                                `json:"ad_raise_version"`
	StartTime       string                                               `json:"start_time"`
	EndTime         string                                               `json:"end_time"`
	Page            int32                                                `json:"page"`
	PageSize        int32                                                `json:"page_size"`
	TimeGranularity string                                               `json:"time_granularity,omitempty"`
	Filed           []string                                             `json:"filed,omitempty"`
	OrderField      string                                               `json:"order_field,omitempty"`
	OrderType       QianchuanToolsSmartBoostAdBoostReportGetV10OrderType `json:"order_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/report/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanToolsSmartBoostAdBoostReportGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AdId(request.AdId).AdRaiseVersion(request.AdRaiseVersion).StartTime(request.StartTime).EndTime(request.EndTime).Page(request.Page).PageSize(request.PageSize).TimeGranularity(request.TimeGranularity).Filed(request.Filed).OrderField(request.OrderField).OrderType(request.OrderType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
