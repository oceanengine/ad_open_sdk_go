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

type ApiOpenApi2FileRebateMaterialDownloadCreateTaskPostRequestExample struct {
	FileRebateMaterialDownloadCreateTaskV2Request FileRebateMaterialDownloadCreateTaskV2Request `json:"FileRebateMaterialDownloadCreateTaskV2Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/file/rebate/material_download/create_task/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2FileRebateMaterialDownloadCreateTaskPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.FileRebateMaterialDownloadCreateTaskV2Api().
		Post(ctx).
		AccessToken(accessToken).
		FileRebateMaterialDownloadCreateTaskV2Request(request.FileRebateMaterialDownloadCreateTaskV2Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
