package intf

import (
	"context"
	"time"
)

// IArticle 文章接口
type IArticle interface {
	// GetList 文章列表
	GetList(ctx context.Context, request *GetListRequest) (*GetListResponse, error)
	// GetDetail 文章详情
	GetDetail(ctx context.Context, request *GetDetailRequest) (*GetDetailResponse, error)
	// GetTotal 文章总数
	GetTotal(ctx context.Context) (int, error)
}

type GetListRequest struct {
	Page     int
	PageSize int
}

type GetListResponse struct {
	Data []Article
}

type GetDetailRequest struct {
	Slug string
}

type GetDetailResponse struct {
	Data Article
}

type Article struct {
	Slug          string // 文章ID
	Title         string // 标题
	Content       string // 文章内容
	Desc          string // 文章简介
	Toc           []H    // 目录
	CommentsCount int    // 评论数
	Hits          int    // 阅读数
	Author        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type H struct {
	ID    string
	Title string
	Level int
}
