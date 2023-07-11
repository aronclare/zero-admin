// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package membermembertagrelationservice

import (
	"context"

	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GrowthChangeHistoryAddReq               = umsclient.GrowthChangeHistoryAddReq
	GrowthChangeHistoryAddResp              = umsclient.GrowthChangeHistoryAddResp
	GrowthChangeHistoryDeleteReq            = umsclient.GrowthChangeHistoryDeleteReq
	GrowthChangeHistoryDeleteResp           = umsclient.GrowthChangeHistoryDeleteResp
	GrowthChangeHistoryListData             = umsclient.GrowthChangeHistoryListData
	GrowthChangeHistoryListReq              = umsclient.GrowthChangeHistoryListReq
	GrowthChangeHistoryListResp             = umsclient.GrowthChangeHistoryListResp
	GrowthChangeHistoryUpdateReq            = umsclient.GrowthChangeHistoryUpdateReq
	GrowthChangeHistoryUpdateResp           = umsclient.GrowthChangeHistoryUpdateResp
	IntegrationChangeHistoryAddReq          = umsclient.IntegrationChangeHistoryAddReq
	IntegrationChangeHistoryAddResp         = umsclient.IntegrationChangeHistoryAddResp
	IntegrationChangeHistoryDeleteReq       = umsclient.IntegrationChangeHistoryDeleteReq
	IntegrationChangeHistoryDeleteResp      = umsclient.IntegrationChangeHistoryDeleteResp
	IntegrationChangeHistoryListData        = umsclient.IntegrationChangeHistoryListData
	IntegrationChangeHistoryListReq         = umsclient.IntegrationChangeHistoryListReq
	IntegrationChangeHistoryListResp        = umsclient.IntegrationChangeHistoryListResp
	IntegrationChangeHistoryUpdateReq       = umsclient.IntegrationChangeHistoryUpdateReq
	IntegrationChangeHistoryUpdateResp      = umsclient.IntegrationChangeHistoryUpdateResp
	IntegrationConsumeSettingAddReq         = umsclient.IntegrationConsumeSettingAddReq
	IntegrationConsumeSettingAddResp        = umsclient.IntegrationConsumeSettingAddResp
	IntegrationConsumeSettingDeleteReq      = umsclient.IntegrationConsumeSettingDeleteReq
	IntegrationConsumeSettingDeleteResp     = umsclient.IntegrationConsumeSettingDeleteResp
	IntegrationConsumeSettingListData       = umsclient.IntegrationConsumeSettingListData
	IntegrationConsumeSettingListReq        = umsclient.IntegrationConsumeSettingListReq
	IntegrationConsumeSettingListResp       = umsclient.IntegrationConsumeSettingListResp
	IntegrationConsumeSettingUpdateReq      = umsclient.IntegrationConsumeSettingUpdateReq
	IntegrationConsumeSettingUpdateResp     = umsclient.IntegrationConsumeSettingUpdateResp
	MemberAddReq                            = umsclient.MemberAddReq
	MemberAddResp                           = umsclient.MemberAddResp
	MemberByIdReq                           = umsclient.MemberByIdReq
	MemberDeleteReq                         = umsclient.MemberDeleteReq
	MemberDeleteResp                        = umsclient.MemberDeleteResp
	MemberLevelAddReq                       = umsclient.MemberLevelAddReq
	MemberLevelAddResp                      = umsclient.MemberLevelAddResp
	MemberLevelDeleteReq                    = umsclient.MemberLevelDeleteReq
	MemberLevelDeleteResp                   = umsclient.MemberLevelDeleteResp
	MemberLevelListData                     = umsclient.MemberLevelListData
	MemberLevelListReq                      = umsclient.MemberLevelListReq
	MemberLevelListResp                     = umsclient.MemberLevelListResp
	MemberLevelUpdateReq                    = umsclient.MemberLevelUpdateReq
	MemberLevelUpdateResp                   = umsclient.MemberLevelUpdateResp
	MemberListData                          = umsclient.MemberListData
	MemberListReq                           = umsclient.MemberListReq
	MemberListResp                          = umsclient.MemberListResp
	MemberLoginLogAddReq                    = umsclient.MemberLoginLogAddReq
	MemberLoginLogAddResp                   = umsclient.MemberLoginLogAddResp
	MemberLoginLogDeleteReq                 = umsclient.MemberLoginLogDeleteReq
	MemberLoginLogDeleteResp                = umsclient.MemberLoginLogDeleteResp
	MemberLoginLogListData                  = umsclient.MemberLoginLogListData
	MemberLoginLogListReq                   = umsclient.MemberLoginLogListReq
	MemberLoginLogListResp                  = umsclient.MemberLoginLogListResp
	MemberLoginLogUpdateReq                 = umsclient.MemberLoginLogUpdateReq
	MemberLoginLogUpdateResp                = umsclient.MemberLoginLogUpdateResp
	MemberLoginReq                          = umsclient.MemberLoginReq
	MemberLoginResp                         = umsclient.MemberLoginResp
	MemberMemberTagRelationAddReq           = umsclient.MemberMemberTagRelationAddReq
	MemberMemberTagRelationAddResp          = umsclient.MemberMemberTagRelationAddResp
	MemberMemberTagRelationDeleteReq        = umsclient.MemberMemberTagRelationDeleteReq
	MemberMemberTagRelationDeleteResp       = umsclient.MemberMemberTagRelationDeleteResp
	MemberMemberTagRelationListData         = umsclient.MemberMemberTagRelationListData
	MemberMemberTagRelationListReq          = umsclient.MemberMemberTagRelationListReq
	MemberMemberTagRelationListResp         = umsclient.MemberMemberTagRelationListResp
	MemberMemberTagRelationUpdateReq        = umsclient.MemberMemberTagRelationUpdateReq
	MemberMemberTagRelationUpdateResp       = umsclient.MemberMemberTagRelationUpdateResp
	MemberProductCategoryRelationAddReq     = umsclient.MemberProductCategoryRelationAddReq
	MemberProductCategoryRelationAddResp    = umsclient.MemberProductCategoryRelationAddResp
	MemberProductCategoryRelationDeleteReq  = umsclient.MemberProductCategoryRelationDeleteReq
	MemberProductCategoryRelationDeleteResp = umsclient.MemberProductCategoryRelationDeleteResp
	MemberProductCategoryRelationListData   = umsclient.MemberProductCategoryRelationListData
	MemberProductCategoryRelationListReq    = umsclient.MemberProductCategoryRelationListReq
	MemberProductCategoryRelationListResp   = umsclient.MemberProductCategoryRelationListResp
	MemberProductCategoryRelationUpdateReq  = umsclient.MemberProductCategoryRelationUpdateReq
	MemberProductCategoryRelationUpdateResp = umsclient.MemberProductCategoryRelationUpdateResp
	MemberProductCollectionAddReq           = umsclient.MemberProductCollectionAddReq
	MemberProductCollectionAddResp          = umsclient.MemberProductCollectionAddResp
	MemberProductCollectionDeleteReq        = umsclient.MemberProductCollectionDeleteReq
	MemberProductCollectionDeleteResp       = umsclient.MemberProductCollectionDeleteResp
	MemberProductCollectionListData         = umsclient.MemberProductCollectionListData
	MemberProductCollectionListReq          = umsclient.MemberProductCollectionListReq
	MemberProductCollectionListResp         = umsclient.MemberProductCollectionListResp
	MemberProductCollectionUpdateReq        = umsclient.MemberProductCollectionUpdateReq
	MemberProductCollectionUpdateResp       = umsclient.MemberProductCollectionUpdateResp
	MemberReadHistoryAddReq                 = umsclient.MemberReadHistoryAddReq
	MemberReadHistoryAddResp                = umsclient.MemberReadHistoryAddResp
	MemberReadHistoryDeleteReq              = umsclient.MemberReadHistoryDeleteReq
	MemberReadHistoryDeleteResp             = umsclient.MemberReadHistoryDeleteResp
	MemberReadHistoryListData               = umsclient.MemberReadHistoryListData
	MemberReadHistoryListReq                = umsclient.MemberReadHistoryListReq
	MemberReadHistoryListResp               = umsclient.MemberReadHistoryListResp
	MemberReadHistoryUpdateReq              = umsclient.MemberReadHistoryUpdateReq
	MemberReadHistoryUpdateResp             = umsclient.MemberReadHistoryUpdateResp
	MemberReceiveAddressAddReq              = umsclient.MemberReceiveAddressAddReq
	MemberReceiveAddressAddResp             = umsclient.MemberReceiveAddressAddResp
	MemberReceiveAddressDeleteReq           = umsclient.MemberReceiveAddressDeleteReq
	MemberReceiveAddressDeleteResp          = umsclient.MemberReceiveAddressDeleteResp
	MemberReceiveAddressListData            = umsclient.MemberReceiveAddressListData
	MemberReceiveAddressListReq             = umsclient.MemberReceiveAddressListReq
	MemberReceiveAddressListResp            = umsclient.MemberReceiveAddressListResp
	MemberReceiveAddressQueryDetailReq      = umsclient.MemberReceiveAddressQueryDetailReq
	MemberReceiveAddressQueryDetailResp     = umsclient.MemberReceiveAddressQueryDetailResp
	MemberReceiveAddressUpdateReq           = umsclient.MemberReceiveAddressUpdateReq
	MemberReceiveAddressUpdateResp          = umsclient.MemberReceiveAddressUpdateResp
	MemberRuleSettingAddReq                 = umsclient.MemberRuleSettingAddReq
	MemberRuleSettingAddResp                = umsclient.MemberRuleSettingAddResp
	MemberRuleSettingDeleteReq              = umsclient.MemberRuleSettingDeleteReq
	MemberRuleSettingDeleteResp             = umsclient.MemberRuleSettingDeleteResp
	MemberRuleSettingListData               = umsclient.MemberRuleSettingListData
	MemberRuleSettingListReq                = umsclient.MemberRuleSettingListReq
	MemberRuleSettingListResp               = umsclient.MemberRuleSettingListResp
	MemberRuleSettingUpdateReq              = umsclient.MemberRuleSettingUpdateReq
	MemberRuleSettingUpdateResp             = umsclient.MemberRuleSettingUpdateResp
	MemberStatisticsInfoAddReq              = umsclient.MemberStatisticsInfoAddReq
	MemberStatisticsInfoAddResp             = umsclient.MemberStatisticsInfoAddResp
	MemberStatisticsInfoDeleteReq           = umsclient.MemberStatisticsInfoDeleteReq
	MemberStatisticsInfoDeleteResp          = umsclient.MemberStatisticsInfoDeleteResp
	MemberStatisticsInfoListData            = umsclient.MemberStatisticsInfoListData
	MemberStatisticsInfoListReq             = umsclient.MemberStatisticsInfoListReq
	MemberStatisticsInfoListResp            = umsclient.MemberStatisticsInfoListResp
	MemberStatisticsInfoUpdateReq           = umsclient.MemberStatisticsInfoUpdateReq
	MemberStatisticsInfoUpdateResp          = umsclient.MemberStatisticsInfoUpdateResp
	MemberTagAddReq                         = umsclient.MemberTagAddReq
	MemberTagAddResp                        = umsclient.MemberTagAddResp
	MemberTagDeleteReq                      = umsclient.MemberTagDeleteReq
	MemberTagDeleteResp                     = umsclient.MemberTagDeleteResp
	MemberTagListData                       = umsclient.MemberTagListData
	MemberTagListReq                        = umsclient.MemberTagListReq
	MemberTagListResp                       = umsclient.MemberTagListResp
	MemberTagUpdateReq                      = umsclient.MemberTagUpdateReq
	MemberTagUpdateResp                     = umsclient.MemberTagUpdateResp
	MemberTaskAddReq                        = umsclient.MemberTaskAddReq
	MemberTaskAddResp                       = umsclient.MemberTaskAddResp
	MemberTaskDeleteReq                     = umsclient.MemberTaskDeleteReq
	MemberTaskDeleteResp                    = umsclient.MemberTaskDeleteResp
	MemberTaskListData                      = umsclient.MemberTaskListData
	MemberTaskListReq                       = umsclient.MemberTaskListReq
	MemberTaskListResp                      = umsclient.MemberTaskListResp
	MemberTaskUpdateReq                     = umsclient.MemberTaskUpdateReq
	MemberTaskUpdateResp                    = umsclient.MemberTaskUpdateResp
	MemberUpdatePasswordReq                 = umsclient.MemberUpdatePasswordReq
	MemberUpdateReq                         = umsclient.MemberUpdateReq
	MemberUpdateResp                        = umsclient.MemberUpdateResp

	MemberMemberTagRelationService interface {
		MemberMemberTagRelationAdd(ctx context.Context, in *MemberMemberTagRelationAddReq, opts ...grpc.CallOption) (*MemberMemberTagRelationAddResp, error)
		MemberMemberTagRelationList(ctx context.Context, in *MemberMemberTagRelationListReq, opts ...grpc.CallOption) (*MemberMemberTagRelationListResp, error)
		MemberMemberTagRelationUpdate(ctx context.Context, in *MemberMemberTagRelationUpdateReq, opts ...grpc.CallOption) (*MemberMemberTagRelationUpdateResp, error)
		MemberMemberTagRelationDelete(ctx context.Context, in *MemberMemberTagRelationDeleteReq, opts ...grpc.CallOption) (*MemberMemberTagRelationDeleteResp, error)
	}

	defaultMemberMemberTagRelationService struct {
		cli zrpc.Client
	}
)

