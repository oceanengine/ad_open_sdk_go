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

type ApiOpenApi2FileRebateCommonDownloadDownloadFileGetRequestExample struct {
	AgentId int64  `json:"agent_id"`
	TaskId  string `json:"task_id"`
}

// url: https://api.oceanengine.com/open_api/2/file/rebate/common_download/download_file/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2FileRebateCommonDownloadDownloadFileGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.FileRebateCommonDownloadDownloadFileV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AgentId(request.AgentId).TaskId(request.TaskId).
		Execute()
	fmt.Println(string(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
