/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

type ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequestExample struct {
	AdvertiserId int64                                  `json:"advertiser_id"`
	AwemeId      int64                                  `json:"aweme_id"`
	Filtering    QianchuanFileVideoAwemeGetV10Filtering `json:"filtering,omitempty"`
	Cursor       int64                                  `json:"cursor,omitempty"`
	Count        int64                                  `json:"count,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/file/video/aweme/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanFileVideoAwemeGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanFileVideoAwemeGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AwemeId(request.AwemeId).Filtering(request.Filtering).Cursor(request.Cursor).Count(request.Count).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
