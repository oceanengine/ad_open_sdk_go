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

type ApiOpenApi2ReportRtaCusExpGetGetRequestExample struct {
	AdvertiserId   int64  `json:"advertiser_id"`
	RtaInterfaceId int64  `json:"rta_interface_id"`
	RtaId          int64  `json:"rta_id"`
	RtaVid         string `json:"rta_vid"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
}

// url: https://api.oceanengine.com/open_api/2/report/rta_cus_exp/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ReportRtaCusExpGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportRtaCusExpGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).RtaInterfaceId(request.RtaInterfaceId).RtaId(request.RtaId).RtaVid(request.RtaVid).StartTime(request.StartTime).EndTime(request.EndTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
