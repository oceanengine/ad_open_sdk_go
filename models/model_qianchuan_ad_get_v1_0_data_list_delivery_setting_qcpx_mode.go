/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10DataListDeliverySettingQcpxMode
type QianchuanAdGetV10DataListDeliverySettingQcpxMode string

// List of qianchuan_ad_get_v1.0_data_list_delivery_setting_qcpx_mode
const (
	QCPX_MODE_DEFAULT_QianchuanAdGetV10DataListDeliverySettingQcpxMode QianchuanAdGetV10DataListDeliverySettingQcpxMode = "QCPX_MODE_DEFAULT"
	QCPX_MODE_OFF_QianchuanAdGetV10DataListDeliverySettingQcpxMode     QianchuanAdGetV10DataListDeliverySettingQcpxMode = "QCPX_MODE_OFF"
	QCPX_MODE_ON_QianchuanAdGetV10DataListDeliverySettingQcpxMode      QianchuanAdGetV10DataListDeliverySettingQcpxMode = "QCPX_MODE_ON"
)

// All allowed values of QianchuanAdGetV10DataListDeliverySettingQcpxMode enum
var AllowedQianchuanAdGetV10DataListDeliverySettingQcpxModeEnumValues = []QianchuanAdGetV10DataListDeliverySettingQcpxMode{
	"QCPX_MODE_DEFAULT",
	"QCPX_MODE_OFF",
	"QCPX_MODE_ON",
}

// NewQianchuanAdGetV10DataListDeliverySettingQcpxModeFromValue returns a pointer to a valid QianchuanAdGetV10DataListDeliverySettingQcpxMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListDeliverySettingQcpxModeFromValue(v string) (*QianchuanAdGetV10DataListDeliverySettingQcpxMode, error) {
	ev := QianchuanAdGetV10DataListDeliverySettingQcpxMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListDeliverySettingQcpxMode: valid values are %v", v, AllowedQianchuanAdGetV10DataListDeliverySettingQcpxModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListDeliverySettingQcpxMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListDeliverySettingQcpxModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_delivery_setting_qcpx_mode value
func (v QianchuanAdGetV10DataListDeliverySettingQcpxMode) Ptr() *QianchuanAdGetV10DataListDeliverySettingQcpxMode {
	return &v
}
