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

type ApiOpenApi2StarDemanderDemanderGetAuthorBindInfoGetRequestExample struct {
	StarId    int64    `json:"star_id"`
	BizUid    []string `json:"biz_uid,omitempty"`
	StartTime int64    `json:"start_time,omitempty"`
	EndTime   int64    `json:"end_time,omitempty"`
	Page      int64    `json:"page,omitempty"`
	Limit     int64    `json:"limit,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/demander/demander_get_author_bind_info/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarDemanderDemanderGetAuthorBindInfoGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarDemanderDemanderGetAuthorBindInfoV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).BizUid(request.BizUid).StartTime(request.StartTime).EndTime(request.EndTime).Page(request.Page).Limit(request.Limit).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
