// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package homerecommendproductservice

import (
	"context"

	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CouponAddReq                                  = smsclient.CouponAddReq
	CouponAddResp                                 = smsclient.CouponAddResp
	CouponCountReq                                = smsclient.CouponCountReq
	CouponCountResp                               = smsclient.CouponCountResp
	CouponDeleteReq                               = smsclient.CouponDeleteReq
	CouponDeleteResp                              = smsclient.CouponDeleteResp
	CouponFindByIdReq                             = smsclient.CouponFindByIdReq
	CouponFindByIdResp                            = smsclient.CouponFindByIdResp
	CouponFindByIdsReq                            = smsclient.CouponFindByIdsReq
	CouponFindByIdsResp                           = smsclient.CouponFindByIdsResp
	CouponFindByProductIdAndProductCategoryIdReq  = smsclient.CouponFindByProductIdAndProductCategoryIdReq
	CouponFindByProductIdAndProductCategoryIdResp = smsclient.CouponFindByProductIdAndProductCategoryIdResp
	CouponHistoryAddReq                           = smsclient.CouponHistoryAddReq
	CouponHistoryAddResp                          = smsclient.CouponHistoryAddResp
	CouponHistoryDeleteReq                        = smsclient.CouponHistoryDeleteReq
	CouponHistoryDeleteResp                       = smsclient.CouponHistoryDeleteResp
	CouponHistoryListData                         = smsclient.CouponHistoryListData
	CouponHistoryListReq                          = smsclient.CouponHistoryListReq
	CouponHistoryListResp                         = smsclient.CouponHistoryListResp
	CouponHistoryUpdateReq                        = smsclient.CouponHistoryUpdateReq
	CouponHistoryUpdateResp                       = smsclient.CouponHistoryUpdateResp
	CouponListData                                = smsclient.CouponListData
	CouponListReq                                 = smsclient.CouponListReq
	CouponListResp                                = smsclient.CouponListResp
	CouponProductCategoryRelationAddReq           = smsclient.CouponProductCategoryRelationAddReq
	CouponProductCategoryRelationAddResp          = smsclient.CouponProductCategoryRelationAddResp
	CouponProductCategoryRelationDeleteReq        = smsclient.CouponProductCategoryRelationDeleteReq
	CouponProductCategoryRelationDeleteResp       = smsclient.CouponProductCategoryRelationDeleteResp
	CouponProductCategoryRelationListData         = smsclient.CouponProductCategoryRelationListData
	CouponProductCategoryRelationListReq          = smsclient.CouponProductCategoryRelationListReq
	CouponProductCategoryRelationListResp         = smsclient.CouponProductCategoryRelationListResp
	CouponProductCategoryRelationUpdateReq        = smsclient.CouponProductCategoryRelationUpdateReq
	CouponProductCategoryRelationUpdateResp       = smsclient.CouponProductCategoryRelationUpdateResp
	CouponProductRelationAddReq                   = smsclient.CouponProductRelationAddReq
	CouponProductRelationAddResp                  = smsclient.CouponProductRelationAddResp
	CouponProductRelationDeleteReq                = smsclient.CouponProductRelationDeleteReq
	CouponProductRelationDeleteResp               = smsclient.CouponProductRelationDeleteResp
	CouponProductRelationListData                 = smsclient.CouponProductRelationListData
	CouponProductRelationListReq                  = smsclient.CouponProductRelationListReq
	CouponProductRelationListResp                 = smsclient.CouponProductRelationListResp
	CouponProductRelationUpdateReq                = smsclient.CouponProductRelationUpdateReq
	CouponProductRelationUpdateResp               = smsclient.CouponProductRelationUpdateResp
	CouponUpdateReq                               = smsclient.CouponUpdateReq
	CouponUpdateResp                              = smsclient.CouponUpdateResp
	FlashPromotionAddReq                          = smsclient.FlashPromotionAddReq
	FlashPromotionAddResp                         = smsclient.FlashPromotionAddResp
	FlashPromotionDeleteReq                       = smsclient.FlashPromotionDeleteReq
	FlashPromotionDeleteResp                      = smsclient.FlashPromotionDeleteResp
	FlashPromotionListByDateReq                   = smsclient.FlashPromotionListByDateReq
	FlashPromotionListByDateResp                  = smsclient.FlashPromotionListByDateResp
	FlashPromotionListData                        = smsclient.FlashPromotionListData
	FlashPromotionListReq                         = smsclient.FlashPromotionListReq
	FlashPromotionListResp                        = smsclient.FlashPromotionListResp
	FlashPromotionLogAddReq                       = smsclient.FlashPromotionLogAddReq
	FlashPromotionLogAddResp                      = smsclient.FlashPromotionLogAddResp
	FlashPromotionLogDeleteReq                    = smsclient.FlashPromotionLogDeleteReq
	FlashPromotionLogDeleteResp                   = smsclient.FlashPromotionLogDeleteResp
	FlashPromotionLogListData                     = smsclient.FlashPromotionLogListData
	FlashPromotionLogListReq                      = smsclient.FlashPromotionLogListReq
	FlashPromotionLogListResp                     = smsclient.FlashPromotionLogListResp
	FlashPromotionLogUpdateReq                    = smsclient.FlashPromotionLogUpdateReq
	FlashPromotionLogUpdateResp                   = smsclient.FlashPromotionLogUpdateResp
	FlashPromotionProductRelationAddReq           = smsclient.FlashPromotionProductRelationAddReq
	FlashPromotionProductRelationAddResp          = smsclient.FlashPromotionProductRelationAddResp
	FlashPromotionProductRelationDeleteReq        = smsclient.FlashPromotionProductRelationDeleteReq
	FlashPromotionProductRelationDeleteResp       = smsclient.FlashPromotionProductRelationDeleteResp
	FlashPromotionProductRelationListData         = smsclient.FlashPromotionProductRelationListData
	FlashPromotionProductRelationListReq          = smsclient.FlashPromotionProductRelationListReq
	FlashPromotionProductRelationListResp         = smsclient.FlashPromotionProductRelationListResp
	FlashPromotionProductRelationUpdateReq        = smsclient.FlashPromotionProductRelationUpdateReq
	FlashPromotionProductRelationUpdateResp       = smsclient.FlashPromotionProductRelationUpdateResp
	FlashPromotionSessionAddReq                   = smsclient.FlashPromotionSessionAddReq
	FlashPromotionSessionAddResp                  = smsclient.FlashPromotionSessionAddResp
	FlashPromotionSessionByTimeReq                = smsclient.FlashPromotionSessionByTimeReq
	FlashPromotionSessionByTimeResp               = smsclient.FlashPromotionSessionByTimeResp
	FlashPromotionSessionDeleteReq                = smsclient.FlashPromotionSessionDeleteReq
	FlashPromotionSessionDeleteResp               = smsclient.FlashPromotionSessionDeleteResp
	FlashPromotionSessionListData                 = smsclient.FlashPromotionSessionListData
	FlashPromotionSessionListReq                  = smsclient.FlashPromotionSessionListReq
	FlashPromotionSessionListResp                 = smsclient.FlashPromotionSessionListResp
	FlashPromotionSessionUpdateReq                = smsclient.FlashPromotionSessionUpdateReq
	FlashPromotionSessionUpdateResp               = smsclient.FlashPromotionSessionUpdateResp
	FlashPromotionUpdateReq                       = smsclient.FlashPromotionUpdateReq
	FlashPromotionUpdateResp                      = smsclient.FlashPromotionUpdateResp
	HomeAdvertiseAddReq                           = smsclient.HomeAdvertiseAddReq
	HomeAdvertiseAddResp                          = smsclient.HomeAdvertiseAddResp
	HomeAdvertiseDeleteReq                        = smsclient.HomeAdvertiseDeleteReq
	HomeAdvertiseDeleteResp                       = smsclient.HomeAdvertiseDeleteResp
	HomeAdvertiseListData                         = smsclient.HomeAdvertiseListData
	HomeAdvertiseListReq                          = smsclient.HomeAdvertiseListReq
	HomeAdvertiseListResp                         = smsclient.HomeAdvertiseListResp
	HomeAdvertiseUpdateReq                        = smsclient.HomeAdvertiseUpdateReq
	HomeAdvertiseUpdateResp                       = smsclient.HomeAdvertiseUpdateResp
	HomeBrandAddData                              = smsclient.HomeBrandAddData
	HomeBrandAddReq                               = smsclient.HomeBrandAddReq
	HomeBrandAddResp                              = smsclient.HomeBrandAddResp
	HomeBrandDeleteReq                            = smsclient.HomeBrandDeleteReq
	HomeBrandDeleteResp                           = smsclient.HomeBrandDeleteResp
	HomeBrandListData                             = smsclient.HomeBrandListData
	HomeBrandListReq                              = smsclient.HomeBrandListReq
	HomeBrandListResp                             = smsclient.HomeBrandListResp
	HomeBrandUpdateReq                            = smsclient.HomeBrandUpdateReq
	HomeBrandUpdateResp                           = smsclient.HomeBrandUpdateResp
	HomeNewProductAddData                         = smsclient.HomeNewProductAddData
	HomeNewProductAddReq                          = smsclient.HomeNewProductAddReq
	HomeNewProductAddResp                         = smsclient.HomeNewProductAddResp
	HomeNewProductDeleteReq                       = smsclient.HomeNewProductDeleteReq
	HomeNewProductDeleteResp                      = smsclient.HomeNewProductDeleteResp
	HomeNewProductListData                        = smsclient.HomeNewProductListData
	HomeNewProductListReq                         = smsclient.HomeNewProductListReq
	HomeNewProductListResp                        = smsclient.HomeNewProductListResp
	HomeNewProductUpdateReq                       = smsclient.HomeNewProductUpdateReq
	HomeNewProductUpdateResp                      = smsclient.HomeNewProductUpdateResp
	HomeRecommendProductAddData                   = smsclient.HomeRecommendProductAddData
	HomeRecommendProductAddReq                    = smsclient.HomeRecommendProductAddReq
	HomeRecommendProductAddResp                   = smsclient.HomeRecommendProductAddResp
	HomeRecommendProductDeleteReq                 = smsclient.HomeRecommendProductDeleteReq
	HomeRecommendProductDeleteResp                = smsclient.HomeRecommendProductDeleteResp
	HomeRecommendProductListData                  = smsclient.HomeRecommendProductListData
	HomeRecommendProductListReq                   = smsclient.HomeRecommendProductListReq
	HomeRecommendProductListResp                  = smsclient.HomeRecommendProductListResp
	HomeRecommendProductUpdateReq                 = smsclient.HomeRecommendProductUpdateReq
	HomeRecommendProductUpdateResp                = smsclient.HomeRecommendProductUpdateResp
	HomeRecommendSubjectAddData                   = smsclient.HomeRecommendSubjectAddData
	HomeRecommendSubjectAddReq                    = smsclient.HomeRecommendSubjectAddReq
	HomeRecommendSubjectAddResp                   = smsclient.HomeRecommendSubjectAddResp
	HomeRecommendSubjectDeleteReq                 = smsclient.HomeRecommendSubjectDeleteReq
	HomeRecommendSubjectDeleteResp                = smsclient.HomeRecommendSubjectDeleteResp
	HomeRecommendSubjectListData                  = smsclient.HomeRecommendSubjectListData
	HomeRecommendSubjectListReq                   = smsclient.HomeRecommendSubjectListReq
	HomeRecommendSubjectListResp                  = smsclient.HomeRecommendSubjectListResp
	HomeRecommendSubjectUpdateReq                 = smsclient.HomeRecommendSubjectUpdateReq
	HomeRecommendSubjectUpdateResp                = smsclient.HomeRecommendSubjectUpdateResp
	QueryFlashPromotionByProductReq               = smsclient.QueryFlashPromotionByProductReq
	QueryFlashPromotionByProductResp              = smsclient.QueryFlashPromotionByProductResp
	QueryMemberCouponListReq                      = smsclient.QueryMemberCouponListReq
	QueryMemberCouponListResp                     = smsclient.QueryMemberCouponListResp
	UpdateCouponStatusReq                         = smsclient.UpdateCouponStatusReq
	UpdateCouponStatusResp                        = smsclient.UpdateCouponStatusResp

	HomeRecommendProductService interface {
		HomeRecommendProductAdd(ctx context.Context, in *HomeRecommendProductAddReq, opts ...grpc.CallOption) (*HomeRecommendProductAddResp, error)
		HomeRecommendProductList(ctx context.Context, in *HomeRecommendProductListReq, opts ...grpc.CallOption) (*HomeRecommendProductListResp, error)
		HomeRecommendProductUpdate(ctx context.Context, in *HomeRecommendProductUpdateReq, opts ...grpc.CallOption) (*HomeRecommendProductUpdateResp, error)
		HomeRecommendProductDelete(ctx context.Context, in *HomeRecommendProductDeleteReq, opts ...grpc.CallOption) (*HomeRecommendProductDeleteResp, error)
	}

	defaultHomeRecommendProductService struct {
		cli zrpc.Client
	}
)

