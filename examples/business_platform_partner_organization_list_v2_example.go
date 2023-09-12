/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "code.byted.org/ad/ad_open_sdk_go"
	"code.byted.org/ad/ad_open_sdk_go/config"
	. "code.byted.org/ad/ad_open_sdk_go/models"
)

type ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequestExample struct {
	OrganizationId int64                                              `json:"organization_id"`
	Page           int32                                              `json:"page,omitempty"`
	PageSize       int32                                              `json:"page_size,omitempty"`
	Filtering      BusinessPlatformPartnerOrganizationListV2Filtering `json:"filtering,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/business_platform/partner_organization/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.BusinessPlatformPartnerOrganizationListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		OrganizationId(request.OrganizationId).Page(request.Page).PageSize(request.PageSize).Filtering(request.Filtering).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
