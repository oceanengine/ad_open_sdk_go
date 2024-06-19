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

type ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequestExample struct {
	AssetType      ToolsBpAssetManagementShareGetV30AssetType `json:"asset_type"`
	InstanceId     int64                                      `json:"instance_id"`
	OrganizationId int64                                      `json:"organization_id,omitempty"`
	Page           int32                                      `json:"page,omitempty"`
	PageSize       int32                                      `json:"page_size,omitempty"`
	ShareType      ToolsBpAssetManagementShareGetV30ShareType `json:"share_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/tools/bp_asset_management/share/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsBpAssetManagementShareGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AssetType(request.AssetType).InstanceId(request.InstanceId).OrganizationId(request.OrganizationId).Page(request.Page).PageSize(request.PageSize).ShareType(request.ShareType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
