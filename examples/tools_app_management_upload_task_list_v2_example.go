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

type ApiOpenApi2ToolsAppManagementUploadTaskListGetRequestExample struct {
	AccountId   int64                                         `json:"account_id"`
	AccountType ToolsAppManagementUploadTaskListV2AccountType `json:"account_type"`
	Filtering   ToolsAppManagementUploadTaskListV2Filtering   `json:"filtering"`
	Version     string                                        `json:"version"`
}

// url: https://api.oceanengine.com/open_api/2/tools/app_management/upload_task/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAppManagementUploadTaskListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAppManagementUploadTaskListV2Api().
		Get(ctx, version).
		AccessToken(accessToken).
		AccountId(request.AccountId).AccountType(request.AccountType).Filtering(request.Filtering).Version(request.Version).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
