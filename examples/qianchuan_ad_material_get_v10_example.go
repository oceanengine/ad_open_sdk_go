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

type ApiOpenApiV10QianchuanAdMaterialGetGetRequestExample struct {
	AdvertiserId int64                              `json:"advertiser_id"`
	AdId         int64                              `json:"ad_id"`
	Filtering    QianchuanAdMaterialGetV10Filtering `json:"filtering"`
	Page         int32                              `json:"page,omitempty"`
	PageSize     QianchuanAdMaterialGetV10PageSize  `json:"page_size,omitempty"`
	OrderType    QianchuanAdMaterialGetV10OrderType `json:"order_type,omitempty"`
	OrderField   string                             `json:"order_field,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/ad/material/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanAdMaterialGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanAdMaterialGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AdId(request.AdId).Filtering(request.Filtering).Page(request.Page).PageSize(request.PageSize).OrderType(request.OrderType).OrderField(request.OrderField).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
