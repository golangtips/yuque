package transport

import (
	"html/template"
	"log"
	"net/http"

	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/service"
	"github.com/gorilla/mux"
)

var _ http.Handler = (*ArticleDetail)(nil)

type ArticleDetail struct {
	Template *template.Template
	Service  service.IArticle
	Config   *config.Toml
}

func (h *ArticleDetail) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	detail, _ := h.Service.GetDetail(r.Context(), &service.GetDetailRequest{
		Slug: slug,
	})

	err := h.Template.ExecuteTemplate(w, "article/detail", map[string]interface{}{
		"article": detail.Data,
		"site":    h.Config.Site,
	})
	if err != nil {
		log.Println(err.Error())
	}
}
