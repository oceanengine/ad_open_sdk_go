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

type ApiOpenApiV30YuntuAudienceLabelGetGetRequestExample struct {
	YuntuBrandId      int64 `json:"yuntu_brand_id"`
	ServiceProviderId int64 `json:"service_provider_id"`
	LabelId           int64 `json:"label_id"`
}

// url: https://api.oceanengine.com/open_api/v3.0/yuntu/audience_label/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30YuntuAudienceLabelGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.YuntuAudienceLabelGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		YuntuBrandId(request.YuntuBrandId).ServiceProviderId(request.ServiceProviderId).LabelId(request.LabelId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
