/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

type ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequestExample struct {
	AdvertiserId int64                                            `json:"advertiser_id"`
	Filter       ToolsUnionFlowPackagePromotionReportV30Filter    `json:"filter"`
	OrderField   string                                           `json:"order_field,omitempty"`
	OrderType    ToolsUnionFlowPackagePromotionReportV30OrderType `json:"order_type,omitempty"`
	Page         int64                                            `json:"page,omitempty"`
	PageSize     int64                                            `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/tools/union/flow_package/promotion/report/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsUnionFlowPackagePromotionReportV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Filter(request.Filter).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
