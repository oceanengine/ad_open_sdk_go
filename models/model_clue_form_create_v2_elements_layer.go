/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormCreateV2ElementsLayer
type ClueFormCreateV2ElementsLayer string

// List of clue_form_create_v2_elements_layer
const (
	Enum_0_ClueFormCreateV2ElementsLayer ClueFormCreateV2ElementsLayer = "0"
	Enum_1_ClueFormCreateV2ElementsLayer ClueFormCreateV2ElementsLayer = "1"
)

// All allowed values of ClueFormCreateV2ElementsLayer enum
var AllowedClueFormCreateV2ElementsLayerEnumValues = []ClueFormCreateV2ElementsLayer{
	"0",
	"1",
}

// NewClueFormCreateV2ElementsLayerFromValue returns a pointer to a valid ClueFormCreateV2ElementsLayer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormCreateV2ElementsLayerFromValue(v string) (*ClueFormCreateV2ElementsLayer, error) {
	ev := ClueFormCreateV2ElementsLayer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormCreateV2ElementsLayer: valid values are %v", v, AllowedClueFormCreateV2ElementsLayerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormCreateV2ElementsLayer) IsValid() bool {
	for _, existing := range AllowedClueFormCreateV2ElementsLayerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_create_v2_elements_layer value
func (v ClueFormCreateV2ElementsLayer) Ptr() *ClueFormCreateV2ElementsLayer {
	return &v
}
