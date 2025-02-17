/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueLifeCallbackV2ClueConvertState
type ToolsClueLifeCallbackV2ClueConvertState string

// List of tools_clue_life_callback_v2_clue_convert_state
const (
	ARRIVAL_ToolsClueLifeCallbackV2ClueConvertState             ToolsClueLifeCallbackV2ClueConvertState = "ARRIVAL"
	CLUE_CONFIRM_ToolsClueLifeCallbackV2ClueConvertState        ToolsClueLifeCallbackV2ClueConvertState = "CLUE_CONFIRM"
	CLUE_HIGH_INTENTION_ToolsClueLifeCallbackV2ClueConvertState ToolsClueLifeCallbackV2ClueConvertState = "CLUE_HIGH_INTENTION"
	CONVERSION_CLASS_ToolsClueLifeCallbackV2ClueConvertState    ToolsClueLifeCallbackV2ClueConvertState = "CONVERSION_CLASS"
	INVALID_EVENT_ToolsClueLifeCallbackV2ClueConvertState       ToolsClueLifeCallbackV2ClueConvertState = "INVALID_EVENT"
)

// Ptr returns reference to tools_clue_life_callback_v2_clue_convert_state value
func (v ToolsClueLifeCallbackV2ClueConvertState) Ptr() *ToolsClueLifeCallbackV2ClueConvertState {
	return &v
}
