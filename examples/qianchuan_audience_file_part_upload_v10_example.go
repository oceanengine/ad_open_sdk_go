/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
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

type ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequestExample struct {
	AdvertiserId int64         `json:"advertiser_id"`
	File         *FormFileInfo `json:"file"`
	IsFinished   int32         `json:"is_finished"`
	PartNum      int64         `json:"part_num"`
	FileKey      string        `json:"file_key,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/audience_file/part_upload/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanAudienceFilePartUploadPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanAudienceFilePartUploadV10Api().
		Post(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).File(request.File).IsFinished(request.IsFinished).PartNum(request.PartNum).FileKey(request.FileKey).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
