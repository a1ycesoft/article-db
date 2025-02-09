// Code generated by trpc-go/trpc-cmdline v1.0.7. DO NOT EDIT.
// source: article-db.proto

package pb

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// ArticleDbService defines service.
type ArticleDbService interface {
	Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error)

	GetArticleById(ctx context.Context, req *GetArticleByIdRequest) (*GetArticleByIdResponse, error)

	InsertArticle(ctx context.Context, req *InsertArticleRequest) (*InsertArticleResponse, error)

	QueryArticleByKeyword(ctx context.Context, req *QueryArticleByKeywordRequest) (*QueryArticleByKeywordResponse, error)
}

func ArticleDbService_Hello_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &HelloRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(ArticleDbService).Hello(ctx, reqbody.(*HelloRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func ArticleDbService_GetArticleById_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &GetArticleByIdRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(ArticleDbService).GetArticleById(ctx, reqbody.(*GetArticleByIdRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func ArticleDbService_InsertArticle_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &InsertArticleRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(ArticleDbService).InsertArticle(ctx, reqbody.(*InsertArticleRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func ArticleDbService_QueryArticleByKeyword_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QueryArticleByKeywordRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(ArticleDbService).QueryArticleByKeyword(ctx, reqbody.(*QueryArticleByKeywordRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// ArticleDbServer_ServiceDesc descriptor for server.RegisterService.
var ArticleDbServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "arisu.ArticleDb",
	HandlerType: ((*ArticleDbService)(nil)),
	Methods: []server.Method{
		{
			Name: "/arisu.ArticleDb/Hello",
			Func: ArticleDbService_Hello_Handler,
		},
		{
			Name: "/arisu.ArticleDb/GetArticleById",
			Func: ArticleDbService_GetArticleById_Handler,
		},
		{
			Name: "/arisu.ArticleDb/InsertArticle",
			Func: ArticleDbService_InsertArticle_Handler,
		},
		{
			Name: "/arisu.ArticleDb/QueryArticleByKeyword",
			Func: ArticleDbService_QueryArticleByKeyword_Handler,
		},
	},
}

// RegisterArticleDbService registers service.
func RegisterArticleDbService(s server.Service, svr ArticleDbService) {
	if err := s.Register(&ArticleDbServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("ArticleDb register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedArticleDb struct{}

func (s *UnimplementedArticleDb) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, errors.New("rpc Hello of service ArticleDb is not implemented")
}
func (s *UnimplementedArticleDb) GetArticleById(ctx context.Context, req *GetArticleByIdRequest) (*GetArticleByIdResponse, error) {
	return nil, errors.New("rpc GetArticleById of service ArticleDb is not implemented")
}
func (s *UnimplementedArticleDb) InsertArticle(ctx context.Context, req *InsertArticleRequest) (*InsertArticleResponse, error) {
	return nil, errors.New("rpc InsertArticle of service ArticleDb is not implemented")
}
func (s *UnimplementedArticleDb) QueryArticleByKeyword(ctx context.Context, req *QueryArticleByKeywordRequest) (*QueryArticleByKeywordResponse, error) {
	return nil, errors.New("rpc QueryArticleByKeyword of service ArticleDb is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// ArticleDbClientProxy defines service client proxy
type ArticleDbClientProxy interface {
	Hello(ctx context.Context, req *HelloRequest, opts ...client.Option) (rsp *HelloResponse, err error)

	GetArticleById(ctx context.Context, req *GetArticleByIdRequest, opts ...client.Option) (rsp *GetArticleByIdResponse, err error)

	InsertArticle(ctx context.Context, req *InsertArticleRequest, opts ...client.Option) (rsp *InsertArticleResponse, err error)

	QueryArticleByKeyword(ctx context.Context, req *QueryArticleByKeywordRequest, opts ...client.Option) (rsp *QueryArticleByKeywordResponse, err error)
}

type ArticleDbClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewArticleDbClientProxy = func(opts ...client.Option) ArticleDbClientProxy {
	return &ArticleDbClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *ArticleDbClientProxyImpl) Hello(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/arisu.ArticleDb/Hello")
	msg.WithCalleeServiceName(ArticleDbServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("ArticleDb")
	msg.WithCalleeMethod("Hello")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &HelloResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *ArticleDbClientProxyImpl) GetArticleById(ctx context.Context, req *GetArticleByIdRequest, opts ...client.Option) (*GetArticleByIdResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/arisu.ArticleDb/GetArticleById")
	msg.WithCalleeServiceName(ArticleDbServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("ArticleDb")
	msg.WithCalleeMethod("GetArticleById")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &GetArticleByIdResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *ArticleDbClientProxyImpl) InsertArticle(ctx context.Context, req *InsertArticleRequest, opts ...client.Option) (*InsertArticleResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/arisu.ArticleDb/InsertArticle")
	msg.WithCalleeServiceName(ArticleDbServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("ArticleDb")
	msg.WithCalleeMethod("InsertArticle")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &InsertArticleResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *ArticleDbClientProxyImpl) QueryArticleByKeyword(ctx context.Context, req *QueryArticleByKeywordRequest, opts ...client.Option) (*QueryArticleByKeywordResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/arisu.ArticleDb/QueryArticleByKeyword")
	msg.WithCalleeServiceName(ArticleDbServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("ArticleDb")
	msg.WithCalleeMethod("QueryArticleByKeyword")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QueryArticleByKeywordResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
