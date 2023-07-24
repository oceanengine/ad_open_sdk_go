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

func (c *Client) CreativeStrategyListV2Api() *api.CreativeStrategyListV2ApiService {
	return c.ApiClient.CreativeStrategyListV2Api
}

func (c *Client) DouplusOrderListV30Api() *api.DouplusOrderListV30ApiService {
	return c.ApiClient.DouplusOrderListV30Api
}

func (c *Client) DouplusOrderReportV30Api() *api.DouplusOrderReportV30ApiService {
	return c.ApiClient.DouplusOrderReportV30Api
}

func (c *Client) MaterialStatusUpdateV30Api() *api.MaterialStatusUpdateV30ApiService {
	return c.ApiClient.MaterialStatusUpdateV30Api
}

func (c *Client) Oauth2AccessTokenApi() *api.Oauth2AccessTokenApiService {
	return c.ApiClient.Oauth2AccessTokenApi
}

func (c *Client) Oauth2RefreshTokenApi() *api.Oauth2RefreshTokenApiService {
	return c.ApiClient.Oauth2RefreshTokenApi
}

func (c *Client) ProjectBudgetUpdateV30Api() *api.ProjectBudgetUpdateV30ApiService {
	return c.ApiClient.ProjectBudgetUpdateV30Api
}

func (c *Client) ProjectCreateV30Api() *api.ProjectCreateV30ApiService {
	return c.ApiClient.ProjectCreateV30Api
}

func (c *Client) ProjectDeleteV30Api() *api.ProjectDeleteV30ApiService {
	return c.ApiClient.ProjectDeleteV30Api
}

func (c *Client) ProjectListV30Api() *api.ProjectListV30ApiService {
	return c.ApiClient.ProjectListV30Api
}

func (c *Client) ProjectStatusUpdateV30Api() *api.ProjectStatusUpdateV30ApiService {
	return c.ApiClient.ProjectStatusUpdateV30Api
}

func (c *Client) ProjectUpdateV30Api() *api.ProjectUpdateV30ApiService {
	return c.ApiClient.ProjectUpdateV30Api
}

func (c *Client) PromotionAutoGenerateConfigCreateV30Api() *api.PromotionAutoGenerateConfigCreateV30ApiService {
	return c.ApiClient.PromotionAutoGenerateConfigCreateV30Api
}

func (c *Client) PromotionAutoGenerateConfigGetV30Api() *api.PromotionAutoGenerateConfigGetV30ApiService {
	return c.ApiClient.PromotionAutoGenerateConfigGetV30Api
}

func (c *Client) PromotionBidUpdateV30Api() *api.PromotionBidUpdateV30ApiService {
	return c.ApiClient.PromotionBidUpdateV30Api
}

func (c *Client) PromotionBudgetUpdateV30Api() *api.PromotionBudgetUpdateV30ApiService {
	return c.ApiClient.PromotionBudgetUpdateV30Api
}

func (c *Client) PromotionCostProtectStatusGetV30Api() *api.PromotionCostProtectStatusGetV30ApiService {
	return c.ApiClient.PromotionCostProtectStatusGetV30Api
}

func (c *Client) PromotionCreateV30Api() *api.PromotionCreateV30ApiService {
	return c.ApiClient.PromotionCreateV30Api
}

func (c *Client) PromotionDeepbidUpdateV30Api() *api.PromotionDeepbidUpdateV30ApiService {
	return c.ApiClient.PromotionDeepbidUpdateV30Api
}

func (c *Client) PromotionDeleteV30Api() *api.PromotionDeleteV30ApiService {
	return c.ApiClient.PromotionDeleteV30Api
}

func (c *Client) PromotionListV30Api() *api.PromotionListV30ApiService {
	return c.ApiClient.PromotionListV30Api
}

func (c *Client) PromotionRejectReasonGetV30Api() *api.PromotionRejectReasonGetV30ApiService {
	return c.ApiClient.PromotionRejectReasonGetV30Api
}

func (c *Client) PromotionStatusUpdateV30Api() *api.PromotionStatusUpdateV30ApiService {
	return c.ApiClient.PromotionStatusUpdateV30Api
}

func (c *Client) PromotionUpdateV30Api() *api.PromotionUpdateV30ApiService {
	return c.ApiClient.PromotionUpdateV30Api
}

func (c *Client) ReportCustomConfigGetV30Api() *api.ReportCustomConfigGetV30ApiService {
	return c.ApiClient.ReportCustomConfigGetV30Api
}

func (c *Client) ReportCustomGetV30Api() *api.ReportCustomGetV30ApiService {
	return c.ApiClient.ReportCustomGetV30Api
}

func (c *Client) SuggWordsV30Api() *api.SuggWordsV30ApiService {
	return c.ApiClient.SuggWordsV30Api
}

func (c *Client) ToolsKeywordsBidRatioCreateV30Api() *api.ToolsKeywordsBidRatioCreateV30ApiService {
	return c.ApiClient.ToolsKeywordsBidRatioCreateV30Api
}

func (c *Client) ToolsKeywordsBidRatioDeleteV30Api() *api.ToolsKeywordsBidRatioDeleteV30ApiService {
	return c.ApiClient.ToolsKeywordsBidRatioDeleteV30Api
}

func (c *Client) ToolsKeywordsBidRatioGetV30Api() *api.ToolsKeywordsBidRatioGetV30ApiService {
	return c.ApiClient.ToolsKeywordsBidRatioGetV30Api
}

func (c *Client) ToolsKeywordsBidRatioUpdateV30Api() *api.ToolsKeywordsBidRatioUpdateV30ApiService {
	return c.ApiClient.ToolsKeywordsBidRatioUpdateV30Api
}

func (c *Client) ToolsKeywordsProjectInfoGetV30Api() *api.ToolsKeywordsProjectInfoGetV30ApiService {
	return c.ApiClient.ToolsKeywordsProjectInfoGetV30Api
}

func (c *Client) ToolsPrivativeWordBatchGetV30Api() *api.ToolsPrivativeWordBatchGetV30ApiService {
	return c.ApiClient.ToolsPrivativeWordBatchGetV30Api
}

func (c *Client) ToolsPrivativeWordProjectAddV30Api() *api.ToolsPrivativeWordProjectAddV30ApiService {
	return c.ApiClient.ToolsPrivativeWordProjectAddV30Api
}

func (c *Client) ToolsPrivativeWordProjectUpdateV30Api() *api.ToolsPrivativeWordProjectUpdateV30ApiService {
	return c.ApiClient.ToolsPrivativeWordProjectUpdateV30Api
}

func (c *Client) ToolsPrivativeWordPromotionAddV30Api() *api.ToolsPrivativeWordPromotionAddV30ApiService {
	return c.ApiClient.ToolsPrivativeWordPromotionAddV30Api
}

func (c *Client) ToolsPrivativeWordPromotionUpdateV30Api() *api.ToolsPrivativeWordPromotionUpdateV30ApiService {
	return c.ApiClient.ToolsPrivativeWordPromotionUpdateV30Api
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
