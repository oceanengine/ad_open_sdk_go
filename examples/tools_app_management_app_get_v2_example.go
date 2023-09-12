/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "code.byted.org/ad/ad_open_sdk_go"
	"code.byted.org/ad/ad_open_sdk_go/config"
	. "code.byted.org/ad/ad_open_sdk_go/models"
)

type ApiOpenApi2ToolsAppManagementAppGetGetRequestExample struct {
	AdvertiserId int64                                 `json:"advertiser_id"`
	Page         int32                                 `json:"page,omitempty"`
	PageSize     int32                                 `json:"page_size,omitempty"`
	SearchKey    string                                `json:"search_key,omitempty"`
	SearchType   ToolsAppManagementAppGetV2SearchType  `json:"search_type,omitempty"`
	Status       ToolsAppManagementAppGetV2Status      `json:"status,omitempty"`
	PublishTime  ToolsAppManagementAppGetV2PublishTime `json:"publish_time,omitempty"`
	CreateTime   ToolsAppManagementAppGetV2CreateTime  `json:"create_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/app_management/app/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAppManagementAppGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAppManagementAppGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Page(request.Page).PageSize(request.PageSize).SearchKey(request.SearchKey).SearchType(request.SearchType).Status(request.Status).PublishTime(request.PublishTime).CreateTime(request.CreateTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
