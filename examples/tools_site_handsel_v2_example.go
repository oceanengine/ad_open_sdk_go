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

type ApiOpenApi2ToolsSiteHandselPostRequestExample struct {
	XOrangeCaller             string                    `json:"X-Orange-Caller,omitempty"`
	ToolsSiteHandselV2Request ToolsSiteHandselV2Request `json:"ToolsSiteHandselV2Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/site/handsel/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsSiteHandselPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsSiteHandselV2Api().
		Post(ctx).
		AccessToken(accessToken).
		XOrangeCaller(request.XOrangeCaller).ToolsSiteHandselV2Request(request.ToolsSiteHandselV2Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
