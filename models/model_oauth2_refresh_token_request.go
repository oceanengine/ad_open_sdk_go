/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2RefreshTokenRequest struct for Oauth2RefreshTokenRequest
type Oauth2RefreshTokenRequest struct {
	//
	AppId *int64 `json:"app_id,omitempty"`
	//
	RefreshToken string `json:"refresh_token"`
	//
	Secret string `json:"secret"`
}
