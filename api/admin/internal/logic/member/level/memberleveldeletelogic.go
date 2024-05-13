package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLevelDeleteLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/13 13:38
*/
type MemberLevelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLevelDeleteLogic {
	return MemberLevelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MemberLevelDelete 删除会员等级
func (l *MemberLevelDeleteLogic) MemberLevelDelete(req types.DeleteMemberLevelReq) (*types.DeleteMemberLevelResp, error) {
	_, err := l.svcCtx.MemberLevelService.MemberLevelDelete(l.ctx, &umsclient.MemberLevelDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除会员等级异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除会员等级失败")
	}
	return &types.DeleteMemberLevelResp{
		Code:    "000000",
		Message: "添加会员等级成功",
	}, nil
}
