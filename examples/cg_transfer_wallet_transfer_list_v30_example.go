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

type ApiOpenApiV30CgTransferWalletTransferListGetRequestExample struct {
	AccountId         int64                                      `json:"account_id"`
	AccountType       CgTransferWalletTransferListV30AccountType `json:"account_type"`
	BizRequestNo      string                                     `json:"biz_request_no"`
	QueryBeginTime    string                                     `json:"query_begin_time"`
	QueryEndTime      string                                     `json:"query_end_time"`
	QueryWalletIdList []int64                                    `json:"query_wallet_id_list"`
	PageInfo          CgTransferWalletTransferListV30PageInfo    `json:"page_info"`
	PayeeId           int64                                      `json:"payee_id,omitempty"`
	RemitterId        int64                                      `json:"remitter_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/cg_transfer/wallet/transfer/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30CgTransferWalletTransferListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.CgTransferWalletTransferListV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AccountId(request.AccountId).AccountType(request.AccountType).BizRequestNo(request.BizRequestNo).QueryBeginTime(request.QueryBeginTime).QueryEndTime(request.QueryEndTime).QueryWalletIdList(request.QueryWalletIdList).PageInfo(request.PageInfo).PayeeId(request.PayeeId).RemitterId(request.RemitterId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
