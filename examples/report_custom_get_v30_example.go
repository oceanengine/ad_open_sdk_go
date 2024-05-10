/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

type ApiOpenApiV30ReportCustomGetGetRequestExample struct {
	Dimensions   []string                          `json:"dimensions"`
	AdvertiserId int64                             `json:"advertiser_id"`
	Metrics      []string                          `json:"metrics"`
	Filters      []*ReportCustomGetV30FiltersInner `json:"filters"`
	StartTime    string                            `json:"start_time"`
	EndTime      string                            `json:"end_time"`
	OrderBy      []*ReportCustomGetV30OrderByInner `json:"order_by"`
	Page         int32                             `json:"page,omitempty"`
	PageSize     int32                             `json:"page_size,omitempty"`
	DataTopic    ReportCustomGetV30DataTopic       `json:"data_topic,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/report/custom/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ReportCustomGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportCustomGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		Dimensions(request.Dimensions).AdvertiserId(request.AdvertiserId).Metrics(request.Metrics).Filters(request.Filters).StartTime(request.StartTime).EndTime(request.EndTime).OrderBy(request.OrderBy).Page(request.Page).PageSize(request.PageSize).DataTopic(request.DataTopic).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
