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

type ApiOpenApi2DiagnosisTaskAgentListGetRequestExample struct {
	AgentId       int64                              `json:"agent_id"`
	Results       []*DiagnosisTaskAgentListV2Results `json:"results,omitempty"`
	Status        []*DiagnosisTaskAgentListV2Status  `json:"status,omitempty"`
	StartTime     string                             `json:"start_time,omitempty"`
	EndTime       string                             `json:"end_time,omitempty"`
	Page          int32                              `json:"page,omitempty"`
	PageSize      int32                              `json:"page_size,omitempty"`
	StartDateTime string                             `json:"start_date_time,omitempty"`
	EndDateTime   string                             `json:"end_date_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/diagnosis_task/agent/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2DiagnosisTaskAgentListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.DiagnosisTaskAgentListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AgentId(request.AgentId).Results(request.Results).Status(request.Status).StartTime(request.StartTime).EndTime(request.EndTime).Page(request.Page).PageSize(request.PageSize).StartDateTime(request.StartDateTime).EndDateTime(request.EndDateTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
