/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
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

type ApiOpenApi2StarReportOrderUserDistributionGetGetRequestExample struct {
	FanType         StarReportOrderUserDistributionGetV2FanType         `json:"fan_type,omitempty"`
	InteractiveType StarReportOrderUserDistributionGetV2InteractiveType `json:"interactive_type,omitempty"`
	OrderId         int64                                               `json:"order_id,omitempty"`
	StarId          int64                                               `json:"star_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/report/order_user_distribution/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarReportOrderUserDistributionGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarReportOrderUserDistributionGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		FanType(request.FanType).InteractiveType(request.InteractiveType).OrderId(request.OrderId).StarId(request.StarId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
