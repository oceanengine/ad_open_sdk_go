/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SecurityOpenMaterialAuditV30BusinessType
type SecurityOpenMaterialAuditV30BusinessType string

// List of security_open_material_audit_v3.0_business_type
const (
	AD_SecurityOpenMaterialAuditV30BusinessType         SecurityOpenMaterialAuditV30BusinessType = "AD"
	QIAN_CHUAN_SecurityOpenMaterialAuditV30BusinessType SecurityOpenMaterialAuditV30BusinessType = "QIAN_CHUAN"
	LOCAL_LIFE_SecurityOpenMaterialAuditV30BusinessType SecurityOpenMaterialAuditV30BusinessType = "LOCAL_LIFE"
)

// Ptr returns reference to security_open_material_audit_v3.0_business_type value
func (v SecurityOpenMaterialAuditV30BusinessType) Ptr() *SecurityOpenMaterialAuditV30BusinessType {
	return &v
}
