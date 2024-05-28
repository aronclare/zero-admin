// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package dictservice

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

	DictService interface {
		// 添加字典表
		AddDict(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error)
		// 删除字典表
		DeleteDict(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error)
		// 更新字典表
		UpdateDict(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error)
		// 根据条件查询单条字典表记录
		QueryDict(ctx context.Context, in *DictReq, opts ...grpc.CallOption) (*DictResp, error)
		// 查询字典表列表
		QueryDictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error)
	}

	defaultDictService struct {
		cli zrpc.Client
	}
)

func NewDictService(cli zrpc.Client) DictService {
	return &defaultDictService{
		cli: cli,
	}
}

// 添加字典表
func (m *defaultDictService) AddDict(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error) {
	client := sysclient.NewDictServiceClient(m.cli.Conn())
	return client.AddDict(ctx, in, opts...)
}

// 删除字典表
func (m *defaultDictService) DeleteDict(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error) {
	client := sysclient.NewDictServiceClient(m.cli.Conn())
	return client.DeleteDict(ctx, in, opts...)
}

// 更新字典表
func (m *defaultDictService) UpdateDict(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error) {
	client := sysclient.NewDictServiceClient(m.cli.Conn())
	return client.UpdateDict(ctx, in, opts...)
}

// 根据条件查询单条字典表记录
func (m *defaultDictService) QueryDict(ctx context.Context, in *DictReq, opts ...grpc.CallOption) (*DictResp, error) {
	client := sysclient.NewDictServiceClient(m.cli.Conn())
	return client.QueryDict(ctx, in, opts...)
}

// 查询字典表列表
func (m *defaultDictService) QueryDictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error) {
	client := sysclient.NewDictServiceClient(m.cli.Conn())
	return client.QueryDictList(ctx, in, opts...)
}
