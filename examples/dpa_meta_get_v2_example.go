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

type ApiOpenApi2DpaMetaGetGetRequestExample struct {
	AdvertiserId int64 `json:"advertiser_id"`
	PlatformId   int64 `json:"platform_id"`
	Indexable    int64 `json:"indexable,omitempty"`
	Extractable  int64 `json:"extractable,omitempty"`
	Industry     int64 `json:"industry,omitempty"`
	Status       int64 `json:"status,omitempty"`
	MediaType    int64 `json:"mediaType,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/dpa/meta/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2DpaMetaGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.DpaMetaGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).PlatformId(request.PlatformId).Indexable(request.Indexable).Extractable(request.Extractable).Industry(request.Industry).Status(request.Status).MediaType(request.MediaType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
