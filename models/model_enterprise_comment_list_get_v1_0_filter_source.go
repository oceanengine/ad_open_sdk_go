/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentListGetV10FilterSource
type EnterpriseCommentListGetV10FilterSource string

// List of enterprise_comment_list_get_v1.0_filter_source
const (
	FROM_PERFORM_EnterpriseCommentListGetV10FilterSource EnterpriseCommentListGetV10FilterSource = "FROM_PERFORM"
	FROM_NATURAL_EnterpriseCommentListGetV10FilterSource EnterpriseCommentListGetV10FilterSource = "FROM_NATURAL"
	FROM_OTHER_EnterpriseCommentListGetV10FilterSource   EnterpriseCommentListGetV10FilterSource = "FROM_OTHER"
	FROM_DOUPLUS_EnterpriseCommentListGetV10FilterSource EnterpriseCommentListGetV10FilterSource = "FROM_DOUPLUS"
	FROM_BRAND_EnterpriseCommentListGetV10FilterSource   EnterpriseCommentListGetV10FilterSource = "FROM_BRAND"
)

// Ptr returns reference to enterprise_comment_list_get_v1.0_filter_source value
func (v EnterpriseCommentListGetV10FilterSource) Ptr() *EnterpriseCommentListGetV10FilterSource {
	return &v
}
