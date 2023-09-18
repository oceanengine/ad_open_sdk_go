/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
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

type ApiOpenApi2StarClueGetGetRequestExample struct {
	DemandId int64 `json:"demand_id,omitempty"`
	OrderId  int64 `json:"order_id,omitempty"`
	Page     int64 `json:"page,omitempty"`
	PageSize int64 `json:"page_size,omitempty"`
	StarId   int64 `json:"star_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/clue/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarClueGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarClueGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		DemandId(request.DemandId).OrderId(request.OrderId).Page(request.Page).PageSize(request.PageSize).StarId(request.StarId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}