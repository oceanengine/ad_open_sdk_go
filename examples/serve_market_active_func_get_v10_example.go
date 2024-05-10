/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

type ApiOpenApiV10ServeMarketActiveFuncGetGetRequestExample struct {
	UseUid   int64    `json:"use_uid"`
	AppId    int64    `json:"app_id"`
	Page     int64    `json:"page"`
	PageSize int64    `json:"page_size"`
	FuncKeys []string `json:"func_keys,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/serve_market/active_func/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10ServeMarketActiveFuncGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ServeMarketActiveFuncGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		UseUid(request.UseUid).AppId(request.AppId).Page(request.Page).PageSize(request.PageSize).FuncKeys(request.FuncKeys).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
