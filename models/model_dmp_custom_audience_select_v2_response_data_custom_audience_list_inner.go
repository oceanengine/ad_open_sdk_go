/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpCustomAudienceSelectV2ResponseDataCustomAudienceListInner struct for DmpCustomAudienceSelectV2ResponseDataCustomAudienceListInner
type DmpCustomAudienceSelectV2ResponseDataCustomAudienceListInner struct {
	//
	CoverNum *int64 `json:"cover_num,omitempty"`
	//
	CustomAudienceId *int64 `json:"custom_audience_id,omitempty"`
	//
	DataSourceId   *string                                                        `json:"data_source_id,omitempty"`
	DeliveryStatus *DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus `json:"delivery_status,omitempty"`
	//
	Isdel *int64 `json:"isdel,omitempty"`
	//
	Name   *string                                                `json:"name,omitempty"`
	Source *DmpCustomAudienceSelectV2DataCustomAudienceListSource `json:"source,omitempty"`
	//
	Status *int64 `json:"status,omitempty"`
	//
	Tag *string `json:"tag,omitempty"`
	//
	UploadNum *int64 `json:"upload_num,omitempty"`
}
