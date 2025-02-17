/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "ck " + string(c)
}

const Version = "1.1.38"

var (
	// ContextAccessToken takes a string access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")
	// ContextEnableLog  takes a bool as a flag to enable request response log
	ContextEnableLog = contextKey("enable_log")
)

// Configuration stores the configuration of the API client
type Configuration struct {
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	LogEnable     bool              `json:"log_enable,omitempty"`
	UseLogMw      bool              `json:"use_log_mw,omitempty"`
	HTTPClient    *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "Bytedance Ads Openapi SDK",
		Host:          "api.oceanengine.com",
		Scheme:        "https",
		LogEnable:     false,
		UseLogMw:      true,
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) GetBasePath() string {
	return c.Scheme + "://" + c.Host
}
