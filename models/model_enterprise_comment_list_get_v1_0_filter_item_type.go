/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentListGetV10FilterItemType
type EnterpriseCommentListGetV10FilterItemType string

// List of enterprise_comment_list_get_v1.0_filter_item_type
const (
	ITEM_AD_EnterpriseCommentListGetV10FilterItemType      EnterpriseCommentListGetV10FilterItemType = "ITEM_AD"
	ITEM_CONTENT_EnterpriseCommentListGetV10FilterItemType EnterpriseCommentListGetV10FilterItemType = "ITEM_CONTENT"
)

// Ptr returns reference to enterprise_comment_list_get_v1.0_filter_item_type value
func (v EnterpriseCommentListGetV10FilterItemType) Ptr() *EnterpriseCommentListGetV10FilterItemType {
	return &v
}
