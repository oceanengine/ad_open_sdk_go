/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2AppAccessTokenResponse struct for Oauth2AppAccessTokenResponse
type Oauth2AppAccessTokenResponse struct {
	//
	Code *int64                            `json:"code,omitempty"`
	Data *Oauth2AppAccessTokenResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
