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

type ApiOpenApi2ToolsAdRaiseVersionGetGetRequestExample struct {
	AdId         int64  `json:"ad_id"`
	AdvertiserId int64  `json:"advertiser_id"`
	Version      string `json:"version"`
	PageSize     int64  `json:"page_size,omitempty"`
	Page         int64  `json:"page,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/ad_raise_version/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAdRaiseVersionGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAdRaiseVersionGetV2Api().
		Get(ctx, version).
		AccessToken(accessToken).
		AdId(request.AdId).AdvertiserId(request.AdvertiserId).Version(request.Version).PageSize(request.PageSize).Page(request.Page).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
