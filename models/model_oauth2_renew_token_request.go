/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2RenewTokenRequest struct for Oauth2RenewTokenRequest
type Oauth2RenewTokenRequest struct {
	//
	AppId int64 `json:"app_id"`
	//
	RefreshToken string `json:"refresh_token"`
	//
	Secret string `json:"secret"`
}
