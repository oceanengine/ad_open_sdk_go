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

type ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequestExample struct {
	AppId        int64                                             `json:"app_id"`
	DataFileType ToolsPioneerProgramAttachmentUploadV2DataFileType `json:"data_file_type"`
	FileData     *FormFileInfo                                     `json:"file_data"`
	FileIndex    int64                                             `json:"file_index"`
	FileTotalNum int64                                             `json:"file_total_num"`
	PDate        string                                            `json:"p_date"`
	Platform     ToolsPioneerProgramAttachmentUploadV2Platform     `json:"platform"`
	DebugMode    bool                                              `json:"debug_mode,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/pioneer_program/attachment/upload/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsPioneerProgramAttachmentUploadV2Api().
		Post(ctx).
		AccessToken(accessToken).
		AppId(request.AppId).DataFileType(request.DataFileType).FileData(request.FileData).FileIndex(request.FileIndex).FileTotalNum(request.FileTotalNum).PDate(request.PDate).Platform(request.Platform).DebugMode(request.DebugMode).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
