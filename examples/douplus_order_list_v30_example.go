/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
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

type ApiOpenApiV30DouplusOrderListGetRequestExample struct {
	AwemeSecUid string                    `json:"aweme_sec_uid,omitempty"`
	Filter      DouplusOrderListV30Filter `json:"filter,omitempty"`
	PageSize    int64                     `json:"page_size,omitempty"`
	Page        int64                     `json:"page,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/douplus/order/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30DouplusOrderListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.DouplusOrderListV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AwemeSecUid(request.AwemeSecUid).Filter(request.Filter).PageSize(request.PageSize).Page(request.Page).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
