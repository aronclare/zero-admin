// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package cms

import (
	"context"

	"zero-admin/rpc/cms/cmsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PrefrenceAreaAddReq                    = cmsclient.PrefrenceAreaAddReq
	PrefrenceAreaAddResp                   = cmsclient.PrefrenceAreaAddResp
	PrefrenceAreaDeleteReq                 = cmsclient.PrefrenceAreaDeleteReq
	PrefrenceAreaDeleteResp                = cmsclient.PrefrenceAreaDeleteResp
	PrefrenceAreaListData                  = cmsclient.PrefrenceAreaListData
	PrefrenceAreaListReq                   = cmsclient.PrefrenceAreaListReq
	PrefrenceAreaListResp                  = cmsclient.PrefrenceAreaListResp
	PrefrenceAreaProductRelationAddReq     = cmsclient.PrefrenceAreaProductRelationAddReq
	PrefrenceAreaProductRelationAddResp    = cmsclient.PrefrenceAreaProductRelationAddResp
	PrefrenceAreaProductRelationDeleteReq  = cmsclient.PrefrenceAreaProductRelationDeleteReq
	PrefrenceAreaProductRelationDeleteResp = cmsclient.PrefrenceAreaProductRelationDeleteResp
	PrefrenceAreaProductRelationListData   = cmsclient.PrefrenceAreaProductRelationListData
	PrefrenceAreaProductRelationListReq    = cmsclient.PrefrenceAreaProductRelationListReq
	PrefrenceAreaProductRelationListResp   = cmsclient.PrefrenceAreaProductRelationListResp
	PrefrenceAreaProductRelationUpdateReq  = cmsclient.PrefrenceAreaProductRelationUpdateReq
	PrefrenceAreaProductRelationUpdateResp = cmsclient.PrefrenceAreaProductRelationUpdateResp
	PrefrenceAreaUpdateReq                 = cmsclient.PrefrenceAreaUpdateReq
	PrefrenceAreaUpdateResp                = cmsclient.PrefrenceAreaUpdateResp
	SubjectAddReq                          = cmsclient.SubjectAddReq
	SubjectAddResp                         = cmsclient.SubjectAddResp
	SubjectDeleteReq                       = cmsclient.SubjectDeleteReq
	SubjectDeleteResp                      = cmsclient.SubjectDeleteResp
	SubjectListByIdsReq                    = cmsclient.SubjectListByIdsReq
	SubjectListData                        = cmsclient.SubjectListData
	SubjectListReq                         = cmsclient.SubjectListReq
	SubjectListResp                        = cmsclient.SubjectListResp
	SubjectProductRelationAddReq           = cmsclient.SubjectProductRelationAddReq
	SubjectProductRelationAddResp          = cmsclient.SubjectProductRelationAddResp
	SubjectProductRelationDeleteReq        = cmsclient.SubjectProductRelationDeleteReq
	SubjectProductRelationDeleteResp       = cmsclient.SubjectProductRelationDeleteResp
	SubjectProductRelationListData         = cmsclient.SubjectProductRelationListData
	SubjectProductRelationListReq          = cmsclient.SubjectProductRelationListReq
	SubjectProductRelationListResp         = cmsclient.SubjectProductRelationListResp
	SubjectProductRelationUpdateReq        = cmsclient.SubjectProductRelationUpdateReq
	SubjectProductRelationUpdateResp       = cmsclient.SubjectProductRelationUpdateResp
	SubjectUpdateReq                       = cmsclient.SubjectUpdateReq
	SubjectUpdateResp                      = cmsclient.SubjectUpdateResp

	Cms interface {
		// 专题
		SubjectAdd(ctx context.Context, in *SubjectAddReq, opts ...grpc.CallOption) (*SubjectAddResp, error)
		SubjectDelete(ctx context.Context, in *SubjectDeleteReq, opts ...grpc.CallOption) (*SubjectDeleteResp, error)
		SubjectUpdate(ctx context.Context, in *SubjectUpdateReq, opts ...grpc.CallOption) (*SubjectUpdateResp, error)
		SubjectList(ctx context.Context, in *SubjectListReq, opts ...grpc.CallOption) (*SubjectListResp, error)
		SubjectListByIds(ctx context.Context, in *SubjectListByIdsReq, opts ...grpc.CallOption) (*SubjectListResp, error)
		// 专题关联
		SubjectProductRelationAdd(ctx context.Context, in *SubjectProductRelationAddReq, opts ...grpc.CallOption) (*SubjectProductRelationAddResp, error)
		SubjectProductRelationDelete(ctx context.Context, in *SubjectProductRelationDeleteReq, opts ...grpc.CallOption) (*SubjectProductRelationDeleteResp, error)
		SubjectProductRelationUpdate(ctx context.Context, in *SubjectProductRelationUpdateReq, opts ...grpc.CallOption) (*SubjectProductRelationUpdateResp, error)
		SubjectProductRelationList(ctx context.Context, in *SubjectProductRelationListReq, opts ...grpc.CallOption) (*SubjectProductRelationListResp, error)
		// 商品优选
		PrefrenceAreaAdd(ctx context.Context, in *PrefrenceAreaAddReq, opts ...grpc.CallOption) (*PrefrenceAreaAddResp, error)
		PrefrenceAreaDelete(ctx context.Context, in *PrefrenceAreaDeleteReq, opts ...grpc.CallOption) (*PrefrenceAreaDeleteResp, error)
		PrefrenceAreaUpdate(ctx context.Context, in *PrefrenceAreaUpdateReq, opts ...grpc.CallOption) (*PrefrenceAreaUpdateResp, error)
		PrefrenceAreaList(ctx context.Context, in *PrefrenceAreaListReq, opts ...grpc.CallOption) (*PrefrenceAreaListResp, error)
		// 优选商品关联
		PrefrenceAreaProductRelationAdd(ctx context.Context, in *PrefrenceAreaProductRelationAddReq, opts ...grpc.CallOption) (*PrefrenceAreaProductRelationAddResp, error)
		PrefrenceAreaProductRelationDelete(ctx context.Context, in *PrefrenceAreaProductRelationDeleteReq, opts ...grpc.CallOption) (*PrefrenceAreaProductRelationDeleteResp, error)
		PrefrenceAreaProductRelationUpdate(ctx context.Context, in *PrefrenceAreaProductRelationUpdateReq, opts ...grpc.CallOption) (*PrefrenceAreaProductRelationUpdateResp, error)
		PrefrenceAreaProductRelationList(ctx context.Context, in *PrefrenceAreaProductRelationListReq, opts ...grpc.CallOption) (*PrefrenceAreaProductRelationListResp, error)
	}

	defaultCms struct {
		cli zrpc.Client
	}
)

