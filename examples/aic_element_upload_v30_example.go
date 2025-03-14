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

type ApiOpenApiV30AicElementUploadPostRequestExample struct {
	AccountId   int64                          `json:"account_id"`
	AccountType AicElementUploadV30AccountType `json:"account_type"`
	ElementType AicElementUploadV30ElementType `json:"element_type"`
	ElementName string                         `json:"element_name,omitempty"`
	File        *FormFileInfo                  `json:"file,omitempty"`
	PathName    string                         `json:"path_name,omitempty"`
	UseAs       AicElementUploadV30UseAs       `json:"use_as,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/aic/element/upload/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30AicElementUploadPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.AicElementUploadV30Api().
		Post(ctx).
		AccessToken(accessToken).
		AccountId(request.AccountId).AccountType(request.AccountType).ElementType(request.ElementType).ElementName(request.ElementName).File(request.File).PathName(request.PathName).UseAs(request.UseAs).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
