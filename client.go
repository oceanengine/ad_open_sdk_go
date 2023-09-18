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

func (c *Client) AdlabGroupDetailV30Api() *api.AdlabGroupDetailV30ApiService {
	return c.ApiClient.AdlabGroupDetailV30Api
}

func (c *Client) AdlabGroupListV30Api() *api.AdlabGroupListV30ApiService {
	return c.ApiClient.AdlabGroupListV30Api
}

func (c *Client) AdlabGroupUpdateStatusV30Api() *api.AdlabGroupUpdateStatusV30ApiService {
	return c.ApiClient.AdlabGroupUpdateStatusV30Api
}

func (c *Client) AdvertiserAvatarGetV2Api() *api.AdvertiserAvatarGetV2ApiService {
	return c.ApiClient.AdvertiserAvatarGetV2Api
}

func (c *Client) AdvertiserAvatarSubmitV2Api() *api.AdvertiserAvatarSubmitV2ApiService {
	return c.ApiClient.AdvertiserAvatarSubmitV2Api
}

func (c *Client) AdvertiserBudgetGetV2Api() *api.AdvertiserBudgetGetV2ApiService {
	return c.ApiClient.AdvertiserBudgetGetV2Api
}

func (c *Client) AdvertiserDeliveryQualificationListV30Api() *api.AdvertiserDeliveryQualificationListV30ApiService {
	return c.ApiClient.AdvertiserDeliveryQualificationListV30Api
}

func (c *Client) AdvertiserDeliveryQualificationSubmitV30Api() *api.AdvertiserDeliveryQualificationSubmitV30ApiService {
	return c.ApiClient.AdvertiserDeliveryQualificationSubmitV30Api
}

func (c *Client) AdvertiserFundDailyStatV2Api() *api.AdvertiserFundDailyStatV2ApiService {
	return c.ApiClient.AdvertiserFundDailyStatV2Api
}

func (c *Client) AdvertiserFundTransactionGetV2Api() *api.AdvertiserFundTransactionGetV2ApiService {
	return c.ApiClient.AdvertiserFundTransactionGetV2Api
}

func (c *Client) AdvertiserInfoV2Api() *api.AdvertiserInfoV2ApiService {
	return c.ApiClient.AdvertiserInfoV2Api
}

func (c *Client) AdvertiserPublicInfoV2Api() *api.AdvertiserPublicInfoV2ApiService {
	return c.ApiClient.AdvertiserPublicInfoV2Api
}

func (c *Client) AdvertiserQualificationCreateV2V2Api() *api.AdvertiserQualificationCreateV2V2ApiService {
	return c.ApiClient.AdvertiserQualificationCreateV2V2Api
}

func (c *Client) AdvertiserQualificationGetV30Api() *api.AdvertiserQualificationGetV30ApiService {
	return c.ApiClient.AdvertiserQualificationGetV30Api
}

func (c *Client) AdvertiserQualificationSelectV2V2Api() *api.AdvertiserQualificationSelectV2V2ApiService {
	return c.ApiClient.AdvertiserQualificationSelectV2V2Api
}

func (c *Client) AdvertiserQualificationSubmitV30Api() *api.AdvertiserQualificationSubmitV30ApiService {
	return c.ApiClient.AdvertiserQualificationSubmitV30Api
}

func (c *Client) AdvertiserTransferableFundGetV2Api() *api.AdvertiserTransferableFundGetV2ApiService {
	return c.ApiClient.AdvertiserTransferableFundGetV2Api
}

func (c *Client) AdvertiserUpdateBudgetV2Api() *api.AdvertiserUpdateBudgetV2ApiService {
	return c.ApiClient.AdvertiserUpdateBudgetV2Api
}

func (c *Client) AgentAdvertiserSelectV2Api() *api.AgentAdvertiserSelectV2ApiService {
	return c.ApiClient.AgentAdvertiserSelectV2Api
}

func (c *Client) AgentAdvertiserUpdateV2Api() *api.AgentAdvertiserUpdateV2ApiService {
	return c.ApiClient.AgentAdvertiserUpdateV2Api
}

func (c *Client) AgentChildAgentSelectV2Api() *api.AgentChildAgentSelectV2ApiService {
	return c.ApiClient.AgentChildAgentSelectV2Api
}

func (c *Client) AgentFundTransferSeqCommitV2Api() *api.AgentFundTransferSeqCommitV2ApiService {
	return c.ApiClient.AgentFundTransferSeqCommitV2Api
}

func (c *Client) AgentFundTransferSeqCreateV2Api() *api.AgentFundTransferSeqCreateV2ApiService {
	return c.ApiClient.AgentFundTransferSeqCreateV2Api
}

func (c *Client) AgentInfoV2Api() *api.AgentInfoV2ApiService {
	return c.ApiClient.AgentInfoV2Api
}

func (c *Client) AgentRefundTransferSeqCommitV2Api() *api.AgentRefundTransferSeqCommitV2ApiService {
	return c.ApiClient.AgentRefundTransferSeqCommitV2Api
}

func (c *Client) AgentRefundTransferSeqCreateV2Api() *api.AgentRefundTransferSeqCreateV2ApiService {
	return c.ApiClient.AgentRefundTransferSeqCreateV2Api
}

func (c *Client) AssetsCreativeComponentCreateV2Api() *api.AssetsCreativeComponentCreateV2ApiService {
	return c.ApiClient.AssetsCreativeComponentCreateV2Api
}

func (c *Client) AssetsCreativeComponentGetV2Api() *api.AssetsCreativeComponentGetV2ApiService {
	return c.ApiClient.AssetsCreativeComponentGetV2Api
}

func (c *Client) AssetsCreativeComponentUpdateV2Api() *api.AssetsCreativeComponentUpdateV2ApiService {
	return c.ApiClient.AssetsCreativeComponentUpdateV2Api
}

func (c *Client) AsyncTaskCreateV2Api() *api.AsyncTaskCreateV2ApiService {
	return c.ApiClient.AsyncTaskCreateV2Api
}

func (c *Client) AsyncTaskDownloadV2Api() *api.AsyncTaskDownloadV2ApiService {
	return c.ApiClient.AsyncTaskDownloadV2Api
}

func (c *Client) AsyncTaskGetV2Api() *api.AsyncTaskGetV2ApiService {
	return c.ApiClient.AsyncTaskGetV2Api
}

func (c *Client) AudiencePackageDeleteV2Api() *api.AudiencePackageDeleteV2ApiService {
	return c.ApiClient.AudiencePackageDeleteV2Api
}

func (c *Client) AudiencePackageUpdateV2Api() *api.AudiencePackageUpdateV2ApiService {
	return c.ApiClient.AudiencePackageUpdateV2Api
}

func (c *Client) BusinessPlatformCompanyAccountGetV30Api() *api.BusinessPlatformCompanyAccountGetV30ApiService {
	return c.ApiClient.BusinessPlatformCompanyAccountGetV30Api
}

func (c *Client) BusinessPlatformCompanyInfoGetV30Api() *api.BusinessPlatformCompanyInfoGetV30ApiService {
	return c.ApiClient.BusinessPlatformCompanyInfoGetV30Api
}

func (c *Client) BusinessPlatformPartnerOrganizationListV2Api() *api.BusinessPlatformPartnerOrganizationListV2ApiService {
	return c.ApiClient.BusinessPlatformPartnerOrganizationListV2Api
}

func (c *Client) CdpBrandGetV30Api() *api.CdpBrandGetV30ApiService {
	return c.ApiClient.CdpBrandGetV30Api
}

func (c *Client) ClueCouponCodeConsumeV2Api() *api.ClueCouponCodeConsumeV2ApiService {
	return c.ApiClient.ClueCouponCodeConsumeV2Api
}

func (c *Client) ClueCouponCodeGetV2Api() *api.ClueCouponCodeGetV2ApiService {
	return c.ApiClient.ClueCouponCodeGetV2Api
}

func (c *Client) ClueCouponCreateV2Api() *api.ClueCouponCreateV2ApiService {
	return c.ApiClient.ClueCouponCreateV2Api
}

func (c *Client) ClueCouponDetailV2Api() *api.ClueCouponDetailV2ApiService {
	return c.ApiClient.ClueCouponDetailV2Api
}

func (c *Client) ClueCouponEmployeeCreateV2Api() *api.ClueCouponEmployeeCreateV2ApiService {
	return c.ApiClient.ClueCouponEmployeeCreateV2Api
}

func (c *Client) ClueCouponEmployeeDeleteV2Api() *api.ClueCouponEmployeeDeleteV2ApiService {
	return c.ApiClient.ClueCouponEmployeeDeleteV2Api
}

func (c *Client) ClueCouponEmployeeGetV2Api() *api.ClueCouponEmployeeGetV2ApiService {
	return c.ApiClient.ClueCouponEmployeeGetV2Api
}