func NewCms(cli zrpc.Client) Cms {
	return &defaultCms{
		cli: cli,
	}
}

// 专题
func (m *defaultCms) SubjectAdd(ctx context.Context, in *SubjectAddReq, opts ...grpc.CallOption) (*SubjectAddResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectAdd(ctx, in, opts...)
}

func (m *defaultCms) SubjectDelete(ctx context.Context, in *SubjectDeleteReq, opts ...grpc.CallOption) (*SubjectDeleteResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectDelete(ctx, in, opts...)
}

func (m *defaultCms) SubjectUpdate(ctx context.Context, in *SubjectUpdateReq, opts ...grpc.CallOption) (*SubjectUpdateResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectUpdate(ctx, in, opts...)
}

func (m *defaultCms) SubjectList(ctx context.Context, in *SubjectListReq, opts ...grpc.CallOption) (*SubjectListResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectList(ctx, in, opts...)
}

func (m *defaultCms) SubjectListByIds(ctx context.Context, in *SubjectListByIdsReq, opts ...grpc.CallOption) (*SubjectListResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectListByIds(ctx, in, opts...)
}

// 专题关联
func (m *defaultCms) SubjectProductRelationAdd(ctx context.Context, in *SubjectProductRelationAddReq, opts ...grpc.CallOption) (*SubjectProductRelationAddResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectProductRelationAdd(ctx, in, opts...)
}

func (m *defaultCms) SubjectProductRelationDelete(ctx context.Context, in *SubjectProductRelationDeleteReq, opts ...grpc.CallOption) (*SubjectProductRelationDeleteResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectProductRelationDelete(ctx, in, opts...)
}

func (m *defaultCms) SubjectProductRelationUpdate(ctx context.Context, in *SubjectProductRelationUpdateReq, opts ...grpc.CallOption) (*SubjectProductRelationUpdateResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectProductRelationUpdate(ctx, in, opts...)
}

func (m *defaultCms) SubjectProductRelationList(ctx context.Context, in *SubjectProductRelationListReq, opts ...grpc.CallOption) (*SubjectProductRelationListResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectProductRelationList(ctx, in, opts...)
}

// 商品优选
func (m *defaultCms) PrefrenceAreaAdd(ctx context.Context, in *PrefrenceAreaAddReq, opts ...grpc.CallOption) (*PrefrenceAreaAddResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.PrefrenceAreaAdd(ctx, in, opts...)
}

func (m *defaultCms) PrefrenceAreaDelete(ctx context.Context, in *PrefrenceAreaDeleteReq, opts ...grpc.CallOption) (*PrefrenceAreaDeleteResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.PrefrenceAreaDelete(ctx, in, opts...)
}

func (m *defaultCms) PrefrenceAreaUpdate(ctx context.Context, in *PrefrenceAreaUpdateReq, opts ...grpc.CallOption) (*PrefrenceAreaUpdateResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.PrefrenceAreaUpdate(ctx, in, opts...)
}

func (m *defaultCms) PrefrenceAreaList(ctx context.Context, in *PrefrenceAreaListReq, opts ...grpc.CallOption) (*PrefrenceAreaListResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.PrefrenceAreaList(ctx, in, opts...)
}

// 优选商品关联
func (m *defaultCms) PrefrenceAreaProductRelationAdd(ctx context.Context, in *PrefrenceAreaProductRelationAddReq, opts ...grpc.CallOption) (*PrefrenceAreaProductRelationAddResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.PrefrenceAreaProductRelationAdd(ctx, in, opts...)
}

func (m *defaultCms) PrefrenceAreaProductRelationDelete(ctx context.Context, in *PrefrenceAreaProductRelationDeleteReq, opts ...grpc.CallOption) (*PrefrenceAreaProductRelationDeleteResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.PrefrenceAreaProductRelationDelete(ctx, in, opts...)
}

func (m *defaultCms) PrefrenceAreaProductRelationUpdate(ctx context.Context, in *PrefrenceAreaProductRelationUpdateReq, opts ...grpc.CallOption) (*PrefrenceAreaProductRelationUpdateResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.PrefrenceAreaProductRelationUpdate(ctx, in, opts...)
}

func (m *defaultCms) PrefrenceAreaProductRelationList(ctx context.Context, in *PrefrenceAreaProductRelationListReq, opts ...grpc.CallOption) (*PrefrenceAreaProductRelationListResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.PrefrenceAreaProductRelationList(ctx, in, opts...)
}
