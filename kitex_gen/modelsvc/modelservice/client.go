// Code generated by Kitex v0.10.0. DO NOT EDIT.

package modelservice

import (
	"context"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/cloudwego/kitex/client/streamclient"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GenToken(ctx context.Context, req *modelsvc.GenTokenReq, callOptions ...callopt.Option) (r *modelsvc.GenTokenResp, err error)
	CreateModel(ctx context.Context, req *modelsvc.CreateModelReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	UpdateModel(ctx context.Context, req *modelsvc.UpdateModelReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	DeleteModel(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	GetModelByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *modelsvc.Model, err error)
	ListModel(ctx context.Context, req *modelsvc.ListModelReq, callOptions ...callopt.Option) (r *modelsvc.ListModelResp, err error)
}

// StreamClient is designed to provide Interface for Streaming APIs.
type StreamClient interface {
	Chat(ctx context.Context, req *modelsvc.ChatReq, callOptions ...streamcall.Option) (stream ModelService_ChatClient, err error)
}

type ModelService_ChatClient interface {
	streaming.Stream
	Recv() (*modelsvc.ChatResp, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kModelServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kModelServiceClient struct {
	*kClient
}

func (p *kModelServiceClient) GenToken(ctx context.Context, req *modelsvc.GenTokenReq, callOptions ...callopt.Option) (r *modelsvc.GenTokenResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GenToken(ctx, req)
}

func (p *kModelServiceClient) CreateModel(ctx context.Context, req *modelsvc.CreateModelReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateModel(ctx, req)
}

func (p *kModelServiceClient) UpdateModel(ctx context.Context, req *modelsvc.UpdateModelReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateModel(ctx, req)
}

func (p *kModelServiceClient) DeleteModel(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteModel(ctx, req)
}

func (p *kModelServiceClient) GetModelByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *modelsvc.Model, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetModelByID(ctx, req)
}

func (p *kModelServiceClient) ListModel(ctx context.Context, req *modelsvc.ListModelReq, callOptions ...callopt.Option) (r *modelsvc.ListModelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListModel(ctx, req)
}

// NewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
func NewStreamClient(destService string, opts ...streamclient.Option) (StreamClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithTransportProtocol(transport.GRPC))
	options = append(options, streamclient.GetClientOptions(opts)...)

	kc, err := client.NewClient(serviceInfoForStreamClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kModelServiceStreamClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
// It panics if any error occurs.
func MustNewStreamClient(destService string, opts ...streamclient.Option) StreamClient {
	kc, err := NewStreamClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kModelServiceStreamClient struct {
	*kClient
}

func (p *kModelServiceStreamClient) Chat(ctx context.Context, req *modelsvc.ChatReq, callOptions ...streamcall.Option) (stream ModelService_ChatClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.Chat(ctx, req)
}