func (c *Client) ClueCouponGetV2Api() *api.ClueCouponGetV2ApiService {
	return c.ApiClient.ClueCouponGetV2Api
}

func (c *Client) ClueCouponUpdateV2Api() *api.ClueCouponUpdateV2ApiService {
	return c.ApiClient.ClueCouponUpdateV2Api
}

func (c *Client) ClueFormCreateV2Api() *api.ClueFormCreateV2ApiService {
	return c.ApiClient.ClueFormCreateV2Api
}

func (c *Client) ClueFormDeleteV2Api() *api.ClueFormDeleteV2ApiService {
	return c.ApiClient.ClueFormDeleteV2Api
}

func (c *Client) ClueFormDetailV2Api() *api.ClueFormDetailV2ApiService {
	return c.ApiClient.ClueFormDetailV2Api
}

func (c *Client) ClueFormListV2Api() *api.ClueFormListV2ApiService {
	return c.ApiClient.ClueFormListV2Api
}

func (c *Client) ClueFormUpdateV2Api() *api.ClueFormUpdateV2ApiService {
	return c.ApiClient.ClueFormUpdateV2Api
}

func (c *Client) ClueSmartphoneCreateV2Api() *api.ClueSmartphoneCreateV2ApiService {
	return c.ApiClient.ClueSmartphoneCreateV2Api
}

func (c *Client) ClueSmartphoneDeleteV2Api() *api.ClueSmartphoneDeleteV2ApiService {
	return c.ApiClient.ClueSmartphoneDeleteV2Api
}

func (c *Client) ClueSmartphoneGetV2Api() *api.ClueSmartphoneGetV2ApiService {
	return c.ApiClient.ClueSmartphoneGetV2Api
}

func (c *Client) ClueSmartphoneRecordV2Api() *api.ClueSmartphoneRecordV2ApiService {
	return c.ApiClient.ClueSmartphoneRecordV2Api
}

func (c *Client) ClueWechatInstanceDetailV2Api() *api.ClueWechatInstanceDetailV2ApiService {
	return c.ApiClient.ClueWechatInstanceDetailV2Api
}

func (c *Client) ClueWechatInstanceListV2Api() *api.ClueWechatInstanceListV2ApiService {
	return c.ApiClient.ClueWechatInstanceListV2Api
}

func (c *Client) ClueWechatInstanceUpdateV2Api() *api.ClueWechatInstanceUpdateV2ApiService {
	return c.ApiClient.ClueWechatInstanceUpdateV2Api
}

func (c *Client) ClueWechatPoolListV2Api() *api.ClueWechatPoolListV2ApiService {
	return c.ApiClient.ClueWechatPoolListV2Api
}

func (c *Client) CreativeStrategyListV2Api() *api.CreativeStrategyListV2ApiService {
	return c.ApiClient.CreativeStrategyListV2Api
}

func (c *Client) CustomerCenterAdvertiserTransferableListV2Api() *api.CustomerCenterAdvertiserTransferableListV2ApiService {
	return c.ApiClient.CustomerCenterAdvertiserTransferableListV2Api
}

func (c *Client) CustomerCenterFundTransferSeqCommitV2Api() *api.CustomerCenterFundTransferSeqCommitV2ApiService {
	return c.ApiClient.CustomerCenterFundTransferSeqCommitV2Api
}

func (c *Client) DecorationCouponGetV30Api() *api.DecorationCouponGetV30ApiService {
	return c.ApiClient.DecorationCouponGetV30Api
}

func (c *Client) DmpBrandGetV2Api() *api.DmpBrandGetV2ApiService {
	return c.ApiClient.DmpBrandGetV2Api
}

func (c *Client) DmpCustomAudienceCopyV2Api() *api.DmpCustomAudienceCopyV2ApiService {
	return c.ApiClient.DmpCustomAudienceCopyV2Api
}

func (c *Client) DmpCustomAudienceDeleteV2Api() *api.DmpCustomAudienceDeleteV2ApiService {
	return c.ApiClient.DmpCustomAudienceDeleteV2Api
}

func (c *Client) DmpCustomAudiencePublishV2Api() *api.DmpCustomAudiencePublishV2ApiService {
	return c.ApiClient.DmpCustomAudiencePublishV2Api
}

func (c *Client) DmpCustomAudiencePushV2V2Api() *api.DmpCustomAudiencePushV2V2ApiService {
	return c.ApiClient.DmpCustomAudiencePushV2V2Api
}

func (c *Client) DmpCustomAudienceReadV2Api() *api.DmpCustomAudienceReadV2ApiService {
	return c.ApiClient.DmpCustomAudienceReadV2Api
}

func (c *Client) DmpCustomAudienceSelectV2Api() *api.DmpCustomAudienceSelectV2ApiService {
	return c.ApiClient.DmpCustomAudienceSelectV2Api
}

func (c *Client) DmpDataSourceReadV2Api() *api.DmpDataSourceReadV2ApiService {
	return c.ApiClient.DmpDataSourceReadV2Api
}

func (c *Client) DmpDataSourceUpdateV2Api() *api.DmpDataSourceUpdateV2ApiService {
	return c.ApiClient.DmpDataSourceUpdateV2Api
}

func (c *Client) DouplusOrderListV30Api() *api.DouplusOrderListV30ApiService {
	return c.ApiClient.DouplusOrderListV30Api
}

func (c *Client) DouplusOrderReportV30Api() *api.DouplusOrderReportV30ApiService {
	return c.ApiClient.DouplusOrderReportV30Api
}

func (c *Client) DpaAssetsDetailReadV2Api() *api.DpaAssetsDetailReadV2ApiService {
	return c.ApiClient.DpaAssetsDetailReadV2Api
}

func (c *Client) DpaAssetsListV2Api() *api.DpaAssetsListV2ApiService {
	return c.ApiClient.DpaAssetsListV2Api
}

func (c *Client) DpaCategoryGetV2Api() *api.DpaCategoryGetV2ApiService {
	return c.ApiClient.DpaCategoryGetV2Api
}

func (c *Client) DpaDetailGetV2Api() *api.DpaDetailGetV2ApiService {
	return c.ApiClient.DpaDetailGetV2Api
}

func (c *Client) DpaDictGetV2Api() *api.DpaDictGetV2ApiService {
	return c.ApiClient.DpaDictGetV2Api
}

func (c *Client) DpaMetaGetV2Api() *api.DpaMetaGetV2ApiService {
	return c.ApiClient.DpaMetaGetV2Api
}

func (c *Client) DpaProductAvailablesV2Api() *api.DpaProductAvailablesV2ApiService {
	return c.ApiClient.DpaProductAvailablesV2Api
}

func (c *Client) DpaProductCreateV2Api() *api.DpaProductCreateV2ApiService {
	return c.ApiClient.DpaProductCreateV2Api
}

func (c *Client) DpaProductDeleteV2Api() *api.DpaProductDeleteV2ApiService {
	return c.ApiClient.DpaProductDeleteV2Api
}

func (c *Client) DpaProductDetailGetV2Api() *api.DpaProductDetailGetV2ApiService {
	return c.ApiClient.DpaProductDetailGetV2Api
}

func (c *Client) DpaProductStatusBatchUpdateV2Api() *api.DpaProductStatusBatchUpdateV2ApiService {
	return c.ApiClient.DpaProductStatusBatchUpdateV2Api
}

func (c *Client) DpaProductUpdateV2Api() *api.DpaProductUpdateV2ApiService {
	return c.ApiClient.DpaProductUpdateV2Api
}

func (c *Client) DpaTemplateGetV2Api() *api.DpaTemplateGetV2ApiService {
	return c.ApiClient.DpaTemplateGetV2Api
}

func (c *Client) DpaVideoGetV2Api() *api.DpaVideoGetV2ApiService {
	return c.ApiClient.DpaVideoGetV2Api
}

func (c *Client) EnterpriseBindListGetV10Api() *api.EnterpriseBindListGetV10ApiService {
	return c.ApiClient.EnterpriseBindListGetV10Api
}

func (c *Client) EnterpriseCommentDetailV10Api() *api.EnterpriseCommentDetailV10ApiService {
	return c.ApiClient.EnterpriseCommentDetailV10Api
}

func (c *Client) EnterpriseCommentReplyListV10Api() *api.EnterpriseCommentReplyListV10ApiService {
	return c.ApiClient.EnterpriseCommentReplyListV10Api
}

func (c *Client) EnterpriseCommentReplyV10Api() *api.EnterpriseCommentReplyV10ApiService {
	return c.ApiClient.EnterpriseCommentReplyV10Api
}

func (c *Client) EnterpriseItemListV10Api() *api.EnterpriseItemListV10ApiService {
	return c.ApiClient.EnterpriseItemListV10Api
}

