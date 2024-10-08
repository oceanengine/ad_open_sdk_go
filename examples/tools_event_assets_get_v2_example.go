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

type ApiOpenApi2ToolsEventAssetsGetGetRequestExample struct {
	AdvertiserId int64                          `json:"advertiser_id"`
	AssetType    ToolsEventAssetsGetV2AssetType `json:"asset_type"`
	Filtering    ToolsEventAssetsGetV2Filtering `json:"filtering,omitempty"`
	SortType     ToolsEventAssetsGetV2SortType  `json:"sort_type,omitempty"`
	Page         int64                          `json:"page,omitempty"`
	PageSize     int64                          `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/event/assets/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsEventAssetsGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsEventAssetsGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AssetType(request.AssetType).Filtering(request.Filtering).SortType(request.SortType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
