// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/logic/roleservice"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
)

type RoleServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedRoleServiceServer
}

func NewRoleServiceServer(svcCtx *svc.ServiceContext) *RoleServiceServer {
	return &RoleServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加角色信息表
func (s *RoleServiceServer) AddRole(ctx context.Context, in *sysclient.AddRoleReq) (*sysclient.AddRoleResp, error) {
	l := roleservicelogic.NewAddRoleLogic(ctx, s.svcCtx)
	return l.AddRole(in)
}

// 删除角色信息表
func (s *RoleServiceServer) DeleteRole(ctx context.Context, in *sysclient.DeleteRoleReq) (*sysclient.DeleteRoleResp, error) {
	l := roleservicelogic.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

// 更新角色信息表
func (s *RoleServiceServer) UpdateRole(ctx context.Context, in *sysclient.UpdateRoleReq) (*sysclient.UpdateRoleResp, error) {
	l := roleservicelogic.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

// 更新角色信息表状态
func (s *RoleServiceServer) UpdateRoleStatus(ctx context.Context, in *sysclient.UpdateRoleStatusReq) (*sysclient.UpdateRoleStatusResp, error) {
	l := roleservicelogic.NewUpdateRoleStatusLogic(ctx, s.svcCtx)
	return l.UpdateRoleStatus(in)
}

// 查询角色信息表详情
func (s *RoleServiceServer) QueryRoleDetail(ctx context.Context, in *sysclient.QueryRoleDetailReq) (*sysclient.QueryRoleDetailResp, error) {
	l := roleservicelogic.NewQueryRoleDetailLogic(ctx, s.svcCtx)
	return l.QueryRoleDetail(in)
}

// 查询角色信息表列表
func (s *RoleServiceServer) QueryRoleList(ctx context.Context, in *sysclient.QueryRoleListReq) (*sysclient.QueryRoleListResp, error) {
	l := roleservicelogic.NewQueryRoleListLogic(ctx, s.svcCtx)
	return l.QueryRoleList(in)
}

// 查询用户与角色的关联
func (s *RoleServiceServer) QueryRoleMenuList(ctx context.Context, in *sysclient.QueryRoleMenuListReq) (*sysclient.QueryRoleMenuListResp, error) {
	l := roleservicelogic.NewQueryRoleMenuListLogic(ctx, s.svcCtx)
	return l.QueryRoleMenuList(in)
}

// 更新用户与角色的关联
func (s *RoleServiceServer) UpdateMenuRoleList(ctx context.Context, in *sysclient.UpdateMenuRoleReq) (*sysclient.UpdateMenuRoleResp, error) {
	l := roleservicelogic.NewUpdateMenuRoleListLogic(ctx, s.svcCtx)
	return l.UpdateMenuRoleList(in)
}

// 查询角色的用户关联
func (s *RoleServiceServer) QueryRoleUserList(ctx context.Context, in *sysclient.QueryRoleUserListReq) (*sysclient.QueryRoleUserListResp, error) {
	l := roleservicelogic.NewQueryRoleUserListLogic(ctx, s.svcCtx)
	return l.QueryRoleUserList(in)
}

// 更新角色的用户关联
func (s *RoleServiceServer) UpdateRoleUserList(ctx context.Context, in *sysclient.UpdateRoleUserListReq) (*sysclient.UpdateRoleUserListResp, error) {
	l := roleservicelogic.NewUpdateRoleUserListLogic(ctx, s.svcCtx)
	return l.UpdateRoleUserList(in)
}

// 取消授权
func (s *RoleServiceServer) CancelAuthorization(ctx context.Context, in *sysclient.CancelAuthorizationReq) (*sysclient.CancelAuthorizationResp, error) {
	l := roleservicelogic.NewCancelAuthorizationLogic(ctx, s.svcCtx)
	return l.CancelAuthorization(in)
}
