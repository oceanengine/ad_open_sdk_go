/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerCommonWalletInfo 通用钱包信息
type SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerCommonWalletInfo struct {
	// 钱包创建时间
	CreateTime         *string                                                                       `json:"create_time,omitempty"`
	SharedWalletStatus *SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus `json:"shared_wallet_status,omitempty"`
	// 钱包描述
	WalletDescription *string `json:"wallet_description,omitempty"`
	// 钱包标签
	WalletLabel []string `json:"wallet_label,omitempty"`
	// 钱包名称
	WalletName *string `json:"wallet_name,omitempty"`
}
