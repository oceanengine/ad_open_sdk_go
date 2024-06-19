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

type ApiOpenApi2StarReportCustomDataTopicReportGetRequestExample struct {
	StarId   int64                                      `json:"star_id"`
	WorkId   int64                                      `json:"work_id"`
	DemandId int64                                      `json:"demand_id"`
	Topics   []*StarReportCustomDataTopicReportV2Topics `json:"topics"`
}

// url: https://api.oceanengine.com/open_api/2/star/report/custom_data_topic_report/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarReportCustomDataTopicReportGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarReportCustomDataTopicReportV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).WorkId(request.WorkId).DemandId(request.DemandId).Topics(request.Topics).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
