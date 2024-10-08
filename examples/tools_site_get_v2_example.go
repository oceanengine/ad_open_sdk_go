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

type ApiOpenApi2ToolsSiteGetGetRequestExample struct {
	AdvertiserId  int64                   `json:"advertiser_id"`
	XOrangeCaller string                  `json:"X-Orange-Caller,omitempty"`
	Page          int64                   `json:"page,omitempty"`
	PageSize      int64                   `json:"page_size,omitempty"`
	Status        ToolsSiteGetV2Status    `json:"status,omitempty"`
	ShareType     ToolsSiteGetV2ShareType `json:"share_type,omitempty"`
	Filtering     ToolsSiteGetV2Filtering `json:"filtering,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/site/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsSiteGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsSiteGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).XOrangeCaller(request.XOrangeCaller).Page(request.Page).PageSize(request.PageSize).Status(request.Status).ShareType(request.ShareType).Filtering(request.Filtering).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
