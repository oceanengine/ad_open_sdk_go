/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandRegionGetV30ResponseDataRegionsInner struct for BrandRegionGetV30ResponseDataRegionsInner
type BrandRegionGetV30ResponseDataRegionsInner struct {
	//
	Id *int64 `json:"id,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SubRegions []*BrandRegionGetV30ResponseDataRegionsInnerSubRegionsInner `json:"sub_regions,omitempty"`
}
