package transport

import (
	"html/template"
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
	h.Template.ExecuteTemplate(w, "about", map[string]interface{}{
		"content": "About",
		"site":    h.Config.Site,
	})
}
