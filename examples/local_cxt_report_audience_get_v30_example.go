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

type ApiOpenApiV30LocalCxtReportAudienceGetGetRequestExample struct {
	LocalAccountId int64                                 `json:"local_account_id"`
	StartDate      string                                `json:"start_date"`
	EndDate        string                                `json:"end_date"`
	Metrics        []string                              `json:"metrics"`
	Dimension      LocalCxtReportAudienceGetV30Dimension `json:"dimension"`
}

// url: https://api.oceanengine.com/open_api/v3.0/local/cxt/report/audience/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30LocalCxtReportAudienceGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.LocalCxtReportAudienceGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		LocalAccountId(request.LocalAccountId).StartDate(request.StartDate).EndDate(request.EndDate).Metrics(request.Metrics).Dimension(request.Dimension).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
