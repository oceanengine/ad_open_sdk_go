/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

type ApiOpenApi2DmpCustomAudiencePushV2PostRequestExample struct {
	DmpCustomAudiencePushV2V2Request DmpCustomAudiencePushV2V2Request `json:"DmpCustomAudiencePushV2V2Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/dmp/custom_audience/push_v2/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2DmpCustomAudiencePushV2PostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.DmpCustomAudiencePushV2V2Api().
		Post(ctx).
		AccessToken(accessToken).
		DmpCustomAudiencePushV2V2Request(request.DmpCustomAudiencePushV2V2Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
