/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

type ApiOpenApiV10QianchuanReportSearchWordGetGetRequestExample struct {
	AdvertiserId    int64                                          `json:"advertiser_id"`
	StartDate       string                                         `json:"start_date"`
	EndDate         string                                         `json:"end_date"`
	Fields          []string                                       `json:"fields"`
	WordType        QianchuanReportSearchWordGetV10WordType        `json:"word_type"`
	Filtering       QianchuanReportSearchWordGetV10Filtering       `json:"filtering"`
	TimeGranularity QianchuanReportSearchWordGetV10TimeGranularity `json:"time_granularity,omitempty"`
	OrderField      string                                         `json:"order_field,omitempty"`
	OrderType       QianchuanReportSearchWordGetV10OrderType       `json:"order_type,omitempty"`
	Page            int32                                          `json:"page,omitempty"`
	PageSize        int32                                          `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/report/search_word/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanReportSearchWordGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanReportSearchWordGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).StartDate(request.StartDate).EndDate(request.EndDate).Fields(request.Fields).WordType(request.WordType).Filtering(request.Filtering).TimeGranularity(request.TimeGranularity).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
