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

type ApiOpenApi2ReportCampaignGetGetRequestExample struct {
	AdvertiserId    int64                              `json:"advertiser_id,omitempty"`
	EndDate         *string                            `json:"end_date,omitempty"`
	Fields          []string                           `json:"fields,omitempty"`
	Filtering       ReportCampaignGetV2Filtering       `json:"filtering,omitempty"`
	GroupBy         []*ReportCampaignGetV2GroupBy      `json:"group_by,omitempty"`
	OrderField      ReportCampaignGetV2OrderField      `json:"order_field,omitempty"`
	OrderType       ReportCampaignGetV2OrderType       `json:"order_type,omitempty"`
	Page            int64                              `json:"page,omitempty"`
	PageSize        int64                              `json:"page_size,omitempty"`
	StartDate       *string                            `json:"start_date,omitempty"`
	TimeGranularity ReportCampaignGetV2TimeGranularity `json:"time_granularity,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/report/campaign/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ReportCampaignGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportCampaignGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).EndDate(request.EndDate).Fields(request.Fields).Filtering(request.Filtering).GroupBy(request.GroupBy).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).StartDate(request.StartDate).TimeGranularity(request.TimeGranularity).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
