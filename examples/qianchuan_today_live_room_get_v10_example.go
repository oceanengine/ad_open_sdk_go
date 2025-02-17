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

type ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequestExample struct {
	AdvertiserId int64                                  `json:"advertiser_id"`
	AwemeId      int64                                  `json:"aweme_id"`
	DateTime     string                                 `json:"date_time"`
	Fields       []string                               `json:"fields"`
	RoomStatus   QianchuanTodayLiveRoomGetV10RoomStatus `json:"room_status,omitempty"`
	AdStatus     QianchuanTodayLiveRoomGetV10AdStatus   `json:"ad_status,omitempty"`
	Page         int32                                  `json:"page,omitempty"`
	PageSize     int32                                  `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/today_live/room/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanTodayLiveRoomGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AwemeId(request.AwemeId).DateTime(request.DateTime).Fields(request.Fields).RoomStatus(request.RoomStatus).AdStatus(request.AdStatus).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
