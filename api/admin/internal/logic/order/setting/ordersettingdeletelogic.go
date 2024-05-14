package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderSettingDeleteLogic 订单设置
/*
Author: LiuFeiHua
Date: 2024/5/14 11:18
*/
type OrderSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderSettingDeleteLogic {
	return OrderSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// OrderSettingDelete 删除订单设置
func (l *OrderSettingDeleteLogic) OrderSettingDelete(req types.DeleteOrderSettingReq) (*types.DeleteOrderSettingResp, error) {
	_, err := l.svcCtx.OrderSettingService.OrderSettingDelete(l.ctx, &omsclient.OrderSettingDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除订单设置异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除订单设置失败")
	}

	return &types.DeleteOrderSettingResp{
		Code:    "000000",
		Message: "",
	}, nil
}
