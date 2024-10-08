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

type ApiOpenApi2StarChallengeGetCustomTaskDataGetRequestExample struct {
	StarId          int64 `json:"star_id"`
	ChallengeTaskId int64 `json:"challenge_task_id"`
	Page            int32 `json:"page"`
	PageSize        int32 `json:"page_size"`
}

// url: https://api.oceanengine.com/open_api/2/star/challenge/get_custom_task_data/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarChallengeGetCustomTaskDataGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarChallengeGetCustomTaskDataV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).ChallengeTaskId(request.ChallengeTaskId).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