func NewHomeRecommendProductService(cli zrpc.Client) HomeRecommendProductService {
	return &defaultHomeRecommendProductService{
		cli: cli,
	}
}

func (m *defaultHomeRecommendProductService) HomeRecommendProductAdd(ctx context.Context, in *HomeRecommendProductAddReq, opts ...grpc.CallOption) (*HomeRecommendProductAddResp, error) {
	client := smsclient.NewHomeRecommendProductServiceClient(m.cli.Conn())
	return client.HomeRecommendProductAdd(ctx, in, opts...)
}

func (m *defaultHomeRecommendProductService) HomeRecommendProductList(ctx context.Context, in *HomeRecommendProductListReq, opts ...grpc.CallOption) (*HomeRecommendProductListResp, error) {
	client := smsclient.NewHomeRecommendProductServiceClient(m.cli.Conn())
	return client.HomeRecommendProductList(ctx, in, opts...)
}

func (m *defaultHomeRecommendProductService) HomeRecommendProductUpdate(ctx context.Context, in *HomeRecommendProductUpdateReq, opts ...grpc.CallOption) (*HomeRecommendProductUpdateResp, error) {
	client := smsclient.NewHomeRecommendProductServiceClient(m.cli.Conn())
	return client.HomeRecommendProductUpdate(ctx, in, opts...)
}

func (m *defaultHomeRecommendProductService) HomeRecommendProductDelete(ctx context.Context, in *HomeRecommendProductDeleteReq, opts ...grpc.CallOption) (*HomeRecommendProductDeleteResp, error) {
	client := smsclient.NewHomeRecommendProductServiceClient(m.cli.Conn())
	return client.HomeRecommendProductDelete(ctx, in, opts...)
}
