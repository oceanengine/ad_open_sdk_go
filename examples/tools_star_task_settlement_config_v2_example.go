/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

type ApiOpenApi2ToolsStarTaskSettlementConfigGetRequestExample struct {
	AdvertiserId           int64                                                 `json:"advertiser_id"`
	FirstIndustryId        int64                                                 `json:"first_industry_id"`
	SecondIndustryId       int64                                                 `json:"second_industry_id"`
	StarMaterialFirstType  int64                                                 `json:"star_material_first_type"`
	StarMaterialSecondType int64                                                 `json:"star_material_second_type"`
	StarTaskExternalAction ToolsStarTaskSettlementConfigV2StarTaskExternalAction `json:"star_task_external_action"`
}

// url: https://api.oceanengine.com/open_api/2/tools/star_task/settlement_config/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsStarTaskSettlementConfigGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsStarTaskSettlementConfigV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).FirstIndustryId(request.FirstIndustryId).SecondIndustryId(request.SecondIndustryId).StarMaterialFirstType(request.StarMaterialFirstType).StarMaterialSecondType(request.StarMaterialSecondType).StarTaskExternalAction(request.StarTaskExternalAction).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
