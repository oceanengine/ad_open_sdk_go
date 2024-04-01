/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserDeliveryQualificationListV30DataListQualificationType
type AdvertiserDeliveryQualificationListV30DataListQualificationType string

// List of advertiser_delivery_qualification_list_v3.0_data_list_qualification_type
const (
	AGENT_ARRANGEMENT_AdvertiserDeliveryQualificationListV30DataListQualificationType                           AdvertiserDeliveryQualificationListV30DataListQualificationType = "AGENT_ARRANGEMENT"
	AUTHORIZATION_CONTRACT_AdvertiserDeliveryQualificationListV30DataListQualificationType                      AdvertiserDeliveryQualificationListV30DataListQualificationType = "AUTHORIZATION_CONTRACT"
	DISTRIBUTION_AUTHORIZATION_AdvertiserDeliveryQualificationListV30DataListQualificationType                  AdvertiserDeliveryQualificationListV30DataListQualificationType = "DISTRIBUTION_AUTHORIZATION"
	ICP_RECORD_AUTHORIZATION_AdvertiserDeliveryQualificationListV30DataListQualificationType                    AdvertiserDeliveryQualificationListV30DataListQualificationType = "ICP_RECORD_AUTHORIZATION"
	OTHER_CERTIFICATION_AdvertiserDeliveryQualificationListV30DataListQualificationType                         AdvertiserDeliveryQualificationListV30DataListQualificationType = "OTHER_CERTIFICATION"
	PATENT_CERTIFICATE_AdvertiserDeliveryQualificationListV30DataListQualificationType                          AdvertiserDeliveryQualificationListV30DataListQualificationType = "PATENT_CERTIFICATE"
	PORTRAIT_AUTHORIZATION_AdvertiserDeliveryQualificationListV30DataListQualificationType                      AdvertiserDeliveryQualificationListV30DataListQualificationType = "PORTRAIT_AUTHORIZATION"
	QUALITY_REPORT_AdvertiserDeliveryQualificationListV30DataListQualificationType                              AdvertiserDeliveryQualificationListV30DataListQualificationType = "QUALITY_REPORT"
	SOFTWARE_COPYRIGHT_REGISTRATION_CERTIFICATE_AdvertiserDeliveryQualificationListV30DataListQualificationType AdvertiserDeliveryQualificationListV30DataListQualificationType = "SOFTWARE_COPYRIGHT_REGISTRATION_CERTIFICATE"
	TRADEMARK_REGISTRATION_CERTIFICATE_AdvertiserDeliveryQualificationListV30DataListQualificationType          AdvertiserDeliveryQualificationListV30DataListQualificationType = "TRADEMARK_REGISTRATION_CERTIFICATE"
	VIDEO_MATERIAL_PRODUCTION_IP_AUTHORIZATION_AdvertiserDeliveryQualificationListV30DataListQualificationType  AdvertiserDeliveryQualificationListV30DataListQualificationType = "VIDEO_MATERIAL_PRODUCTION_IP_AUTHORIZATION"
)

// All allowed values of AdvertiserDeliveryQualificationListV30DataListQualificationType enum
var AllowedAdvertiserDeliveryQualificationListV30DataListQualificationTypeEnumValues = []AdvertiserDeliveryQualificationListV30DataListQualificationType{
	"AGENT_ARRANGEMENT",
	"AUTHORIZATION_CONTRACT",
	"DISTRIBUTION_AUTHORIZATION",
	"ICP_RECORD_AUTHORIZATION",
	"OTHER_CERTIFICATION",
	"PATENT_CERTIFICATE",
	"PORTRAIT_AUTHORIZATION",
	"QUALITY_REPORT",
	"SOFTWARE_COPYRIGHT_REGISTRATION_CERTIFICATE",
	"TRADEMARK_REGISTRATION_CERTIFICATE",
	"VIDEO_MATERIAL_PRODUCTION_IP_AUTHORIZATION",
}

// NewAdvertiserDeliveryQualificationListV30DataListQualificationTypeFromValue returns a pointer to a valid AdvertiserDeliveryQualificationListV30DataListQualificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserDeliveryQualificationListV30DataListQualificationTypeFromValue(v string) (*AdvertiserDeliveryQualificationListV30DataListQualificationType, error) {
	ev := AdvertiserDeliveryQualificationListV30DataListQualificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserDeliveryQualificationListV30DataListQualificationType: valid values are %v", v, AllowedAdvertiserDeliveryQualificationListV30DataListQualificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserDeliveryQualificationListV30DataListQualificationType) IsValid() bool {
	for _, existing := range AllowedAdvertiserDeliveryQualificationListV30DataListQualificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_delivery_qualification_list_v3.0_data_list_qualification_type value
func (v AdvertiserDeliveryQualificationListV30DataListQualificationType) Ptr() *AdvertiserDeliveryQualificationListV30DataListQualificationType {
	return &v
}
