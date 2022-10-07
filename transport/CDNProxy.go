package transport

import (
	"html/template"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/service"
)

var _ http.Handler = (*CDNProxy)(nil)

// CDNProxy 反向代理，解决跨域图片加载问题
type CDNProxy struct {
	Template *template.Template
	Service  service.IArticle
	Config   *config.Toml
}

func (h *CDNProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
