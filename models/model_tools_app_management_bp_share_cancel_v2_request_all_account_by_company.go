/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBpShareCancelV2RequestAllAccountByCompany
type ToolsAppManagementBpShareCancelV2RequestAllAccountByCompany struct {
	//
	AccountType []*ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType `json:"account_type"`
	//
	CompanyId int64 `json:"company_id"`
}
