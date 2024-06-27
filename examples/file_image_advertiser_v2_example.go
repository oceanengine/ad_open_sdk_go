/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

type ApiOpenApi2FileImageAdvertiserPostRequestExample struct {
	AdvertiserId   int64                           `json:"advertiser_id"`
	UploadTo       FileImageAdvertiserV2UploadTo   `json:"upload_to"`
	UploadType     FileImageAdvertiserV2UploadType `json:"upload_type"`
	ImageFile      *FormFileInfo                   `json:"image_file,omitempty"`
	ImageSignature string                          `json:"image_signature,omitempty"`
	ImageUrl       string                          `json:"image_url,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/file/image/advertiser/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2FileImageAdvertiserPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.FileImageAdvertiserV2Api().
		Post(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).UploadTo(request.UploadTo).UploadType(request.UploadType).ImageFile(request.ImageFile).ImageSignature(request.ImageSignature).ImageUrl(request.ImageUrl).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
