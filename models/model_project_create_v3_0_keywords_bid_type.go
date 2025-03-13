/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30KeywordsBidType
type ProjectCreateV30KeywordsBidType string

// List of project_create_v3.0_keywords_bid_type
const (
	CUSTOM_ProjectCreateV30KeywordsBidType         ProjectCreateV30KeywordsBidType = "CUSTOM"
	FEED_TO_SEARCH_ProjectCreateV30KeywordsBidType ProjectCreateV30KeywordsBidType = "FEED_TO_SEARCH"
	WITH_PROMOTION_ProjectCreateV30KeywordsBidType ProjectCreateV30KeywordsBidType = "WITH_PROMOTION"
)

// Ptr returns reference to project_create_v3.0_keywords_bid_type value
func (v ProjectCreateV30KeywordsBidType) Ptr() *ProjectCreateV30KeywordsBidType {
	return &v
}
