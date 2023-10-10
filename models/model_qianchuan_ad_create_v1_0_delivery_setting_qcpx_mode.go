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

// QianchuanAdCreateV10DeliverySettingQcpxMode
type QianchuanAdCreateV10DeliverySettingQcpxMode string

// List of qianchuan_ad_create_v1.0_delivery_setting_qcpx_mode
const (
	QCPX_MODE_DEFAULT_QianchuanAdCreateV10DeliverySettingQcpxMode QianchuanAdCreateV10DeliverySettingQcpxMode = "QCPX_MODE_DEFAULT"
	QCPX_MODE_OFF_QianchuanAdCreateV10DeliverySettingQcpxMode     QianchuanAdCreateV10DeliverySettingQcpxMode = "QCPX_MODE_OFF"
	QCPX_MODE_ON_QianchuanAdCreateV10DeliverySettingQcpxMode      QianchuanAdCreateV10DeliverySettingQcpxMode = "QCPX_MODE_ON"
)

// All allowed values of QianchuanAdCreateV10DeliverySettingQcpxMode enum
var AllowedQianchuanAdCreateV10DeliverySettingQcpxModeEnumValues = []QianchuanAdCreateV10DeliverySettingQcpxMode{
	"QCPX_MODE_DEFAULT",
	"QCPX_MODE_OFF",
	"QCPX_MODE_ON",
}

// NewQianchuanAdCreateV10DeliverySettingQcpxModeFromValue returns a pointer to a valid QianchuanAdCreateV10DeliverySettingQcpxMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10DeliverySettingQcpxModeFromValue(v string) (*QianchuanAdCreateV10DeliverySettingQcpxMode, error) {
	ev := QianchuanAdCreateV10DeliverySettingQcpxMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10DeliverySettingQcpxMode: valid values are %v", v, AllowedQianchuanAdCreateV10DeliverySettingQcpxModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10DeliverySettingQcpxMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10DeliverySettingQcpxModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_qcpx_mode value
func (v QianchuanAdCreateV10DeliverySettingQcpxMode) Ptr() *QianchuanAdCreateV10DeliverySettingQcpxMode {
	return &v
}
