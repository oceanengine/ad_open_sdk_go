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

type ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequestExample struct {
	StarId                  int64 `json:"star_id"`
	Page                    int32 `json:"page"`
	PageSize                int32 `json:"page_size"`
	ProviderOrderTaskStatus int32 `json:"provider_order_task_status,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/mcn/provider_get_participated_task/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarMcnProviderGetParticipatedTaskV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).Page(request.Page).PageSize(request.PageSize).ProviderOrderTaskStatus(request.ProviderOrderTaskStatus).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
