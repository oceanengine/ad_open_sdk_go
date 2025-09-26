/*
API version: 1.1.71
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

func (c *Client) AccountFundGetV30Api() *api.AccountFundGetV30ApiService {
	return c.ApiClient.AccountFundGetV30Api
}

func (c *Client) AccountUpdateV30Api() *api.AccountUpdateV30ApiService {
	return c.ApiClient.AccountUpdateV30Api
}

func (c *Client) AdCostProtectStatusGetV2Api() *api.AdCostProtectStatusGetV2ApiService {
	return c.ApiClient.AdCostProtectStatusGetV2Api
}

func (c *Client) AdGetV2Api() *api.AdGetV2ApiService {
	return c.ApiClient.AdGetV2Api
}

func (c *Client) AdRejectReasonV2Api() *api.AdRejectReasonV2ApiService {
	return c.ApiClient.AdRejectReasonV2Api
}

func (c *Client) AdShopInfoUpdateV30Api() *api.AdShopInfoUpdateV30ApiService {
	return c.ApiClient.AdShopInfoUpdateV30Api
}

func (c *Client) AdUdUpdateV2Api() *api.AdUdUpdateV2ApiService {
	return c.ApiClient.AdUdUpdateV2Api
}

func (c *Client) AdUpdateBidV2Api() *api.AdUpdateBidV2ApiService {
	return c.ApiClient.AdUpdateBidV2Api
}

func (c *Client) AdUpdateBudgetV2Api() *api.AdUpdateBudgetV2ApiService {
	return c.ApiClient.AdUpdateBudgetV2Api
}

func (c *Client) AdUpdateStatusV2Api() *api.AdUpdateStatusV2ApiService {
	return c.ApiClient.AdUpdateStatusV2Api
}

func (c *Client) AdvConvertOleConvertV2Api() *api.AdvConvertOleConvertV2ApiService {
	return c.ApiClient.AdvConvertOleConvertV2Api
}

func (c *Client) AdvertiserAttachmentUploadV30Api() *api.AdvertiserAttachmentUploadV30ApiService {
	return c.ApiClient.AdvertiserAttachmentUploadV30Api
}

func (c *Client) AdvertiserAvatarGetV2Api() *api.AdvertiserAvatarGetV2ApiService {
	return c.ApiClient.AdvertiserAvatarGetV2Api
}

func (c *Client) AdvertiserAvatarSubmitV2Api() *api.AdvertiserAvatarSubmitV2ApiService {
	return c.ApiClient.AdvertiserAvatarSubmitV2Api
}

func (c *Client) AdvertiserAvatarUploadV2Api() *api.AdvertiserAvatarUploadV2ApiService {
	return c.ApiClient.AdvertiserAvatarUploadV2Api
}

func (c *Client) AdvertiserBudgetGetV2Api() *api.AdvertiserBudgetGetV2ApiService {
	return c.ApiClient.AdvertiserBudgetGetV2Api
}

func (c *Client) AdvertiserDeliveryPkgConfigV30Api() *api.AdvertiserDeliveryPkgConfigV30ApiService {
	return c.ApiClient.AdvertiserDeliveryPkgConfigV30Api
}

func (c *Client) AdvertiserDeliveryPkgDeleteV30Api() *api.AdvertiserDeliveryPkgDeleteV30ApiService {
	return c.ApiClient.AdvertiserDeliveryPkgDeleteV30Api
}

func (c *Client) AdvertiserDeliveryPkgGetV30Api() *api.AdvertiserDeliveryPkgGetV30ApiService {
	return c.ApiClient.AdvertiserDeliveryPkgGetV30Api
}

func (c *Client) AdvertiserDeliveryPkgSubmitV30Api() *api.AdvertiserDeliveryPkgSubmitV30ApiService {
	return c.ApiClient.AdvertiserDeliveryPkgSubmitV30Api
}

func (c *Client) AdvertiserDeliveryQualificationDeleteV30Api() *api.AdvertiserDeliveryQualificationDeleteV30ApiService {
	return c.ApiClient.AdvertiserDeliveryQualificationDeleteV30Api
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

func (c *Client) AdvertiserFundDetailGrantV2Api() *api.AdvertiserFundDetailGrantV2ApiService {
	return c.ApiClient.AdvertiserFundDetailGrantV2Api
}

func (c *Client) AdvertiserFundGetV2Api() *api.AdvertiserFundGetV2ApiService {
	return c.ApiClient.AdvertiserFundGetV2Api
}

func (c *Client) AdvertiserFundGrantTransactionGetV2Api() *api.AdvertiserFundGrantTransactionGetV2ApiService {
	return c.ApiClient.AdvertiserFundGrantTransactionGetV2Api
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

func (c *Client) AdvertiserVerifyInfoGetV30Api() *api.AdvertiserVerifyInfoGetV30ApiService {
	return c.ApiClient.AdvertiserVerifyInfoGetV30Api
}

func (c *Client) AgentAdvAdvertiserUpdateSaleV2Api() *api.AgentAdvAdvertiserUpdateSaleV2ApiService {
	return c.ApiClient.AgentAdvAdvertiserUpdateSaleV2Api
}

func (c *Client) AgentAdvBiddingListQueryV2Api() *api.AgentAdvBiddingListQueryV2ApiService {
	return c.ApiClient.AgentAdvBiddingListQueryV2Api
}

func (c *Client) AgentAdvBrandListQueryV2Api() *api.AgentAdvBrandListQueryV2ApiService {
	return c.ApiClient.AgentAdvBrandListQueryV2Api
}

func (c *Client) AgentAdvCostReportListQueryV2Api() *api.AgentAdvCostReportListQueryV2ApiService {
	return c.ApiClient.AgentAdvCostReportListQueryV2Api
}

func (c *Client) AgentAdvPerenniallyPunishHistoryQueryV2Api() *api.AgentAdvPerenniallyPunishHistoryQueryV2ApiService {
	return c.ApiClient.AgentAdvPerenniallyPunishHistoryQueryV2Api
}

func (c *Client) AgentAdvPerenniallyPunishV2Api() *api.AgentAdvPerenniallyPunishV2ApiService {
	return c.ApiClient.AgentAdvPerenniallyPunishV2Api
}

func (c *Client) AgentAdvRechargeRechargeRecordV2Api() *api.AgentAdvRechargeRechargeRecordV2ApiService {
	return c.ApiClient.AgentAdvRechargeRechargeRecordV2Api
}

func (c *Client) AgentAdvTemporaryPunishV2Api() *api.AgentAdvTemporaryPunishV2ApiService {
	return c.ApiClient.AgentAdvTemporaryPunishV2Api
}

func (c *Client) AgentAdvertiserAssignV2Api() *api.AgentAdvertiserAssignV2ApiService {
	return c.ApiClient.AgentAdvertiserAssignV2Api
}

func (c *Client) AgentAdvertiserCopyV2Api() *api.AgentAdvertiserCopyV2ApiService {
	return c.ApiClient.AgentAdvertiserCopyV2Api
}

func (c *Client) AgentAdvertiserInfoQueryV2Api() *api.AgentAdvertiserInfoQueryV2ApiService {
	return c.ApiClient.AgentAdvertiserInfoQueryV2Api
}

func (c *Client) AgentAdvertiserSelectV2Api() *api.AgentAdvertiserSelectV2ApiService {
	return c.ApiClient.AgentAdvertiserSelectV2Api
}

func (c *Client) AgentAdvertiserUnassignV2Api() *api.AgentAdvertiserUnassignV2ApiService {
	return c.ApiClient.AgentAdvertiserUnassignV2Api
}

func (c *Client) AgentAdvertiserUpdateV2Api() *api.AgentAdvertiserUpdateV2ApiService {
	return c.ApiClient.AgentAdvertiserUpdateV2Api
}

func (c *Client) AgentChargeVerifyV2Api() *api.AgentChargeVerifyV2ApiService {
	return c.ApiClient.AgentChargeVerifyV2Api
}

func (c *Client) AgentChildAgentSelectV2Api() *api.AgentChildAgentSelectV2ApiService {
	return c.ApiClient.AgentChildAgentSelectV2Api
}

func (c *Client) AgentCompanyBiddingListQueryV2Api() *api.AgentCompanyBiddingListQueryV2ApiService {
	return c.ApiClient.AgentCompanyBiddingListQueryV2Api
}

func (c *Client) AgentCompanyBrandListQueryV2Api() *api.AgentCompanyBrandListQueryV2ApiService {
	return c.ApiClient.AgentCompanyBrandListQueryV2Api
}

func (c *Client) AgentCreditChargeSubmitV2Api() *api.AgentCreditChargeSubmitV2ApiService {
	return c.ApiClient.AgentCreditChargeSubmitV2Api
}

func (c *Client) AgentInfoV2Api() *api.AgentInfoV2ApiService {
	return c.ApiClient.AgentInfoV2Api
}

func (c *Client) AgentPrepayChargeGenerateRemittanceCodeV2Api() *api.AgentPrepayChargeGenerateRemittanceCodeV2ApiService {
	return c.ApiClient.AgentPrepayChargeGenerateRemittanceCodeV2Api
}

func (c *Client) AgentQueryRiskPromotionListV2Api() *api.AgentQueryRiskPromotionListV2ApiService {
	return c.ApiClient.AgentQueryRiskPromotionListV2Api
}

func (c *Client) AgentTransferTransactionRecordV2Api() *api.AgentTransferTransactionRecordV2ApiService {
	return c.ApiClient.AgentTransferTransactionRecordV2Api
}

func (c *Client) AicElementDeleteV30Api() *api.AicElementDeleteV30ApiService {
	return c.ApiClient.AicElementDeleteV30Api
}

func (c *Client) AicElementGetV30Api() *api.AicElementGetV30ApiService {
	return c.ApiClient.AicElementGetV30Api
}

func (c *Client) AicElementUpdateV30Api() *api.AicElementUpdateV30ApiService {
	return c.ApiClient.AicElementUpdateV30Api
}

func (c *Client) AicElementUploadV30Api() *api.AicElementUploadV30ApiService {
	return c.ApiClient.AicElementUploadV30Api
}

func (c *Client) AicImageMixcutCreateV30Api() *api.AicImageMixcutCreateV30ApiService {
	return c.ApiClient.AicImageMixcutCreateV30Api
}

func (c *Client) AicMaterialGetV30Api() *api.AicMaterialGetV30ApiService {
	return c.ApiClient.AicMaterialGetV30Api
}

func (c *Client) AicMaterialPushV30Api() *api.AicMaterialPushV30ApiService {
	return c.ApiClient.AicMaterialPushV30Api
}

func (c *Client) AicMixcutTaskResultGetV30Api() *api.AicMixcutTaskResultGetV30ApiService {
	return c.ApiClient.AicMixcutTaskResultGetV30Api
}

func (c *Client) AicMixcutTaskSaveV30Api() *api.AicMixcutTaskSaveV30ApiService {
	return c.ApiClient.AicMixcutTaskSaveV30Api
}

func (c *Client) AicVideoMixcutCreateV30Api() *api.AicVideoMixcutCreateV30ApiService {
	return c.ApiClient.AicVideoMixcutCreateV30Api
}

func (c *Client) AnalyticsAttributionV30Api() *api.AnalyticsAttributionV30ApiService {
	return c.ApiClient.AnalyticsAttributionV30Api
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

func (c *Client) AudiencePackageBindinfoGetV30Api() *api.AudiencePackageBindinfoGetV30ApiService {
	return c.ApiClient.AudiencePackageBindinfoGetV30Api
}

func (c *Client) AudiencePackageCreateV2Api() *api.AudiencePackageCreateV2ApiService {
	return c.ApiClient.AudiencePackageCreateV2Api
}

func (c *Client) AudiencePackageDeleteV2Api() *api.AudiencePackageDeleteV2ApiService {
	return c.ApiClient.AudiencePackageDeleteV2Api
}

func (c *Client) AudiencePackageGetV30Api() *api.AudiencePackageGetV30ApiService {
	return c.ApiClient.AudiencePackageGetV30Api
}

func (c *Client) AudiencePackageUpdateV2Api() *api.AudiencePackageUpdateV2ApiService {
	return c.ApiClient.AudiencePackageUpdateV2Api
}

func (c *Client) BrandActionCategoryV30Api() *api.BrandActionCategoryV30ApiService {
	return c.ApiClient.BrandActionCategoryV30Api
}

func (c *Client) BrandAdCancelDeleteV30Api() *api.BrandAdCancelDeleteV30ApiService {
	return c.ApiClient.BrandAdCancelDeleteV30Api
}

func (c *Client) BrandAdCreateV30Api() *api.BrandAdCreateV30ApiService {
	return c.ApiClient.BrandAdCreateV30Api
}

func (c *Client) BrandAdDeleteV30Api() *api.BrandAdDeleteV30ApiService {
	return c.ApiClient.BrandAdDeleteV30Api
}

func (c *Client) BrandAdGetV30Api() *api.BrandAdGetV30ApiService {
	return c.ApiClient.BrandAdGetV30Api
}

func (c *Client) BrandAdUpdateBaseInfoV30Api() *api.BrandAdUpdateBaseInfoV30ApiService {
	return c.ApiClient.BrandAdUpdateBaseInfoV30Api
}

func (c *Client) BrandAdUpdateDeliveryInfoV30Api() *api.BrandAdUpdateDeliveryInfoV30ApiService {
	return c.ApiClient.BrandAdUpdateDeliveryInfoV30Api
}

func (c *Client) BrandAnchorListV30Api() *api.BrandAnchorListV30ApiService {
	return c.ApiClient.BrandAnchorListV30Api
}

func (c *Client) BrandAwemeListV30Api() *api.BrandAwemeListV30ApiService {
	return c.ApiClient.BrandAwemeListV30Api
}

func (c *Client) BrandCampaignAddV30Api() *api.BrandCampaignAddV30ApiService {
	return c.ApiClient.BrandCampaignAddV30Api
}

func (c *Client) BrandCampaignCreateV30Api() *api.BrandCampaignCreateV30ApiService {
	return c.ApiClient.BrandCampaignCreateV30Api
}

func (c *Client) BrandCampaignDeleteV30Api() *api.BrandCampaignDeleteV30ApiService {
	return c.ApiClient.BrandCampaignDeleteV30Api
}

func (c *Client) BrandCampaignEditV30Api() *api.BrandCampaignEditV30ApiService {
	return c.ApiClient.BrandCampaignEditV30Api
}

func (c *Client) BrandCampaignGetV30Api() *api.BrandCampaignGetV30ApiService {
	return c.ApiClient.BrandCampaignGetV30Api
}

func (c *Client) BrandCampaignListV30Api() *api.BrandCampaignListV30ApiService {
	return c.ApiClient.BrandCampaignListV30Api
}

func (c *Client) BrandCampaignModifyV30Api() *api.BrandCampaignModifyV30ApiService {
	return c.ApiClient.BrandCampaignModifyV30Api
}

func (c *Client) BrandCampaignOperateV30Api() *api.BrandCampaignOperateV30ApiService {
	return c.ApiClient.BrandCampaignOperateV30Api
}

func (c *Client) BrandCampaignRemoveV30Api() *api.BrandCampaignRemoveV30ApiService {
	return c.ApiClient.BrandCampaignRemoveV30Api
}

func (c *Client) BrandCampaignRevokeModifyV30Api() *api.BrandCampaignRevokeModifyV30ApiService {
	return c.ApiClient.BrandCampaignRevokeModifyV30Api
}

func (c *Client) BrandCampaignSubmitV30Api() *api.BrandCampaignSubmitV30ApiService {
	return c.ApiClient.BrandCampaignSubmitV30Api
}

func (c *Client) BrandCampaignUpdateV30Api() *api.BrandCampaignUpdateV30ApiService {
	return c.ApiClient.BrandCampaignUpdateV30Api
}

func (c *Client) BrandContractGetV30Api() *api.BrandContractGetV30ApiService {
	return c.ApiClient.BrandContractGetV30Api
}

func (c *Client) BrandCreativeCreateV30Api() *api.BrandCreativeCreateV30ApiService {
	return c.ApiClient.BrandCreativeCreateV30Api
}

func (c *Client) BrandCreativeDeleteV30Api() *api.BrandCreativeDeleteV30ApiService {
	return c.ApiClient.BrandCreativeDeleteV30Api
}

func (c *Client) BrandCreativeGetV30Api() *api.BrandCreativeGetV30ApiService {
	return c.ApiClient.BrandCreativeGetV30Api
}

func (c *Client) BrandCreativeUpdateV30Api() *api.BrandCreativeUpdateV30ApiService {
	return c.ApiClient.BrandCreativeUpdateV30Api
}

func (c *Client) BrandCustomAudienceListV30Api() *api.BrandCustomAudienceListV30ApiService {
	return c.ApiClient.BrandCustomAudienceListV30Api
}

func (c *Client) BrandFileVideoUploadV30Api() *api.BrandFileVideoUploadV30ApiService {
	return c.ApiClient.BrandFileVideoUploadV30Api
}

func (c *Client) BrandMaterialCreateV30Api() *api.BrandMaterialCreateV30ApiService {
	return c.ApiClient.BrandMaterialCreateV30Api
}

func (c *Client) BrandMaterialListV30Api() *api.BrandMaterialListV30ApiService {
	return c.ApiClient.BrandMaterialListV30Api
}

func (c *Client) BrandMaterialUpdateV30Api() *api.BrandMaterialUpdateV30ApiService {
	return c.ApiClient.BrandMaterialUpdateV30Api
}

func (c *Client) BrandOperationLogQueryV30Api() *api.BrandOperationLogQueryV30ApiService {
	return c.ApiClient.BrandOperationLogQueryV30Api
}

func (c *Client) BrandOrderCancelDeleteV30Api() *api.BrandOrderCancelDeleteV30ApiService {
	return c.ApiClient.BrandOrderCancelDeleteV30Api
}

func (c *Client) BrandOrderCreateV30Api() *api.BrandOrderCreateV30ApiService {
	return c.ApiClient.BrandOrderCreateV30Api
}

func (c *Client) BrandOrderDeleteV30Api() *api.BrandOrderDeleteV30ApiService {
	return c.ApiClient.BrandOrderDeleteV30Api
}

func (c *Client) BrandOrderListV30Api() *api.BrandOrderListV30ApiService {
	return c.ApiClient.BrandOrderListV30Api
}

func (c *Client) BrandOrderUpdateV30Api() *api.BrandOrderUpdateV30ApiService {
	return c.ApiClient.BrandOrderUpdateV30Api
}

func (c *Client) BrandPolicyListV30Api() *api.BrandPolicyListV30ApiService {
	return c.ApiClient.BrandPolicyListV30Api
}

func (c *Client) BrandQueryStockV30Api() *api.BrandQueryStockV30ApiService {
	return c.ApiClient.BrandQueryStockV30Api
}

func (c *Client) BrandQueryYuntu5aBrandCategoryV30Api() *api.BrandQueryYuntu5aBrandCategoryV30ApiService {
	return c.ApiClient.BrandQueryYuntu5aBrandCategoryV30Api
}

func (c *Client) BrandRegionGetV30Api() *api.BrandRegionGetV30ApiService {
	return c.ApiClient.BrandRegionGetV30Api
}

func (c *Client) BrandToolCreativePreviewV30Api() *api.BrandToolCreativePreviewV30ApiService {
	return c.ApiClient.BrandToolCreativePreviewV30Api
}

func (c *Client) BrandToolMaterialPreviewV30Api() *api.BrandToolMaterialPreviewV30ApiService {
	return c.ApiClient.BrandToolMaterialPreviewV30Api
}

func (c *Client) BrandToolQueryPublishPriceV30Api() *api.BrandToolQueryPublishPriceV30ApiService {
	return c.ApiClient.BrandToolQueryPublishPriceV30Api
}

func (c *Client) BrandToolQueryStockBalanceV30Api() *api.BrandToolQueryStockBalanceV30ApiService {
	return c.ApiClient.BrandToolQueryStockBalanceV30Api
}

func (c *Client) BrandUploadImageV30Api() *api.BrandUploadImageV30ApiService {
	return c.ApiClient.BrandUploadImageV30Api
}

func (c *Client) BudgetGroupCreateV30Api() *api.BudgetGroupCreateV30ApiService {
	return c.ApiClient.BudgetGroupCreateV30Api
}

func (c *Client) BudgetGroupDeleteV30Api() *api.BudgetGroupDeleteV30ApiService {
	return c.ApiClient.BudgetGroupDeleteV30Api
}

func (c *Client) BudgetGroupListV30Api() *api.BudgetGroupListV30ApiService {
	return c.ApiClient.BudgetGroupListV30Api
}

func (c *Client) BudgetGroupUpdateV30Api() *api.BudgetGroupUpdateV30ApiService {
	return c.ApiClient.BudgetGroupUpdateV30Api
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

func (c *Client) CampaignCreateV2Api() *api.CampaignCreateV2ApiService {
	return c.ApiClient.CampaignCreateV2Api
}

func (c *Client) CampaignGetV2Api() *api.CampaignGetV2ApiService {
	return c.ApiClient.CampaignGetV2Api
}

func (c *Client) CampaignUpdateStatusV2Api() *api.CampaignUpdateStatusV2ApiService {
	return c.ApiClient.CampaignUpdateStatusV2Api
}

func (c *Client) CampaignUpdateV2Api() *api.CampaignUpdateV2ApiService {
	return c.ApiClient.CampaignUpdateV2Api
}

func (c *Client) CarouselAdGetV2Api() *api.CarouselAdGetV2ApiService {
	return c.ApiClient.CarouselAdGetV2Api
}

func (c *Client) CarouselCreateV2Api() *api.CarouselCreateV2ApiService {
	return c.ApiClient.CarouselCreateV2Api
}

func (c *Client) CarouselDeleteV2Api() *api.CarouselDeleteV2ApiService {
	return c.ApiClient.CarouselDeleteV2Api
}

func (c *Client) CarouselListV2Api() *api.CarouselListV2ApiService {
	return c.ApiClient.CarouselListV2Api
}

func (c *Client) CarouselUpdateV2Api() *api.CarouselUpdateV2ApiService {
	return c.ApiClient.CarouselUpdateV2Api
}

func (c *Client) CdpBrandGetV30Api() *api.CdpBrandGetV30ApiService {
	return c.ApiClient.CdpBrandGetV30Api
}

func (c *Client) CgTransferCanTransferBalanceGetV30Api() *api.CgTransferCanTransferBalanceGetV30ApiService {
	return c.ApiClient.CgTransferCanTransferBalanceGetV30Api
}

func (c *Client) CgTransferCanTransferTargetListV30Api() *api.CgTransferCanTransferTargetListV30ApiService {
	return c.ApiClient.CgTransferCanTransferTargetListV30Api
}

func (c *Client) CgTransferCreateTransferV30Api() *api.CgTransferCreateTransferV30ApiService {
	return c.ApiClient.CgTransferCreateTransferV30Api
}

func (c *Client) CgTransferQueryCanTransferBalanceV30Api() *api.CgTransferQueryCanTransferBalanceV30ApiService {
	return c.ApiClient.CgTransferQueryCanTransferBalanceV30Api
}

func (c *Client) CgTransferQueryTransferBalanceV30Api() *api.CgTransferQueryTransferBalanceV30ApiService {
	return c.ApiClient.CgTransferQueryTransferBalanceV30Api
}

func (c *Client) CgTransferQueryTransferDetailV30Api() *api.CgTransferQueryTransferDetailV30ApiService {
	return c.ApiClient.CgTransferQueryTransferDetailV30Api
}

func (c *Client) CgTransferTransferBalanceGetV30Api() *api.CgTransferTransferBalanceGetV30ApiService {
	return c.ApiClient.CgTransferTransferBalanceGetV30Api
}

func (c *Client) CgTransferTransferCreateV30Api() *api.CgTransferTransferCreateV30ApiService {
	return c.ApiClient.CgTransferTransferCreateV30Api
}

func (c *Client) CgTransferTransferDetailGetV30Api() *api.CgTransferTransferDetailGetV30ApiService {
	return c.ApiClient.CgTransferTransferDetailGetV30Api
}

func (c *Client) CgTransferWalletTransferCanTransferBalanceV30Api() *api.CgTransferWalletTransferCanTransferBalanceV30ApiService {
	return c.ApiClient.CgTransferWalletTransferCanTransferBalanceV30Api
}

func (c *Client) CgTransferWalletTransferCreateV30Api() *api.CgTransferWalletTransferCreateV30ApiService {
	return c.ApiClient.CgTransferWalletTransferCreateV30Api
}

func (c *Client) CgTransferWalletTransferDetailV30Api() *api.CgTransferWalletTransferDetailV30ApiService {
	return c.ApiClient.CgTransferWalletTransferDetailV30Api
}

func (c *Client) CgTransferWalletTransferListV30Api() *api.CgTransferWalletTransferListV30ApiService {
	return c.ApiClient.CgTransferWalletTransferListV30Api
}

func (c *Client) ChargeListV30Api() *api.ChargeListV30ApiService {
	return c.ApiClient.ChargeListV30Api
}

func (c *Client) ChargeResultV30Api() *api.ChargeResultV30ApiService {
	return c.ApiClient.ChargeResultV30Api
}

func (c *Client) ClueCaCreateV2Api() *api.ClueCaCreateV2ApiService {
	return c.ApiClient.ClueCaCreateV2Api
}

func (c *Client) ClueCaInterfaceCreateV2Api() *api.ClueCaInterfaceCreateV2ApiService {
	return c.ApiClient.ClueCaInterfaceCreateV2Api
}

func (c *Client) ClueCaInterfaceUpdateV2Api() *api.ClueCaInterfaceUpdateV2ApiService {
	return c.ApiClient.ClueCaInterfaceUpdateV2Api
}

func (c *Client) ClueCaUpdateV2Api() *api.ClueCaUpdateV2ApiService {
	return c.ApiClient.ClueCaUpdateV2Api
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

func (c *Client) ClueWechatDataGetV2Api() *api.ClueWechatDataGetV2ApiService {
	return c.ApiClient.ClueWechatDataGetV2Api
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

func (c *Client) CreateProjectInvoiceV2Api() *api.CreateProjectInvoiceV2ApiService {
	return c.ApiClient.CreateProjectInvoiceV2Api
}

func (c *Client) CreateStatementInvoiceV2Api() *api.CreateStatementInvoiceV2ApiService {
	return c.ApiClient.CreateStatementInvoiceV2Api
}

func (c *Client) CreateStatementV2Api() *api.CreateStatementV2ApiService {
	return c.ApiClient.CreateStatementV2Api
}

func (c *Client) CreativeCustomCreativeCreateV2Api() *api.CreativeCustomCreativeCreateV2ApiService {
	return c.ApiClient.CreativeCustomCreativeCreateV2Api
}

func (c *Client) CreativeCustomCreativeUpdateV2Api() *api.CreativeCustomCreativeUpdateV2ApiService {
	return c.ApiClient.CreativeCustomCreativeUpdateV2Api
}

func (c *Client) CreativeDetailGetV30Api() *api.CreativeDetailGetV30ApiService {
	return c.ApiClient.CreativeDetailGetV30Api
}

func (c *Client) CreativeGetV2Api() *api.CreativeGetV2ApiService {
	return c.ApiClient.CreativeGetV2Api
}

func (c *Client) CreativeProceduralCreativeCreateV2Api() *api.CreativeProceduralCreativeCreateV2ApiService {
	return c.ApiClient.CreativeProceduralCreativeCreateV2Api
}

func (c *Client) CreativeProceduralCreativeUpdateV2Api() *api.CreativeProceduralCreativeUpdateV2ApiService {
	return c.ApiClient.CreativeProceduralCreativeUpdateV2Api
}

func (c *Client) CreativeRejectReasonV2Api() *api.CreativeRejectReasonV2ApiService {
	return c.ApiClient.CreativeRejectReasonV2Api
}

func (c *Client) CreativeStrategyListV2Api() *api.CreativeStrategyListV2ApiService {
	return c.ApiClient.CreativeStrategyListV2Api
}

func (c *Client) CustomerCenterAccountListV30Api() *api.CustomerCenterAccountListV30ApiService {
	return c.ApiClient.CustomerCenterAccountListV30Api
}

func (c *Client) CustomerCenterAccountOfflineListV30Api() *api.CustomerCenterAccountOfflineListV30ApiService {
	return c.ApiClient.CustomerCenterAccountOfflineListV30Api
}

func (c *Client) CustomerCenterAdvertiserCopyV2Api() *api.CustomerCenterAdvertiserCopyV2ApiService {
	return c.ApiClient.CustomerCenterAdvertiserCopyV2Api
}

func (c *Client) CustomerCenterAdvertiserListV2Api() *api.CustomerCenterAdvertiserListV2ApiService {
	return c.ApiClient.CustomerCenterAdvertiserListV2Api
}

func (c *Client) CustomerCenterAdvertiserTransferableListV2Api() *api.CustomerCenterAdvertiserTransferableListV2ApiService {
	return c.ApiClient.CustomerCenterAdvertiserTransferableListV2Api
}

func (c *Client) CustomerCenterFundTransferSeqCommitV2Api() *api.CustomerCenterFundTransferSeqCommitV2ApiService {
	return c.ApiClient.CustomerCenterFundTransferSeqCommitV2Api
}

func (c *Client) CustomerCenterFundTransferSeqCreateV2Api() *api.CustomerCenterFundTransferSeqCreateV2ApiService {
	return c.ApiClient.CustomerCenterFundTransferSeqCreateV2Api
}

func (c *Client) DcdChargeSubmitV30Api() *api.DcdChargeSubmitV30ApiService {
	return c.ApiClient.DcdChargeSubmitV30Api
}

func (c *Client) DecorationCouponGetV30Api() *api.DecorationCouponGetV30ApiService {
	return c.ApiClient.DecorationCouponGetV30Api
}

func (c *Client) DiagnosisTaskAdvCreateV2Api() *api.DiagnosisTaskAdvCreateV2ApiService {
	return c.ApiClient.DiagnosisTaskAdvCreateV2Api
}

func (c *Client) DiagnosisTaskAdvGetV2Api() *api.DiagnosisTaskAdvGetV2ApiService {
	return c.ApiClient.DiagnosisTaskAdvGetV2Api
}

func (c *Client) DiagnosisTaskAdvListV2Api() *api.DiagnosisTaskAdvListV2ApiService {
	return c.ApiClient.DiagnosisTaskAdvListV2Api
}

func (c *Client) DiagnosisTaskAgentCreateV2Api() *api.DiagnosisTaskAgentCreateV2ApiService {
	return c.ApiClient.DiagnosisTaskAgentCreateV2Api
}

func (c *Client) DiagnosisTaskAgentGetV2Api() *api.DiagnosisTaskAgentGetV2ApiService {
	return c.ApiClient.DiagnosisTaskAgentGetV2Api
}

func (c *Client) DiagnosisTaskAgentListV2Api() *api.DiagnosisTaskAgentListV2ApiService {
	return c.ApiClient.DiagnosisTaskAgentListV2Api
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

func (c *Client) DmpDataSourceCreateV2Api() *api.DmpDataSourceCreateV2ApiService {
	return c.ApiClient.DmpDataSourceCreateV2Api
}

func (c *Client) DmpDataSourceFileUploadV2Api() *api.DmpDataSourceFileUploadV2ApiService {
	return c.ApiClient.DmpDataSourceFileUploadV2Api
}

func (c *Client) DmpDataSourceReadV2Api() *api.DmpDataSourceReadV2ApiService {
	return c.ApiClient.DmpDataSourceReadV2Api
}

func (c *Client) DmpDataSourceUpdateV2Api() *api.DmpDataSourceUpdateV2ApiService {
	return c.ApiClient.DmpDataSourceUpdateV2Api
}

func (c *Client) DouplusOptionalItemsListV30Api() *api.DouplusOptionalItemsListV30ApiService {
	return c.ApiClient.DouplusOptionalItemsListV30Api
}

func (c *Client) DouplusOptionalTargetsListV30Api() *api.DouplusOptionalTargetsListV30ApiService {
	return c.ApiClient.DouplusOptionalTargetsListV30Api
}

func (c *Client) DouplusOrderCloseV30Api() *api.DouplusOrderCloseV30ApiService {
	return c.ApiClient.DouplusOrderCloseV30Api
}

func (c *Client) DouplusOrderCreateV30Api() *api.DouplusOrderCreateV30ApiService {
	return c.ApiClient.DouplusOrderCreateV30Api
}

func (c *Client) DouplusOrderListV30Api() *api.DouplusOrderListV30ApiService {
	return c.ApiClient.DouplusOrderListV30Api
}

func (c *Client) DouplusOrderRenewV30Api() *api.DouplusOrderRenewV30ApiService {
	return c.ApiClient.DouplusOrderRenewV30Api
}

func (c *Client) DouplusOrderReportV30Api() *api.DouplusOrderReportV30ApiService {
	return c.ApiClient.DouplusOrderReportV30Api
}

func (c *Client) DouplusRtaGetInfoV30Api() *api.DouplusRtaGetInfoV30ApiService {
	return c.ApiClient.DouplusRtaGetInfoV30Api
}

func (c *Client) DouplusRtaSetScopeV30Api() *api.DouplusRtaSetScopeV30ApiService {
	return c.ApiClient.DouplusRtaSetScopeV30Api
}

func (c *Client) DownloadStatementEsignFileV2Api() *api.DownloadStatementEsignFileV2ApiService {
	return c.ApiClient.DownloadStatementEsignFileV2Api
}

func (c *Client) DownloadStatementV2Api() *api.DownloadStatementV2ApiService {
	return c.ApiClient.DownloadStatementV2Api
}

func (c *Client) DpaAlbumCreateV30Api() *api.DpaAlbumCreateV30ApiService {
	return c.ApiClient.DpaAlbumCreateV30Api
}

func (c *Client) DpaAlbumStatusGetV30Api() *api.DpaAlbumStatusGetV30ApiService {
	return c.ApiClient.DpaAlbumStatusGetV30Api
}

func (c *Client) DpaAssetV2DetailReadV2Api() *api.DpaAssetV2DetailReadV2ApiService {
	return c.ApiClient.DpaAssetV2DetailReadV2Api
}

func (c *Client) DpaAssetV2ListV2Api() *api.DpaAssetV2ListV2ApiService {
	return c.ApiClient.DpaAssetV2ListV2Api
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

func (c *Client) DpaCheckIndexEntryProgressV2Api() *api.DpaCheckIndexEntryProgressV2ApiService {
	return c.ApiClient.DpaCheckIndexEntryProgressV2Api
}

func (c *Client) DpaClueProductDeleteV2Api() *api.DpaClueProductDeleteV2ApiService {
	return c.ApiClient.DpaClueProductDeleteV2Api
}

func (c *Client) DpaClueProductDetailV2Api() *api.DpaClueProductDetailV2ApiService {
	return c.ApiClient.DpaClueProductDetailV2Api
}

func (c *Client) DpaClueProductListV2Api() *api.DpaClueProductListV2ApiService {
	return c.ApiClient.DpaClueProductListV2Api
}

func (c *Client) DpaClueProductSaveV2Api() *api.DpaClueProductSaveV2ApiService {
	return c.ApiClient.DpaClueProductSaveV2Api
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

func (c *Client) DpaPlayletAuthGetV2Api() *api.DpaPlayletAuthGetV2ApiService {
	return c.ApiClient.DpaPlayletAuthGetV2Api
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

func (c *Client) EbpAdvertiserListV2Api() *api.EbpAdvertiserListV2ApiService {
	return c.ApiClient.EbpAdvertiserListV2Api
}

func (c *Client) EbpAdvertiserTaskCreateV2Api() *api.EbpAdvertiserTaskCreateV2ApiService {
	return c.ApiClient.EbpAdvertiserTaskCreateV2Api
}

func (c *Client) EbpAdvertiserTaskDownloadV2Api() *api.EbpAdvertiserTaskDownloadV2ApiService {
	return c.ApiClient.EbpAdvertiserTaskDownloadV2Api
}

func (c *Client) EbpAdvertiserTaskListV2Api() *api.EbpAdvertiserTaskListV2ApiService {
	return c.ApiClient.EbpAdvertiserTaskListV2Api
}

func (c *Client) EbpLevelGetV2Api() *api.EbpLevelGetV2ApiService {
	return c.ApiClient.EbpLevelGetV2Api
}

func (c *Client) EnterpriseBindListGetV10Api() *api.EnterpriseBindListGetV10ApiService {
	return c.ApiClient.EnterpriseBindListGetV10Api
}

func (c *Client) EventManagerAbnormalAssetsGetV30Api() *api.EventManagerAbnormalAssetsGetV30ApiService {
	return c.ApiClient.EventManagerAbnormalAssetsGetV30Api
}

func (c *Client) EventManagerAssetsCreateV2Api() *api.EventManagerAssetsCreateV2ApiService {
	return c.ApiClient.EventManagerAssetsCreateV2Api
}

func (c *Client) EventManagerAuthAddPublicKeyV2Api() *api.EventManagerAuthAddPublicKeyV2ApiService {
	return c.ApiClient.EventManagerAuthAddPublicKeyV2Api
}

func (c *Client) EventManagerAuthDelPublicKeyV2Api() *api.EventManagerAuthDelPublicKeyV2ApiService {
	return c.ApiClient.EventManagerAuthDelPublicKeyV2Api
}

func (c *Client) EventManagerAuthDisableV2Api() *api.EventManagerAuthDisableV2ApiService {
	return c.ApiClient.EventManagerAuthDisableV2Api
}

func (c *Client) EventManagerAuthEnableV2Api() *api.EventManagerAuthEnableV2ApiService {
	return c.ApiClient.EventManagerAuthEnableV2Api
}

func (c *Client) EventManagerAuthGetAllPublicKeysV2Api() *api.EventManagerAuthGetAllPublicKeysV2ApiService {
	return c.ApiClient.EventManagerAuthGetAllPublicKeysV2Api
}

func (c *Client) EventManagerAuthGetAuthStatusV2Api() *api.EventManagerAuthGetAuthStatusV2ApiService {
	return c.ApiClient.EventManagerAuthGetAuthStatusV2Api
}

func (c *Client) EventManagerAuthGetPublicKeyV2Api() *api.EventManagerAuthGetPublicKeyV2ApiService {
	return c.ApiClient.EventManagerAuthGetPublicKeyV2Api
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

func (c *Client) FileAudioAdV2Api() *api.FileAudioAdV2ApiService {
	return c.ApiClient.FileAudioAdV2Api
}

func (c *Client) FileAudioGetV2Api() *api.FileAudioGetV2ApiService {
	return c.ApiClient.FileAudioGetV2Api
}

func (c *Client) FileAutoGenerateSourceGetV2Api() *api.FileAutoGenerateSourceGetV2ApiService {
	return c.ApiClient.FileAutoGenerateSourceGetV2Api
}

func (c *Client) FileCarouselAwemeGetV30Api() *api.FileCarouselAwemeGetV30ApiService {
	return c.ApiClient.FileCarouselAwemeGetV30Api
}

func (c *Client) FileImageAdGetV2Api() *api.FileImageAdGetV2ApiService {
	return c.ApiClient.FileImageAdGetV2Api
}

func (c *Client) FileImageAdV2Api() *api.FileImageAdV2ApiService {
	return c.ApiClient.FileImageAdV2Api
}

func (c *Client) FileImageAdvertiserV2Api() *api.FileImageAdvertiserV2ApiService {
	return c.ApiClient.FileImageAdvertiserV2Api
}

func (c *Client) FileImageDeleteV30Api() *api.FileImageDeleteV30ApiService {
	return c.ApiClient.FileImageDeleteV30Api
}

func (c *Client) FileImageGetV2Api() *api.FileImageGetV2ApiService {
	return c.ApiClient.FileImageGetV2Api
}

func (c *Client) FileMaterialAttributesListV2Api() *api.FileMaterialAttributesListV2ApiService {
	return c.ApiClient.FileMaterialAttributesListV2Api
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

func (c *Client) FilePlayableCreateV30Api() *api.FilePlayableCreateV30ApiService {
	return c.ApiClient.FilePlayableCreateV30Api
}

func (c *Client) FilePlayableListV30Api() *api.FilePlayableListV30ApiService {
	return c.ApiClient.FilePlayableListV30Api
}

func (c *Client) FilePreauditGetV30Api() *api.FilePreauditGetV30ApiService {
	return c.ApiClient.FilePreauditGetV30Api
}

func (c *Client) FilePreauditSubmitV30Api() *api.FilePreauditSubmitV30ApiService {
	return c.ApiClient.FilePreauditSubmitV30Api
}

func (c *Client) FileQualityGetV30Api() *api.FileQualityGetV30ApiService {
	return c.ApiClient.FileQualityGetV30Api
}

func (c *Client) FileQualitySubmitV30Api() *api.FileQualitySubmitV30ApiService {
	return c.ApiClient.FileQualitySubmitV30Api
}

func (c *Client) FileRebateCommonDownloadCreateTaskV2Api() *api.FileRebateCommonDownloadCreateTaskV2ApiService {
	return c.ApiClient.FileRebateCommonDownloadCreateTaskV2Api
}

func (c *Client) FileRebateCommonDownloadDownloadFileV2Api() *api.FileRebateCommonDownloadDownloadFileV2ApiService {
	return c.ApiClient.FileRebateCommonDownloadDownloadFileV2Api
}

func (c *Client) FileRebateCommonDownloadGetDownloadTaskListV2Api() *api.FileRebateCommonDownloadGetDownloadTaskListV2ApiService {
	return c.ApiClient.FileRebateCommonDownloadGetDownloadTaskListV2Api
}

func (c *Client) FileRebateMaterialDownloadCreateTaskV2Api() *api.FileRebateMaterialDownloadCreateTaskV2ApiService {
	return c.ApiClient.FileRebateMaterialDownloadCreateTaskV2Api
}

func (c *Client) FileRebateMaterialDownloadDownloadFileV2Api() *api.FileRebateMaterialDownloadDownloadFileV2ApiService {
	return c.ApiClient.FileRebateMaterialDownloadDownloadFileV2Api
}

func (c *Client) FileRebateMaterialDownloadGetDownloadTaskListV2Api() *api.FileRebateMaterialDownloadGetDownloadTaskListV2ApiService {
	return c.ApiClient.FileRebateMaterialDownloadGetDownloadTaskListV2Api
}

func (c *Client) FileRebateRebateDownloadCreateTaskV2Api() *api.FileRebateRebateDownloadCreateTaskV2ApiService {
	return c.ApiClient.FileRebateRebateDownloadCreateTaskV2Api
}

func (c *Client) FileUploadTaskCreateV2Api() *api.FileUploadTaskCreateV2ApiService {
	return c.ApiClient.FileUploadTaskCreateV2Api
}

func (c *Client) FileVideoAdGetV2Api() *api.FileVideoAdGetV2ApiService {
	return c.ApiClient.FileVideoAdGetV2Api
}

func (c *Client) FileVideoAdV2Api() *api.FileVideoAdV2ApiService {
	return c.ApiClient.FileVideoAdV2Api
}

func (c *Client) FileVideoAgentGetV2Api() *api.FileVideoAgentGetV2ApiService {
	return c.ApiClient.FileVideoAgentGetV2Api
}

func (c *Client) FileVideoAgentV2Api() *api.FileVideoAgentV2ApiService {
	return c.ApiClient.FileVideoAgentV2Api
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

func (c *Client) FileVideoPauseV2Api() *api.FileVideoPauseV2ApiService {
	return c.ApiClient.FileVideoPauseV2Api
}

func (c *Client) FileVideoUpdateV2Api() *api.FileVideoUpdateV2ApiService {
	return c.ApiClient.FileVideoUpdateV2Api
}

func (c *Client) FileVideoUploadTaskListV2Api() *api.FileVideoUploadTaskListV2ApiService {
	return c.ApiClient.FileVideoUploadTaskListV2Api
}

func (c *Client) FundSharedWalletBalanceGetV2Api() *api.FundSharedWalletBalanceGetV2ApiService {
	return c.ApiClient.FundSharedWalletBalanceGetV2Api
}

func (c *Client) GameAddictionIdGetV30Api() *api.GameAddictionIdGetV30ApiService {
	return c.ApiClient.GameAddictionIdGetV30Api
}

func (c *Client) KeywordCreateV2V2Api() *api.KeywordCreateV2V2ApiService {
	return c.ApiClient.KeywordCreateV2V2Api
}

func (c *Client) KeywordCreateV30Api() *api.KeywordCreateV30ApiService {
	return c.ApiClient.KeywordCreateV30Api
}

func (c *Client) KeywordDeleteV2V2Api() *api.KeywordDeleteV2V2ApiService {
	return c.ApiClient.KeywordDeleteV2V2Api
}

func (c *Client) KeywordDeleteV30Api() *api.KeywordDeleteV30ApiService {
	return c.ApiClient.KeywordDeleteV30Api
}

func (c *Client) KeywordFeedadsSuggestV2Api() *api.KeywordFeedadsSuggestV2ApiService {
	return c.ApiClient.KeywordFeedadsSuggestV2Api
}

func (c *Client) KeywordGetV2Api() *api.KeywordGetV2ApiService {
	return c.ApiClient.KeywordGetV2Api
}

func (c *Client) KeywordListV30Api() *api.KeywordListV30ApiService {
	return c.ApiClient.KeywordListV30Api
}

func (c *Client) KeywordUpdateV2V2Api() *api.KeywordUpdateV2V2ApiService {
	return c.ApiClient.KeywordUpdateV2V2Api
}

func (c *Client) KeywordUpdateV30Api() *api.KeywordUpdateV30ApiService {
	return c.ApiClient.KeywordUpdateV30Api
}

func (c *Client) LocalAwemeAuthorizedGetV30Api() *api.LocalAwemeAuthorizedGetV30ApiService {
	return c.ApiClient.LocalAwemeAuthorizedGetV30Api
}

func (c *Client) LocalChargeListV30Api() *api.LocalChargeListV30ApiService {
	return c.ApiClient.LocalChargeListV30Api
}

func (c *Client) LocalChargeResultV30Api() *api.LocalChargeResultV30ApiService {
	return c.ApiClient.LocalChargeResultV30Api
}

func (c *Client) LocalChargeSubmitV30Api() *api.LocalChargeSubmitV30ApiService {
	return c.ApiClient.LocalChargeSubmitV30Api
}

func (c *Client) LocalCustomAudienceGetV30Api() *api.LocalCustomAudienceGetV30ApiService {
	return c.ApiClient.LocalCustomAudienceGetV30Api
}

func (c *Client) LocalFileUploadTaskCreateV30Api() *api.LocalFileUploadTaskCreateV30ApiService {
	return c.ApiClient.LocalFileUploadTaskCreateV30Api
}

func (c *Client) LocalFileVideoAwemeGetV30Api() *api.LocalFileVideoAwemeGetV30ApiService {
	return c.ApiClient.LocalFileVideoAwemeGetV30Api
}

func (c *Client) LocalFileVideoGetV30Api() *api.LocalFileVideoGetV30ApiService {
	return c.ApiClient.LocalFileVideoGetV30Api
}

func (c *Client) LocalFileVideoUploadTaskListV30Api() *api.LocalFileVideoUploadTaskListV30ApiService {
	return c.ApiClient.LocalFileVideoUploadTaskListV30Api
}

func (c *Client) LocalFileVideoUploadV30Api() *api.LocalFileVideoUploadV30ApiService {
	return c.ApiClient.LocalFileVideoUploadV30Api
}

func (c *Client) LocalImQueryMsgV30Api() *api.LocalImQueryMsgV30ApiService {
	return c.ApiClient.LocalImQueryMsgV30Api
}

func (c *Client) LocalImSendMsgV30Api() *api.LocalImSendMsgV30ApiService {
	return c.ApiClient.LocalImSendMsgV30Api
}

func (c *Client) LocalLifeAdvertiserCreateV30Api() *api.LocalLifeAdvertiserCreateV30ApiService {
	return c.ApiClient.LocalLifeAdvertiserCreateV30Api
}

func (c *Client) LocalLifeAdvertiserListV30Api() *api.LocalLifeAdvertiserListV30ApiService {
	return c.ApiClient.LocalLifeAdvertiserListV30Api
}

func (c *Client) LocalMultiPoiIdPoiIdsGetV30Api() *api.LocalMultiPoiIdPoiIdsGetV30ApiService {
	return c.ApiClient.LocalMultiPoiIdPoiIdsGetV30Api
}

func (c *Client) LocalPoiGetV30Api() *api.LocalPoiGetV30ApiService {
	return c.ApiClient.LocalPoiGetV30Api
}

func (c *Client) LocalProductGetByPoiidsV30Api() *api.LocalProductGetByPoiidsV30ApiService {
	return c.ApiClient.LocalProductGetByPoiidsV30Api
}

func (c *Client) LocalProductGetV30Api() *api.LocalProductGetV30ApiService {
	return c.ApiClient.LocalProductGetV30Api
}

func (c *Client) LocalProjectCreateV30Api() *api.LocalProjectCreateV30ApiService {
	return c.ApiClient.LocalProjectCreateV30Api
}

func (c *Client) LocalProjectDetailV30Api() *api.LocalProjectDetailV30ApiService {
	return c.ApiClient.LocalProjectDetailV30Api
}

func (c *Client) LocalProjectListV30Api() *api.LocalProjectListV30ApiService {
	return c.ApiClient.LocalProjectListV30Api
}

func (c *Client) LocalProjectStatusUpdateV30Api() *api.LocalProjectStatusUpdateV30ApiService {
	return c.ApiClient.LocalProjectStatusUpdateV30Api
}

func (c *Client) LocalProjectUpdateV30Api() *api.LocalProjectUpdateV30ApiService {
	return c.ApiClient.LocalProjectUpdateV30Api
}

func (c *Client) LocalPromotionCreateV30Api() *api.LocalPromotionCreateV30ApiService {
	return c.ApiClient.LocalPromotionCreateV30Api
}

func (c *Client) LocalPromotionDetailV30Api() *api.LocalPromotionDetailV30ApiService {
	return c.ApiClient.LocalPromotionDetailV30Api
}

func (c *Client) LocalPromotionListV30Api() *api.LocalPromotionListV30ApiService {
	return c.ApiClient.LocalPromotionListV30Api
}

func (c *Client) LocalPromotionStatusUpdateV30Api() *api.LocalPromotionStatusUpdateV30ApiService {
	return c.ApiClient.LocalPromotionStatusUpdateV30Api
}

func (c *Client) LocalPromotionUpdateV30Api() *api.LocalPromotionUpdateV30ApiService {
	return c.ApiClient.LocalPromotionUpdateV30Api
}

func (c *Client) LocalReportAccountGetV30Api() *api.LocalReportAccountGetV30ApiService {
	return c.ApiClient.LocalReportAccountGetV30Api
}

func (c *Client) LocalReportAudienceGetV30Api() *api.LocalReportAudienceGetV30ApiService {
	return c.ApiClient.LocalReportAudienceGetV30Api
}

func (c *Client) LocalReportMaterialGetV30Api() *api.LocalReportMaterialGetV30ApiService {
	return c.ApiClient.LocalReportMaterialGetV30Api
}

func (c *Client) LocalReportProjectGetV30Api() *api.LocalReportProjectGetV30ApiService {
	return c.ApiClient.LocalReportProjectGetV30Api
}

func (c *Client) LocalReportPromotionGetV30Api() *api.LocalReportPromotionGetV30ApiService {
	return c.ApiClient.LocalReportPromotionGetV30Api
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

func (c *Client) NativeAnchorDeleteV30Api() *api.NativeAnchorDeleteV30ApiService {
	return c.ApiClient.NativeAnchorDeleteV30Api
}

func (c *Client) NativeAnchorGetDetailV30Api() *api.NativeAnchorGetDetailV30ApiService {
	return c.ApiClient.NativeAnchorGetDetailV30Api
}

func (c *Client) NativeAnchorGetV30Api() *api.NativeAnchorGetV30ApiService {
	return c.ApiClient.NativeAnchorGetV30Api
}

func (c *Client) NativeAnchorQrcodePreviewGetV30Api() *api.NativeAnchorQrcodePreviewGetV30ApiService {
	return c.ApiClient.NativeAnchorQrcodePreviewGetV30Api
}

func (c *Client) NativeAnchorUpdateV30Api() *api.NativeAnchorUpdateV30ApiService {
	return c.ApiClient.NativeAnchorUpdateV30Api
}

func (c *Client) Oauth2AccessTokenApi() *api.Oauth2AccessTokenApiService {
	return c.ApiClient.Oauth2AccessTokenApi
}

func (c *Client) Oauth2AdvertiserGetApi() *api.Oauth2AdvertiserGetApiService {
	return c.ApiClient.Oauth2AdvertiserGetApi
}

func (c *Client) Oauth2AppAccessTokenApi() *api.Oauth2AppAccessTokenApiService {
	return c.ApiClient.Oauth2AppAccessTokenApi
}

func (c *Client) Oauth2RefreshTokenApi() *api.Oauth2RefreshTokenApiService {
	return c.ApiClient.Oauth2RefreshTokenApi
}

func (c *Client) Oauth2RenewTokenApi() *api.Oauth2RenewTokenApiService {
	return c.ApiClient.Oauth2RenewTokenApi
}

func (c *Client) OpenMaterialAuditProGetV30Api() *api.OpenMaterialAuditProGetV30ApiService {
	return c.ApiClient.OpenMaterialAuditProGetV30Api
}

func (c *Client) OpenMaterialAuditProSubmitV30Api() *api.OpenMaterialAuditProSubmitV30ApiService {
	return c.ApiClient.OpenMaterialAuditProSubmitV30Api
}

func (c *Client) PenaltyTaskGetV30Api() *api.PenaltyTaskGetV30ApiService {
	return c.ApiClient.PenaltyTaskGetV30Api
}

func (c *Client) ProjectBudgetUpdateV30Api() *api.ProjectBudgetUpdateV30ApiService {
	return c.ApiClient.ProjectBudgetUpdateV30Api
}

func (c *Client) ProjectCostProtectStatusGetV30Api() *api.ProjectCostProtectStatusGetV30ApiService {
	return c.ApiClient.ProjectCostProtectStatusGetV30Api
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

func (c *Client) ProjectNameUpdateV30Api() *api.ProjectNameUpdateV30ApiService {
	return c.ApiClient.ProjectNameUpdateV30Api
}

func (c *Client) ProjectRoigoalUpdateV30Api() *api.ProjectRoigoalUpdateV30ApiService {
	return c.ApiClient.ProjectRoigoalUpdateV30Api
}

func (c *Client) ProjectScheduleTimeUpdateV30Api() *api.ProjectScheduleTimeUpdateV30ApiService {
	return c.ApiClient.ProjectScheduleTimeUpdateV30Api
}

func (c *Client) ProjectStatusUpdateV30Api() *api.ProjectStatusUpdateV30ApiService {
	return c.ApiClient.ProjectStatusUpdateV30Api
}

func (c *Client) ProjectUpdateV30Api() *api.ProjectUpdateV30ApiService {
	return c.ApiClient.ProjectUpdateV30Api
}

func (c *Client) ProjectWeekScheduleUpdateV30Api() *api.ProjectWeekScheduleUpdateV30ApiService {
	return c.ApiClient.ProjectWeekScheduleUpdateV30Api
}

func (c *Client) PromotionAidGetV30Api() *api.PromotionAidGetV30ApiService {
	return c.ApiClient.PromotionAidGetV30Api
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

func (c *Client) PromotionEasyKeepDeliverySwitchUpdateV30Api() *api.PromotionEasyKeepDeliverySwitchUpdateV30ApiService {
	return c.ApiClient.PromotionEasyKeepDeliverySwitchUpdateV30Api
}

func (c *Client) PromotionEasyUpdateV30Api() *api.PromotionEasyUpdateV30ApiService {
	return c.ApiClient.PromotionEasyUpdateV30Api
}

func (c *Client) PromotionListV30Api() *api.PromotionListV30ApiService {
	return c.ApiClient.PromotionListV30Api
}

func (c *Client) PromotionMaterialDeleteV30Api() *api.PromotionMaterialDeleteV30ApiService {
	return c.ApiClient.PromotionMaterialDeleteV30Api
}

func (c *Client) PromotionNameUpdateV30Api() *api.PromotionNameUpdateV30ApiService {
	return c.ApiClient.PromotionNameUpdateV30Api
}

func (c *Client) PromotionNewcustomerCreateV30Api() *api.PromotionNewcustomerCreateV30ApiService {
	return c.ApiClient.PromotionNewcustomerCreateV30Api
}

func (c *Client) PromotionNewcustomerTypeGetV30Api() *api.PromotionNewcustomerTypeGetV30ApiService {
	return c.ApiClient.PromotionNewcustomerTypeGetV30Api
}

func (c *Client) PromotionRejectReasonGetV30Api() *api.PromotionRejectReasonGetV30ApiService {
	return c.ApiClient.PromotionRejectReasonGetV30Api
}

func (c *Client) PromotionScheduleTimeUpdateV30Api() *api.PromotionScheduleTimeUpdateV30ApiService {
	return c.ApiClient.PromotionScheduleTimeUpdateV30Api
}

func (c *Client) PromotionShopInfoUpdateV30Api() *api.PromotionShopInfoUpdateV30ApiService {
	return c.ApiClient.PromotionShopInfoUpdateV30Api
}

func (c *Client) PromotionStatusUpdateV30Api() *api.PromotionStatusUpdateV30ApiService {
	return c.ApiClient.PromotionStatusUpdateV30Api
}

func (c *Client) PromotionUpdateV30Api() *api.PromotionUpdateV30ApiService {
	return c.ApiClient.PromotionUpdateV30Api
}

func (c *Client) QianchuanAccountBalanceGetV10Api() *api.QianchuanAccountBalanceGetV10ApiService {
	return c.ApiClient.QianchuanAccountBalanceGetV10Api
}

func (c *Client) QianchuanAccountBudgetGetV10Api() *api.QianchuanAccountBudgetGetV10ApiService {
	return c.ApiClient.QianchuanAccountBudgetGetV10Api
}

func (c *Client) QianchuanAccountBudgetUpdateV10Api() *api.QianchuanAccountBudgetUpdateV10ApiService {
	return c.ApiClient.QianchuanAccountBudgetUpdateV10Api
}

func (c *Client) QianchuanAdBidUpdateV10Api() *api.QianchuanAdBidUpdateV10ApiService {
	return c.ApiClient.QianchuanAdBidUpdateV10Api
}

func (c *Client) QianchuanAdBudgetUpdateV10Api() *api.QianchuanAdBudgetUpdateV10ApiService {
	return c.ApiClient.QianchuanAdBudgetUpdateV10Api
}

func (c *Client) QianchuanAdCompensateStatusGetV10Api() *api.QianchuanAdCompensateStatusGetV10ApiService {
	return c.ApiClient.QianchuanAdCompensateStatusGetV10Api
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

func (c *Client) QianchuanAdLearingStatusGetV10Api() *api.QianchuanAdLearingStatusGetV10ApiService {
	return c.ApiClient.QianchuanAdLearingStatusGetV10Api
}

func (c *Client) QianchuanAdMaterialDeleteV10Api() *api.QianchuanAdMaterialDeleteV10ApiService {
	return c.ApiClient.QianchuanAdMaterialDeleteV10Api
}

func (c *Client) QianchuanAdMaterialGetV10Api() *api.QianchuanAdMaterialGetV10ApiService {
	return c.ApiClient.QianchuanAdMaterialGetV10Api
}

func (c *Client) QianchuanAdMaterialSuggestionV10Api() *api.QianchuanAdMaterialSuggestionV10ApiService {
	return c.ApiClient.QianchuanAdMaterialSuggestionV10Api
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

func (c *Client) QianchuanAdRegionUpdateV10Api() *api.QianchuanAdRegionUpdateV10ApiService {
	return c.ApiClient.QianchuanAdRegionUpdateV10Api
}

func (c *Client) QianchuanAdRejectReasonV10Api() *api.QianchuanAdRejectReasonV10ApiService {
	return c.ApiClient.QianchuanAdRejectReasonV10Api
}

func (c *Client) QianchuanAdScheduleDateUpdateV10Api() *api.QianchuanAdScheduleDateUpdateV10ApiService {
	return c.ApiClient.QianchuanAdScheduleDateUpdateV10Api
}

func (c *Client) QianchuanAdScheduleFixedRangeUpdateV10Api() *api.QianchuanAdScheduleFixedRangeUpdateV10ApiService {
	return c.ApiClient.QianchuanAdScheduleFixedRangeUpdateV10Api
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

func (c *Client) QianchuanAudienceCreateByFileV10Api() *api.QianchuanAudienceCreateByFileV10ApiService {
	return c.ApiClient.QianchuanAudienceCreateByFileV10Api
}

func (c *Client) QianchuanAudienceDeleteV10Api() *api.QianchuanAudienceDeleteV10ApiService {
	return c.ApiClient.QianchuanAudienceDeleteV10Api
}

func (c *Client) QianchuanAudienceFilePartUploadV10Api() *api.QianchuanAudienceFilePartUploadV10ApiService {
	return c.ApiClient.QianchuanAudienceFilePartUploadV10Api
}

func (c *Client) QianchuanAudienceFileUploadV10Api() *api.QianchuanAudienceFileUploadV10ApiService {
	return c.ApiClient.QianchuanAudienceFileUploadV10Api
}

func (c *Client) QianchuanAudienceGroupGetV10Api() *api.QianchuanAudienceGroupGetV10ApiService {
	return c.ApiClient.QianchuanAudienceGroupGetV10Api
}

func (c *Client) QianchuanAudienceListGetV10Api() *api.QianchuanAudienceListGetV10ApiService {
	return c.ApiClient.QianchuanAudienceListGetV10Api
}

func (c *Client) QianchuanAudiencePushV10Api() *api.QianchuanAudiencePushV10ApiService {
	return c.ApiClient.QianchuanAudiencePushV10Api
}

func (c *Client) QianchuanAwemeAuthListGetV10Api() *api.QianchuanAwemeAuthListGetV10ApiService {
	return c.ApiClient.QianchuanAwemeAuthListGetV10Api
}

func (c *Client) QianchuanAwemeAuthorizedGetV10Api() *api.QianchuanAwemeAuthorizedGetV10ApiService {
	return c.ApiClient.QianchuanAwemeAuthorizedGetV10Api
}

func (c *Client) QianchuanAwemeEstimateProfitV10Api() *api.QianchuanAwemeEstimateProfitV10ApiService {
	return c.ApiClient.QianchuanAwemeEstimateProfitV10Api
}

func (c *Client) QianchuanAwemeInterestActionInterestKeywordV10Api() *api.QianchuanAwemeInterestActionInterestKeywordV10ApiService {
	return c.ApiClient.QianchuanAwemeInterestActionInterestKeywordV10Api
}

func (c *Client) QianchuanAwemeOrderBudgetAddV10Api() *api.QianchuanAwemeOrderBudgetAddV10ApiService {
	return c.ApiClient.QianchuanAwemeOrderBudgetAddV10Api
}

func (c *Client) QianchuanAwemeOrderCreateV10Api() *api.QianchuanAwemeOrderCreateV10ApiService {
	return c.ApiClient.QianchuanAwemeOrderCreateV10Api
}

func (c *Client) QianchuanAwemeOrderDetailGetV10Api() *api.QianchuanAwemeOrderDetailGetV10ApiService {
	return c.ApiClient.QianchuanAwemeOrderDetailGetV10Api
}

func (c *Client) QianchuanAwemeOrderGetV10Api() *api.QianchuanAwemeOrderGetV10ApiService {
	return c.ApiClient.QianchuanAwemeOrderGetV10Api
}

func (c *Client) QianchuanAwemeOrderQuotaGetV10Api() *api.QianchuanAwemeOrderQuotaGetV10ApiService {
	return c.ApiClient.QianchuanAwemeOrderQuotaGetV10Api
}

func (c *Client) QianchuanAwemeOrderSuggestDeliveryTimeGetV10Api() *api.QianchuanAwemeOrderSuggestDeliveryTimeGetV10ApiService {
	return c.ApiClient.QianchuanAwemeOrderSuggestDeliveryTimeGetV10Api
}

func (c *Client) QianchuanAwemeOrderTerminateV10Api() *api.QianchuanAwemeOrderTerminateV10ApiService {
	return c.ApiClient.QianchuanAwemeOrderTerminateV10Api
}

func (c *Client) QianchuanAwemeProductAvailableGetV10Api() *api.QianchuanAwemeProductAvailableGetV10ApiService {
	return c.ApiClient.QianchuanAwemeProductAvailableGetV10Api
}

func (c *Client) QianchuanAwemeReportOrderGetV10Api() *api.QianchuanAwemeReportOrderGetV10ApiService {
	return c.ApiClient.QianchuanAwemeReportOrderGetV10Api
}

func (c *Client) QianchuanAwemeSuggestBidV10Api() *api.QianchuanAwemeSuggestBidV10ApiService {
	return c.ApiClient.QianchuanAwemeSuggestBidV10Api
}

func (c *Client) QianchuanAwemeSuggestRoiGoalV10Api() *api.QianchuanAwemeSuggestRoiGoalV10ApiService {
	return c.ApiClient.QianchuanAwemeSuggestRoiGoalV10Api
}

func (c *Client) QianchuanAwemeUniPromotionAdMaterialGetV10Api() *api.QianchuanAwemeUniPromotionAdMaterialGetV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionAdMaterialGetV10Api
}

func (c *Client) QianchuanAwemeUniPromotionEstimateEffectV10Api() *api.QianchuanAwemeUniPromotionEstimateEffectV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionEstimateEffectV10Api
}

func (c *Client) QianchuanAwemeUniPromotionOrderBudgetAddV10Api() *api.QianchuanAwemeUniPromotionOrderBudgetAddV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionOrderBudgetAddV10Api
}

func (c *Client) QianchuanAwemeUniPromotionOrderCreateV10Api() *api.QianchuanAwemeUniPromotionOrderCreateV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionOrderCreateV10Api
}

func (c *Client) QianchuanAwemeUniPromotionOrderDetailV10Api() *api.QianchuanAwemeUniPromotionOrderDetailV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionOrderDetailV10Api
}

func (c *Client) QianchuanAwemeUniPromotionOrderGetV10Api() *api.QianchuanAwemeUniPromotionOrderGetV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionOrderGetV10Api
}

func (c *Client) QianchuanAwemeUniPromotionOrderReportGetV10Api() *api.QianchuanAwemeUniPromotionOrderReportGetV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionOrderReportGetV10Api
}

func (c *Client) QianchuanAwemeUniPromotionOrderSuggestDeliveryTimeGetV10Api() *api.QianchuanAwemeUniPromotionOrderSuggestDeliveryTimeGetV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionOrderSuggestDeliveryTimeGetV10Api
}

func (c *Client) QianchuanAwemeUniPromotionReportV10Api() *api.QianchuanAwemeUniPromotionReportV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionReportV10Api
}

func (c *Client) QianchuanAwemeUniPromotionSuggestRoiV10Api() *api.QianchuanAwemeUniPromotionSuggestRoiV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionSuggestRoiV10Api
}

func (c *Client) QianchuanAwemeUniPromotionSuggestV10Api() *api.QianchuanAwemeUniPromotionSuggestV10ApiService {
	return c.ApiClient.QianchuanAwemeUniPromotionSuggestV10Api
}

func (c *Client) QianchuanAwemeVideoGetV10Api() *api.QianchuanAwemeVideoGetV10ApiService {
	return c.ApiClient.QianchuanAwemeVideoGetV10Api
}

func (c *Client) QianchuanBatchCampaignStatusUpdateV10Api() *api.QianchuanBatchCampaignStatusUpdateV10ApiService {
	return c.ApiClient.QianchuanBatchCampaignStatusUpdateV10Api
}

func (c *Client) QianchuanBrandAuthorizedGetV10Api() *api.QianchuanBrandAuthorizedGetV10ApiService {
	return c.ApiClient.QianchuanBrandAuthorizedGetV10Api
}

func (c *Client) QianchuanBrandReportAdGetV10Api() *api.QianchuanBrandReportAdGetV10ApiService {
	return c.ApiClient.QianchuanBrandReportAdGetV10Api
}

func (c *Client) QianchuanBrandReportAdvertiserGetV10Api() *api.QianchuanBrandReportAdvertiserGetV10ApiService {
	return c.ApiClient.QianchuanBrandReportAdvertiserGetV10Api
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

func (c *Client) QianchuanCarouselAwemeGetV10Api() *api.QianchuanCarouselAwemeGetV10ApiService {
	return c.ApiClient.QianchuanCarouselAwemeGetV10Api
}

func (c *Client) QianchuanCarouselGetV10Api() *api.QianchuanCarouselGetV10ApiService {
	return c.ApiClient.QianchuanCarouselGetV10Api
}

func (c *Client) QianchuanDmpAudiencesGetV10Api() *api.QianchuanDmpAudiencesGetV10ApiService {
	return c.ApiClient.QianchuanDmpAudiencesGetV10Api
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

func (c *Client) QianchuanKeywordCheckV10Api() *api.QianchuanKeywordCheckV10ApiService {
	return c.ApiClient.QianchuanKeywordCheckV10Api
}

func (c *Client) QianchuanKeywordPackageGetV10Api() *api.QianchuanKeywordPackageGetV10ApiService {
	return c.ApiClient.QianchuanKeywordPackageGetV10Api
}

func (c *Client) QianchuanLqAdGetV10Api() *api.QianchuanLqAdGetV10ApiService {
	return c.ApiClient.QianchuanLqAdGetV10Api
}

func (c *Client) QianchuanMaterialAdGetV10Api() *api.QianchuanMaterialAdGetV10ApiService {
	return c.ApiClient.QianchuanMaterialAdGetV10Api
}

func (c *Client) QianchuanMaterialGetV10Api() *api.QianchuanMaterialGetV10ApiService {
	return c.ApiClient.QianchuanMaterialGetV10Api
}

func (c *Client) QianchuanOrientationPackageGetV10Api() *api.QianchuanOrientationPackageGetV10ApiService {
	return c.ApiClient.QianchuanOrientationPackageGetV10Api
}

func (c *Client) QianchuanProductAnalyseCompareCreativeV10Api() *api.QianchuanProductAnalyseCompareCreativeV10ApiService {
	return c.ApiClient.QianchuanProductAnalyseCompareCreativeV10Api
}

func (c *Client) QianchuanProductAnalyseCompareStatsDataV10Api() *api.QianchuanProductAnalyseCompareStatsDataV10ApiService {
	return c.ApiClient.QianchuanProductAnalyseCompareStatsDataV10Api
}

func (c *Client) QianchuanProductAnalyseListV10Api() *api.QianchuanProductAnalyseListV10ApiService {
	return c.ApiClient.QianchuanProductAnalyseListV10Api
}

func (c *Client) QianchuanProductAvailableGetV10Api() *api.QianchuanProductAvailableGetV10ApiService {
	return c.ApiClient.QianchuanProductAvailableGetV10Api
}

func (c *Client) QianchuanQianchuanReportTodayLiveRoomConfigGetV10Api() *api.QianchuanQianchuanReportTodayLiveRoomConfigGetV10ApiService {
	return c.ApiClient.QianchuanQianchuanReportTodayLiveRoomConfigGetV10Api
}

func (c *Client) QianchuanQianchuanVideoStarGetV10Api() *api.QianchuanQianchuanVideoStarGetV10ApiService {
	return c.ApiClient.QianchuanQianchuanVideoStarGetV10Api
}

func (c *Client) QianchuanReportAdGetV10Api() *api.QianchuanReportAdGetV10ApiService {
	return c.ApiClient.QianchuanReportAdGetV10Api
}

func (c *Client) QianchuanReportAdMaterialGetV10Api() *api.QianchuanReportAdMaterialGetV10ApiService {
	return c.ApiClient.QianchuanReportAdMaterialGetV10Api
}

func (c *Client) QianchuanReportAdvertiserGetV10Api() *api.QianchuanReportAdvertiserGetV10ApiService {
	return c.ApiClient.QianchuanReportAdvertiserGetV10Api
}

func (c *Client) QianchuanReportCustomConfigGetV10Api() *api.QianchuanReportCustomConfigGetV10ApiService {
	return c.ApiClient.QianchuanReportCustomConfigGetV10Api
}

func (c *Client) QianchuanReportCustomGetV10Api() *api.QianchuanReportCustomGetV10ApiService {
	return c.ApiClient.QianchuanReportCustomGetV10Api
}

func (c *Client) QianchuanReportLiveGetV10Api() *api.QianchuanReportLiveGetV10ApiService {
	return c.ApiClient.QianchuanReportLiveGetV10Api
}

func (c *Client) QianchuanReportLongTransferOrderConfigGetV10Api() *api.QianchuanReportLongTransferOrderConfigGetV10ApiService {
	return c.ApiClient.QianchuanReportLongTransferOrderConfigGetV10Api
}

func (c *Client) QianchuanReportLongTransferOrderDataGetV10Api() *api.QianchuanReportLongTransferOrderDataGetV10ApiService {
	return c.ApiClient.QianchuanReportLongTransferOrderDataGetV10Api
}

func (c *Client) QianchuanReportLongTransferOrderGetV10Api() *api.QianchuanReportLongTransferOrderGetV10ApiService {
	return c.ApiClient.QianchuanReportLongTransferOrderGetV10Api
}

func (c *Client) QianchuanReportMaterialGetV10Api() *api.QianchuanReportMaterialGetV10ApiService {
	return c.ApiClient.QianchuanReportMaterialGetV10Api
}

func (c *Client) QianchuanReportSearchWordGetV10Api() *api.QianchuanReportSearchWordGetV10ApiService {
	return c.ApiClient.QianchuanReportSearchWordGetV10Api
}

func (c *Client) QianchuanReportTodayLiveGetV10Api() *api.QianchuanReportTodayLiveGetV10ApiService {
	return c.ApiClient.QianchuanReportTodayLiveGetV10Api
}

func (c *Client) QianchuanReportTodayLiveRoomConfigGetV10Api() *api.QianchuanReportTodayLiveRoomConfigGetV10ApiService {
	return c.ApiClient.QianchuanReportTodayLiveRoomConfigGetV10Api
}

func (c *Client) QianchuanReportTodayLiveRoomDataGetV10Api() *api.QianchuanReportTodayLiveRoomDataGetV10ApiService {
	return c.ApiClient.QianchuanReportTodayLiveRoomDataGetV10Api
}

func (c *Client) QianchuanReportUniPromotionConfigGetV10Api() *api.QianchuanReportUniPromotionConfigGetV10ApiService {
	return c.ApiClient.QianchuanReportUniPromotionConfigGetV10Api
}

func (c *Client) QianchuanReportUniPromotionDataGetV10Api() *api.QianchuanReportUniPromotionDataGetV10ApiService {
	return c.ApiClient.QianchuanReportUniPromotionDataGetV10Api
}

func (c *Client) QianchuanReportUniPromotionDimensionDataAuthorGetV10Api() *api.QianchuanReportUniPromotionDimensionDataAuthorGetV10ApiService {
	return c.ApiClient.QianchuanReportUniPromotionDimensionDataAuthorGetV10Api
}

func (c *Client) QianchuanReportUniPromotionDimensionDataRoomGetV10Api() *api.QianchuanReportUniPromotionDimensionDataRoomGetV10ApiService {
	return c.ApiClient.QianchuanReportUniPromotionDimensionDataRoomGetV10Api
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

func (c *Client) QianchuanShopAuthorizedGetV10Api() *api.QianchuanShopAuthorizedGetV10ApiService {
	return c.ApiClient.QianchuanShopAuthorizedGetV10Api
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

func (c *Client) QianchuanTodayLiveRoomDetailGetV10Api() *api.QianchuanTodayLiveRoomDetailGetV10ApiService {
	return c.ApiClient.QianchuanTodayLiveRoomDetailGetV10Api
}

func (c *Client) QianchuanTodayLiveRoomFlowPerformanceGetV10Api() *api.QianchuanTodayLiveRoomFlowPerformanceGetV10ApiService {
	return c.ApiClient.QianchuanTodayLiveRoomFlowPerformanceGetV10Api
}

func (c *Client) QianchuanTodayLiveRoomGetV10Api() *api.QianchuanTodayLiveRoomGetV10ApiService {
	return c.ApiClient.QianchuanTodayLiveRoomGetV10Api
}

func (c *Client) QianchuanTodayLiveRoomProductListGetV10Api() *api.QianchuanTodayLiveRoomProductListGetV10ApiService {
	return c.ApiClient.QianchuanTodayLiveRoomProductListGetV10Api
}

func (c *Client) QianchuanTodayLiveRoomUserGetV10Api() *api.QianchuanTodayLiveRoomUserGetV10ApiService {
	return c.ApiClient.QianchuanTodayLiveRoomUserGetV10Api
}

func (c *Client) QianchuanTodayLiveV2RoomUserGetV10Api() *api.QianchuanTodayLiveV2RoomUserGetV10ApiService {
	return c.ApiClient.QianchuanTodayLiveV2RoomUserGetV10Api
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

func (c *Client) QianchuanToolsLogSearchV10Api() *api.QianchuanToolsLogSearchV10ApiService {
	return c.ApiClient.QianchuanToolsLogSearchV10Api
}

func (c *Client) QianchuanToolsShopAuthV10Api() *api.QianchuanToolsShopAuthV10ApiService {
	return c.ApiClient.QianchuanToolsShopAuthV10Api
}

func (c *Client) QianchuanToolsSmartBoostAdBoostReportGetV10Api() *api.QianchuanToolsSmartBoostAdBoostReportGetV10ApiService {
	return c.ApiClient.QianchuanToolsSmartBoostAdBoostReportGetV10Api
}

func (c *Client) QianchuanToolsSmartBoostAdBoostSetV10Api() *api.QianchuanToolsSmartBoostAdBoostSetV10ApiService {
	return c.ApiClient.QianchuanToolsSmartBoostAdBoostSetV10Api
}

func (c *Client) QianchuanToolsSmartBoostAdBoostStatusGetV10Api() *api.QianchuanToolsSmartBoostAdBoostStatusGetV10ApiService {
	return c.ApiClient.QianchuanToolsSmartBoostAdBoostStatusGetV10Api
}

func (c *Client) QianchuanToolsSmartBoostAdBoostVersionGetV10Api() *api.QianchuanToolsSmartBoostAdBoostVersionGetV10ApiService {
	return c.ApiClient.QianchuanToolsSmartBoostAdBoostVersionGetV10Api
}

func (c *Client) QianchuanTrackUrlCheckV10Api() *api.QianchuanTrackUrlCheckV10ApiService {
	return c.ApiClient.QianchuanTrackUrlCheckV10Api
}

func (c *Client) QianchuanUniAwemeAdCreateV10Api() *api.QianchuanUniAwemeAdCreateV10ApiService {
	return c.ApiClient.QianchuanUniAwemeAdCreateV10Api
}

func (c *Client) QianchuanUniAwemeAdUpdateV10Api() *api.QianchuanUniAwemeAdUpdateV10ApiService {
	return c.ApiClient.QianchuanUniAwemeAdUpdateV10Api
}

func (c *Client) QianchuanUniAwemeAuthorizedGetV10Api() *api.QianchuanUniAwemeAuthorizedGetV10ApiService {
	return c.ApiClient.QianchuanUniAwemeAuthorizedGetV10Api
}

func (c *Client) QianchuanUniAwemeSuggestBudgetV10Api() *api.QianchuanUniAwemeSuggestBudgetV10ApiService {
	return c.ApiClient.QianchuanUniAwemeSuggestBudgetV10Api
}

func (c *Client) QianchuanUniPromotionAdBudgetUpdateV10Api() *api.QianchuanUniPromotionAdBudgetUpdateV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdBudgetUpdateV10Api
}

func (c *Client) QianchuanUniPromotionAdControlTaskCreateV10Api() *api.QianchuanUniPromotionAdControlTaskCreateV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdControlTaskCreateV10Api
}

func (c *Client) QianchuanUniPromotionAdControlTaskListV10Api() *api.QianchuanUniPromotionAdControlTaskListV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdControlTaskListV10Api
}

func (c *Client) QianchuanUniPromotionAdControlTaskStatusUpdateV10Api() *api.QianchuanUniPromotionAdControlTaskStatusUpdateV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdControlTaskStatusUpdateV10Api
}

func (c *Client) QianchuanUniPromotionAdDetailV10Api() *api.QianchuanUniPromotionAdDetailV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdDetailV10Api
}

func (c *Client) QianchuanUniPromotionAdMaterialAddV10Api() *api.QianchuanUniPromotionAdMaterialAddV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdMaterialAddV10Api
}

func (c *Client) QianchuanUniPromotionAdMaterialDeleteV10Api() *api.QianchuanUniPromotionAdMaterialDeleteV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdMaterialDeleteV10Api
}

func (c *Client) QianchuanUniPromotionAdMaterialGetV10Api() *api.QianchuanUniPromotionAdMaterialGetV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdMaterialGetV10Api
}

func (c *Client) QianchuanUniPromotionAdNameUpdateV10Api() *api.QianchuanUniPromotionAdNameUpdateV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdNameUpdateV10Api
}

func (c *Client) QianchuanUniPromotionAdProductDeleteV10Api() *api.QianchuanUniPromotionAdProductDeleteV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdProductDeleteV10Api
}

func (c *Client) QianchuanUniPromotionAdProductGetV10Api() *api.QianchuanUniPromotionAdProductGetV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdProductGetV10Api
}

func (c *Client) QianchuanUniPromotionAdRoi2GoalUpdateV10Api() *api.QianchuanUniPromotionAdRoi2GoalUpdateV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdRoi2GoalUpdateV10Api
}

func (c *Client) QianchuanUniPromotionAdScheduleDateUpdateV10Api() *api.QianchuanUniPromotionAdScheduleDateUpdateV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdScheduleDateUpdateV10Api
}

func (c *Client) QianchuanUniPromotionAdStatusUpdateV10Api() *api.QianchuanUniPromotionAdStatusUpdateV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdStatusUpdateV10Api
}

func (c *Client) QianchuanUniPromotionAdSuggestionV10Api() *api.QianchuanUniPromotionAdSuggestionV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAdSuggestionV10Api
}

func (c *Client) QianchuanUniPromotionAuthInitV10Api() *api.QianchuanUniPromotionAuthInitV10ApiService {
	return c.ApiClient.QianchuanUniPromotionAuthInitV10Api
}

func (c *Client) QianchuanUniPromotionBlockMaterialGetV10Api() *api.QianchuanUniPromotionBlockMaterialGetV10ApiService {
	return c.ApiClient.QianchuanUniPromotionBlockMaterialGetV10Api
}

func (c *Client) QianchuanUniPromotionListV10Api() *api.QianchuanUniPromotionListV10ApiService {
	return c.ApiClient.QianchuanUniPromotionListV10Api
}

func (c *Client) QianchuanUniPromotionProductAwemeGetV10Api() *api.QianchuanUniPromotionProductAwemeGetV10ApiService {
	return c.ApiClient.QianchuanUniPromotionProductAwemeGetV10Api
}

func (c *Client) QianchuanUniPromotionProductGetV10Api() *api.QianchuanUniPromotionProductGetV10ApiService {
	return c.ApiClient.QianchuanUniPromotionProductGetV10Api
}

func (c *Client) QianchuanVideoByAwemeGetV10Api() *api.QianchuanVideoByAwemeGetV10ApiService {
	return c.ApiClient.QianchuanVideoByAwemeGetV10Api
}

func (c *Client) QianchuanVideoGetV10Api() *api.QianchuanVideoGetV10ApiService {
	return c.ApiClient.QianchuanVideoGetV10Api
}

func (c *Client) QueryBookingBusinessEntityIdGetV2Api() *api.QueryBookingBusinessEntityIdGetV2ApiService {
	return c.ApiClient.QueryBookingBusinessEntityIdGetV2Api
}

func (c *Client) QueryInvoiceDetailSelfV2Api() *api.QueryInvoiceDetailSelfV2ApiService {
	return c.ApiClient.QueryInvoiceDetailSelfV2Api
}

func (c *Client) QueryInvoiceDetailV2Api() *api.QueryInvoiceDetailV2ApiService {
	return c.ApiClient.QueryInvoiceDetailV2Api
}

func (c *Client) QueryInvoiceElectronicUrlSelfV2Api() *api.QueryInvoiceElectronicUrlSelfV2ApiService {
	return c.ApiClient.QueryInvoiceElectronicUrlSelfV2Api
}

func (c *Client) QueryInvoiceElectronicUrlV2Api() *api.QueryInvoiceElectronicUrlV2ApiService {
	return c.ApiClient.QueryInvoiceElectronicUrlV2Api
}

func (c *Client) QueryInvoiceSelfV2Api() *api.QueryInvoiceSelfV2ApiService {
	return c.ApiClient.QueryInvoiceSelfV2Api
}

func (c *Client) QueryInvoiceTaxV2Api() *api.QueryInvoiceTaxV2ApiService {
	return c.ApiClient.QueryInvoiceTaxV2Api
}

func (c *Client) QueryInvoiceV2Api() *api.QueryInvoiceV2ApiService {
	return c.ApiClient.QueryInvoiceV2Api
}

func (c *Client) QueryProjectV2Api() *api.QueryProjectV2ApiService {
	return c.ApiClient.QueryProjectV2Api
}

func (c *Client) QueryProjectV30Api() *api.QueryProjectV30ApiService {
	return c.ApiClient.QueryProjectV30Api
}

func (c *Client) QueryRebateAccountingInfoV2Api() *api.QueryRebateAccountingInfoV2ApiService {
	return c.ApiClient.QueryRebateAccountingInfoV2Api
}

func (c *Client) QueryRebateBalanceV2Api() *api.QueryRebateBalanceV2ApiService {
	return c.ApiClient.QueryRebateBalanceV2Api
}

func (c *Client) QueryStatementV2Api() *api.QueryStatementV2ApiService {
	return c.ApiClient.QueryStatementV2Api
}

func (c *Client) RecommendVideoListV30Api() *api.RecommendVideoListV30ApiService {
	return c.ApiClient.RecommendVideoListV30Api
}

func (c *Client) RejectMaterialAiRepairAcceptTaskCreateV30Api() *api.RejectMaterialAiRepairAcceptTaskCreateV30ApiService {
	return c.ApiClient.RejectMaterialAiRepairAcceptTaskCreateV30Api
}

func (c *Client) RejectMaterialAiRepairAcceptTaskListV30Api() *api.RejectMaterialAiRepairAcceptTaskListV30ApiService {
	return c.ApiClient.RejectMaterialAiRepairAcceptTaskListV30Api
}

func (c *Client) RejectMaterialAiRepairCrossAccountGetV30Api() *api.RejectMaterialAiRepairCrossAccountGetV30ApiService {
	return c.ApiClient.RejectMaterialAiRepairCrossAccountGetV30Api
}

func (c *Client) RejectMaterialAiRepairGetV30Api() *api.RejectMaterialAiRepairGetV30ApiService {
	return c.ApiClient.RejectMaterialAiRepairGetV30Api
}

func (c *Client) RemittanceCodeListV30Api() *api.RemittanceCodeListV30ApiService {
	return c.ApiClient.RemittanceCodeListV30Api
}

func (c *Client) ReportAdGetV2Api() *api.ReportAdGetV2ApiService {
	return c.ApiClient.ReportAdGetV2Api
}

func (c *Client) ReportAdvertiserGetV2Api() *api.ReportAdvertiserGetV2ApiService {
	return c.ApiClient.ReportAdvertiserGetV2Api
}

func (c *Client) ReportAgentGetV2V2Api() *api.ReportAgentGetV2V2ApiService {
	return c.ApiClient.ReportAgentGetV2V2Api
}

func (c *Client) ReportAudienceAgeV2Api() *api.ReportAudienceAgeV2ApiService {
	return c.ApiClient.ReportAudienceAgeV2Api
}

func (c *Client) ReportAudienceAwemeListV2Api() *api.ReportAudienceAwemeListV2ApiService {
	return c.ApiClient.ReportAudienceAwemeListV2Api
}

func (c *Client) ReportAudienceCityV2Api() *api.ReportAudienceCityV2ApiService {
	return c.ApiClient.ReportAudienceCityV2Api
}

func (c *Client) ReportAudienceGenderV2Api() *api.ReportAudienceGenderV2ApiService {
	return c.ApiClient.ReportAudienceGenderV2Api
}

func (c *Client) ReportAudienceInterestActionListV2Api() *api.ReportAudienceInterestActionListV2ApiService {
	return c.ApiClient.ReportAudienceInterestActionListV2Api
}

func (c *Client) ReportAudienceProvinceV2Api() *api.ReportAudienceProvinceV2ApiService {
	return c.ApiClient.ReportAudienceProvinceV2Api
}

func (c *Client) ReportBrandAdGetV30Api() *api.ReportBrandAdGetV30ApiService {
	return c.ApiClient.ReportBrandAdGetV30Api
}

func (c *Client) ReportBrandAdvertiserGetV30Api() *api.ReportBrandAdvertiserGetV30ApiService {
	return c.ApiClient.ReportBrandAdvertiserGetV30Api
}

func (c *Client) ReportBrandAgentDataV30Api() *api.ReportBrandAgentDataV30ApiService {
	return c.ApiClient.ReportBrandAgentDataV30Api
}

func (c *Client) ReportBrandCampaignGetV30Api() *api.ReportBrandCampaignGetV30ApiService {
	return c.ApiClient.ReportBrandCampaignGetV30Api
}

func (c *Client) ReportBrandCreativeGetV30Api() *api.ReportBrandCreativeGetV30ApiService {
	return c.ApiClient.ReportBrandCreativeGetV30Api
}

func (c *Client) ReportBrandDataV30Api() *api.ReportBrandDataV30ApiService {
	return c.ApiClient.ReportBrandDataV30Api
}

func (c *Client) ReportBusinessPlatformStardeliveryTaskVideoDataGetV30Api() *api.ReportBusinessPlatformStardeliveryTaskVideoDataGetV30ApiService {
	return c.ApiClient.ReportBusinessPlatformStardeliveryTaskVideoDataGetV30Api
}

func (c *Client) ReportCampaignGetV2Api() *api.ReportCampaignGetV2ApiService {
	return c.ApiClient.ReportCampaignGetV2Api
}

func (c *Client) ReportCreativeGetV2Api() *api.ReportCreativeGetV2ApiService {
	return c.ApiClient.ReportCreativeGetV2Api
}

func (c *Client) ReportCustomAsyncTaskCreateV30Api() *api.ReportCustomAsyncTaskCreateV30ApiService {
	return c.ApiClient.ReportCustomAsyncTaskCreateV30Api
}

func (c *Client) ReportCustomAsyncTaskDownloadV30Api() *api.ReportCustomAsyncTaskDownloadV30ApiService {
	return c.ApiClient.ReportCustomAsyncTaskDownloadV30Api
}

func (c *Client) ReportCustomAsyncTaskGetV30Api() *api.ReportCustomAsyncTaskGetV30ApiService {
	return c.ApiClient.ReportCustomAsyncTaskGetV30Api
}

func (c *Client) ReportCustomConfigGetV30Api() *api.ReportCustomConfigGetV30ApiService {
	return c.ApiClient.ReportCustomConfigGetV30Api
}

func (c *Client) ReportCustomCreativeGetV30Api() *api.ReportCustomCreativeGetV30ApiService {
	return c.ApiClient.ReportCustomCreativeGetV30Api
}

func (c *Client) ReportCustomGetV30Api() *api.ReportCustomGetV30ApiService {
	return c.ApiClient.ReportCustomGetV30Api
}

func (c *Client) ReportLiveRoomAnalysisGetV2Api() *api.ReportLiveRoomAnalysisGetV2ApiService {
	return c.ApiClient.ReportLiveRoomAnalysisGetV2Api
}

func (c *Client) ReportLiveRoomAttributeGetV2Api() *api.ReportLiveRoomAttributeGetV2ApiService {
	return c.ApiClient.ReportLiveRoomAttributeGetV2Api
}

func (c *Client) ReportLiveRoomAudiencePortraitGetV2Api() *api.ReportLiveRoomAudiencePortraitGetV2ApiService {
	return c.ApiClient.ReportLiveRoomAudiencePortraitGetV2Api
}

func (c *Client) ReportLiveRoomFlowCategoryGetV2Api() *api.ReportLiveRoomFlowCategoryGetV2ApiService {
	return c.ApiClient.ReportLiveRoomFlowCategoryGetV2Api
}

func (c *Client) ReportProductAsyncTaskDownloadV30Api() *api.ReportProductAsyncTaskDownloadV30ApiService {
	return c.ApiClient.ReportProductAsyncTaskDownloadV30Api
}

func (c *Client) ReportProductAsyncTaskGetV30Api() *api.ReportProductAsyncTaskGetV30ApiService {
	return c.ApiClient.ReportProductAsyncTaskGetV30Api
}

func (c *Client) ReportProductDailyAsyncTaskCreateV30Api() *api.ReportProductDailyAsyncTaskCreateV30ApiService {
	return c.ApiClient.ReportProductDailyAsyncTaskCreateV30Api
}

func (c *Client) ReportProductHourlyAsyncTaskCreateV30Api() *api.ReportProductHourlyAsyncTaskCreateV30ApiService {
	return c.ApiClient.ReportProductHourlyAsyncTaskCreateV30Api
}

func (c *Client) ReportRtaCusExpGetV2Api() *api.ReportRtaCusExpGetV2ApiService {
	return c.ApiClient.ReportRtaCusExpGetV2Api
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

func (c *Client) ReportRtaGetV2Api() *api.ReportRtaGetV2ApiService {
	return c.ApiClient.ReportRtaGetV2Api
}

func (c *Client) ReportRubeexGetV2Api() *api.ReportRubeexGetV2ApiService {
	return c.ApiClient.ReportRubeexGetV2Api
}

func (c *Client) ReportSitePageV2Api() *api.ReportSitePageV2ApiService {
	return c.ApiClient.ReportSitePageV2Api
}

func (c *Client) ReportStardeliveryTaskDataGetV30Api() *api.ReportStardeliveryTaskDataGetV30ApiService {
	return c.ApiClient.ReportStardeliveryTaskDataGetV30Api
}

func (c *Client) ReportStardeliveryTaskVideoDataGetV30Api() *api.ReportStardeliveryTaskVideoDataGetV30ApiService {
	return c.ApiClient.ReportStardeliveryTaskVideoDataGetV30Api
}

func (c *Client) ReportVideoFrameGetV2Api() *api.ReportVideoFrameGetV2ApiService {
	return c.ApiClient.ReportVideoFrameGetV2Api
}

func (c *Client) SecurityAuditResultsV30Api() *api.SecurityAuditResultsV30ApiService {
	return c.ApiClient.SecurityAuditResultsV30Api
}

func (c *Client) SecurityOpenMaterialAuditV30Api() *api.SecurityOpenMaterialAuditV30ApiService {
	return c.ApiClient.SecurityOpenMaterialAuditV30Api
}

func (c *Client) SecurityScoreDisposalInfoGetV30Api() *api.SecurityScoreDisposalInfoGetV30ApiService {
	return c.ApiClient.SecurityScoreDisposalInfoGetV30Api
}

func (c *Client) SecurityScoreTotalGetV30Api() *api.SecurityScoreTotalGetV30ApiService {
	return c.ApiClient.SecurityScoreTotalGetV30Api
}

func (c *Client) SecurityScoreViolationEventGetV30Api() *api.SecurityScoreViolationEventGetV30ApiService {
	return c.ApiClient.SecurityScoreViolationEventGetV30Api
}

func (c *Client) ServeMarketActiveFuncGetV10Api() *api.ServeMarketActiveFuncGetV10ApiService {
	return c.ApiClient.ServeMarketActiveFuncGetV10Api
}

func (c *Client) ServeMarketCidVerifyTokenV10Api() *api.ServeMarketCidVerifyTokenV10ApiService {
	return c.ApiClient.ServeMarketCidVerifyTokenV10Api
}

func (c *Client) ServeMarketOrderGetV10Api() *api.ServeMarketOrderGetV10ApiService {
	return c.ApiClient.ServeMarketOrderGetV10Api
}

func (c *Client) SharedWalletAccountRelationGetV30Api() *api.SharedWalletAccountRelationGetV30ApiService {
	return c.ApiClient.SharedWalletAccountRelationGetV30Api
}

func (c *Client) SharedWalletBudgetGetV30Api() *api.SharedWalletBudgetGetV30ApiService {
	return c.ApiClient.SharedWalletBudgetGetV30Api
}

func (c *Client) SharedWalletBudgetSubmitV30Api() *api.SharedWalletBudgetSubmitV30ApiService {
	return c.ApiClient.SharedWalletBudgetSubmitV30Api
}

func (c *Client) SharedWalletDailyStatGetV30Api() *api.SharedWalletDailyStatGetV30ApiService {
	return c.ApiClient.SharedWalletDailyStatGetV30Api
}

func (c *Client) SharedWalletMainWalletGetV30Api() *api.SharedWalletMainWalletGetV30ApiService {
	return c.ApiClient.SharedWalletMainWalletGetV30Api
}

func (c *Client) SharedWalletTransactionDetailGetV30Api() *api.SharedWalletTransactionDetailGetV30ApiService {
	return c.ApiClient.SharedWalletTransactionDetailGetV30Api
}

func (c *Client) SharedWalletWalletBalanceGetV30Api() *api.SharedWalletWalletBalanceGetV30ApiService {
	return c.ApiClient.SharedWalletWalletBalanceGetV30Api
}

func (c *Client) SharedWalletWalletInfoGetV30Api() *api.SharedWalletWalletInfoGetV30ApiService {
	return c.ApiClient.SharedWalletWalletInfoGetV30Api
}

func (c *Client) SharedWalletWalletRelationGetV30Api() *api.SharedWalletWalletRelationGetV30ApiService {
	return c.ApiClient.SharedWalletWalletRelationGetV30Api
}

func (c *Client) SharedWalletWatchRuleGetV30Api() *api.SharedWalletWatchRuleGetV30ApiService {
	return c.ApiClient.SharedWalletWatchRuleGetV30Api
}

func (c *Client) SharedWalletWatchRuleSubmitV30Api() *api.SharedWalletWatchRuleSubmitV30ApiService {
	return c.ApiClient.SharedWalletWatchRuleSubmitV30Api
}

func (c *Client) ShopBonusCreateV30Api() *api.ShopBonusCreateV30ApiService {
	return c.ApiClient.ShopBonusCreateV30Api
}

func (c *Client) ShopBonusSuccessGetV30Api() *api.ShopBonusSuccessGetV30ApiService {
	return c.ApiClient.ShopBonusSuccessGetV30Api
}

func (c *Client) SpiTaskGetV2Api() *api.SpiTaskGetV2ApiService {
	return c.ApiClient.SpiTaskGetV2Api
}

func (c *Client) StarAttachmentUploadV2Api() *api.StarAttachmentUploadV2ApiService {
	return c.ApiClient.StarAttachmentUploadV2Api
}

func (c *Client) StarAttributeItemEventFeedbackV2Api() *api.StarAttributeItemEventFeedbackV2ApiService {
	return c.ApiClient.StarAttributeItemEventFeedbackV2Api
}

func (c *Client) StarAttributeJdOverflowConvertV2Api() *api.StarAttributeJdOverflowConvertV2ApiService {
	return c.ApiClient.StarAttributeJdOverflowConvertV2Api
}

func (c *Client) StarAttributeUserEventFeedbackV2Api() *api.StarAttributeUserEventFeedbackV2ApiService {
	return c.ApiClient.StarAttributeUserEventFeedbackV2Api
}

func (c *Client) StarAttributeUserInfoFeedbackV2Api() *api.StarAttributeUserInfoFeedbackV2ApiService {
	return c.ApiClient.StarAttributeUserInfoFeedbackV2Api
}

func (c *Client) StarBillGetPendingV2Api() *api.StarBillGetPendingV2ApiService {
	return c.ApiClient.StarBillGetPendingV2Api
}

func (c *Client) StarBillPayV2Api() *api.StarBillPayV2ApiService {
	return c.ApiClient.StarBillPayV2Api
}

func (c *Client) StarBrandCategoryListV2Api() *api.StarBrandCategoryListV2ApiService {
	return c.ApiClient.StarBrandCategoryListV2Api
}

func (c *Client) StarBrandListV2Api() *api.StarBrandListV2ApiService {
	return c.ApiClient.StarBrandListV2Api
}

func (c *Client) StarCampaignListV2Api() *api.StarCampaignListV2ApiService {
	return c.ApiClient.StarCampaignListV2Api
}

func (c *Client) StarChallengeAddBudgetV2Api() *api.StarChallengeAddBudgetV2ApiService {
	return c.ApiClient.StarChallengeAddBudgetV2Api
}

func (c *Client) StarChallengeAuthorListV2Api() *api.StarChallengeAuthorListV2ApiService {
	return c.ApiClient.StarChallengeAuthorListV2Api
}

func (c *Client) StarChallengeCancelV2Api() *api.StarChallengeCancelV2ApiService {
	return c.ApiClient.StarChallengeCancelV2Api
}

func (c *Client) StarChallengeChooseTaskItemWithRewardV2Api() *api.StarChallengeChooseTaskItemWithRewardV2ApiService {
	return c.ApiClient.StarChallengeChooseTaskItemWithRewardV2Api
}

func (c *Client) StarChallengeExpandRangeV2Api() *api.StarChallengeExpandRangeV2ApiService {
	return c.ApiClient.StarChallengeExpandRangeV2Api
}

func (c *Client) StarChallengeGetCustomTaskDataV2Api() *api.StarChallengeGetCustomTaskDataV2ApiService {
	return c.ApiClient.StarChallengeGetCustomTaskDataV2Api
}

func (c *Client) StarChallengeGetCustomTaskListV2Api() *api.StarChallengeGetCustomTaskListV2ApiService {
	return c.ApiClient.StarChallengeGetCustomTaskListV2Api
}

func (c *Client) StarChallengeGetPushAdResultsV2Api() *api.StarChallengeGetPushAdResultsV2ApiService {
	return c.ApiClient.StarChallengeGetPushAdResultsV2Api
}

func (c *Client) StarChallengeInfoV2Api() *api.StarChallengeInfoV2ApiService {
	return c.ApiClient.StarChallengeInfoV2Api
}

func (c *Client) StarChallengeItemsDataV2Api() *api.StarChallengeItemsDataV2ApiService {
	return c.ApiClient.StarChallengeItemsDataV2Api
}

func (c *Client) StarChallengeListV2Api() *api.StarChallengeListV2ApiService {
	return c.ApiClient.StarChallengeListV2Api
}

func (c *Client) StarChallengePushItemsToAdV2Api() *api.StarChallengePushItemsToAdV2ApiService {
	return c.ApiClient.StarChallengePushItemsToAdV2Api
}

func (c *Client) StarChallengeSyncItemToAdV2Api() *api.StarChallengeSyncItemToAdV2ApiService {
	return c.ApiClient.StarChallengeSyncItemToAdV2Api
}

func (c *Client) StarClueGetV2Api() *api.StarClueGetV2ApiService {
	return c.ApiClient.StarClueGetV2Api
}

func (c *Client) StarComponentCreateCommonComponentV2Api() *api.StarComponentCreateCommonComponentV2ApiService {
	return c.ApiClient.StarComponentCreateCommonComponentV2Api
}

func (c *Client) StarComponentCreateLinkV2Api() *api.StarComponentCreateLinkV2ApiService {
	return c.ApiClient.StarComponentCreateLinkV2Api
}

func (c *Client) StarComponentQueryIndustryAnchorV2Api() *api.StarComponentQueryIndustryAnchorV2ApiService {
	return c.ApiClient.StarComponentQueryIndustryAnchorV2Api
}

func (c *Client) StarComponentQueryLinkV2Api() *api.StarComponentQueryLinkV2ApiService {
	return c.ApiClient.StarComponentQueryLinkV2Api
}

func (c *Client) StarComponentUpdateCommonComponentV2Api() *api.StarComponentUpdateCommonComponentV2ApiService {
	return c.ApiClient.StarComponentUpdateCommonComponentV2Api
}

func (c *Client) StarComponentUpdateLinkV2Api() *api.StarComponentUpdateLinkV2ApiService {
	return c.ApiClient.StarComponentUpdateLinkV2Api
}

func (c *Client) StarCreateProjectV2Api() *api.StarCreateProjectV2ApiService {
	return c.ApiClient.StarCreateProjectV2Api
}

func (c *Client) StarDataTaskTimelineReportV2Api() *api.StarDataTaskTimelineReportV2ApiService {
	return c.ApiClient.StarDataTaskTimelineReportV2Api
}

func (c *Client) StarDemandCreateAssignV2Api() *api.StarDemandCreateAssignV2ApiService {
	return c.ApiClient.StarDemandCreateAssignV2Api
}

func (c *Client) StarDemandCreateChallengeV2Api() *api.StarDemandCreateChallengeV2ApiService {
	return c.ApiClient.StarDemandCreateChallengeV2Api
}

func (c *Client) StarDemandGetEpisodeLimitV2Api() *api.StarDemandGetEpisodeLimitV2ApiService {
	return c.ApiClient.StarDemandGetEpisodeLimitV2Api
}

func (c *Client) StarDemandGetResellCodeV2Api() *api.StarDemandGetResellCodeV2ApiService {
	return c.ApiClient.StarDemandGetResellCodeV2Api
}

func (c *Client) StarDemandListV2Api() *api.StarDemandListV2ApiService {
	return c.ApiClient.StarDemandListV2Api
}

func (c *Client) StarDemandOmCreateChallengeV2Api() *api.StarDemandOmCreateChallengeV2ApiService {
	return c.ApiClient.StarDemandOmCreateChallengeV2Api
}

func (c *Client) StarDemandOmExpandChallengeProviderV2Api() *api.StarDemandOmExpandChallengeProviderV2ApiService {
	return c.ApiClient.StarDemandOmExpandChallengeProviderV2Api
}

func (c *Client) StarDemandOmExpandChallengeV2Api() *api.StarDemandOmExpandChallengeV2ApiService {
	return c.ApiClient.StarDemandOmExpandChallengeV2Api
}

func (c *Client) StarDemandOmGetChallengeDispatchedProviderListV2Api() *api.StarDemandOmGetChallengeDispatchedProviderListV2ApiService {
	return c.ApiClient.StarDemandOmGetChallengeDispatchedProviderListV2Api
}

func (c *Client) StarDemandOmGetChallengeItemsDataV2Api() *api.StarDemandOmGetChallengeItemsDataV2ApiService {
	return c.ApiClient.StarDemandOmGetChallengeItemsDataV2Api
}

func (c *Client) StarDemandOmGetChallengeV2Api() *api.StarDemandOmGetChallengeV2ApiService {
	return c.ApiClient.StarDemandOmGetChallengeV2Api
}

func (c *Client) StarDemandOmGetDemandListV2Api() *api.StarDemandOmGetDemandListV2ApiService {
	return c.ApiClient.StarDemandOmGetDemandListV2Api
}

func (c *Client) StarDemandOmUpdateChallengeV2Api() *api.StarDemandOmUpdateChallengeV2ApiService {
	return c.ApiClient.StarDemandOmUpdateChallengeV2Api
}

func (c *Client) StarDemandOrderListV2Api() *api.StarDemandOrderListV2ApiService {
	return c.ApiClient.StarDemandOrderListV2Api
}

func (c *Client) StarDemandSearchWordModifyV2Api() *api.StarDemandSearchWordModifyV2ApiService {
	return c.ApiClient.StarDemandSearchWordModifyV2Api
}

func (c *Client) StarDemanderDemanderGetAuthorBindInfoV2Api() *api.StarDemanderDemanderGetAuthorBindInfoV2ApiService {
	return c.ApiClient.StarDemanderDemanderGetAuthorBindInfoV2Api
}

func (c *Client) StarDemanderDemanderGetBindShareCodeV2Api() *api.StarDemanderDemanderGetBindShareCodeV2ApiService {
	return c.ApiClient.StarDemanderDemanderGetBindShareCodeV2Api
}

func (c *Client) StarDemanderGetCarBrandListV2Api() *api.StarDemanderGetCarBrandListV2ApiService {
	return c.ApiClient.StarDemanderGetCarBrandListV2Api
}

func (c *Client) StarDemanderGetRegisteredIpActsV2Api() *api.StarDemanderGetRegisteredIpActsV2ApiService {
	return c.ApiClient.StarDemanderGetRegisteredIpActsV2Api
}

func (c *Client) StarDemanderQueryCpsTaskSalesV2Api() *api.StarDemanderQueryCpsTaskSalesV2ApiService {
	return c.ApiClient.StarDemanderQueryCpsTaskSalesV2Api
}

func (c *Client) StarDemanderUploadCpsTaskSalesV2Api() *api.StarDemanderUploadCpsTaskSalesV2ApiService {
	return c.ApiClient.StarDemanderUploadCpsTaskSalesV2Api
}

func (c *Client) StarFeCommonServiceTalentDayGetCheckInCodeV2Api() *api.StarFeCommonServiceTalentDayGetCheckInCodeV2ApiService {
	return c.ApiClient.StarFeCommonServiceTalentDayGetCheckInCodeV2Api
}

func (c *Client) StarFeCommonServiceTalentDayJudgeCheckedInV2Api() *api.StarFeCommonServiceTalentDayJudgeCheckedInV2ApiService {
	return c.ApiClient.StarFeCommonServiceTalentDayJudgeCheckedInV2Api
}

func (c *Client) StarGetCreateChallengeDataDictV2Api() *api.StarGetCreateChallengeDataDictV2ApiService {
	return c.ApiClient.StarGetCreateChallengeDataDictV2Api
}

func (c *Client) StarInfoV2Api() *api.StarInfoV2ApiService {
	return c.ApiClient.StarInfoV2Api
}

func (c *Client) StarMcnContractChallengeV2Api() *api.StarMcnContractChallengeV2ApiService {
	return c.ApiClient.StarMcnContractChallengeV2Api
}

func (c *Client) StarMcnGetAuthorListV2Api() *api.StarMcnGetAuthorListV2ApiService {
	return c.ApiClient.StarMcnGetAuthorListV2Api
}

func (c *Client) StarMcnGetContractChallengeAuthorItemListV2V2Api() *api.StarMcnGetContractChallengeAuthorItemListV2V2ApiService {
	return c.ApiClient.StarMcnGetContractChallengeAuthorItemListV2V2Api
}

func (c *Client) StarMcnGetContractedChallengeListV2Api() *api.StarMcnGetContractedChallengeListV2ApiService {
	return c.ApiClient.StarMcnGetContractedChallengeListV2Api
}

func (c *Client) StarMcnGetContractedChallengeUrlV2Api() *api.StarMcnGetContractedChallengeUrlV2ApiService {
	return c.ApiClient.StarMcnGetContractedChallengeUrlV2Api
}

func (c *Client) StarMcnGetUnparticipatedTaskV2Api() *api.StarMcnGetUnparticipatedTaskV2ApiService {
	return c.ApiClient.StarMcnGetUnparticipatedTaskV2Api
}

func (c *Client) StarMcnProviderGetParticipatedTaskV2Api() *api.StarMcnProviderGetParticipatedTaskV2ApiService {
	return c.ApiClient.StarMcnProviderGetParticipatedTaskV2Api
}

func (c *Client) StarMcnProviderGetTagOnTaskV2Api() *api.StarMcnProviderGetTagOnTaskV2ApiService {
	return c.ApiClient.StarMcnProviderGetTagOnTaskV2Api
}

func (c *Client) StarMcnProviderGetTaskAuthorListV2Api() *api.StarMcnProviderGetTaskAuthorListV2ApiService {
	return c.ApiClient.StarMcnProviderGetTaskAuthorListV2Api
}

func (c *Client) StarMcnProviderGetTaskDetailV2Api() *api.StarMcnProviderGetTaskDetailV2ApiService {
	return c.ApiClient.StarMcnProviderGetTaskDetailV2Api
}

func (c *Client) StarMcnProviderGetTaskItemListV2Api() *api.StarMcnProviderGetTaskItemListV2ApiService {
	return c.ApiClient.StarMcnProviderGetTaskItemListV2Api
}

func (c *Client) StarMcnProviderGetTaskShareCodeV2Api() *api.StarMcnProviderGetTaskShareCodeV2ApiService {
	return c.ApiClient.StarMcnProviderGetTaskShareCodeV2Api
}

func (c *Client) StarMcnProviderGetUnparticipatedTaskV2Api() *api.StarMcnProviderGetUnparticipatedTaskV2ApiService {
	return c.ApiClient.StarMcnProviderGetUnparticipatedTaskV2Api
}

func (c *Client) StarMcnProviderHandleJoinedAuthorV2Api() *api.StarMcnProviderHandleJoinedAuthorV2ApiService {
	return c.ApiClient.StarMcnProviderHandleJoinedAuthorV2Api
}

func (c *Client) StarMcnProviderUploadSearchItemV2Api() *api.StarMcnProviderUploadSearchItemV2ApiService {
	return c.ApiClient.StarMcnProviderUploadSearchItemV2Api
}

func (c *Client) StarOrderApproveResourceV2Api() *api.StarOrderApproveResourceV2ApiService {
	return c.ApiClient.StarOrderApproveResourceV2Api
}

func (c *Client) StarOrderDemanderCancelV2Api() *api.StarOrderDemanderCancelV2ApiService {
	return c.ApiClient.StarOrderDemanderCancelV2Api
}

func (c *Client) StarOrderDetailV2Api() *api.StarOrderDetailV2ApiService {
	return c.ApiClient.StarOrderDetailV2Api
}

func (c *Client) StarOrderFinishV2Api() *api.StarOrderFinishV2ApiService {
	return c.ApiClient.StarOrderFinishV2Api
}

func (c *Client) StarOrderGetCancelAmountV2Api() *api.StarOrderGetCancelAmountV2ApiService {
	return c.ApiClient.StarOrderGetCancelAmountV2Api
}

func (c *Client) StarOrderGetComponentV2Api() *api.StarOrderGetComponentV2ApiService {
	return c.ApiClient.StarOrderGetComponentV2Api
}

func (c *Client) StarOrderGetInfoV2Api() *api.StarOrderGetInfoV2ApiService {
	return c.ApiClient.StarOrderGetInfoV2Api
}

func (c *Client) StarOrderGetLiveV2Api() *api.StarOrderGetLiveV2ApiService {
	return c.ApiClient.StarOrderGetLiveV2Api
}

func (c *Client) StarOrderGetScriptV2Api() *api.StarOrderGetScriptV2ApiService {
	return c.ApiClient.StarOrderGetScriptV2Api
}

func (c *Client) StarOrderGetVideoV2Api() *api.StarOrderGetVideoV2ApiService {
	return c.ApiClient.StarOrderGetVideoV2Api
}

func (c *Client) StarOrderListByCampaignV2Api() *api.StarOrderListByCampaignV2ApiService {
	return c.ApiClient.StarOrderListByCampaignV2Api
}

func (c *Client) StarOrderPublishResourceV2Api() *api.StarOrderPublishResourceV2ApiService {
	return c.ApiClient.StarOrderPublishResourceV2Api
}

func (c *Client) StarOrderPushResourceV2Api() *api.StarOrderPushResourceV2ApiService {
	return c.ApiClient.StarOrderPushResourceV2Api
}

func (c *Client) StarOrderRejectResourceV2Api() *api.StarOrderRejectResourceV2ApiService {
	return c.ApiClient.StarOrderRejectResourceV2Api
}

func (c *Client) StarOrderReplyAuthorCancelV2Api() *api.StarOrderReplyAuthorCancelV2ApiService {
	return c.ApiClient.StarOrderReplyAuthorCancelV2Api
}

func (c *Client) StarOrderUpdateV2Api() *api.StarOrderUpdateV2ApiService {
	return c.ApiClient.StarOrderUpdateV2Api
}

func (c *Client) StarProjectListV2Api() *api.StarProjectListV2ApiService {
	return c.ApiClient.StarProjectListV2Api
}

func (c *Client) StarReportCustomDataTopicDailyReportV2Api() *api.StarReportCustomDataTopicDailyReportV2ApiService {
	return c.ApiClient.StarReportCustomDataTopicDailyReportV2Api
}

func (c *Client) StarReportCustomDataTopicReportV2Api() *api.StarReportCustomDataTopicReportV2ApiService {
	return c.ApiClient.StarReportCustomDataTopicReportV2Api
}

func (c *Client) StarReportDataTopicConfigV2Api() *api.StarReportDataTopicConfigV2ApiService {
	return c.ApiClient.StarReportDataTopicConfigV2Api
}

func (c *Client) StarReportOrderOverviewGetV2Api() *api.StarReportOrderOverviewGetV2ApiService {
	return c.ApiClient.StarReportOrderOverviewGetV2Api
}

func (c *Client) StarReportOrderOverviewV2Api() *api.StarReportOrderOverviewV2ApiService {
	return c.ApiClient.StarReportOrderOverviewV2Api
}

func (c *Client) StarReportOrderUserDistributionGetV2Api() *api.StarReportOrderUserDistributionGetV2ApiService {
	return c.ApiClient.StarReportOrderUserDistributionGetV2Api
}

func (c *Client) StarStarAdUniteTaskDetailV2Api() *api.StarStarAdUniteTaskDetailV2ApiService {
	return c.ApiClient.StarStarAdUniteTaskDetailV2Api
}

func (c *Client) StarStarAdUniteTaskItemListV2Api() *api.StarStarAdUniteTaskItemListV2ApiService {
	return c.ApiClient.StarStarAdUniteTaskItemListV2Api
}

func (c *Client) StarStarAdUniteTaskListV2Api() *api.StarStarAdUniteTaskListV2ApiService {
	return c.ApiClient.StarStarAdUniteTaskListV2Api
}

func (c *Client) StarTaskBindProjectV2Api() *api.StarTaskBindProjectV2ApiService {
	return c.ApiClient.StarTaskBindProjectV2Api
}

func (c *Client) StarTaskListByProjectV2Api() *api.StarTaskListByProjectV2ApiService {
	return c.ApiClient.StarTaskListByProjectV2Api
}

func (c *Client) StarUpdateProjectV2Api() *api.StarUpdateProjectV2ApiService {
	return c.ApiClient.StarUpdateProjectV2Api
}

func (c *Client) StarUserGetAwemeAuthorIdV2Api() *api.StarUserGetAwemeAuthorIdV2ApiService {
	return c.ApiClient.StarUserGetAwemeAuthorIdV2Api
}

func (c *Client) StarUserGetStarIdV2Api() *api.StarUserGetStarIdV2ApiService {
	return c.ApiClient.StarUserGetStarIdV2Api
}

func (c *Client) StarVasAppendOrderToBoostItemGroupV2Api() *api.StarVasAppendOrderToBoostItemGroupV2ApiService {
	return c.ApiClient.StarVasAppendOrderToBoostItemGroupV2Api
}

func (c *Client) StarVasCancelBoostItemGroupV2Api() *api.StarVasCancelBoostItemGroupV2ApiService {
	return c.ApiClient.StarVasCancelBoostItemGroupV2Api
}

func (c *Client) StarVasCreateBoostItemGroupV2Api() *api.StarVasCreateBoostItemGroupV2ApiService {
	return c.ApiClient.StarVasCreateBoostItemGroupV2Api
}

func (c *Client) StarVasGetBoostGroupListV2Api() *api.StarVasGetBoostGroupListV2ApiService {
	return c.ApiClient.StarVasGetBoostGroupListV2Api
}

func (c *Client) StarVasGetBoostItemGroupDetailV2Api() *api.StarVasGetBoostItemGroupDetailV2ApiService {
	return c.ApiClient.StarVasGetBoostItemGroupDetailV2Api
}

func (c *Client) StarVasGetCommonAuthorPackageListV2Api() *api.StarVasGetCommonAuthorPackageListV2ApiService {
	return c.ApiClient.StarVasGetCommonAuthorPackageListV2Api
}

func (c *Client) StarVasGetExportBoostItemGroupResultV2Api() *api.StarVasGetExportBoostItemGroupResultV2ApiService {
	return c.ApiClient.StarVasGetExportBoostItemGroupResultV2Api
}

func (c *Client) StarVasSubmitExportBoostItemGroupDataV2Api() *api.StarVasSubmitExportBoostItemGroupDataV2ApiService {
	return c.ApiClient.StarVasSubmitExportBoostItemGroupDataV2Api
}

func (c *Client) StardeliveryTaskAuthorDetailV30Api() *api.StardeliveryTaskAuthorDetailV30ApiService {
	return c.ApiClient.StardeliveryTaskAuthorDetailV30Api
}

func (c *Client) StardeliveryTaskAuthorVideoAuditV30Api() *api.StardeliveryTaskAuthorVideoAuditV30ApiService {
	return c.ApiClient.StardeliveryTaskAuthorVideoAuditV30Api
}

func (c *Client) StardeliveryTaskAuthorVideoDetailV30Api() *api.StardeliveryTaskAuthorVideoDetailV30ApiService {
	return c.ApiClient.StardeliveryTaskAuthorVideoDetailV30Api
}

func (c *Client) StardeliveryTaskBudgetUpdateV30Api() *api.StardeliveryTaskBudgetUpdateV30ApiService {
	return c.ApiClient.StardeliveryTaskBudgetUpdateV30Api
}

func (c *Client) StardeliveryTaskCancelV30Api() *api.StardeliveryTaskCancelV30ApiService {
	return c.ApiClient.StardeliveryTaskCancelV30Api
}

func (c *Client) StardeliveryTaskDetailV30Api() *api.StardeliveryTaskDetailV30ApiService {
	return c.ApiClient.StardeliveryTaskDetailV30Api
}

func (c *Client) StardeliveryTaskListV30Api() *api.StardeliveryTaskListV30ApiService {
	return c.ApiClient.StardeliveryTaskListV30Api
}

func (c *Client) StardeliveryTaskPostEndTimeUpdateV30Api() *api.StardeliveryTaskPostEndTimeUpdateV30ApiService {
	return c.ApiClient.StardeliveryTaskPostEndTimeUpdateV30Api
}

func (c *Client) StardeliveryTaskShareV30Api() *api.StardeliveryTaskShareV30ApiService {
	return c.ApiClient.StardeliveryTaskShareV30Api
}

func (c *Client) StardeliveryTaskShareableListV30Api() *api.StardeliveryTaskShareableListV30ApiService {
	return c.ApiClient.StardeliveryTaskShareableListV30Api
}

func (c *Client) StardeliveryTaskSharingListV30Api() *api.StardeliveryTaskSharingListV30ApiService {
	return c.ApiClient.StardeliveryTaskSharingListV30Api
}

func (c *Client) StardeliveryTaskUnshareV30Api() *api.StardeliveryTaskUnshareV30ApiService {
	return c.ApiClient.StardeliveryTaskUnshareV30Api
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

func (c *Client) ToolsAbTestInfoGetV2Api() *api.ToolsAbTestInfoGetV2ApiService {
	return c.ApiClient.ToolsAbTestInfoGetV2Api
}

func (c *Client) ToolsAbTestListGetV2Api() *api.ToolsAbTestListGetV2ApiService {
	return c.ApiClient.ToolsAbTestListGetV2Api
}

func (c *Client) ToolsAdPreviewQrcodeGetV30Api() *api.ToolsAdPreviewQrcodeGetV30ApiService {
	return c.ApiClient.ToolsAdPreviewQrcodeGetV30Api
}

func (c *Client) ToolsAdRaiseStatusGetV2Api() *api.ToolsAdRaiseStatusGetV2ApiService {
	return c.ApiClient.ToolsAdRaiseStatusGetV2Api
}

func (c *Client) ToolsAdRaiseVersionGetV2Api() *api.ToolsAdRaiseVersionGetV2ApiService {
	return c.ApiClient.ToolsAdRaiseVersionGetV2Api
}

func (c *Client) ToolsAdminInfoV2Api() *api.ToolsAdminInfoV2ApiService {
	return c.ApiClient.ToolsAdminInfoV2Api
}

func (c *Client) ToolsAdvertiserStoreSearchV2Api() *api.ToolsAdvertiserStoreSearchV2ApiService {
	return c.ApiClient.ToolsAdvertiserStoreSearchV2Api
}

func (c *Client) ToolsAipThirdSiteCreateV2Api() *api.ToolsAipThirdSiteCreateV2ApiService {
	return c.ApiClient.ToolsAipThirdSiteCreateV2Api
}

func (c *Client) ToolsAipThirdSiteGetV2Api() *api.ToolsAipThirdSiteGetV2ApiService {
	return c.ApiClient.ToolsAipThirdSiteGetV2Api
}

func (c *Client) ToolsAipThirdSiteUpdateV2Api() *api.ToolsAipThirdSiteUpdateV2ApiService {
	return c.ApiClient.ToolsAipThirdSiteUpdateV2Api
}

func (c *Client) ToolsAppIosListV2Api() *api.ToolsAppIosListV2ApiService {
	return c.ApiClient.ToolsAppIosListV2Api
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

func (c *Client) ToolsAppManagementExtendPackageListV2Api() *api.ToolsAppManagementExtendPackageListV2ApiService {
	return c.ApiClient.ToolsAppManagementExtendPackageListV2Api
}

func (c *Client) ToolsAppManagementExtendPackageListV2V2Api() *api.ToolsAppManagementExtendPackageListV2V2ApiService {
	return c.ApiClient.ToolsAppManagementExtendPackageListV2V2Api
}

func (c *Client) ToolsAppManagementExtendPackageUpdateV2Api() *api.ToolsAppManagementExtendPackageUpdateV2ApiService {
	return c.ApiClient.ToolsAppManagementExtendPackageUpdateV2Api
}

func (c *Client) ToolsAppManagementExtendPackageUpdateV2V2Api() *api.ToolsAppManagementExtendPackageUpdateV2V2ApiService {
	return c.ApiClient.ToolsAppManagementExtendPackageUpdateV2V2Api
}

func (c *Client) ToolsAppManagementHarmonyAppListV2Api() *api.ToolsAppManagementHarmonyAppListV2ApiService {
	return c.ApiClient.ToolsAppManagementHarmonyAppListV2Api
}

func (c *Client) ToolsAppManagementIndustryInfoListV2Api() *api.ToolsAppManagementIndustryInfoListV2ApiService {
	return c.ApiClient.ToolsAppManagementIndustryInfoListV2Api
}

func (c *Client) ToolsAppManagementShareAccountListV2Api() *api.ToolsAppManagementShareAccountListV2ApiService {
	return c.ApiClient.ToolsAppManagementShareAccountListV2Api
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

func (c *Client) ToolsAssetLinkListV30Api() *api.ToolsAssetLinkListV30ApiService {
	return c.ApiClient.ToolsAssetLinkListV30Api
}

func (c *Client) ToolsAwemeAuthAuthShareAdShareV2Api() *api.ToolsAwemeAuthAuthShareAdShareV2ApiService {
	return c.ApiClient.ToolsAwemeAuthAuthShareAdShareV2Api
}

func (c *Client) ToolsAwemeAuthCancelV2Api() *api.ToolsAwemeAuthCancelV2ApiService {
	return c.ApiClient.ToolsAwemeAuthCancelV2Api
}

func (c *Client) ToolsAwemeAuthListV2Api() *api.ToolsAwemeAuthListV2ApiService {
	return c.ApiClient.ToolsAwemeAuthListV2Api
}

func (c *Client) ToolsAwemeAuthRenewalV2Api() *api.ToolsAwemeAuthRenewalV2ApiService {
	return c.ApiClient.ToolsAwemeAuthRenewalV2Api
}

func (c *Client) ToolsAwemeAuthV2Api() *api.ToolsAwemeAuthV2ApiService {
	return c.ApiClient.ToolsAwemeAuthV2Api
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

func (c *Client) ToolsBlueFlowKeywordListV30Api() *api.ToolsBlueFlowKeywordListV30ApiService {
	return c.ApiClient.ToolsBlueFlowKeywordListV30Api
}

func (c *Client) ToolsBlueFlowPackageListV30Api() *api.ToolsBlueFlowPackageListV30ApiService {
	return c.ApiClient.ToolsBlueFlowPackageListV30Api
}

func (c *Client) ToolsBpAssetManagementShareCancelV30Api() *api.ToolsBpAssetManagementShareCancelV30ApiService {
	return c.ApiClient.ToolsBpAssetManagementShareCancelV30Api
}

func (c *Client) ToolsBpAssetManagementShareGetV30Api() *api.ToolsBpAssetManagementShareGetV30ApiService {
	return c.ApiClient.ToolsBpAssetManagementShareGetV30Api
}

func (c *Client) ToolsBpAssetManagementShareV30Api() *api.ToolsBpAssetManagementShareV30ApiService {
	return c.ApiClient.ToolsBpAssetManagementShareV30Api
}

func (c *Client) ToolsClueBridgeCallCreateV2Api() *api.ToolsClueBridgeCallCreateV2ApiService {
	return c.ApiClient.ToolsClueBridgeCallCreateV2Api
}

func (c *Client) ToolsClueCallCreateV2Api() *api.ToolsClueCallCreateV2ApiService {
	return c.ApiClient.ToolsClueCallCreateV2Api
}

func (c *Client) ToolsClueCallVirtualNumberGetV2Api() *api.ToolsClueCallVirtualNumberGetV2ApiService {
	return c.ApiClient.ToolsClueCallVirtualNumberGetV2Api
}

func (c *Client) ToolsClueCallVirtualNumberRefundDetailGetV2Api() *api.ToolsClueCallVirtualNumberRefundDetailGetV2ApiService {
	return c.ApiClient.ToolsClueCallVirtualNumberRefundDetailGetV2Api
}

func (c *Client) ToolsClueCallbackV2Api() *api.ToolsClueCallbackV2ApiService {
	return c.ApiClient.ToolsClueCallbackV2Api
}

func (c *Client) ToolsClueClueOverviewQueryV2Api() *api.ToolsClueClueOverviewQueryV2ApiService {
	return c.ApiClient.ToolsClueClueOverviewQueryV2Api
}

func (c *Client) ToolsClueContactLogListV2Api() *api.ToolsClueContactLogListV2ApiService {
	return c.ApiClient.ToolsClueContactLogListV2Api
}

func (c *Client) ToolsClueContactLogOverviewQueryV2Api() *api.ToolsClueContactLogOverviewQueryV2ApiService {
	return c.ApiClient.ToolsClueContactLogOverviewQueryV2Api
}

func (c *Client) ToolsClueContactLogRecordUrlGetV2Api() *api.ToolsClueContactLogRecordUrlGetV2ApiService {
	return c.ApiClient.ToolsClueContactLogRecordUrlGetV2Api
}

func (c *Client) ToolsClueExtInfoCallbackV2Api() *api.ToolsClueExtInfoCallbackV2ApiService {
	return c.ApiClient.ToolsClueExtInfoCallbackV2Api
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

func (c *Client) ToolsClueInfoGetV2Api() *api.ToolsClueInfoGetV2ApiService {
	return c.ApiClient.ToolsClueInfoGetV2Api
}

func (c *Client) ToolsClueInfoUpdateV2Api() *api.ToolsClueInfoUpdateV2ApiService {
	return c.ApiClient.ToolsClueInfoUpdateV2Api
}

func (c *Client) ToolsClueLifeCallbackV2Api() *api.ToolsClueLifeCallbackV2ApiService {
	return c.ApiClient.ToolsClueLifeCallbackV2Api
}

func (c *Client) ToolsClueLifeGetV2Api() *api.ToolsClueLifeGetV2ApiService {
	return c.ApiClient.ToolsClueLifeGetV2Api
}

func (c *Client) ToolsClueLiteContactGetV2Api() *api.ToolsClueLiteContactGetV2ApiService {
	return c.ApiClient.ToolsClueLiteContactGetV2Api
}

func (c *Client) ToolsClueLiteContactRecordV2Api() *api.ToolsClueLiteContactRecordV2ApiService {
	return c.ApiClient.ToolsClueLiteContactRecordV2Api
}

func (c *Client) ToolsCluePrivateMessageCallbackV2Api() *api.ToolsCluePrivateMessageCallbackV2ApiService {
	return c.ApiClient.ToolsCluePrivateMessageCallbackV2Api
}

func (c *Client) ToolsClueRefundDetailGetV2Api() *api.ToolsClueRefundDetailGetV2ApiService {
	return c.ApiClient.ToolsClueRefundDetailGetV2Api
}

func (c *Client) ToolsClueRefundInfoQueryV2Api() *api.ToolsClueRefundInfoQueryV2ApiService {
	return c.ApiClient.ToolsClueRefundInfoQueryV2Api
}

func (c *Client) ToolsClueRefundReportGetV2Api() *api.ToolsClueRefundReportGetV2ApiService {
	return c.ApiClient.ToolsClueRefundReportGetV2Api
}

func (c *Client) ToolsClueRefundViewGetV2Api() *api.ToolsClueRefundViewGetV2ApiService {
	return c.ApiClient.ToolsClueRefundViewGetV2Api
}

func (c *Client) ToolsClueRobotScriptQueryV2Api() *api.ToolsClueRobotScriptQueryV2ApiService {
	return c.ApiClient.ToolsClueRobotScriptQueryV2Api
}

func (c *Client) ToolsClueRobotTaskCancelV2Api() *api.ToolsClueRobotTaskCancelV2ApiService {
	return c.ApiClient.ToolsClueRobotTaskCancelV2Api
}

func (c *Client) ToolsClueRobotTaskCreateV2Api() *api.ToolsClueRobotTaskCreateV2ApiService {
	return c.ApiClient.ToolsClueRobotTaskCreateV2Api
}

func (c *Client) ToolsClueSmartPhoneGetV2Api() *api.ToolsClueSmartPhoneGetV2ApiService {
	return c.ApiClient.ToolsClueSmartPhoneGetV2Api
}

func (c *Client) ToolsClueWebrtcCreateV2V30Api() *api.ToolsClueWebrtcCreateV2V30ApiService {
	return c.ApiClient.ToolsClueWebrtcCreateV2V30Api
}

func (c *Client) ToolsClueWebrtcTokenGetV2Api() *api.ToolsClueWebrtcTokenGetV2ApiService {
	return c.ApiClient.ToolsClueWebrtcTokenGetV2Api
}

func (c *Client) ToolsClueWebrtcTokenGetV2V30Api() *api.ToolsClueWebrtcTokenGetV2V30ApiService {
	return c.ApiClient.ToolsClueWebrtcTokenGetV2V30Api
}

func (c *Client) ToolsCommentGetV30Api() *api.ToolsCommentGetV30ApiService {
	return c.ApiClient.ToolsCommentGetV30Api
}

func (c *Client) ToolsCommentHideV30Api() *api.ToolsCommentHideV30ApiService {
	return c.ApiClient.ToolsCommentHideV30Api
}

func (c *Client) ToolsCommentMetricsGetV30Api() *api.ToolsCommentMetricsGetV30ApiService {
	return c.ApiClient.ToolsCommentMetricsGetV30Api
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

func (c *Client) ToolsEventAllAssetsDetailV2Api() *api.ToolsEventAllAssetsDetailV2ApiService {
	return c.ApiClient.ToolsEventAllAssetsDetailV2Api
}

func (c *Client) ToolsEventAllAssetsListV2Api() *api.ToolsEventAllAssetsListV2ApiService {
	return c.ApiClient.ToolsEventAllAssetsListV2Api
}

func (c *Client) ToolsEventAssetsGetV2Api() *api.ToolsEventAssetsGetV2ApiService {
	return c.ApiClient.ToolsEventAssetsGetV2Api
}

func (c *Client) ToolsEventConvertOptimizedGoalGetV30Api() *api.ToolsEventConvertOptimizedGoalGetV30ApiService {
	return c.ApiClient.ToolsEventConvertOptimizedGoalGetV30Api
}

func (c *Client) ToolsForbiddenLinkGreyGetV30Api() *api.ToolsForbiddenLinkGreyGetV30ApiService {
	return c.ApiClient.ToolsForbiddenLinkGreyGetV30Api
}

func (c *Client) ToolsGrayGetV30Api() *api.ToolsGrayGetV30ApiService {
	return c.ApiClient.ToolsGrayGetV30Api
}

func (c *Client) ToolsHotMaterialDeriveAdoptV30Api() *api.ToolsHotMaterialDeriveAdoptV30ApiService {
	return c.ApiClient.ToolsHotMaterialDeriveAdoptV30Api
}

func (c *Client) ToolsHotMaterialDeriveGetV30Api() *api.ToolsHotMaterialDeriveGetV30ApiService {
	return c.ApiClient.ToolsHotMaterialDeriveGetV30Api
}

func (c *Client) ToolsHotMaterialDeriveListV30Api() *api.ToolsHotMaterialDeriveListV30ApiService {
	return c.ApiClient.ToolsHotMaterialDeriveListV30Api
}

func (c *Client) ToolsHotMaterialDeriveSubmitV30Api() *api.ToolsHotMaterialDeriveSubmitV30ApiService {
	return c.ApiClient.ToolsHotMaterialDeriveSubmitV30Api
}

func (c *Client) ToolsInactiveAdvertiserListV30Api() *api.ToolsInactiveAdvertiserListV30ApiService {
	return c.ApiClient.ToolsInactiveAdvertiserListV30Api
}

func (c *Client) ToolsIndustryGetV2Api() *api.ToolsIndustryGetV2ApiService {
	return c.ApiClient.ToolsIndustryGetV2Api
}

func (c *Client) ToolsInterestActionActionKeywordV2Api() *api.ToolsInterestActionActionKeywordV2ApiService {
	return c.ApiClient.ToolsInterestActionActionKeywordV2Api
}

func (c *Client) ToolsInterestActionId2wordV2Api() *api.ToolsInterestActionId2wordV2ApiService {
	return c.ApiClient.ToolsInterestActionId2wordV2Api
}

func (c *Client) ToolsInterestActionInterestKeywordV2Api() *api.ToolsInterestActionInterestKeywordV2ApiService {
	return c.ApiClient.ToolsInterestActionInterestKeywordV2Api
}

func (c *Client) ToolsInterestActionKeywordSuggestV2Api() *api.ToolsInterestActionKeywordSuggestV2ApiService {
	return c.ApiClient.ToolsInterestActionKeywordSuggestV2Api
}

func (c *Client) ToolsIsSupportUniversalGetV2Api() *api.ToolsIsSupportUniversalGetV2ApiService {
	return c.ApiClient.ToolsIsSupportUniversalGetV2Api
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

func (c *Client) ToolsMaterialRaiseCreateV30Api() *api.ToolsMaterialRaiseCreateV30ApiService {
	return c.ApiClient.ToolsMaterialRaiseCreateV30Api
}

func (c *Client) ToolsMaterialRaiseGetV30Api() *api.ToolsMaterialRaiseGetV30ApiService {
	return c.ApiClient.ToolsMaterialRaiseGetV30Api
}

func (c *Client) ToolsMaterialRaiseMaterialIdsGetV30Api() *api.ToolsMaterialRaiseMaterialIdsGetV30ApiService {
	return c.ApiClient.ToolsMaterialRaiseMaterialIdsGetV30Api
}

func (c *Client) ToolsMaterialRaiseRecordsGetV30Api() *api.ToolsMaterialRaiseRecordsGetV30ApiService {
	return c.ApiClient.ToolsMaterialRaiseRecordsGetV30Api
}

func (c *Client) ToolsMaterialRaiseStatusGetV30Api() *api.ToolsMaterialRaiseStatusGetV30ApiService {
	return c.ApiClient.ToolsMaterialRaiseStatusGetV30Api
}

func (c *Client) ToolsMaterialRaiseStopV30Api() *api.ToolsMaterialRaiseStopV30ApiService {
	return c.ApiClient.ToolsMaterialRaiseStopV30Api
}

func (c *Client) ToolsMicroAppCreateV30Api() *api.ToolsMicroAppCreateV30ApiService {
	return c.ApiClient.ToolsMicroAppCreateV30Api
}

func (c *Client) ToolsMicroAppListV30Api() *api.ToolsMicroAppListV30ApiService {
	return c.ApiClient.ToolsMicroAppListV30Api
}

func (c *Client) ToolsMicroAppUpdateV30Api() *api.ToolsMicroAppUpdateV30ApiService {
	return c.ApiClient.ToolsMicroAppUpdateV30Api
}

func (c *Client) ToolsMicroGameConvertWindowGetV30Api() *api.ToolsMicroGameConvertWindowGetV30ApiService {
	return c.ApiClient.ToolsMicroGameConvertWindowGetV30Api
}

func (c *Client) ToolsMicroGameConvertWindowUpdateV30Api() *api.ToolsMicroGameConvertWindowUpdateV30ApiService {
	return c.ApiClient.ToolsMicroGameConvertWindowUpdateV30Api
}

func (c *Client) ToolsMicroGameCreateV30Api() *api.ToolsMicroGameCreateV30ApiService {
	return c.ApiClient.ToolsMicroGameCreateV30Api
}

func (c *Client) ToolsMicroGameListV30Api() *api.ToolsMicroGameListV30ApiService {
	return c.ApiClient.ToolsMicroGameListV30Api
}

func (c *Client) ToolsMicroGameUpdateV30Api() *api.ToolsMicroGameUpdateV30ApiService {
	return c.ApiClient.ToolsMicroGameUpdateV30Api
}

func (c *Client) ToolsNoBidSuggestBidV2Api() *api.ToolsNoBidSuggestBidV2ApiService {
	return c.ApiClient.ToolsNoBidSuggestBidV2Api
}

func (c *Client) ToolsOrangeSiteGetV30Api() *api.ToolsOrangeSiteGetV30ApiService {
	return c.ApiClient.ToolsOrangeSiteGetV30Api
}

func (c *Client) ToolsPioneerProgramAttachmentUploadV2Api() *api.ToolsPioneerProgramAttachmentUploadV2ApiService {
	return c.ApiClient.ToolsPioneerProgramAttachmentUploadV2Api
}

func (c *Client) ToolsPlayableCloudGameListV2Api() *api.ToolsPlayableCloudGameListV2ApiService {
	return c.ApiClient.ToolsPlayableCloudGameListV2Api
}

func (c *Client) ToolsPlayableCreateV2Api() *api.ToolsPlayableCreateV2ApiService {
	return c.ApiClient.ToolsPlayableCreateV2Api
}

func (c *Client) ToolsPlayableGrantResultV2Api() *api.ToolsPlayableGrantResultV2ApiService {
	return c.ApiClient.ToolsPlayableGrantResultV2Api
}

func (c *Client) ToolsPlayableGrantV2Api() *api.ToolsPlayableGrantV2ApiService {
	return c.ApiClient.ToolsPlayableGrantV2Api
}

func (c *Client) ToolsPlayableListGetV2Api() *api.ToolsPlayableListGetV2ApiService {
	return c.ApiClient.ToolsPlayableListGetV2Api
}

func (c *Client) ToolsPlayableSaveV2Api() *api.ToolsPlayableSaveV2ApiService {
	return c.ApiClient.ToolsPlayableSaveV2Api
}

func (c *Client) ToolsPlayableUploadV2Api() *api.ToolsPlayableUploadV2ApiService {
	return c.ApiClient.ToolsPlayableUploadV2Api
}

func (c *Client) ToolsPlayableValidateV2Api() *api.ToolsPlayableValidateV2ApiService {
	return c.ApiClient.ToolsPlayableValidateV2Api
}

func (c *Client) ToolsPreAuditGetV2Api() *api.ToolsPreAuditGetV2ApiService {
	return c.ApiClient.ToolsPreAuditGetV2Api
}

func (c *Client) ToolsPreAuditSendV2Api() *api.ToolsPreAuditSendV2ApiService {
	return c.ApiClient.ToolsPreAuditSendV2Api
}

func (c *Client) ToolsPrivativeWordAdAddV2Api() *api.ToolsPrivativeWordAdAddV2ApiService {
	return c.ApiClient.ToolsPrivativeWordAdAddV2Api
}

func (c *Client) ToolsPrivativeWordAdUpdateV2Api() *api.ToolsPrivativeWordAdUpdateV2ApiService {
	return c.ApiClient.ToolsPrivativeWordAdUpdateV2Api
}

func (c *Client) ToolsPrivativeWordBatchGetV30Api() *api.ToolsPrivativeWordBatchGetV30ApiService {
	return c.ApiClient.ToolsPrivativeWordBatchGetV30Api
}

func (c *Client) ToolsPrivativeWordCampaignAddV2Api() *api.ToolsPrivativeWordCampaignAddV2ApiService {
	return c.ApiClient.ToolsPrivativeWordCampaignAddV2Api
}

func (c *Client) ToolsPrivativeWordCampaignUpdateV2Api() *api.ToolsPrivativeWordCampaignUpdateV2ApiService {
	return c.ApiClient.ToolsPrivativeWordCampaignUpdateV2Api
}

func (c *Client) ToolsPrivativeWordGetV2Api() *api.ToolsPrivativeWordGetV2ApiService {
	return c.ApiClient.ToolsPrivativeWordGetV2Api
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

func (c *Client) ToolsPromotionCardRecommendTitleGetV2Api() *api.ToolsPromotionCardRecommendTitleGetV2ApiService {
	return c.ApiClient.ToolsPromotionCardRecommendTitleGetV2Api
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

func (c *Client) ToolsRtaGetInfoTmpV2Api() *api.ToolsRtaGetInfoTmpV2ApiService {
	return c.ApiClient.ToolsRtaGetInfoTmpV2Api
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

func (c *Client) ToolsRubeexPlayableAdListV2Api() *api.ToolsRubeexPlayableAdListV2ApiService {
	return c.ApiClient.ToolsRubeexPlayableAdListV2Api
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

func (c *Client) ToolsSiteGetV2Api() *api.ToolsSiteGetV2ApiService {
	return c.ApiClient.ToolsSiteGetV2Api
}

func (c *Client) ToolsSiteHandselV2Api() *api.ToolsSiteHandselV2ApiService {
	return c.ApiClient.ToolsSiteHandselV2Api
}

func (c *Client) ToolsSitePreviewV2Api() *api.ToolsSitePreviewV2ApiService {
	return c.ApiClient.ToolsSitePreviewV2Api
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

func (c *Client) ToolsStarTaskSettlementConfigV2Api() *api.ToolsStarTaskSettlementConfigV2ApiService {
	return c.ApiClient.ToolsStarTaskSettlementConfigV2Api
}

func (c *Client) ToolsStarTaskTitleTopicGetV2Api() *api.ToolsStarTaskTitleTopicGetV2ApiService {
	return c.ApiClient.ToolsStarTaskTitleTopicGetV2Api
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

func (c *Client) ToolsUnionFlowPackagePromotionReportV30Api() *api.ToolsUnionFlowPackagePromotionReportV30ApiService {
	return c.ApiClient.ToolsUnionFlowPackagePromotionReportV30Api
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

func (c *Client) UploadStatementV2Api() *api.UploadStatementV2ApiService {
	return c.ApiClient.UploadStatementV2Api
}

func (c *Client) UserInfoV2Api() *api.UserInfoV2ApiService {
	return c.ApiClient.UserInfoV2Api
}

func (c *Client) YuntuAudienceInfoCreateV30Api() *api.YuntuAudienceInfoCreateV30ApiService {
	return c.ApiClient.YuntuAudienceInfoCreateV30Api
}

func (c *Client) YuntuAudienceInfoDeleteV30Api() *api.YuntuAudienceInfoDeleteV30ApiService {
	return c.ApiClient.YuntuAudienceInfoDeleteV30Api
}

func (c *Client) YuntuAudienceInfoGetV30Api() *api.YuntuAudienceInfoGetV30ApiService {
	return c.ApiClient.YuntuAudienceInfoGetV30Api
}

func (c *Client) YuntuAudienceLabelCreateV30Api() *api.YuntuAudienceLabelCreateV30ApiService {
	return c.ApiClient.YuntuAudienceLabelCreateV30Api
}

func (c *Client) YuntuAudienceLabelDeleteV30Api() *api.YuntuAudienceLabelDeleteV30ApiService {
	return c.ApiClient.YuntuAudienceLabelDeleteV30Api
}

func (c *Client) YuntuAudienceLabelGetV30Api() *api.YuntuAudienceLabelGetV30ApiService {
	return c.ApiClient.YuntuAudienceLabelGetV30Api
}

func (c *Client) YuntuBrandInfoGetV30Api() *api.YuntuBrandInfoGetV30ApiService {
	return c.ApiClient.YuntuBrandInfoGetV30Api
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
