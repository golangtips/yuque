package internal

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"

	sdk "github.com/golangtips/yuque/sdk/intf"

	"github.com/golangtips/yuque/service/intf"
)

var _ intf.IArticle = (*Article)(nil)

type Article struct {
	YQ sdk.IYuQue
}

func NewArticle(yq sdk.IYuQue) (*Article, error) {
	return &Article{
		YQ: yq,
	}, nil
}

func (s *Article) GetList(ctx context.Context, request *intf.GetListRequest) (*intf.GetListResponse, error) {
	log.Println("GetList")
	response, err := s.YQ.GetRepoDocList(ctx, &sdk.GetRepoDocListRequest{
		Offset: (request.Page - 1) * request.PageSize,
		Limit:  request.PageSize,
	})

	if err != nil {
		return nil, fmt.Errorf("yq GetRepoDocList: %w", err)
	}

	var articles []intf.Article
	for _, item := range response.Data {
		articles = append(articles, intf.Article{
			Slug:          item.Slug,
			Title:         item.Title,
			Desc:          item.Description,
			CommentsCount: item.CommentsCount,
			Hits:          item.ReadCount,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			Author:        item.LastEditor.Name,
		})
	}

	return &intf.GetListResponse{
		Data: articles,
	}, nil
}

func (s *Article) GetDetail(ctx context.Context, request *intf.GetDetailRequest) (*intf.GetDetailResponse, error) {
	log.Println("GetDetail")
	response, err := s.YQ.GetRepoDocDetail(ctx, &sdk.GetRepoDocDetailRequest{
		Slug: request.Slug,
		Raw:  1,
	})

	if err != nil {
		return nil, fmt.Errorf("yq GetDetail: %w", err)
	}

	detail := response.Data

	// 文章内容
	content := detail.BodyHtml
	// 替换html中的cdn链接进行反向代理
	content = strings.Replace(content, "https://cdn.nlark.com/", "/", -1)
	// 文章目录
	toc := s.getToc(ctx, content)

	return &intf.GetDetailResponse{
		Data: intf.Article{
			Slug:          detail.Slug,
			Title:         detail.Title,
			Content:       content,
			Desc:          detail.Description,
			Toc:           toc,
			CommentsCount: detail.CommentsCount,
			Hits:          detail.Hits,
			CreatedAt:     detail.CreatedAt,
			UpdatedAt:     detail.UpdatedAt,
			Author:        detail.Creator.Name,
		},
	}, nil
}

func (s *Article) GetTotal(ctx context.Context) (int, error) {
	log.Println("GetTotal")
	// 获取文章总数
	all, err := s.YQ.GetRepoDocList(context.TODO(), &sdk.GetRepoDocListRequest{})
	if err != nil {
		return 0, err
	}
	return len(all.Data), nil
}

// getTocList 获取目录列表
func (s *Article) getToc(ctx context.Context, content string) []intf.H {

	var tocList []intf.H
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	doc.Find("h1, h2, h3, h4, h5").Each(func(i int, s *goquery.Selection) {

		// 提取 ID
		id, _ := s.Attr("id")
		id = strings.ReplaceAll(id, "\\", "")
		id = strings.ReplaceAll(id, "\"", "")
		// 提取标题
		title := s.Text()

		// 提取标题等级
		level := 1

		if s.Is("h1") {
			level = 1
		}

		if s.Is("h2") {
			level = 2
		}

		if s.Is("h3") {
			level = 3
		}

		if s.Is("h4") {
			level = 4
		}

		if s.Is("h5") {
			level = 5
		}

		tocList = append(tocList, intf.H{
			ID:    id,
			Title: title,
			Level: level,
		})
	})

	return tocList
}
