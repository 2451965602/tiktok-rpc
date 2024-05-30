// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "tiktokrpc/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error)
	Info(ctx context.Context, req *user.InfoRequest, callOptions ...callopt.Option) (r *user.InfoResponse, err error)
	NameToInfo(ctx context.Context, req *user.NameToInfoRequest, callOptions ...callopt.Option) (r *user.NameToInfoResponse, err error)
	Upload(ctx context.Context, req *user.UploadRequest, callOptions ...callopt.Option) (r *user.UploadResponse, err error)
	MFAGet(ctx context.Context, req *user.MFAGetRequest, callOptions ...callopt.Option) (r *user.MFAGetResponse, err error)
	MFA(ctx context.Context, req *user.MFABindRequest, callOptions ...callopt.Option) (r *user.MFABindResponse, err error)
	MFAStatus(ctx context.Context, req *user.MFAStatusRequest, callOptions ...callopt.Option) (r *user.MFAStatusResponse, err error)
	UploadImages(ctx context.Context, req *user.UploadImagesRequest, callOptions ...callopt.Option) (r *user.UploadImagesResponse, err error)
	SearchImages(ctx context.Context, req *user.SearchImagesRequest, callOptions ...callopt.Option) (r *user.SearchImagesResponse, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kUserServiceClient) Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kUserServiceClient) Info(ctx context.Context, req *user.InfoRequest, callOptions ...callopt.Option) (r *user.InfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Info(ctx, req)
}

func (p *kUserServiceClient) NameToInfo(ctx context.Context, req *user.NameToInfoRequest, callOptions ...callopt.Option) (r *user.NameToInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NameToInfo(ctx, req)
}

func (p *kUserServiceClient) Upload(ctx context.Context, req *user.UploadRequest, callOptions ...callopt.Option) (r *user.UploadResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Upload(ctx, req)
}

func (p *kUserServiceClient) MFAGet(ctx context.Context, req *user.MFAGetRequest, callOptions ...callopt.Option) (r *user.MFAGetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MFAGet(ctx, req)
}

func (p *kUserServiceClient) MFA(ctx context.Context, req *user.MFABindRequest, callOptions ...callopt.Option) (r *user.MFABindResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MFA(ctx, req)
}

func (p *kUserServiceClient) MFAStatus(ctx context.Context, req *user.MFAStatusRequest, callOptions ...callopt.Option) (r *user.MFAStatusResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MFAStatus(ctx, req)
}

func (p *kUserServiceClient) UploadImages(ctx context.Context, req *user.UploadImagesRequest, callOptions ...callopt.Option) (r *user.UploadImagesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadImages(ctx, req)
}

func (p *kUserServiceClient) SearchImages(ctx context.Context, req *user.SearchImagesRequest, callOptions ...callopt.Option) (r *user.SearchImagesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchImages(ctx, req)
}
