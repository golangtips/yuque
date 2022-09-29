package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/golangtips/yuque/config"

	service "github.com/golangtips/yuque/service/intf"
	"github.com/golangtips/yuque/util"
	"github.com/gorilla/mux"
)

// HomePage 首页
func HomePage(t *template.Template, s service.IArticle, c config.Toml) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 分页参数
		query := r.URL.Query()
		page := query.Get("page")
		pageSize := query.Get("pageSize")

		intPage, _ := strconv.Atoi(page)
		intPageSize, _ := strconv.Atoi(pageSize)
		if intPage < 1 {
			intPage = 1
		}
		if intPageSize < 1 {
			intPageSize = 10
		}

		total, _ := s.GetTotal(r.Context())

		// 分页处理
		p := util.NewPagination(total, intPageSize, intPage, 10)
		p.Paginate()

		if intPage > p.TotalPages {
			t.ExecuteTemplate(w, "404", nil)
			return
		}

		articles, _ := s.GetList(r.Context(), &service.GetListRequest{
			Page:     intPage,
			PageSize: intPageSize,
		})

		// 最新文章
		// TODO...

		// 归档
		// TODO...

		// 分类
		// TODO...

		// 标签云
		// TODO...

		// RSS 订阅
		// TODO...

		err := t.ExecuteTemplate(w, "home", map[string]interface{}{
			"articles":  articles.Data,
			"paginator": p,
			"site":      c.Site,
			"pageSize":  intPageSize,
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

// ArticleDetail 文章详情页
func ArticleDetail(t *template.Template, s service.IArticle, c config.Toml) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		slug := params["slug"]

		detail, _ := s.GetDetail(r.Context(), &service.GetDetailRequest{
			Slug: slug,
		})

		err := t.ExecuteTemplate(w, "article/detail", map[string]interface{}{
			"article": detail.Data,
			"site":    c.Site,
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

// Search 搜索结果页
func Search(t *template.Template, s service.IArticle, c config.Toml) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 分页参数
		query := r.URL.Query()
		q := query.Get("q")

		allIndex := s.GetIndex(r.Context())

		// 搜索词拆分
		words, _ := s.Cut(r.Context(), q)

		// 匹配到的文章集合
		slugs := mapset.NewSet[string]()
		for _, word := range words {
			doc, exist := allIndex[word]
			if !exist {
				continue
			}
			slugs = slugs.Union(doc)
		}

		fmt.Println(slugs.ToSlice())

		var articles []*service.Article

		for _, slug := range slugs.ToSlice() {
			article, _ := s.GetDetail(r.Context(), &service.GetDetailRequest{
				Slug: slug,
			})
			articles = append(articles, &article.Data)
		}

		fmt.Println(articles)

		err := t.ExecuteTemplate(w, "search", map[string]interface{}{
			"articles": articles,
			"site":     c.Site,
		})

		if err != nil {
			log.Println(err.Error())
		}
	}
}

func About(t *template.Template, s service.IArticle, c config.Toml) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "about", map[string]interface{}{
			"content": "About",
			"site":    c.Site,
		})
	}
}

// CDNProxy 反向代理，解决跨域图片加载问题
func CDNProxy(w http.ResponseWriter, r *http.Request) {

	remote, err := url.Parse("https://cdn.nlark.com")
	if err != nil {
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	d := proxy.Director
	proxy.Director = func(r *http.Request) {
		r.Header.Set("Referer", "")
		r.Host = remote.Host
		d(r)
	}
	proxy.ServeHTTP(w, r)
}
