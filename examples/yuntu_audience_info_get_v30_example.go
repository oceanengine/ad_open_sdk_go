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

type ApiOpenApiV30YuntuAudienceInfoGetGetRequestExample struct {
	YuntuBrandId      int64 `json:"yuntu_brand_id"`
	ServiceProviderId int64 `json:"service_provider_id"`
	AdvertiserId      int64 `json:"advertiser_id"`
	CustomAudienceId  int64 `json:"custom_audience_id"`
}

// url: https://api.oceanengine.com/open_api/v3.0/yuntu/audience_info/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30YuntuAudienceInfoGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.YuntuAudienceInfoGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		YuntuBrandId(request.YuntuBrandId).ServiceProviderId(request.ServiceProviderId).AdvertiserId(request.AdvertiserId).CustomAudienceId(request.CustomAudienceId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
