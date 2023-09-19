/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
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

type ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequestExample struct {
	QianchuanBatchCampaignStatusUpdateV10Request QianchuanBatchCampaignStatusUpdateV10Request `json:"QianchuanBatchCampaignStatusUpdateV10Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/batch_campaign_status/update/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanBatchCampaignStatusUpdateV10Api().
		Post(ctx).
		AccessToken(accessToken).
		QianchuanBatchCampaignStatusUpdateV10Request(request.QianchuanBatchCampaignStatusUpdateV10Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