func (c *Client) EnterpriseOperationLogGetV10Api() *api.EnterpriseOperationLogGetV10ApiService {
	return c.ApiClient.EnterpriseOperationLogGetV10Api
}

func (c *Client) EventManagerAssetsCreateV2Api() *api.EventManagerAssetsCreateV2ApiService {
	return c.ApiClient.EventManagerAssetsCreateV2Api
}

func (c *Client) EventManagerAvailableEventsGetV2Api() *api.EventManagerAvailableEventsGetV2ApiService {
	return c.ApiClient.EventManagerAvailableEventsGetV2Api
}

func (c *Client) EventManagerDeepBidTypeGetV30Api() *api.EventManagerDeepBidTypeGetV30ApiService {
	return c.ApiClient.EventManagerDeepBidTypeGetV30Api
}

func (c *Client) EventManagerEventConfigsGetV2Api() *api.EventManagerEventConfigsGetV2ApiService {
	return c.ApiClient.EventManagerEventConfigsGetV2Api
}

func (c *Client) EventManagerEventsCreateV2Api() *api.EventManagerEventsCreateV2ApiService {
	return c.ApiClient.EventManagerEventsCreateV2Api
}

func (c *Client) EventManagerOptimizedGoalGetV2V30Api() *api.EventManagerOptimizedGoalGetV2V30ApiService {
	return c.ApiClient.EventManagerOptimizedGoalGetV2V30Api
}

func (c *Client) EventManagerShareCancelV30Api() *api.EventManagerShareCancelV30ApiService {
	return c.ApiClient.EventManagerShareCancelV30Api
}

func (c *Client) EventManagerShareGetV30Api() *api.EventManagerShareGetV30ApiService {
	return c.ApiClient.EventManagerShareGetV30Api
}

func (c *Client) EventManagerShareV30Api() *api.EventManagerShareV30ApiService {
	return c.ApiClient.EventManagerShareV30Api
}

func (c *Client) EventManagerTrackUrlCreateV2Api() *api.EventManagerTrackUrlCreateV2ApiService {
	return c.ApiClient.EventManagerTrackUrlCreateV2Api
}

func (c *Client) EventManagerTrackUrlGetV2Api() *api.EventManagerTrackUrlGetV2ApiService {
	return c.ApiClient.EventManagerTrackUrlGetV2Api
}

func (c *Client) EventManagerTrackUrlUpdateV2Api() *api.EventManagerTrackUrlUpdateV2ApiService {
	return c.ApiClient.EventManagerTrackUrlUpdateV2Api
}

func (c *Client) FileImageAdGetV2Api() *api.FileImageAdGetV2ApiService {
	return c.ApiClient.FileImageAdGetV2Api
}

func (c *Client) FileImageDeleteV30Api() *api.FileImageDeleteV30ApiService {
	return c.ApiClient.FileImageDeleteV30Api
}

func (c *Client) FileImageGetV2Api() *api.FileImageGetV2ApiService {
	return c.ApiClient.FileImageGetV2Api
}

func (c *Client) FileMaterialBindV2Api() *api.FileMaterialBindV2ApiService {
	return c.ApiClient.FileMaterialBindV2Api
}

func (c *Client) FileMaterialDetailV2Api() *api.FileMaterialDetailV2ApiService {
	return c.ApiClient.FileMaterialDetailV2Api
}

func (c *Client) FileMaterialListV2Api() *api.FileMaterialListV2ApiService {
	return c.ApiClient.FileMaterialListV2Api
}

func (c *Client) FileQualityGetV30Api() *api.FileQualityGetV30ApiService {
	return c.ApiClient.FileQualityGetV30Api
}

func (c *Client) FileQualitySubmitV30Api() *api.FileQualitySubmitV30ApiService {
	return c.ApiClient.FileQualitySubmitV30Api
}

func (c *Client) FileVideoAwemeGetV2Api() *api.FileVideoAwemeGetV2ApiService {
	return c.ApiClient.FileVideoAwemeGetV2Api
}

func (c *Client) FileVideoDeleteV2Api() *api.FileVideoDeleteV2ApiService {
	return c.ApiClient.FileVideoDeleteV2Api
}

func (c *Client) FileVideoEfficiencyGetV2Api() *api.FileVideoEfficiencyGetV2ApiService {
	return c.ApiClient.FileVideoEfficiencyGetV2Api
}

func (c *Client) FileVideoGetV2Api() *api.FileVideoGetV2ApiService {
	return c.ApiClient.FileVideoGetV2Api
}

func (c *Client) FileVideoMaterialClearTaskCreateV2Api() *api.FileVideoMaterialClearTaskCreateV2ApiService {
	return c.ApiClient.FileVideoMaterialClearTaskCreateV2Api
}

func (c *Client) FileVideoMaterialClearTaskGetV2Api() *api.FileVideoMaterialClearTaskGetV2ApiService {
	return c.ApiClient.FileVideoMaterialClearTaskGetV2Api
}

func (c *Client) FileVideoMaterialClearTaskResultGetV2Api() *api.FileVideoMaterialClearTaskResultGetV2ApiService {
	return c.ApiClient.FileVideoMaterialClearTaskResultGetV2Api
}

func (c *Client) FileVideoUpdateV2Api() *api.FileVideoUpdateV2ApiService {
	return c.ApiClient.FileVideoUpdateV2Api
}

func (c *Client) FundSharedWalletBalanceGetV2Api() *api.FundSharedWalletBalanceGetV2ApiService {
	return c.ApiClient.FundSharedWalletBalanceGetV2Api
}

func (c *Client) KeywordCreateV30Api() *api.KeywordCreateV30ApiService {
	return c.ApiClient.KeywordCreateV30Api
}

func (c *Client) KeywordDeleteV30Api() *api.KeywordDeleteV30ApiService {
	return c.ApiClient.KeywordDeleteV30Api
}

func (c *Client) KeywordListV30Api() *api.KeywordListV30ApiService {
	return c.ApiClient.KeywordListV30Api
}

func (c *Client) KeywordUpdateV30Api() *api.KeywordUpdateV30ApiService {
	return c.ApiClient.KeywordUpdateV30Api
}

func (c *Client) MajordomoAdvertiserSelectV2Api() *api.MajordomoAdvertiserSelectV2ApiService {
	return c.ApiClient.MajordomoAdvertiserSelectV2Api
}

func (c *Client) MaterialStatusUpdateV30Api() *api.MaterialStatusUpdateV30ApiService {
	return c.ApiClient.MaterialStatusUpdateV30Api
}

func (c *Client) NativeAnchorCreateV30Api() *api.NativeAnchorCreateV30ApiService {
	return c.ApiClient.NativeAnchorCreateV30Api
}

func (c *Client) NativeAnchorGetV30Api() *api.NativeAnchorGetV30ApiService {
	return c.ApiClient.NativeAnchorGetV30Api
}

func (c *Client) Oauth2AccessTokenApi() *api.Oauth2AccessTokenApiService {
	return c.ApiClient.Oauth2AccessTokenApi
}