func NewMemberMemberTagRelationService(cli zrpc.Client) MemberMemberTagRelationService {
	return &defaultMemberMemberTagRelationService{
		cli: cli,
	}
}

func (m *defaultMemberMemberTagRelationService) MemberMemberTagRelationAdd(ctx context.Context, in *MemberMemberTagRelationAddReq, opts ...grpc.CallOption) (*MemberMemberTagRelationAddResp, error) {
	client := umsclient.NewMemberMemberTagRelationServiceClient(m.cli.Conn())
	return client.MemberMemberTagRelationAdd(ctx, in, opts...)
}

func (m *defaultMemberMemberTagRelationService) MemberMemberTagRelationList(ctx context.Context, in *MemberMemberTagRelationListReq, opts ...grpc.CallOption) (*MemberMemberTagRelationListResp, error) {
	client := umsclient.NewMemberMemberTagRelationServiceClient(m.cli.Conn())
	return client.MemberMemberTagRelationList(ctx, in, opts...)
}

func (m *defaultMemberMemberTagRelationService) MemberMemberTagRelationUpdate(ctx context.Context, in *MemberMemberTagRelationUpdateReq, opts ...grpc.CallOption) (*MemberMemberTagRelationUpdateResp, error) {
	client := umsclient.NewMemberMemberTagRelationServiceClient(m.cli.Conn())
	return client.MemberMemberTagRelationUpdate(ctx, in, opts...)
}

func (m *defaultMemberMemberTagRelationService) MemberMemberTagRelationDelete(ctx context.Context, in *MemberMemberTagRelationDeleteReq, opts ...grpc.CallOption) (*MemberMemberTagRelationDeleteResp, error) {
	client := umsclient.NewMemberMemberTagRelationServiceClient(m.cli.Conn())
	return client.MemberMemberTagRelationDelete(ctx, in, opts...)
}
