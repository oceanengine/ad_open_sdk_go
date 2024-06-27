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

type ApiOpenApi2FileAudioAdPostRequestExample struct {
	AdvertiserId   int64                   `json:"advertiser_id"`
	UploadType     FileAudioAdV2UploadType `json:"upload_type"`
	AudioFile      *FormFileInfo           `json:"audio_file,omitempty"`
	AudioSignature string                  `json:"audio_signature,omitempty"`
	AudioUrl       string                  `json:"audio_url,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/file/audio/ad/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2FileAudioAdPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.FileAudioAdV2Api().
		Post(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).UploadType(request.UploadType).AudioFile(request.AudioFile).AudioSignature(request.AudioSignature).AudioUrl(request.AudioUrl).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
