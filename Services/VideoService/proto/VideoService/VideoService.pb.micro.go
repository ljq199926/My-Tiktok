// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/VideoService/VideoService.proto

package VideoService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VideoService service

type VideoService interface {
	PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...client.CallOption) (*DouyinPublishActionResponse, error)
	Feed(ctx context.Context, in *DouyinFeedRequest, opts ...client.CallOption) (*DouyinFeedResponse, error)
	PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...client.CallOption) (*DouyinPublishListResponse, error)
	FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...client.CallOption) (*DouyinFavoriteActionResponse, error)
	FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...client.CallOption) (*DouyinFavoriteListResponse, error)
	CommentAction(ctx context.Context, in *DouyinCommentActionRequest, opts ...client.CallOption) (*DouyinCommentActionResponse, error)
	CommentList(ctx context.Context, in *DouyinCommentListRequest, opts ...client.CallOption) (*DouyinCommentListResponse, error)
}

type videoService struct {
	c    client.Client
	name string
}

func NewVideoService(name string, c client.Client) VideoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.service.VideoService"
	}
	return &videoService{
		c:    c,
		name: name,
	}
}

func (c *videoService) PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...client.CallOption) (*DouyinPublishActionResponse, error) {
	req := c.c.NewRequest(c.name, "VideoService.PublishAction", in)
	out := new(DouyinPublishActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoService) Feed(ctx context.Context, in *DouyinFeedRequest, opts ...client.CallOption) (*DouyinFeedResponse, error) {
	req := c.c.NewRequest(c.name, "VideoService.Feed", in)
	out := new(DouyinFeedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoService) PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...client.CallOption) (*DouyinPublishListResponse, error) {
	req := c.c.NewRequest(c.name, "VideoService.PublishList", in)
	out := new(DouyinPublishListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoService) FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...client.CallOption) (*DouyinFavoriteActionResponse, error) {
	req := c.c.NewRequest(c.name, "VideoService.FavoriteAction", in)
	out := new(DouyinFavoriteActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoService) FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...client.CallOption) (*DouyinFavoriteListResponse, error) {
	req := c.c.NewRequest(c.name, "VideoService.FavoriteList", in)
	out := new(DouyinFavoriteListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoService) CommentAction(ctx context.Context, in *DouyinCommentActionRequest, opts ...client.CallOption) (*DouyinCommentActionResponse, error) {
	req := c.c.NewRequest(c.name, "VideoService.CommentAction", in)
	out := new(DouyinCommentActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoService) CommentList(ctx context.Context, in *DouyinCommentListRequest, opts ...client.CallOption) (*DouyinCommentListResponse, error) {
	req := c.c.NewRequest(c.name, "VideoService.CommentList", in)
	out := new(DouyinCommentListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VideoService service

type VideoServiceHandler interface {
	PublishAction(context.Context, *DouyinPublishActionRequest, *DouyinPublishActionResponse) error
	Feed(context.Context, *DouyinFeedRequest, *DouyinFeedResponse) error
	PublishList(context.Context, *DouyinPublishListRequest, *DouyinPublishListResponse) error
	FavoriteAction(context.Context, *DouyinFavoriteActionRequest, *DouyinFavoriteActionResponse) error
	FavoriteList(context.Context, *DouyinFavoriteListRequest, *DouyinFavoriteListResponse) error
	CommentAction(context.Context, *DouyinCommentActionRequest, *DouyinCommentActionResponse) error
	CommentList(context.Context, *DouyinCommentListRequest, *DouyinCommentListResponse) error
}

func RegisterVideoServiceHandler(s server.Server, hdlr VideoServiceHandler, opts ...server.HandlerOption) error {
	type videoService interface {
		PublishAction(ctx context.Context, in *DouyinPublishActionRequest, out *DouyinPublishActionResponse) error
		Feed(ctx context.Context, in *DouyinFeedRequest, out *DouyinFeedResponse) error
		PublishList(ctx context.Context, in *DouyinPublishListRequest, out *DouyinPublishListResponse) error
		FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, out *DouyinFavoriteActionResponse) error
		FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, out *DouyinFavoriteListResponse) error
		CommentAction(ctx context.Context, in *DouyinCommentActionRequest, out *DouyinCommentActionResponse) error
		CommentList(ctx context.Context, in *DouyinCommentListRequest, out *DouyinCommentListResponse) error
	}
	type VideoService struct {
		videoService
	}
	h := &videoServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&VideoService{h}, opts...))
}

type videoServiceHandler struct {
	VideoServiceHandler
}

func (h *videoServiceHandler) PublishAction(ctx context.Context, in *DouyinPublishActionRequest, out *DouyinPublishActionResponse) error {
	return h.VideoServiceHandler.PublishAction(ctx, in, out)
}

func (h *videoServiceHandler) Feed(ctx context.Context, in *DouyinFeedRequest, out *DouyinFeedResponse) error {
	return h.VideoServiceHandler.Feed(ctx, in, out)
}

func (h *videoServiceHandler) PublishList(ctx context.Context, in *DouyinPublishListRequest, out *DouyinPublishListResponse) error {
	return h.VideoServiceHandler.PublishList(ctx, in, out)
}

func (h *videoServiceHandler) FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, out *DouyinFavoriteActionResponse) error {
	return h.VideoServiceHandler.FavoriteAction(ctx, in, out)
}

func (h *videoServiceHandler) FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, out *DouyinFavoriteListResponse) error {
	return h.VideoServiceHandler.FavoriteList(ctx, in, out)
}

func (h *videoServiceHandler) CommentAction(ctx context.Context, in *DouyinCommentActionRequest, out *DouyinCommentActionResponse) error {
	return h.VideoServiceHandler.CommentAction(ctx, in, out)
}

func (h *videoServiceHandler) CommentList(ctx context.Context, in *DouyinCommentListRequest, out *DouyinCommentListResponse) error {
	return h.VideoServiceHandler.CommentList(ctx, in, out)
}
