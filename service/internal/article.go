package internal

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/golangtips/yuque/service"

	"github.com/wangbin/jiebago"

	"github.com/PuerkitoBio/goquery"

	sdk "github.com/golangtips/yuque/sdk/intf"

	mapset "github.com/deckarep/golang-set/v2"
)

var (
	_ service.IArticle = (*Article)(nil)

	// allIndex 文章索引, key = 搜索词, v = [文章Slug]
	allIndex = make(map[string]mapset.Set[string])
)

type Article struct {
	YQ    sdk.IYuQue
	Jieba *jiebago.Segmenter
}

func NewArticle(yq sdk.IYuQue, jieba *jiebago.Segmenter) (*Article, error) {
	return &Article{
		YQ:    yq,
		Jieba: jieba,
	}, nil
}

func (s *Article) GetList(ctx context.Context, request *service.GetListRequest) (*service.GetListResponse, error) {
	log.Println("GetList")
	response, err := s.YQ.GetRepoDocList(ctx, &sdk.GetRepoDocListRequest{
		Offset: (request.Page - 1) * request.PageSize,
		Limit:  request.PageSize,
	})

	if err != nil {
		return nil, fmt.Errorf("yq GetRepoDocList: %w", err)
	}

	var articles []service.Article
	for _, item := range response.Data {
		articles = append(articles, service.Article{
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

	//fmt.Println(articles)

	return &service.GetListResponse{
		Data: articles,
	}, nil
}

func (s *Article) GetDetail(ctx context.Context, request *service.GetDetailRequest) (*service.GetDetailResponse, error) {
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

	return &service.GetDetailResponse{
		Data: service.Article{
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

func (s *Article) PutIndex(ctx context.Context, article *service.Article) error {

	words := s.Jieba.Cut(TrimHtml(article.Title+article.Desc), true)
	//for _, word := range words {
	for word := range words {
		// 过滤2个字符以下 or 去掉标点符号
		if utf8.RuneCountInString(word) < 2 || regexp.MustCompile(`^\s|。|，|;|；|,|-|、|:|：|\.|\?|？|\(|\)|《|》|"|'$`).MatchString(word) {
			continue
		}

		if _, has := allIndex[word]; has {
			allIndex[word].Add(article.Slug)
		} else {
			set := mapset.NewSet[string]()
			set.Add(article.Slug)
			allIndex[word] = set
		}
	}

	return nil
}

func (s *Article) GetIndex(ctx context.Context) map[string]mapset.Set[string] {
	// ..
	return allIndex
}

func (s *Article) Cut(ctx context.Context, text string) ([]string, error) {
	log.Println("Cut")
	// ...
	words := s.Jieba.Cut(text, true)

	var results []string
	for word := range words {
		// 过滤2个字符以下 or 去掉标点符号
		if utf8.RuneCountInString(word) < 2 || regexp.MustCompile(`^\s|。|，|;|；|,|-|、|:|：|\.|\?|？|\(|\)|《|》|"|'$`).MatchString(word) {
			continue
		}

		word = strings.ToLower(word)
		results = append(results, word)
	}

	return results, nil
}

func (s *Article) BuildAllIndex(ctx context.Context) error {
	articles, err := s.GetList(ctx, &service.GetListRequest{})
	if err != nil {
		return err
	}

	// 构建文章索引
	for _, article := range articles.Data {
		err := s.PutIndex(context.TODO(), &article)
		if err != nil {
			return err
		}
	}

	return nil
}

// getTocList 获取目录列表
func (s *Article) getToc(ctx context.Context, content string) []service.H {

	var tocList []service.H
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

		tocList = append(tocList, service.H{
			ID:    id,
			Title: title,
			Level: level,
		})
	})

	return tocList
}

// TrimHtml 去掉字符串中的html标签
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}
