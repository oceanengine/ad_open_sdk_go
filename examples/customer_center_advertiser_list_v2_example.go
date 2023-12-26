/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
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

type ApiOpenApi2CustomerCenterAdvertiserListGetRequestExample struct {
	AccountSource CustomerCenterAdvertiserListV2AccountSource `json:"account_source,omitempty"`
	CcAccountId   int64                                       `json:"cc_account_id,omitempty"`
	Filtering     CustomerCenterAdvertiserListV2Filtering     `json:"filtering,omitempty"`
	Page          int64                                       `json:"page,omitempty"`
	PageSize      int64                                       `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/customer_center/advertiser/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2CustomerCenterAdvertiserListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.CustomerCenterAdvertiserListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AccountSource(request.AccountSource).CcAccountId(request.CcAccountId).Filtering(request.Filtering).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
