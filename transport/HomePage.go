package transport

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/golangtips/yuque/util"

	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/service"
)

var _ http.Handler = (*HomePage)(nil)

type HomePage struct {
	Template *template.Template
	Service  service.IArticle
	Config   *config.Toml
}

func (h *HomePage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	total, _ := h.Service.GetTotal(r.Context())

	// 分页处理
	p := util.NewPagination(total, intPageSize, intPage, 10)
	p.Paginate()

	if intPage > p.TotalPages {
		h.Template.ExecuteTemplate(w, "404", nil)
		return
	}

	articles, _ := h.Service.GetList(r.Context(), &service.GetListRequest{
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

	err := h.Template.ExecuteTemplate(w, "home", map[string]interface{}{
		"articles":  articles.Data,
		"paginator": p,
		"site":      h.Config.Site,
		"pageSize":  intPageSize,
	})
	if err != nil {
		log.Println(err.Error())
	}
}
