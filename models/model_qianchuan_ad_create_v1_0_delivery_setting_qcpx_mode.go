/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10DeliverySettingQcpxMode
type QianchuanAdCreateV10DeliverySettingQcpxMode string

// List of qianchuan_ad_create_v1.0_delivery_setting_qcpx_mode
const (
	QCPX_MODE_DEFAULT_QianchuanAdCreateV10DeliverySettingQcpxMode QianchuanAdCreateV10DeliverySettingQcpxMode = "QCPX_MODE_DEFAULT"
	QCPX_MODE_OFF_QianchuanAdCreateV10DeliverySettingQcpxMode     QianchuanAdCreateV10DeliverySettingQcpxMode = "QCPX_MODE_OFF"
	QCPX_MODE_ON_QianchuanAdCreateV10DeliverySettingQcpxMode      QianchuanAdCreateV10DeliverySettingQcpxMode = "QCPX_MODE_ON"
)

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_qcpx_mode value
func (v QianchuanAdCreateV10DeliverySettingQcpxMode) Ptr() *QianchuanAdCreateV10DeliverySettingQcpxMode {
	return &v
}
