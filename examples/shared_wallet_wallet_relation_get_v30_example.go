/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

type ApiOpenApiV30SharedWalletWalletRelationGetGetRequestExample struct {
	AccountId      int64                                       `json:"account_id"`
	SharedWalletId int64                                       `json:"shared_wallet_id"`
	AccountType    SharedWalletWalletRelationGetV30AccountType `json:"account_type"`
	Page           int32                                       `json:"page,omitempty"`
	PageSize       int32                                       `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/shared_wallet/wallet_relation/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30SharedWalletWalletRelationGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.SharedWalletWalletRelationGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AccountId(request.AccountId).SharedWalletId(request.SharedWalletId).AccountType(request.AccountType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}