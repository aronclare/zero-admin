package userservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UserAddLogic 新增用户
/*
Author: LiuFeiHua
Date: 2023/12/18 14:13
*/
type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserAdd 新增用户
// 1.根据用户名称查询用户是否已存在
// 2.如果用户已存在,则直接返回
// 3.用户不存在时,则直接添加用户
func (l *UserAddLogic) UserAdd(in *sysclient.UserAddReq) (*sysclient.UserAddResp, error) {
	q := query.SysUser.WithContext(l.ctx)
	// 1.根据用户名称查询用户是否已存在
	name := in.UserName
	count, err := q.Where(query.SysDept.DeptName.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据用户名称：%s,查询用户信息失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("查询用户信息失败"))
	}

	//2.如果用户已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "用户信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("用户：%s,已存在", name))
	}

	//3.用户不存在时,则直接添加用户
	user := &model.SysUser{
		UserName:   in.UserName,
		NickName:   in.NickName,
		Avatar:     in.Avatar,
		Password:   "123456",
		Salt:       "123456",
		Email:      in.Email,
		Mobile:     in.Mobile,
		UserStatus: in.UserStatus,
		DeptID:     in.DeptId,
		CreateBy:   in.CreateBy,
		DelFlag:    1,
		JobID:      in.JobId,
	}

	err = q.Create(user)

	if err != nil {
		logc.Errorf(l.ctx, "新增用户异常,参数:%+v,异常:%s", user, err.Error())
		return nil, errors.New("新增用户异常")
	}

	return &sysclient.UserAddResp{}, nil
}
