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

type ApiOpenApiV30SharedWalletMainWalletGetGetRequestExample struct {
	AccountId    int64                                   `json:"account_id"`
	MainWalletId int64                                   `json:"main_wallet_id"`
	AccountType  SharedWalletMainWalletGetV30AccountType `json:"account_type"`
}

// url: https://api.oceanengine.com/open_api/v3.0/shared_wallet/main_wallet/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30SharedWalletMainWalletGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.SharedWalletMainWalletGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AccountId(request.AccountId).MainWalletId(request.MainWalletId).AccountType(request.AccountType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
