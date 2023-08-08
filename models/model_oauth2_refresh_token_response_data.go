/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2RefreshTokenResponseData
type Oauth2RefreshTokenResponseData struct {
	//
	AccessToken *string `json:"access_token,omitempty"`
	//
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	//
	RefreshToken *string `json:"refresh_token,omitempty"`
	//
	RefreshTokenExpiresIn *int64 `json:"refresh_token_expires_in,omitempty"`
}
