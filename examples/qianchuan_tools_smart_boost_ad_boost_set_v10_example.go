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

type ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequestExample struct {
	QianchuanToolsSmartBoostAdBoostSetV10Request QianchuanToolsSmartBoostAdBoostSetV10Request `json:"QianchuanToolsSmartBoostAdBoostSetV10Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/set/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanToolsSmartBoostAdBoostSetV10Api().
		Post(ctx).
		AccessToken(accessToken).
		QianchuanToolsSmartBoostAdBoostSetV10Request(request.QianchuanToolsSmartBoostAdBoostSetV10Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
