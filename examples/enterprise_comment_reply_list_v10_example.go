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

type ApiOpenApiV10EnterpriseCommentReplyListGetRequestExample struct {
	CcAccountId int64  `json:"cc_account_id,omitempty"`
	CommentId   int64  `json:"comment_id,omitempty"`
	EDouyinId   string `json:"e_douyin_id,omitempty"`
	Page        int64  `json:"page,omitempty"`
	PageSize    int64  `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/enterprise/comment/reply/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10EnterpriseCommentReplyListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.EnterpriseCommentReplyListV10Api().
		Get(ctx).
		AccessToken(accessToken).
		CcAccountId(request.CcAccountId).CommentId(request.CommentId).EDouyinId(request.EDouyinId).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
