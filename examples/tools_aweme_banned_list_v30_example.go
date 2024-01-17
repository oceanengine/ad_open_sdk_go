/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
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

type ApiOpenApiV30ToolsAwemeBannedListGetRequestExample struct {
	AdvertiserId    int64                             `json:"advertiser_id"`
	BannedType      ToolsAwemeBannedListV30BannedType `json:"banned_type,omitempty"`
	AwemeId         string                            `json:"aweme_id,omitempty"`
	IsApplyToAdv    bool                              `json:"is_apply_to_adv,omitempty"`
	NicknameKeyword string                            `json:"nickname_keyword,omitempty"`
	AwemeUserId     string                            `json:"aweme_user_id,omitempty"`
	Page            int64                             `json:"page,omitempty"`
	PageSize        int64                             `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/tools/aweme_banned/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ToolsAwemeBannedListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAwemeBannedListV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).BannedType(request.BannedType).AwemeId(request.AwemeId).IsApplyToAdv(request.IsApplyToAdv).NicknameKeyword(request.NicknameKeyword).AwemeUserId(request.AwemeUserId).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
