/*
API version: 0.0.1
*/
package ad_open_sdk_go

import (
	"github.com/oceanengine/ad_open_sdk_go/api"
	"github.com/oceanengine/ad_open_sdk_go/config"
	"github.com/oceanengine/ad_open_sdk_go/middleware"
)

type Client struct {
	ApiClient *api.APIClient
}

func Init(cfg *config.Configuration) *Client {
	client := &Client{
		ApiClient: api.NewAPIClient(cfg),
	}
	client.ApiClient.Use(
		middleware.AuthMiddleware,
		middleware.HeaderMiddleware,
	)
	if cfg.UseLogMw {
		client.ApiClient.Use(middleware.LogMiddleware)
	}
	return client
}

func (c *Client) Oauth2AccessTokenApi() *api.Oauth2AccessTokenApiService {
	return c.ApiClient.Oauth2AccessTokenApi
}

func (c *Client) Oauth2RefreshTokenApi() *api.Oauth2RefreshTokenApiService {
	return c.ApiClient.Oauth2RefreshTokenApi
}

func (c *Client) ProjectCreateV30Api() *api.ProjectCreateV30ApiService {
	return c.ApiClient.ProjectCreateV30Api
}

func (c *Client) PromotionListV30Api() *api.PromotionListV30ApiService {
	return c.ApiClient.PromotionListV30Api
}

func (c *Client) CommonApi() *api.CommonApiService {
	return c.ApiClient.CommonApi
}

func (c *Client) SetHost(host string) {
	c.ApiClient.Cfg.Host = host
}

func (c *Client) Use(mws ...api.Middleware) {
	c.ApiClient.Use(mws...)
}

func (c *Client) AddDefaultHeader(key string, value string) {
	c.ApiClient.Cfg.AddDefaultHeader(key, value)
}

func (c *Client) SetLogEnable(enable bool) {
	c.ApiClient.Cfg.LogEnable = enable
}
