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

type ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequestExample struct {
	AdvertiserId int64                                                       `json:"advertiser_id"`
	RoomId       int64                                                       `json:"room_id"`
	StartTime    string                                                      `json:"start_time,omitempty"`
	EndTime      string                                                      `json:"end_time,omitempty"`
	Dimension    QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension `json:"dimension,omitempty"`
	Metrics      []string                                                    `json:"metrics,omitempty"`
	OrderField   string                                                      `json:"order_field,omitempty"`
	OrderType    QianchuanReportUniPromotionDimensionDataRoomGetV10OrderType `json:"order_type,omitempty"`
	Page         int32                                                       `json:"page,omitempty"`
	PageSize     int32                                                       `json:"page_size,omitempty"`
	Filtering    QianchuanReportUniPromotionDimensionDataRoomGetV10Filtering `json:"filtering,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/report/uni_promotion/dimension_data/room/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanReportUniPromotionDimensionDataRoomGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanReportUniPromotionDimensionDataRoomGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).RoomId(request.RoomId).StartTime(request.StartTime).EndTime(request.EndTime).Dimension(request.Dimension).Metrics(request.Metrics).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).Filtering(request.Filtering).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
