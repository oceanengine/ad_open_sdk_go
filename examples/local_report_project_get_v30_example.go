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

type ApiOpenApiV30LocalReportProjectGetGetRequestExample struct {
	LocalAccountId  int64                                   `json:"local_account_id"`
	StartDate       string                                  `json:"start_date"`
	EndDate         string                                  `json:"end_date"`
	Metrics         []string                                `json:"metrics"`
	TimeGranularity LocalReportProjectGetV30TimeGranularity `json:"time_granularity,omitempty"`
	OrderType       LocalReportProjectGetV30OrderType       `json:"order_type,omitempty"`
	OrderField      string                                  `json:"order_field,omitempty"`
	Filtering       LocalReportProjectGetV30Filtering       `json:"filtering,omitempty"`
	Page            int32                                   `json:"page,omitempty"`
	PageSize        int32                                   `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/local/report/project/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30LocalReportProjectGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.LocalReportProjectGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		LocalAccountId(request.LocalAccountId).StartDate(request.StartDate).EndDate(request.EndDate).Metrics(request.Metrics).TimeGranularity(request.TimeGranularity).OrderType(request.OrderType).OrderField(request.OrderField).Filtering(request.Filtering).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
