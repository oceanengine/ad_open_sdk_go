/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupGetV2ResponseDataListInnerSitesInner struct for ToolsLandingGroupGetV2ResponseDataListInnerSitesInner
type ToolsLandingGroupGetV2ResponseDataListInnerSitesInner struct {
	//
	MemberId        *int64                                              `json:"member_id,omitempty"`
	SiteAuditStatus *ToolsLandingGroupGetV2DataListSitesSiteAuditStatus `json:"site_audit_status,omitempty"`
	//
	SiteId        *int64                                            `json:"site_id,omitempty"`
	SiteOptStatus *ToolsLandingGroupGetV2DataListSitesSiteOptStatus `json:"site_opt_status,omitempty"`
	//
	SiteUrl *string `json:"site_url,omitempty"`
}
