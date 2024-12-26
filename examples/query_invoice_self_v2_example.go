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

type ApiOpenApi2QueryInvoiceSelfGetRequestExample struct {
	LocalAccountId    int64                                 `json:"local_account_id"`
	PageSize          int64                                 `json:"page_size"`
	Page              int64                                 `json:"page"`
	StatementSerials  []string                              `json:"statement_serials,omitempty"`
	ProjectSerials    []string                              `json:"project_serials,omitempty"`
	InvoiceStatuses   []int64                               `json:"invoice_statuses,omitempty"`
	InvoiceSerialList []string                              `json:"invoice_serial_list,omitempty"`
	ContractSerial    string                                `json:"contract_serial,omitempty"`
	SubmitStartTime   string                                `json:"submit_start_time,omitempty"`
	SubmitEndTime     string                                `json:"submit_end_time,omitempty"`
	InvoiceStartDate  string                                `json:"invoice_start_date,omitempty"`
	InvoiceEndDate    string                                `json:"invoice_end_date,omitempty"`
	InvoiceType       QueryInvoiceSelfV2InvoiceType         `json:"invoice_type,omitempty"`
	DifferenceInvoice QueryInvoiceSelfV2DifferenceInvoice   `json:"difference_invoice,omitempty"`
	RevertStatusList  []*QueryInvoiceSelfV2RevertStatusList `json:"revert_status_list,omitempty"`
	Platform          QueryInvoiceSelfV2Platform            `json:"platform,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/query/invoice/self/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2QueryInvoiceSelfGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QueryInvoiceSelfV2Api().
		Get(ctx).
		AccessToken(accessToken).
		LocalAccountId(request.LocalAccountId).PageSize(request.PageSize).Page(request.Page).StatementSerials(request.StatementSerials).ProjectSerials(request.ProjectSerials).InvoiceStatuses(request.InvoiceStatuses).InvoiceSerialList(request.InvoiceSerialList).ContractSerial(request.ContractSerial).SubmitStartTime(request.SubmitStartTime).SubmitEndTime(request.SubmitEndTime).InvoiceStartDate(request.InvoiceStartDate).InvoiceEndDate(request.InvoiceEndDate).InvoiceType(request.InvoiceType).DifferenceInvoice(request.DifferenceInvoice).RevertStatusList(request.RevertStatusList).Platform(request.Platform).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}