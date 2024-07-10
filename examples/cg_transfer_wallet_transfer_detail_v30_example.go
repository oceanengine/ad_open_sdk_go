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

type ApiOpenApiV30CgTransferWalletTransferDetailGetRequestExample struct {
	AccountId            int64                                        `json:"account_id"`
	AccountType          CgTransferWalletTransferDetailV30AccountType `json:"account_type"`
	BizRequestNo         string                                       `json:"biz_request_no"`
	TransferBizRequestNo string                                       `json:"transfer_biz_request_no,omitempty"`
	TransferSerial       string                                       `json:"transfer_serial,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/cg_transfer/wallet/transfer/detail/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30CgTransferWalletTransferDetailGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.CgTransferWalletTransferDetailV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AccountId(request.AccountId).AccountType(request.AccountType).BizRequestNo(request.BizRequestNo).TransferBizRequestNo(request.TransferBizRequestNo).TransferSerial(request.TransferSerial).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
