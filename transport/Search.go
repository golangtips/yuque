package transport

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/golangtips/yuque/util"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/service"
)

var _ http.Handler = (*Search)(nil)

type Search struct {
	Template *template.Template
	Service  service.IArticle
	Config   *config.Toml
}

func (h *Search) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

	// 查询参数
	q := query.Get("q")

	allIndex := h.Service.GetIndex(r.Context())

	// 搜索词拆分
	words, _ := h.Service.Cut(r.Context(), q)

	// 匹配到的文章集合
	slugs := mapset.NewSet[string]()
	for _, word := range words {
		doc, exist := allIndex[word]
		if !exist {
			continue
		}
		slugs = slugs.Union(doc)
	}

	// 分页处理
	p := util.NewPagination(len(slugs.ToSlice()), intPageSize, intPage, 10)
	p.Paginate()

	var articles []*service.Article

	for _, slug := range slugs.ToSlice() {
		article, _ := h.Service.GetDetail(r.Context(), &service.GetDetailRequest{
			Slug: slug,
		})
		articles = append(articles, &article.Data)
	}

	//fmt.Println(articles)

	err := h.Template.ExecuteTemplate(w, "search", map[string]interface{}{
		"articles":  articles,
		"site":      h.Config.Site,
		"paginator": p,
	})

	if err != nil {
		log.Println(err.Error())
	}
}
