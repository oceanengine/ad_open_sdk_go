/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10AudienceAudienceMode
type QianchuanAdUpdateV10AudienceAudienceMode string

// List of qianchuan_ad_update_v1.0_audience_audience_mode
const (
	AUTO_QianchuanAdUpdateV10AudienceAudienceMode                QianchuanAdUpdateV10AudienceAudienceMode = "AUTO"
	AUTO_ORIENTATION_QianchuanAdUpdateV10AudienceAudienceMode    QianchuanAdUpdateV10AudienceAudienceMode = "AUTO_ORIENTATION"
	CUSTOM_QianchuanAdUpdateV10AudienceAudienceMode              QianchuanAdUpdateV10AudienceAudienceMode = "CUSTOM"
	NONE_QianchuanAdUpdateV10AudienceAudienceMode                QianchuanAdUpdateV10AudienceAudienceMode = "NONE"
	ORIENTATION_PACKAGE_QianchuanAdUpdateV10AudienceAudienceMode QianchuanAdUpdateV10AudienceAudienceMode = "ORIENTATION_PACKAGE"
)

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_audience_mode value
func (v QianchuanAdUpdateV10AudienceAudienceMode) Ptr() *QianchuanAdUpdateV10AudienceAudienceMode {
	return &v
}
