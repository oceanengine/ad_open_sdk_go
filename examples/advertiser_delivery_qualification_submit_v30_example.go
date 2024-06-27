/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

type ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequestExample struct {
	AdvertiserDeliveryQualificationSubmitV30Request AdvertiserDeliveryQualificationSubmitV30Request `json:"AdvertiserDeliveryQualificationSubmitV30Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/advertiser/delivery_qualification/submit/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.AdvertiserDeliveryQualificationSubmitV30Api().
		Post(ctx).
		AccessToken(accessToken).
		AdvertiserDeliveryQualificationSubmitV30Request(request.AdvertiserDeliveryQualificationSubmitV30Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