func (c *Client) Oauth2AppAccessTokenApi() *api.Oauth2AppAccessTokenApiService {
	return c.ApiClient.Oauth2AppAccessTokenApi
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

func (c *Client) PromotionScheduleTimeUpdateV30Api() *api.PromotionScheduleTimeUpdateV30ApiService {
	return c.ApiClient.PromotionScheduleTimeUpdateV30Api
}

func (c *Client) PromotionStatusUpdateV30Api() *api.PromotionStatusUpdateV30ApiService {
	return c.ApiClient.PromotionStatusUpdateV30Api
}

func (c *Client) PromotionUpdateV30Api() *api.PromotionUpdateV30ApiService {
	return c.ApiClient.PromotionUpdateV30Api
}

func (c *Client) QianchuanAdBidUpdateV10Api() *api.QianchuanAdBidUpdateV10ApiService {
	return c.ApiClient.QianchuanAdBidUpdateV10Api
}

func (c *Client) QianchuanAdBudgetUpdateV10Api() *api.QianchuanAdBudgetUpdateV10ApiService {
	return c.ApiClient.QianchuanAdBudgetUpdateV10Api
}

func (c *Client) QianchuanAdCreateV10Api() *api.QianchuanAdCreateV10ApiService {
	return c.ApiClient.QianchuanAdCreateV10Api
}

func (c *Client) QianchuanAdDetailGetV10Api() *api.QianchuanAdDetailGetV10ApiService {
	return c.ApiClient.QianchuanAdDetailGetV10Api
}

func (c *Client) QianchuanAdGetV10Api() *api.QianchuanAdGetV10ApiService {
	return c.ApiClient.QianchuanAdGetV10Api
}

func (c *Client) QianchuanAdKeywordsGetV10Api() *api.QianchuanAdKeywordsGetV10ApiService {
	return c.ApiClient.QianchuanAdKeywordsGetV10Api
}

func (c *Client) QianchuanAdKeywordsUpdateV10Api() *api.QianchuanAdKeywordsUpdateV10ApiService {
	return c.ApiClient.QianchuanAdKeywordsUpdateV10Api
}

func (c *Client) QianchuanAdPivativewordsGetV10Api() *api.QianchuanAdPivativewordsGetV10ApiService {
	return c.ApiClient.QianchuanAdPivativewordsGetV10Api
}

func (c *Client) QianchuanAdPivativewordsUpdateV10Api() *api.QianchuanAdPivativewordsUpdateV10ApiService {
	return c.ApiClient.QianchuanAdPivativewordsUpdateV10Api
}

func (c *Client) QianchuanAdQuotaGetV10Api() *api.QianchuanAdQuotaGetV10ApiService {
	return c.ApiClient.QianchuanAdQuotaGetV10Api
}

func (c *Client) QianchuanAdRecommendKeywordsGetV10Api() *api.QianchuanAdRecommendKeywordsGetV10ApiService {
	return c.ApiClient.QianchuanAdRecommendKeywordsGetV10Api
}

func (c *Client) QianchuanAdRejectReasonV10Api() *api.QianchuanAdRejectReasonV10ApiService {
	return c.ApiClient.QianchuanAdRejectReasonV10Api
}

func (c *Client) QianchuanAdStatusUpdateV10Api() *api.QianchuanAdStatusUpdateV10ApiService {
	return c.ApiClient.QianchuanAdStatusUpdateV10Api
}

func (c *Client) QianchuanAdUpdateV10Api() *api.QianchuanAdUpdateV10ApiService {
	return c.ApiClient.QianchuanAdUpdateV10Api
}

func (c *Client) QianchuanAdvertiserTypeGetV10Api() *api.QianchuanAdvertiserTypeGetV10ApiService {
	return c.ApiClient.QianchuanAdvertiserTypeGetV10Api
}

func (c *Client) QianchuanAwemeAuthorizedGetV10Api() *api.QianchuanAwemeAuthorizedGetV10ApiService {
	return c.ApiClient.QianchuanAwemeAuthorizedGetV10Api
}

func (c *Client) QianchuanAwemeOrderGetV10Api() *api.QianchuanAwemeOrderGetV10ApiService {
	return c.ApiClient.QianchuanAwemeOrderGetV10Api
}

func (c *Client) QianchuanAwemeProductAvailableGetV10Api() *api.QianchuanAwemeProductAvailableGetV10ApiService {
	return c.ApiClient.QianchuanAwemeProductAvailableGetV10Api
}

func (c *Client) QianchuanBatchCampaignStatusUpdateV10Api() *api.QianchuanBatchCampaignStatusUpdateV10ApiService {
	return c.ApiClient.QianchuanBatchCampaignStatusUpdateV10Api
}

func (c *Client) QianchuanCampaignCreateV10Api() *api.QianchuanCampaignCreateV10ApiService {
	return c.ApiClient.QianchuanCampaignCreateV10Api
}

func (c *Client) QianchuanCampaignListGetV10Api() *api.QianchuanCampaignListGetV10ApiService {
	return c.ApiClient.QianchuanCampaignListGetV10Api
}

func (c *Client) QianchuanCampaignUpdateV10Api() *api.QianchuanCampaignUpdateV10ApiService {
	return c.ApiClient.QianchuanCampaignUpdateV10Api
}

func (c *Client) QianchuanCreativeGetV10Api() *api.QianchuanCreativeGetV10ApiService {
	return c.ApiClient.QianchuanCreativeGetV10Api
}

func (c *Client) QianchuanCreativeRejectReasonV10Api() *api.QianchuanCreativeRejectReasonV10ApiService {
	return c.ApiClient.QianchuanCreativeRejectReasonV10Api
}

func (c *Client) QianchuanCreativeStatusUpdateV10Api() *api.QianchuanCreativeStatusUpdateV10ApiService {
	return c.ApiClient.QianchuanCreativeStatusUpdateV10Api
}

func (c *Client) QianchuanEstimateEffectV10Api() *api.QianchuanEstimateEffectV10ApiService {
	return c.ApiClient.QianchuanEstimateEffectV10Api
}

func (c *Client) QianchuanFileImageDeleteV10Api() *api.QianchuanFileImageDeleteV10ApiService {
	return c.ApiClient.QianchuanFileImageDeleteV10Api
}

func (c *Client) QianchuanFileVideoAwemeGetV10Api() *api.QianchuanFileVideoAwemeGetV10ApiService {
	return c.ApiClient.QianchuanFileVideoAwemeGetV10Api
}

func (c *Client) QianchuanFileVideoDeleteV10Api() *api.QianchuanFileVideoDeleteV10ApiService {
	return c.ApiClient.QianchuanFileVideoDeleteV10Api
}

func (c *Client) QianchuanFileVideoEfficiencyGetV10Api() *api.QianchuanFileVideoEfficiencyGetV10ApiService {
	return c.ApiClient.QianchuanFileVideoEfficiencyGetV10Api
}

func (c *Client) QianchuanFileVideoOriginalGetV10Api() *api.QianchuanFileVideoOriginalGetV10ApiService {
	return c.ApiClient.QianchuanFileVideoOriginalGetV10Api
}

func (c *Client) QianchuanFinanceDetailGetV10Api() *api.QianchuanFinanceDetailGetV10ApiService {
	return c.ApiClient.QianchuanFinanceDetailGetV10Api
}

func (c *Client) QianchuanFinanceWalletGetV10Api() *api.QianchuanFinanceWalletGetV10ApiService {
	return c.ApiClient.QianchuanFinanceWalletGetV10Api
}

func (c *Client) QianchuanImageGetV10Api() *api.QianchuanImageGetV10ApiService {
	return c.ApiClient.QianchuanImageGetV10Api
}

func (c *Client) QianchuanLqAdGetV10Api() *api.QianchuanLqAdGetV10ApiService {
	return c.ApiClient.QianchuanLqAdGetV10Api
}

func (c *Client) QianchuanOrientationPackageGetV10Api() *api.QianchuanOrientationPackageGetV10ApiService {
	return c.ApiClient.QianchuanOrientationPackageGetV10Api
}

func (c *Client) QianchuanProductAnalyseCompareStatsDataV10Api() *api.QianchuanProductAnalyseCompareStatsDataV10ApiService {
	return c.ApiClient.QianchuanProductAnalyseCompareStatsDataV10Api
}

func (c *Client) QianchuanProductAvailableGetV10Api() *api.QianchuanProductAvailableGetV10ApiService {
	return c.ApiClient.QianchuanProductAvailableGetV10Api
}

func (c *Client) QianchuanReportAdGetV10Api() *api.QianchuanReportAdGetV10ApiService {
	return c.ApiClient.QianchuanReportAdGetV10Api
}

func (c *Client) QianchuanReportAdvertiserGetV10Api() *api.QianchuanReportAdvertiserGetV10ApiService {
	return c.ApiClient.QianchuanReportAdvertiserGetV10Api
}

func (c *Client) QianchuanReportCreativeGetV10Api() *api.QianchuanReportCreativeGetV10ApiService {
	return c.ApiClient.QianchuanReportCreativeGetV10Api
}

func (c *Client) QianchuanReportLongTransferOrderGetV10Api() *api.QianchuanReportLongTransferOrderGetV10ApiService {
	return c.ApiClient.QianchuanReportLongTransferOrderGetV10Api
}

func (c *Client) QianchuanReportSearchWordGetV10Api() *api.QianchuanReportSearchWordGetV10ApiService {
	return c.ApiClient.QianchuanReportSearchWordGetV10Api
}

func (c *Client) QianchuanReportUniPromotionGetV10Api() *api.QianchuanReportUniPromotionGetV10ApiService {
	return c.ApiClient.QianchuanReportUniPromotionGetV10Api
}

func (c *Client) QianchuanReportVideoUserLoseGetV10Api() *api.QianchuanReportVideoUserLoseGetV10ApiService {
	return c.ApiClient.QianchuanReportVideoUserLoseGetV10Api
}

func (c *Client) QianchuanRoiGoalUpdateV10Api() *api.QianchuanRoiGoalUpdateV10ApiService {
	return c.ApiClient.QianchuanRoiGoalUpdateV10Api
}

func (c *Client) QianchuanShopAdvertiserListV10Api() *api.QianchuanShopAdvertiserListV10ApiService {
	return c.ApiClient.QianchuanShopAdvertiserListV10Api
}

func (c *Client) QianchuanShopGetV10Api() *api.QianchuanShopGetV10ApiService {
	return c.ApiClient.QianchuanShopGetV10Api
}

func (c *Client) QianchuanSuggestBidV10Api() *api.QianchuanSuggestBidV10ApiService {
	return c.ApiClient.QianchuanSuggestBidV10Api
}

func (c *Client) QianchuanSuggestBudgetV10Api() *api.QianchuanSuggestBudgetV10ApiService {
	return c.ApiClient.QianchuanSuggestBudgetV10Api
}

func (c *Client) QianchuanSuggestRoiGoalV10Api() *api.QianchuanSuggestRoiGoalV10ApiService {
	return c.ApiClient.QianchuanSuggestRoiGoalV10Api
}

func (c *Client) QianchuanToolsAllowCouponV10Api() *api.QianchuanToolsAllowCouponV10ApiService {
	return c.ApiClient.QianchuanToolsAllowCouponV10Api
}

func (c *Client) QianchuanToolsAwemeAuthV10Api() *api.QianchuanToolsAwemeAuthV10ApiService {
	return c.ApiClient.QianchuanToolsAwemeAuthV10Api
}

func (c *Client) QianchuanToolsEstimateAudienceV10Api() *api.QianchuanToolsEstimateAudienceV10ApiService {
	return c.ApiClient.QianchuanToolsEstimateAudienceV10Api
}

func (c *Client) QianchuanToolsGrayV10Api() *api.QianchuanToolsGrayV10ApiService {
	return c.ApiClient.QianchuanToolsGrayV10Api
}

func (c *Client) QianchuanUniAwemeAuthorizedGetV10Api() *api.QianchuanUniAwemeAuthorizedGetV10ApiService {
	return c.ApiClient.QianchuanUniAwemeAuthorizedGetV10Api
}

func (c *Client) QianchuanUniPromotionListV10Api() *api.QianchuanUniPromotionListV10ApiService {
	return c.ApiClient.QianchuanUniPromotionListV10Api
}

func (c *Client) QianchuanVideoGetV10Api() *api.QianchuanVideoGetV10ApiService {
	return c.ApiClient.QianchuanVideoGetV10Api
}

func (c *Client) ReportAgentGetV2V2Api() *api.ReportAgentGetV2V2ApiService {
	return c.ApiClient.ReportAgentGetV2V2Api
}

func (c *Client) ReportAudienceAgeV2Api() *api.ReportAudienceAgeV2ApiService {
	return c.ApiClient.ReportAudienceAgeV2Api
}

func (c *Client) ReportAudienceCityV2Api() *api.ReportAudienceCityV2ApiService {
	return c.ApiClient.ReportAudienceCityV2Api
}

func (c *Client) ReportAudienceGenderV2Api() *api.ReportAudienceGenderV2ApiService {
	return c.ApiClient.ReportAudienceGenderV2Api
}

func (c *Client) ReportAudienceProvinceV2Api() *api.ReportAudienceProvinceV2ApiService {
	return c.ApiClient.ReportAudienceProvinceV2Api
}

func (c *Client) ReportCustomConfigGetV30Api() *api.ReportCustomConfigGetV30ApiService {
	return c.ApiClient.ReportCustomConfigGetV30Api
}

func (c *Client) ReportCustomGetV30Api() *api.ReportCustomGetV30ApiService {
	return c.ApiClient.ReportCustomGetV30Api
}

func (c *Client) ReportLiveRoomAttributeGetV2Api() *api.ReportLiveRoomAttributeGetV2ApiService {
	return c.ApiClient.ReportLiveRoomAttributeGetV2Api
}

func (c *Client) ReportLiveRoomAudiencePortraitGetV2Api() *api.ReportLiveRoomAudiencePortraitGetV2ApiService {
	return c.ApiClient.ReportLiveRoomAudiencePortraitGetV2Api
}

func (c *Client) ReportRtaExpGetV2Api() *api.ReportRtaExpGetV2ApiService {
	return c.ApiClient.ReportRtaExpGetV2Api
}

func (c *Client) ReportRtaExpLocalDailyGetV30Api() *api.ReportRtaExpLocalDailyGetV30ApiService {
	return c.ApiClient.ReportRtaExpLocalDailyGetV30Api
}

func (c *Client) ReportRtaExpLocalHourlyGetV30Api() *api.ReportRtaExpLocalHourlyGetV30ApiService {
	return c.ApiClient.ReportRtaExpLocalHourlyGetV30Api
}

func (c *Client) ReportSitePageV2Api() *api.ReportSitePageV2ApiService {
	return c.ApiClient.ReportSitePageV2Api
}

func (c *Client) ReportVideoFrameGetV2Api() *api.ReportVideoFrameGetV2ApiService {
	return c.ApiClient.ReportVideoFrameGetV2Api
}

func (c *Client) SpiTaskGetV2Api() *api.SpiTaskGetV2ApiService {
	return c.ApiClient.SpiTaskGetV2Api
}

func (c *Client) StarClueGetV2Api() *api.StarClueGetV2ApiService {
	return c.ApiClient.StarClueGetV2Api
}

func (c *Client) StarDemandListV2Api() *api.StarDemandListV2ApiService {
	return c.ApiClient.StarDemandListV2Api
}

func (c *Client) StarDemandOrderListV2Api() *api.StarDemandOrderListV2ApiService {
	return c.ApiClient.StarDemandOrderListV2Api
}

func (c *Client) StarReportOrderOverviewGetV2Api() *api.StarReportOrderOverviewGetV2ApiService {
	return c.ApiClient.StarReportOrderOverviewGetV2Api
}

func (c *Client) StarReportOrderUserDistributionGetV2Api() *api.StarReportOrderUserDistributionGetV2ApiService {
	return c.ApiClient.StarReportOrderUserDistributionGetV2Api
}

func (c *Client) SubscribeAccountsAddV30Api() *api.SubscribeAccountsAddV30ApiService {
	return c.ApiClient.SubscribeAccountsAddV30Api
}

func (c *Client) SubscribeAccountsListV30Api() *api.SubscribeAccountsListV30ApiService {
	return c.ApiClient.SubscribeAccountsListV30Api
}

func (c *Client) SubscribeAccountsRemoveV30Api() *api.SubscribeAccountsRemoveV30ApiService {
	return c.ApiClient.SubscribeAccountsRemoveV30Api
}

func (c *Client) SuggWordsV30Api() *api.SuggWordsV30ApiService {
	return c.ApiClient.SuggWordsV30Api
}

func (c *Client) ToolQuickAppManagementQuickAppGetV2Api() *api.ToolQuickAppManagementQuickAppGetV2ApiService {
	return c.ApiClient.ToolQuickAppManagementQuickAppGetV2Api
}

func (c *Client) ToolsAbTestCreateV2Api() *api.ToolsAbTestCreateV2ApiService {
	return c.ApiClient.ToolsAbTestCreateV2Api
}

func (c *Client) ToolsAbTestDeleteV2Api() *api.ToolsAbTestDeleteV2ApiService {
	return c.ApiClient.ToolsAbTestDeleteV2Api
}

func (c *Client) ToolsAbTestInfoGetV2Api() *api.ToolsAbTestInfoGetV2ApiService {
	return c.ApiClient.ToolsAbTestInfoGetV2Api
}

func (c *Client) ToolsAbTestListGetV2Api() *api.ToolsAbTestListGetV2ApiService {
	return c.ApiClient.ToolsAbTestListGetV2Api
}

func (c *Client) ToolsAbTestStopV2Api() *api.ToolsAbTestStopV2ApiService {
	return c.ApiClient.ToolsAbTestStopV2Api
}

func (c *Client) ToolsAbTestUpdateV2Api() *api.ToolsAbTestUpdateV2ApiService {
	return c.ApiClient.ToolsAbTestUpdateV2Api
}

func (c *Client) ToolsActionTextGetV2Api() *api.ToolsActionTextGetV2ApiService {
	return c.ApiClient.ToolsActionTextGetV2Api
}

func (c *Client) ToolsAdPreviewQrcodeGetV30Api() *api.ToolsAdPreviewQrcodeGetV30ApiService {
	return c.ApiClient.ToolsAdPreviewQrcodeGetV30Api
}

func (c *Client) ToolsAdRaiseEstimateGetV2Api() *api.ToolsAdRaiseEstimateGetV2ApiService {
	return c.ApiClient.ToolsAdRaiseEstimateGetV2Api
}

func (c *Client) ToolsAdRaiseStatusGetV2Api() *api.ToolsAdRaiseStatusGetV2ApiService {
	return c.ApiClient.ToolsAdRaiseStatusGetV2Api
}

func (c *Client) ToolsAdRaiseVersionGetV2Api() *api.ToolsAdRaiseVersionGetV2ApiService {
	return c.ApiClient.ToolsAdRaiseVersionGetV2Api
}

func (c *Client) ToolsAdStatExtraInfoGetV2Api() *api.ToolsAdStatExtraInfoGetV2ApiService {
	return c.ApiClient.ToolsAdStatExtraInfoGetV2Api
}

func (c *Client) ToolsAdminInfoV2Api() *api.ToolsAdminInfoV2ApiService {
	return c.ApiClient.ToolsAdminInfoV2Api
}

func (c *Client) ToolsAppManagementAndroidAppListV2Api() *api.ToolsAppManagementAndroidAppListV2ApiService {
	return c.ApiClient.ToolsAppManagementAndroidAppListV2Api
}

func (c *Client) ToolsAppManagementAndroidBasicPackageGetV2Api() *api.ToolsAppManagementAndroidBasicPackageGetV2ApiService {
	return c.ApiClient.ToolsAppManagementAndroidBasicPackageGetV2Api
}

func (c *Client) ToolsAppManagementAndroidBasicPackagePublishV2Api() *api.ToolsAppManagementAndroidBasicPackagePublishV2ApiService {
	return c.ApiClient.ToolsAppManagementAndroidBasicPackagePublishV2Api
}

func (c *Client) ToolsAppManagementAndroidBasicPackageUpdateV2Api() *api.ToolsAppManagementAndroidBasicPackageUpdateV2ApiService {
	return c.ApiClient.ToolsAppManagementAndroidBasicPackageUpdateV2Api
}

func (c *Client) ToolsAppManagementAppGetV2Api() *api.ToolsAppManagementAppGetV2ApiService {
	return c.ApiClient.ToolsAppManagementAppGetV2Api
}

func (c *Client) ToolsAppManagementBookingGetV2Api() *api.ToolsAppManagementBookingGetV2ApiService {
	return c.ApiClient.ToolsAppManagementBookingGetV2Api
}

func (c *Client) ToolsAppManagementBookingRecordsGetV2Api() *api.ToolsAppManagementBookingRecordsGetV2ApiService {
	return c.ApiClient.ToolsAppManagementBookingRecordsGetV2Api
}

func (c *Client) ToolsAppManagementBpShareCancelV2Api() *api.ToolsAppManagementBpShareCancelV2ApiService {
	return c.ApiClient.ToolsAppManagementBpShareCancelV2Api
}

func (c *Client) ToolsAppManagementBpShareV2Api() *api.ToolsAppManagementBpShareV2ApiService {
	return c.ApiClient.ToolsAppManagementBpShareV2Api
}

func (c *Client) ToolsAppManagementExtendPackageCreateV2Api() *api.ToolsAppManagementExtendPackageCreateV2ApiService {
	return c.ApiClient.ToolsAppManagementExtendPackageCreateV2Api
}

func (c *Client) ToolsAppManagementExtendPackageCreateV2V2Api() *api.ToolsAppManagementExtendPackageCreateV2V2ApiService {
	return c.ApiClient.ToolsAppManagementExtendPackageCreateV2V2Api
}

func (c *Client) ToolsAppManagementExtendPackageListV2V2Api() *api.ToolsAppManagementExtendPackageListV2V2ApiService {
	return c.ApiClient.ToolsAppManagementExtendPackageListV2V2Api
}

func (c *Client) ToolsAppManagementExtendPackageUpdateV2Api() *api.ToolsAppManagementExtendPackageUpdateV2ApiService {
	return c.ApiClient.ToolsAppManagementExtendPackageUpdateV2Api
}

func (c *Client) ToolsAppManagementIndustryInfoListV2Api() *api.ToolsAppManagementIndustryInfoListV2ApiService {
	return c.ApiClient.ToolsAppManagementIndustryInfoListV2Api
}

func (c *Client) ToolsAppManagementUpdateAuthorizationV2Api() *api.ToolsAppManagementUpdateAuthorizationV2ApiService {
	return c.ApiClient.ToolsAppManagementUpdateAuthorizationV2Api
}

func (c *Client) ToolsAppManagementUploadTaskCreateV2Api() *api.ToolsAppManagementUploadTaskCreateV2ApiService {
	return c.ApiClient.ToolsAppManagementUploadTaskCreateV2Api
}

func (c *Client) ToolsAppManagementUploadTaskListV2Api() *api.ToolsAppManagementUploadTaskListV2ApiService {
	return c.ApiClient.ToolsAppManagementUploadTaskListV2Api
}

func (c *Client) ToolsAwemeAuthListV2Api() *api.ToolsAwemeAuthListV2ApiService {
	return c.ApiClient.ToolsAwemeAuthListV2Api
}

func (c *Client) ToolsAwemeAuthorInfoGetV2Api() *api.ToolsAwemeAuthorInfoGetV2ApiService {
	return c.ApiClient.ToolsAwemeAuthorInfoGetV2Api
}

func (c *Client) ToolsAwemeBannedCreateV30Api() *api.ToolsAwemeBannedCreateV30ApiService {
	return c.ApiClient.ToolsAwemeBannedCreateV30Api
}

func (c *Client) ToolsAwemeBannedDeleteV30Api() *api.ToolsAwemeBannedDeleteV30ApiService {
	return c.ApiClient.ToolsAwemeBannedDeleteV30Api
}

func (c *Client) ToolsAwemeBannedListV30Api() *api.ToolsAwemeBannedListV30ApiService {
	return c.ApiClient.ToolsAwemeBannedListV30Api
}

func (c *Client) ToolsAwemeCategoryTopAuthorGetV2Api() *api.ToolsAwemeCategoryTopAuthorGetV2ApiService {
	return c.ApiClient.ToolsAwemeCategoryTopAuthorGetV2Api
}

func (c *Client) ToolsAwemeInfoSearchV2Api() *api.ToolsAwemeInfoSearchV2ApiService {
	return c.ApiClient.ToolsAwemeInfoSearchV2Api
}

func (c *Client) ToolsAwemeMultiLevelCategoryGetV2Api() *api.ToolsAwemeMultiLevelCategoryGetV2ApiService {
	return c.ApiClient.ToolsAwemeMultiLevelCategoryGetV2Api
}

func (c *Client) ToolsAwemeSimilarAuthorSearchV2Api() *api.ToolsAwemeSimilarAuthorSearchV2ApiService {
	return c.ApiClient.ToolsAwemeSimilarAuthorSearchV2Api
}

func (c *Client) ToolsBidSuggestV2Api() *api.ToolsBidSuggestV2ApiService {
	return c.ApiClient.ToolsBidSuggestV2Api
}

func (c *Client) ToolsBidsSuggestV30Api() *api.ToolsBidsSuggestV30ApiService {
	return c.ApiClient.ToolsBidsSuggestV30Api
}

func (c *Client) ToolsClueCallbackV2Api() *api.ToolsClueCallbackV2ApiService {
	return c.ApiClient.ToolsClueCallbackV2Api
}

func (c *Client) ToolsClueFormDetailV2Api() *api.ToolsClueFormDetailV2ApiService {
	return c.ApiClient.ToolsClueFormDetailV2Api
}

func (c *Client) ToolsClueFormGetV2Api() *api.ToolsClueFormGetV2ApiService {
	return c.ApiClient.ToolsClueFormGetV2Api
}

func (c *Client) ToolsClueGetV2Api() *api.ToolsClueGetV2ApiService {
	return c.ApiClient.ToolsClueGetV2Api
}

func (c *Client) ToolsClueSmartPhoneGetV2Api() *api.ToolsClueSmartPhoneGetV2ApiService {
	return c.ApiClient.ToolsClueSmartPhoneGetV2Api
}

func (c *Client) ToolsCommentGetV30Api() *api.ToolsCommentGetV30ApiService {
	return c.ApiClient.ToolsCommentGetV30Api
}

func (c *Client) ToolsCommentHideV30Api() *api.ToolsCommentHideV30ApiService {
	return c.ApiClient.ToolsCommentHideV30Api
}

func (c *Client) ToolsCommentMid2itemIdV30Api() *api.ToolsCommentMid2itemIdV30ApiService {
	return c.ApiClient.ToolsCommentMid2itemIdV30Api
}

func (c *Client) ToolsCommentReplyGetV30Api() *api.ToolsCommentReplyGetV30ApiService {
	return c.ApiClient.ToolsCommentReplyGetV30Api
}

func (c *Client) ToolsCommentReplyV30Api() *api.ToolsCommentReplyV30ApiService {
	return c.ApiClient.ToolsCommentReplyV30Api
}

func (c *Client) ToolsCommentStickOnTopV30Api() *api.ToolsCommentStickOnTopV30ApiService {
	return c.ApiClient.ToolsCommentStickOnTopV30Api
}

func (c *Client) ToolsCommentTermsBannedAddV30Api() *api.ToolsCommentTermsBannedAddV30ApiService {
	return c.ApiClient.ToolsCommentTermsBannedAddV30Api
}

func (c *Client) ToolsCommentTermsBannedDeleteV30Api() *api.ToolsCommentTermsBannedDeleteV30ApiService {
	return c.ApiClient.ToolsCommentTermsBannedDeleteV30Api
}

func (c *Client) ToolsCommentTermsBannedGetV30Api() *api.ToolsCommentTermsBannedGetV30ApiService {
	return c.ApiClient.ToolsCommentTermsBannedGetV30Api
}

func (c *Client) ToolsCommentTermsBannedUpdateV30Api() *api.ToolsCommentTermsBannedUpdateV30ApiService {
	return c.ApiClient.ToolsCommentTermsBannedUpdateV30Api
}

func (c *Client) ToolsCountryInfoV2Api() *api.ToolsCountryInfoV2ApiService {
	return c.ApiClient.ToolsCountryInfoV2Api
}

func (c *Client) ToolsCreativeWordSelectV2Api() *api.ToolsCreativeWordSelectV2ApiService {
	return c.ApiClient.ToolsCreativeWordSelectV2Api
}

func (c *Client) ToolsDiagnosisAdGetV2V2Api() *api.ToolsDiagnosisAdGetV2V2ApiService {
	return c.ApiClient.ToolsDiagnosisAdGetV2V2Api
}

func (c *Client) ToolsDiagnosisSuggestionGetV2Api() *api.ToolsDiagnosisSuggestionGetV2ApiService {
	return c.ApiClient.ToolsDiagnosisSuggestionGetV2Api
}

func (c *Client) ToolsDownloadPackageGetV2Api() *api.ToolsDownloadPackageGetV2ApiService {
	return c.ApiClient.ToolsDownloadPackageGetV2Api
}

func (c *Client) ToolsDownloadPackageParseV2Api() *api.ToolsDownloadPackageParseV2ApiService {
	return c.ApiClient.ToolsDownloadPackageParseV2Api
}

func (c *Client) ToolsEstimateAudienceV2Api() *api.ToolsEstimateAudienceV2ApiService {
	return c.ApiClient.ToolsEstimateAudienceV2Api
}

func (c *Client) ToolsEstimatedPriceGetV2Api() *api.ToolsEstimatedPriceGetV2ApiService {
	return c.ApiClient.ToolsEstimatedPriceGetV2Api
}

func (c *Client) ToolsEventAssetsGetV2Api() *api.ToolsEventAssetsGetV2ApiService {
	return c.ApiClient.ToolsEventAssetsGetV2Api
}

func (c *Client) ToolsGrayGetV30Api() *api.ToolsGrayGetV30ApiService {
	return c.ApiClient.ToolsGrayGetV30Api
}

func (c *Client) ToolsIesAccountSearchV2Api() *api.ToolsIesAccountSearchV2ApiService {
	return c.ApiClient.ToolsIesAccountSearchV2Api
}

func (c *Client) ToolsInterestActionActionCategoryV2Api() *api.ToolsInterestActionActionCategoryV2ApiService {
	return c.ApiClient.ToolsInterestActionActionCategoryV2Api
}

func (c *Client) ToolsInterestActionActionKeywordV2Api() *api.ToolsInterestActionActionKeywordV2ApiService {
	return c.ApiClient.ToolsInterestActionActionKeywordV2Api
}

func (c *Client) ToolsInterestActionId2wordV2Api() *api.ToolsInterestActionId2wordV2ApiService {
	return c.ApiClient.ToolsInterestActionId2wordV2Api
}

func (c *Client) ToolsInterestActionInterestCategoryV2Api() *api.ToolsInterestActionInterestCategoryV2ApiService {
	return c.ApiClient.ToolsInterestActionInterestCategoryV2Api
}

func (c *Client) ToolsInterestActionInterestKeywordV2Api() *api.ToolsInterestActionInterestKeywordV2ApiService {
	return c.ApiClient.ToolsInterestActionInterestKeywordV2Api
}

func (c *Client) ToolsInterestActionKeywordSuggestV2Api() *api.ToolsInterestActionKeywordSuggestV2ApiService {
	return c.ApiClient.ToolsInterestActionKeywordSuggestV2Api
}

func (c *Client) ToolsKeyActionGetV2Api() *api.ToolsKeyActionGetV2ApiService {
	return c.ApiClient.ToolsKeyActionGetV2Api
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

func (c *Client) ToolsLandingGroupCreateV2Api() *api.ToolsLandingGroupCreateV2ApiService {
	return c.ApiClient.ToolsLandingGroupCreateV2Api
}

func (c *Client) ToolsLandingGroupGetV2Api() *api.ToolsLandingGroupGetV2ApiService {
	return c.ApiClient.ToolsLandingGroupGetV2Api
}

func (c *Client) ToolsLandingGroupSiteOptStatusUpdateV2Api() *api.ToolsLandingGroupSiteOptStatusUpdateV2ApiService {
	return c.ApiClient.ToolsLandingGroupSiteOptStatusUpdateV2Api
}

func (c *Client) ToolsLandingGroupUpdateV2Api() *api.ToolsLandingGroupUpdateV2ApiService {
	return c.ApiClient.ToolsLandingGroupUpdateV2Api
}

func (c *Client) ToolsLiveAuthorizeListV2Api() *api.ToolsLiveAuthorizeListV2ApiService {
	return c.ApiClient.ToolsLiveAuthorizeListV2Api
}

func (c *Client) ToolsLogSearchV2Api() *api.ToolsLogSearchV2ApiService {
	return c.ApiClient.ToolsLogSearchV2Api
}

func (c *Client) ToolsOrangeSiteGetV30Api() *api.ToolsOrangeSiteGetV30ApiService {
	return c.ApiClient.ToolsOrangeSiteGetV30Api
}

func (c *Client) ToolsPlayableCloudGameListV2Api() *api.ToolsPlayableCloudGameListV2ApiService {
	return c.ApiClient.ToolsPlayableCloudGameListV2Api
}

func (c *Client) ToolsPlayableListGetV2Api() *api.ToolsPlayableListGetV2ApiService {
	return c.ApiClient.ToolsPlayableListGetV2Api
}

func (c *Client) ToolsPreAuditGetV2Api() *api.ToolsPreAuditGetV2ApiService {
	return c.ApiClient.ToolsPreAuditGetV2Api
}

func (c *Client) ToolsPreAuditSendV2Api() *api.ToolsPreAuditSendV2ApiService {
	return c.ApiClient.ToolsPreAuditSendV2Api
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

func (c *Client) ToolsPromotionCardRecommendGetV2Api() *api.ToolsPromotionCardRecommendGetV2ApiService {
	return c.ApiClient.ToolsPromotionCardRecommendGetV2Api
}

func (c *Client) ToolsPromotionDiagnosisSuggestionAcceptV30Api() *api.ToolsPromotionDiagnosisSuggestionAcceptV30ApiService {
	return c.ApiClient.ToolsPromotionDiagnosisSuggestionAcceptV30Api
}

func (c *Client) ToolsPromotionDiagnosisSuggestionGetV30Api() *api.ToolsPromotionDiagnosisSuggestionGetV30ApiService {
	return c.ApiClient.ToolsPromotionDiagnosisSuggestionGetV30Api
}

func (c *Client) ToolsPromotionRaiseSetV30Api() *api.ToolsPromotionRaiseSetV30ApiService {
	return c.ApiClient.ToolsPromotionRaiseSetV30Api
}

func (c *Client) ToolsPromotionRaiseStatusCurrentIdsGetV30Api() *api.ToolsPromotionRaiseStatusCurrentIdsGetV30ApiService {
	return c.ApiClient.ToolsPromotionRaiseStatusCurrentIdsGetV30Api
}

func (c *Client) ToolsPromotionRaiseStatusGetV30Api() *api.ToolsPromotionRaiseStatusGetV30ApiService {
	return c.ApiClient.ToolsPromotionRaiseStatusGetV30Api
}

func (c *Client) ToolsPromotionRaiseStopV30Api() *api.ToolsPromotionRaiseStopV30ApiService {
	return c.ApiClient.ToolsPromotionRaiseStopV30Api
}

func (c *Client) ToolsPromotionRaiseVersionGetV30Api() *api.ToolsPromotionRaiseVersionGetV30ApiService {
	return c.ApiClient.ToolsPromotionRaiseVersionGetV30Api
}

func (c *Client) ToolsQuotaGetV2Api() *api.ToolsQuotaGetV2ApiService {
	return c.ApiClient.ToolsQuotaGetV2Api
}

func (c *Client) ToolsRegionGetV2Api() *api.ToolsRegionGetV2ApiService {
	return c.ApiClient.ToolsRegionGetV2Api
}

func (c *Client) ToolsRtaGetInfoV2Api() *api.ToolsRtaGetInfoV2ApiService {
	return c.ApiClient.ToolsRtaGetInfoV2Api
}

func (c *Client) ToolsRtaGetV2Api() *api.ToolsRtaGetV2ApiService {
	return c.ApiClient.ToolsRtaGetV2Api
}

func (c *Client) ToolsRtaScopeGetV30Api() *api.ToolsRtaScopeGetV30ApiService {
	return c.ApiClient.ToolsRtaScopeGetV30Api
}

func (c *Client) ToolsRtaSetScopeV2Api() *api.ToolsRtaSetScopeV2ApiService {
	return c.ApiClient.ToolsRtaSetScopeV2Api
}

func (c *Client) ToolsRtaStatusUpdateV2Api() *api.ToolsRtaStatusUpdateV2ApiService {
	return c.ApiClient.ToolsRtaStatusUpdateV2Api
}

func (c *Client) ToolsRubeexGetV2Api() *api.ToolsRubeexGetV2ApiService {
	return c.ApiClient.ToolsRubeexGetV2Api
}

func (c *Client) ToolsRubeexPlayableListV2Api() *api.ToolsRubeexPlayableListV2ApiService {
	return c.ApiClient.ToolsRubeexPlayableListV2Api
}

func (c *Client) ToolsRubeexRemarkV2Api() *api.ToolsRubeexRemarkV2ApiService {
	return c.ApiClient.ToolsRubeexRemarkV2Api
}

func (c *Client) ToolsRubeexVersionGetV2Api() *api.ToolsRubeexVersionGetV2ApiService {
	return c.ApiClient.ToolsRubeexVersionGetV2Api
}

func (c *Client) ToolsSearchBidRatioGetV2Api() *api.ToolsSearchBidRatioGetV2ApiService {
	return c.ApiClient.ToolsSearchBidRatioGetV2Api
}

func (c *Client) ToolsSiteCopyV2Api() *api.ToolsSiteCopyV2ApiService {
	return c.ApiClient.ToolsSiteCopyV2Api
}

func (c *Client) ToolsSiteCreateV2Api() *api.ToolsSiteCreateV2ApiService {
	return c.ApiClient.ToolsSiteCreateV2Api
}

func (c *Client) ToolsSiteFormsListV2Api() *api.ToolsSiteFormsListV2ApiService {
	return c.ApiClient.ToolsSiteFormsListV2Api
}

func (c *Client) ToolsSiteHandselV2Api() *api.ToolsSiteHandselV2ApiService {
	return c.ApiClient.ToolsSiteHandselV2Api
}

func (c *Client) ToolsSiteReadV2Api() *api.ToolsSiteReadV2ApiService {
	return c.ApiClient.ToolsSiteReadV2Api
}

func (c *Client) ToolsSiteTemplateCreateV2Api() *api.ToolsSiteTemplateCreateV2ApiService {
	return c.ApiClient.ToolsSiteTemplateCreateV2Api
}

func (c *Client) ToolsSiteTemplateGetV2Api() *api.ToolsSiteTemplateGetV2ApiService {
	return c.ApiClient.ToolsSiteTemplateGetV2Api
}

func (c *Client) ToolsSiteTemplatePicUrlGetV2Api() *api.ToolsSiteTemplatePicUrlGetV2ApiService {
	return c.ApiClient.ToolsSiteTemplatePicUrlGetV2Api
}

func (c *Client) ToolsSiteTemplatePreviewV2Api() *api.ToolsSiteTemplatePreviewV2ApiService {
	return c.ApiClient.ToolsSiteTemplatePreviewV2Api
}

func (c *Client) ToolsSiteTemplateSiteCreateV2Api() *api.ToolsSiteTemplateSiteCreateV2ApiService {
	return c.ApiClient.ToolsSiteTemplateSiteCreateV2Api
}

func (c *Client) ToolsSiteUpdateStatusV2Api() *api.ToolsSiteUpdateStatusV2ApiService {
	return c.ApiClient.ToolsSiteUpdateStatusV2Api
}

func (c *Client) ToolsSiteUpdateV2Api() *api.ToolsSiteUpdateV2ApiService {
	return c.ApiClient.ToolsSiteUpdateV2Api
}

func (c *Client) ToolsSuggestBudgetGetV30Api() *api.ToolsSuggestBudgetGetV30ApiService {
	return c.ApiClient.ToolsSuggestBudgetGetV30Api
}

func (c *Client) ToolsTaskRaiseCreateV2Api() *api.ToolsTaskRaiseCreateV2ApiService {
	return c.ApiClient.ToolsTaskRaiseCreateV2Api
}

func (c *Client) ToolsTaskRaiseDataGetV2Api() *api.ToolsTaskRaiseDataGetV2ApiService {
	return c.ApiClient.ToolsTaskRaiseDataGetV2Api
}

func (c *Client) ToolsTaskRaiseGetV2Api() *api.ToolsTaskRaiseGetV2ApiService {
	return c.ApiClient.ToolsTaskRaiseGetV2Api
}

func (c *Client) ToolsTaskRaiseOptimizationIdsGetV2Api() *api.ToolsTaskRaiseOptimizationIdsGetV2ApiService {
	return c.ApiClient.ToolsTaskRaiseOptimizationIdsGetV2Api
}

func (c *Client) ToolsTaskRaiseStatusStopV2Api() *api.ToolsTaskRaiseStatusStopV2ApiService {
	return c.ApiClient.ToolsTaskRaiseStatusStopV2Api
}

func (c *Client) ToolsThirdSiteCreateV2Api() *api.ToolsThirdSiteCreateV2ApiService {
	return c.ApiClient.ToolsThirdSiteCreateV2Api
}

func (c *Client) ToolsThirdSiteDeleteV2Api() *api.ToolsThirdSiteDeleteV2ApiService {
	return c.ApiClient.ToolsThirdSiteDeleteV2Api
}

func (c *Client) ToolsThirdSiteGetV2Api() *api.ToolsThirdSiteGetV2ApiService {
	return c.ApiClient.ToolsThirdSiteGetV2Api
}

func (c *Client) ToolsThirdSitePreviewV2Api() *api.ToolsThirdSitePreviewV2ApiService {
	return c.ApiClient.ToolsThirdSitePreviewV2Api
}

func (c *Client) ToolsThirdSiteUpdateV2Api() *api.ToolsThirdSiteUpdateV2ApiService {
	return c.ApiClient.ToolsThirdSiteUpdateV2Api
}

func (c *Client) ToolsUnionFlowPackageCreateV2Api() *api.ToolsUnionFlowPackageCreateV2ApiService {
	return c.ApiClient.ToolsUnionFlowPackageCreateV2Api
}

func (c *Client) ToolsUnionFlowPackageDeleteV2Api() *api.ToolsUnionFlowPackageDeleteV2ApiService {
	return c.ApiClient.ToolsUnionFlowPackageDeleteV2Api
}

func (c *Client) ToolsUnionFlowPackageGetV2Api() *api.ToolsUnionFlowPackageGetV2ApiService {
	return c.ApiClient.ToolsUnionFlowPackageGetV2Api
}

func (c *Client) ToolsUnionFlowPackageReportV2Api() *api.ToolsUnionFlowPackageReportV2ApiService {
	return c.ApiClient.ToolsUnionFlowPackageReportV2Api
}

func (c *Client) ToolsUnionFlowPackageUpdateV2Api() *api.ToolsUnionFlowPackageUpdateV2ApiService {
	return c.ApiClient.ToolsUnionFlowPackageUpdateV2Api
}

func (c *Client) ToolsVideoCheckAvailableAnchorV2Api() *api.ToolsVideoCheckAvailableAnchorV2ApiService {
	return c.ApiClient.ToolsVideoCheckAvailableAnchorV2Api
}

func (c *Client) ToolsVideoCoverSuggestV2Api() *api.ToolsVideoCoverSuggestV2ApiService {
	return c.ApiClient.ToolsVideoCoverSuggestV2Api
}

func (c *Client) ToolsWechatAppletCreateV30Api() *api.ToolsWechatAppletCreateV30ApiService {
	return c.ApiClient.ToolsWechatAppletCreateV30Api
}

func (c *Client) ToolsWechatAppletListV30Api() *api.ToolsWechatAppletListV30ApiService {
	return c.ApiClient.ToolsWechatAppletListV30Api
}

func (c *Client) ToolsWechatAppletUpdateV30Api() *api.ToolsWechatAppletUpdateV30ApiService {
	return c.ApiClient.ToolsWechatAppletUpdateV30Api
}

func (c *Client) ToolsWechatGameCreateV30Api() *api.ToolsWechatGameCreateV30ApiService {
	return c.ApiClient.ToolsWechatGameCreateV30Api
}

func (c *Client) ToolsWechatGameListV30Api() *api.ToolsWechatGameListV30ApiService {
	return c.ApiClient.ToolsWechatGameListV30Api
}

func (c *Client) UserInfoV2Api() *api.UserInfoV2ApiService {
	return c.ApiClient.UserInfoV2Api
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
