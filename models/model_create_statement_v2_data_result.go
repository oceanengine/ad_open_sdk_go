/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreateStatementV2DataResult
type CreateStatementV2DataResult string

// List of create_statement_v2_data_result
const (
	SUCCESS_CreateStatementV2DataResult CreateStatementV2DataResult = "SUCCESS"
	FAILED_CreateStatementV2DataResult  CreateStatementV2DataResult = "FAILED"
)

// Ptr returns reference to create_statement_v2_data_result value
func (v CreateStatementV2DataResult) Ptr() *CreateStatementV2DataResult {
	return &v
}
