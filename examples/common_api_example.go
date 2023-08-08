/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/api"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

func getDemo() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	resp, httpRes, err := apiClient.CommonApi().Get(ctx, "/open_api/test").
		AccessToken(accessToken).
		RequestQuery(map[string]interface{}{
			"app_id": 1,
		}).Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)
}

func postDemo() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	resp, httpRes, err := apiClient.CommonApi().Post(ctx, "/open_api/oauth2/access_token/").
		AccessToken(accessToken).
		RequestBody(map[string]interface{}{
			"app_id":    1,
			"secret":    "",
			"auth_code": "",
		}).Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)
}

func postMultipartDemo() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	fByte, err := os.ReadFile("test_file")
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.CommonApi().PostMultipart(ctx, "/open_api/test").
		AccessToken(accessToken).
		RequestForm(map[string]interface{}{
			"app_id": 1,
		}).
		RequestFile(map[string]api.FormFile{
			"test_file": api.FormFile{
				FileBytes: fByte,
				FileName:  "test",
			},
		}).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)
}

func main() {

}
