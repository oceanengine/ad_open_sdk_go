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

type ApiOpenApiV30LocalProductGetGetRequestExample struct {
	LocalAccountId     int64                                `json:"local_account_id"`
	LocalDeliveryScene LocalProductGetV30LocalDeliveryScene `json:"local_delivery_scene"`
	Filtering          LocalProductGetV30Filtering          `json:"filtering,omitempty"`
	Page               int64                                `json:"page,omitempty"`
	PageSize           int64                                `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/local/product/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30LocalProductGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.LocalProductGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		LocalAccountId(request.LocalAccountId).LocalDeliveryScene(request.LocalDeliveryScene).Filtering(request.Filtering).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
