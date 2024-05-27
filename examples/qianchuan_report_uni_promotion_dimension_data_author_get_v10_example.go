/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

type ApiOpenApiV10QianchuanReportUniPromotionDimensionDataAuthorGetGetRequestExample struct {
	AdvertiserId int64                                                         `json:"advertiser_id"`
	AwemeId      int64                                                         `json:"aweme_id"`
	Metrics      []string                                                      `json:"metrics,omitempty"`
	StartTime    string                                                        `json:"start_time,omitempty"`
	EndTime      string                                                        `json:"end_time,omitempty"`
	Dimension    QianchuanReportUniPromotionDimensionDataAuthorGetV10Dimension `json:"dimension,omitempty"`
	OrderType    QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType `json:"order_type,omitempty"`
	OrderField   string                                                        `json:"order_field,omitempty"`
	Page         int32                                                         `json:"page,omitempty"`
	PageSize     int32                                                         `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/report/uni_promotion/dimension_data/author/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanReportUniPromotionDimensionDataAuthorGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanReportUniPromotionDimensionDataAuthorGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AwemeId(request.AwemeId).Metrics(request.Metrics).StartTime(request.StartTime).EndTime(request.EndTime).Dimension(request.Dimension).OrderType(request.OrderType).OrderField(request.OrderField).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
