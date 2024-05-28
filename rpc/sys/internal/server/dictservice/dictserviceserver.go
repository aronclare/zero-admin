// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/logic/dictservice"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
)

type DictServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedDictServiceServer
}

func NewDictServiceServer(svcCtx *svc.ServiceContext) *DictServiceServer {
	return &DictServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加字典表
func (s *DictServiceServer) AddDict(ctx context.Context, in *sysclient.DictAddReq) (*sysclient.DictAddResp, error) {
	l := dictservicelogic.NewAddDictLogic(ctx, s.svcCtx)
	return l.AddDict(in)
}

// 删除字典表
func (s *DictServiceServer) DeleteDict(ctx context.Context, in *sysclient.DictDeleteReq) (*sysclient.DictDeleteResp, error) {
	l := dictservicelogic.NewDeleteDictLogic(ctx, s.svcCtx)
	return l.DeleteDict(in)
}

// 更新字典表
func (s *DictServiceServer) UpdateDict(ctx context.Context, in *sysclient.DictUpdateReq) (*sysclient.DictUpdateResp, error) {
	l := dictservicelogic.NewUpdateDictLogic(ctx, s.svcCtx)
	return l.UpdateDict(in)
}

// 根据条件查询单条字典表记录
func (s *DictServiceServer) QueryDict(ctx context.Context, in *sysclient.DictReq) (*sysclient.DictResp, error) {
	l := dictservicelogic.NewQueryDictLogic(ctx, s.svcCtx)
	return l.QueryDict(in)
}

// 查询字典表列表
func (s *DictServiceServer) QueryDictList(ctx context.Context, in *sysclient.DictListReq) (*sysclient.DictListResp, error) {
	l := dictservicelogic.NewQueryDictListLogic(ctx, s.svcCtx)
	return l.QueryDictList(in)
}
