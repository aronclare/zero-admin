package syslogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SysLogAddLogic 添加操作日志
/*
Author: LiuFeiHua
Date: 2023/12/18 17:08
*/
type SysLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogAddLogic {
	return &SysLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SysLogAdd 添加操作日志
func (l *SysLogAddLogic) SysLogAdd(in *sysclient.SysLogAddReq) (*sysclient.SysLogAddResp, error) {
	err := query.SysOperateLog.WithContext(l.ctx).Create(&model.SysOperateLog{
		UserName:       in.UserName,
		Operation:      in.Operation,
		RequestMethod:  in.Method,
		RequestParams:  in.RequestParams,
		ResponseParams: in.ResponseParams,
		UseTime:        in.Time,
		OperationIP:    in.Ip,
		OperationTime:  time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &sysclient.SysLogAddResp{}, nil
}
