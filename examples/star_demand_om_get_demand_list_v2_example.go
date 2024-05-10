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

type ApiOpenApi2StarDemandOmGetDemandListGetRequestExample struct {
	StarId                  int64                                              `json:"star_id"`
	PageNo                  int32                                              `json:"page_no"`
	PageSize                int32                                              `json:"page_size"`
	UniversalSettlementType StarDemandOmGetDemandListV2UniversalSettlementType `json:"universal_settlement_type,omitempty"`
	MicroAppId              string                                             `json:"micro_app_id,omitempty"`
	CreateStartTime         int64                                              `json:"create_start_time,omitempty"`
	CreateEndTime           int64                                              `json:"create_end_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/demand/om_get_demand_list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarDemandOmGetDemandListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarDemandOmGetDemandListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).PageNo(request.PageNo).PageSize(request.PageSize).UniversalSettlementType(request.UniversalSettlementType).MicroAppId(request.MicroAppId).CreateStartTime(request.CreateStartTime).CreateEndTime(request.CreateEndTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
