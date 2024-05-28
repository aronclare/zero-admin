// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package jobservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeptAddReq             = sysclient.DeptAddReq
	DeptAddResp            = sysclient.DeptAddResp
	DeptDeleteReq          = sysclient.DeptDeleteReq
	DeptDeleteResp         = sysclient.DeptDeleteResp
	DeptListData           = sysclient.DeptListData
	DeptListReq            = sysclient.DeptListReq
	DeptListResp           = sysclient.DeptListResp
	DeptUpdateReq          = sysclient.DeptUpdateReq
	DeptUpdateResp         = sysclient.DeptUpdateResp
	DictAddReq             = sysclient.DictAddReq
	DictAddResp            = sysclient.DictAddResp
	DictDeleteReq          = sysclient.DictDeleteReq
	DictDeleteResp         = sysclient.DictDeleteResp
	DictItemAddReq         = sysclient.DictItemAddReq
	DictItemAddResp        = sysclient.DictItemAddResp
	DictItemDeleteReq      = sysclient.DictItemDeleteReq
	DictItemDeleteResp     = sysclient.DictItemDeleteResp
	DictItemListData       = sysclient.DictItemListData
	DictItemListReq        = sysclient.DictItemListReq
	DictItemListResp       = sysclient.DictItemListResp
	DictItemReq            = sysclient.DictItemReq
	DictItemResp           = sysclient.DictItemResp
	DictItemUpdateReq      = sysclient.DictItemUpdateReq
	DictItemUpdateResp     = sysclient.DictItemUpdateResp
	DictListData           = sysclient.DictListData
	DictListReq            = sysclient.DictListReq
	DictListResp           = sysclient.DictListResp
	DictReq                = sysclient.DictReq
	DictResp               = sysclient.DictResp
	DictUpdateReq          = sysclient.DictUpdateReq
	DictUpdateResp         = sysclient.DictUpdateResp
	InfoReq                = sysclient.InfoReq
	InfoResp               = sysclient.InfoResp
	JobAddReq              = sysclient.JobAddReq
	JobAddResp             = sysclient.JobAddResp
	JobDeleteReq           = sysclient.JobDeleteReq
	JobDeleteResp          = sysclient.JobDeleteResp
	JobListData            = sysclient.JobListData
	JobListReq             = sysclient.JobListReq
	JobListResp            = sysclient.JobListResp
	JobUpdateReq           = sysclient.JobUpdateReq
	JobUpdateResp          = sysclient.JobUpdateResp
	LoginLogAddReq         = sysclient.LoginLogAddReq
	LoginLogAddResp        = sysclient.LoginLogAddResp
	LoginLogDeleteReq      = sysclient.LoginLogDeleteReq
	LoginLogDeleteResp     = sysclient.LoginLogDeleteResp
	LoginLogListData       = sysclient.LoginLogListData
	LoginLogListReq        = sysclient.LoginLogListReq
	LoginLogListResp       = sysclient.LoginLogListResp
	LoginReq               = sysclient.LoginReq
	LoginResp              = sysclient.LoginResp
	MenuAddReq             = sysclient.MenuAddReq
	MenuAddResp            = sysclient.MenuAddResp
	MenuDeleteReq          = sysclient.MenuDeleteReq
	MenuDeleteResp         = sysclient.MenuDeleteResp
	MenuListData           = sysclient.MenuListData
	MenuListReq            = sysclient.MenuListReq
	MenuListResp           = sysclient.MenuListResp
	MenuListTree           = sysclient.MenuListTree
	MenuUpdateReq          = sysclient.MenuUpdateReq
	MenuUpdateResp         = sysclient.MenuUpdateResp
	QueryRoleMenuListReq   = sysclient.QueryRoleMenuListReq
	QueryRoleMenuListResp  = sysclient.QueryRoleMenuListResp
	QueryUserRoleListReq   = sysclient.QueryUserRoleListReq
	QueryUserRoleListResp  = sysclient.QueryUserRoleListResp
	ReSetPasswordReq       = sysclient.ReSetPasswordReq
	ReSetPasswordResp      = sysclient.ReSetPasswordResp
	RoleAddReq             = sysclient.RoleAddReq
	RoleAddResp            = sysclient.RoleAddResp
	RoleDeleteReq          = sysclient.RoleDeleteReq
	RoleDeleteResp         = sysclient.RoleDeleteResp
	RoleListData           = sysclient.RoleListData
	RoleListReq            = sysclient.RoleListReq
	RoleListResp           = sysclient.RoleListResp
	RoleUpdateReq          = sysclient.RoleUpdateReq
	RoleUpdateResp         = sysclient.RoleUpdateResp
	StatisticsLoginLogReq  = sysclient.StatisticsLoginLogReq
	StatisticsLoginLogResp = sysclient.StatisticsLoginLogResp
	SysLogAddReq           = sysclient.SysLogAddReq
	SysLogAddResp          = sysclient.SysLogAddResp
	SysLogDeleteReq        = sysclient.SysLogDeleteReq
	SysLogDeleteResp       = sysclient.SysLogDeleteResp
	SysLogListData         = sysclient.SysLogListData
	SysLogListReq          = sysclient.SysLogListReq
	SysLogListResp         = sysclient.SysLogListResp
	UpdateMenuRoleReq      = sysclient.UpdateMenuRoleReq
	UpdateMenuRoleResp     = sysclient.UpdateMenuRoleResp
	UpdateUserRoleListReq  = sysclient.UpdateUserRoleListReq
	UpdateUserRoleListResp = sysclient.UpdateUserRoleListResp
	UserAddReq             = sysclient.UserAddReq
	UserAddResp            = sysclient.UserAddResp
	UserDeleteReq          = sysclient.UserDeleteReq
	UserDeleteResp         = sysclient.UserDeleteResp
	UserListData           = sysclient.UserListData
	UserListReq            = sysclient.UserListReq
	UserListResp           = sysclient.UserListResp
	UserStatusReq          = sysclient.UserStatusReq
	UserStatusResp         = sysclient.UserStatusResp
	UserUpdateReq          = sysclient.UserUpdateReq
	UserUpdateResp         = sysclient.UserUpdateResp

	JobService interface {
		JobAdd(ctx context.Context, in *JobAddReq, opts ...grpc.CallOption) (*JobAddResp, error)
		JobList(ctx context.Context, in *JobListReq, opts ...grpc.CallOption) (*JobListResp, error)
		JobUpdate(ctx context.Context, in *JobUpdateReq, opts ...grpc.CallOption) (*JobUpdateResp, error)
		JobDelete(ctx context.Context, in *JobDeleteReq, opts ...grpc.CallOption) (*JobDeleteResp, error)
	}

	defaultJobService struct {
		cli zrpc.Client
	}
)

func NewJobService(cli zrpc.Client) JobService {
	return &defaultJobService{
		cli: cli,
	}
}

func (m *defaultJobService) JobAdd(ctx context.Context, in *JobAddReq, opts ...grpc.CallOption) (*JobAddResp, error) {
	client := sysclient.NewJobServiceClient(m.cli.Conn())
	return client.JobAdd(ctx, in, opts...)
}

func (m *defaultJobService) JobList(ctx context.Context, in *JobListReq, opts ...grpc.CallOption) (*JobListResp, error) {
	client := sysclient.NewJobServiceClient(m.cli.Conn())
	return client.JobList(ctx, in, opts...)
}

func (m *defaultJobService) JobUpdate(ctx context.Context, in *JobUpdateReq, opts ...grpc.CallOption) (*JobUpdateResp, error) {
	client := sysclient.NewJobServiceClient(m.cli.Conn())
	return client.JobUpdate(ctx, in, opts...)
}

func (m *defaultJobService) JobDelete(ctx context.Context, in *JobDeleteReq, opts ...grpc.CallOption) (*JobDeleteResp, error) {
	client := sysclient.NewJobServiceClient(m.cli.Conn())
	return client.JobDelete(ctx, in, opts...)
}
