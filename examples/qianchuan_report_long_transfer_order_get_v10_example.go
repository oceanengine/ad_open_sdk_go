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

type ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequestExample struct {
	AdvertiserId          int64                                                       `json:"advertiser_id"`
	StartDate             string                                                      `json:"start_date"`
	EndDate               string                                                      `json:"end_date"`
	OptimizationTimeScope QianchuanReportLongTransferOrderGetV10OptimizationTimeScope `json:"optimization_time_scope"`
	Filtering             QianchuanReportLongTransferOrderGetV10Filtering             `json:"filtering"`
	MarketingGoal         QianchuanReportLongTransferOrderGetV10MarketingGoal         `json:"marketing_goal"`
	OrderType             QianchuanReportLongTransferOrderGetV10OrderType             `json:"order_type,omitempty"`
	Page                  int32                                                       `json:"page,omitempty"`
	PageSize              int32                                                       `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/report/long_transfer/order/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanReportLongTransferOrderGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).StartDate(request.StartDate).EndDate(request.EndDate).OptimizationTimeScope(request.OptimizationTimeScope).Filtering(request.Filtering).MarketingGoal(request.MarketingGoal).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
