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

type ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequestExample struct {
	AdvertiserId int64                                       `json:"advertiser_id"`
	AwemeId      string                                      `json:"aweme_id"`
	Behaviors    []*ToolsAwemeSimilarAuthorSearchV2Behaviors `json:"behaviors,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/aweme_similar_author_search/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAwemeSimilarAuthorSearchV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AwemeId(request.AwemeId).Behaviors(request.Behaviors).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
