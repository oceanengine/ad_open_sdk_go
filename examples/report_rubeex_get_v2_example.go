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

type ApiOpenApi2ReportRubeexGetGetRequestExample struct {
	AdvertiserId int64                      `json:"advertiser_id,omitempty"`
	Dimensions   []string                   `json:"dimensions,omitempty"`
	Filtering    ReportRubeexGetV2Filtering `json:"filtering,omitempty"`
	Metrics      []string                   `json:"metrics,omitempty"`
	Order        ReportRubeexGetV2Order     `json:"order,omitempty"`
	Page         int64                      `json:"page,omitempty"`
	PageSize     int64                      `json:"page_size,omitempty"`
	ProjectId    int64                      `json:"project_id,omitempty"`
	StatTimeDay  []string                   `json:"stat_time_day,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/report/rubeex/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ReportRubeexGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportRubeexGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Dimensions(request.Dimensions).Filtering(request.Filtering).Metrics(request.Metrics).Order(request.Order).Page(request.Page).PageSize(request.PageSize).ProjectId(request.ProjectId).StatTimeDay(request.StatTimeDay).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
