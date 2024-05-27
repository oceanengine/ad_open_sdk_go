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

type ApiOpenApiV30CgTransferQueryCanTransferBalanceGetRequestExample struct {
	BizRequestNo        string                                                `json:"biz_request_no"`
	AgentId             int64                                                 `json:"agent_id"`
	AccountId           int64                                                 `json:"account_id"`
	TargetAccountIdList []int64                                               `json:"target_account_id_list"`
	TransferDirection   CgTransferQueryCanTransferBalanceV30TransferDirection `json:"transfer_direction"`
}

// url: https://api.oceanengine.com/open_api/v3.0/cg_transfer/query_can_transfer_balance/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30CgTransferQueryCanTransferBalanceGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.CgTransferQueryCanTransferBalanceV30Api().
		Get(ctx).
		AccessToken(accessToken).
		BizRequestNo(request.BizRequestNo).AgentId(request.AgentId).AccountId(request.AccountId).TargetAccountIdList(request.TargetAccountIdList).TransferDirection(request.TransferDirection).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
