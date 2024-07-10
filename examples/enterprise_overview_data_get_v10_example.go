/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

type ApiOpenApiV10EnterpriseOverviewDataGetGetRequestExample struct {
	AdvertiserId  int64                                     `json:"advertiser_id,omitempty"`
	EndTime       *string                                   `json:"end_time,omitempty"`
	Fields        []string                                  `json:"fields,omitempty"`
	Filter        EnterpriseOverviewDataGetV10Filter        `json:"filter,omitempty"`
	LastEndTime   *string                                   `json:"last_end_time,omitempty"`
	LastStartTime *string                                   `json:"last_start_time,omitempty"`
	Limit         int64                                     `json:"limit,omitempty"`
	Offset        int64                                     `json:"offset,omitempty"`
	OrderField    string                                    `json:"order_field,omitempty"`
	OrderType     EnterpriseOverviewDataGetV10OrderType     `json:"order_type,omitempty"`
	RatioFields   []string                                  `json:"ratio_fields,omitempty"`
	StartTime     *string                                   `json:"start_time,omitempty"`
	TimeDimension EnterpriseOverviewDataGetV10TimeDimension `json:"time_dimension,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/enterprise/overview/data/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10EnterpriseOverviewDataGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.EnterpriseOverviewDataGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).EndTime(request.EndTime).Fields(request.Fields).Filter(request.Filter).LastEndTime(request.LastEndTime).LastStartTime(request.LastStartTime).Limit(request.Limit).Offset(request.Offset).OrderField(request.OrderField).OrderType(request.OrderType).RatioFields(request.RatioFields).StartTime(request.StartTime).TimeDimension(request.TimeDimension).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
