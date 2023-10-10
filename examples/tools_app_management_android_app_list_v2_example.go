/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
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

type ApiOpenApi2ToolsAppManagementAndroidAppListGetRequestExample struct {
	AccountId   int64                                         `json:"account_id"`
	AccountType ToolsAppManagementAndroidAppListV2AccountType `json:"account_type"`
	Page        int32                                         `json:"page,omitempty"`
	PageSize    int32                                         `json:"page_size,omitempty"`
	Filtering   ToolsAppManagementAndroidAppListV2Filtering   `json:"filtering,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/app_management/android_app/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAppManagementAndroidAppListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAppManagementAndroidAppListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AccountId(request.AccountId).AccountType(request.AccountType).Page(request.Page).PageSize(request.PageSize).Filtering(request.Filtering).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
