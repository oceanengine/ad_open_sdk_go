/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

type ApiOpenApi2ToolsAppManagementShareAccountListGetRequestExample struct {
	OrganizationId int64                                          `json:"organization_id,omitempty"`
	PackageId      string                                         `json:"package_id,omitempty"`
	Page           int64                                          `json:"page,omitempty"`
	PageSize       int64                                          `json:"page_size,omitempty"`
	SearchType     ToolsAppManagementShareAccountListV2SearchType `json:"search_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/app_management/share_account/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAppManagementShareAccountListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAppManagementShareAccountListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		OrganizationId(request.OrganizationId).PackageId(request.PackageId).Page(request.Page).PageSize(request.PageSize).SearchType(request.SearchType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
