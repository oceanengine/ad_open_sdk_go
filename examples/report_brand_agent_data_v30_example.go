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

type ApiOpenApiV30ReportBrandAgentDataGetRequestExample struct {
	AdvertiserId   int64                                 `json:"advertiser_id"`
	DataReportType ReportBrandAgentDataV30DataReportType `json:"data_report_type"`
	StartTime      string                                `json:"start_time"`
	EndTime        string                                `json:"end_time"`
	TimeDimension  ReportBrandAgentDataV30TimeDimension  `json:"time_dimension,omitempty"`
	PageType       ReportBrandAgentDataV30PageType       `json:"page_type,omitempty"`
	Filter         ReportBrandAgentDataV30Filter         `json:"filter,omitempty"`
	Metrics        []string                              `json:"metrics,omitempty"`
	Page           ReportBrandAgentDataV30Page           `json:"page,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/report/brand/agent/data/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ReportBrandAgentDataGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportBrandAgentDataV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).DataReportType(request.DataReportType).StartTime(request.StartTime).EndTime(request.EndTime).TimeDimension(request.TimeDimension).PageType(request.PageType).Filter(request.Filter).Metrics(request.Metrics).Page(request.Page).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}