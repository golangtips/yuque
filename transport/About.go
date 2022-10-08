package transport

import (
	"html/template"
	"log"
	"net/http"

	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/service"
)

var _ http.Handler = (*About)(nil)

type About struct {
	Template *template.Template
	Service  service.IArticle
	Config   *config.Toml
}

func (h *About) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	detail, _ := h.Service.GetDetail(r.Context(), &service.GetDetailRequest{
		Slug: "about",
	})

	err := h.Template.ExecuteTemplate(w, "about", map[string]interface{}{
		"article": detail.Data,
		"site":    h.Config.Site,
	})
	
	if err != nil {
		log.Println(err.Error())
	}
}
