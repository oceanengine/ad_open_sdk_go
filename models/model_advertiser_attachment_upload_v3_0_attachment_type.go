/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserAttachmentUploadV30AttachmentType
type AdvertiserAttachmentUploadV30AttachmentType string

// List of advertiser_attachment_upload_v3.0_attachment_type
const (
	DELIVERY_ATTACHMENT_AdvertiserAttachmentUploadV30AttachmentType AdvertiserAttachmentUploadV30AttachmentType = "DELIVERY_ATTACHMENT"
)

// All allowed values of AdvertiserAttachmentUploadV30AttachmentType enum
var AllowedAdvertiserAttachmentUploadV30AttachmentTypeEnumValues = []AdvertiserAttachmentUploadV30AttachmentType{
	"DELIVERY_ATTACHMENT",
}

// NewAdvertiserAttachmentUploadV30AttachmentTypeFromValue returns a pointer to a valid AdvertiserAttachmentUploadV30AttachmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserAttachmentUploadV30AttachmentTypeFromValue(v string) (*AdvertiserAttachmentUploadV30AttachmentType, error) {
	ev := AdvertiserAttachmentUploadV30AttachmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserAttachmentUploadV30AttachmentType: valid values are %v", v, AllowedAdvertiserAttachmentUploadV30AttachmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserAttachmentUploadV30AttachmentType) IsValid() bool {
	for _, existing := range AllowedAdvertiserAttachmentUploadV30AttachmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_attachment_upload_v3.0_attachment_type value
func (v AdvertiserAttachmentUploadV30AttachmentType) Ptr() *AdvertiserAttachmentUploadV30AttachmentType {
	return &v
}
