/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCreateV2SiteType
type ToolsSiteCreateV2SiteType string

// List of tools_site_create_v2_site_type
const (
	STREAMING_ToolsSiteCreateV2SiteType     ToolsSiteCreateV2SiteType = "STREAMING"
	ECOMMERCEPAGE_ToolsSiteCreateV2SiteType ToolsSiteCreateV2SiteType = "ECOMMERCEPAGE"
)

// Ptr returns reference to tools_site_create_v2_site_type value
func (v ToolsSiteCreateV2SiteType) Ptr() *ToolsSiteCreateV2SiteType {
	return &v
}
