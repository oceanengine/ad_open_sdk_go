/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicMaterialPushV30AccountType
type AicMaterialPushV30AccountType string

// List of aic_material_push_v3.0_account_type
const (
	AGENT_AicMaterialPushV30AccountType AicMaterialPushV30AccountType = "AGENT"
	BP_AicMaterialPushV30AccountType    AicMaterialPushV30AccountType = "BP"
)

// Ptr returns reference to aic_material_push_v3.0_account_type value
func (v AicMaterialPushV30AccountType) Ptr() *AicMaterialPushV30AccountType {
	return &v
}
