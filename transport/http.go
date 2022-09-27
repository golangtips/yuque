package transport

import (
	"html/template"
	"net/http"

	"github.com/golangtips/yuque/config"

	"github.com/golangtips/yuque/handler"
	"github.com/golangtips/yuque/service"
	"github.com/gorilla/mux"
)

func MakeHandler(t *template.Template, s *service.Set, c config.Toml) http.Handler {
	r := mux.NewRouter()

	// 静态文件
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// 首页
	r.Methods("GET").Path("/").HandlerFunc(handler.HomePage(t, s.Article, c))

	// 文章详情页
	r.Methods("GET").Path("/articles/{slug:[a-zA-Z0-9.]+}").HandlerFunc(handler.ArticleDetail(t, s.Article, c))

	// 文章搜索页
	r.Methods("GET").Path("/search").HandlerFunc(handler.Search(t, s.Article, c))

	r.Methods("GET").Path("/about").HandlerFunc(handler.About(t, s.Article, c))

	// 反向代理，解决跨域图片加载问题
	r.PathPrefix("/yuque/").HandlerFunc(handler.CDNProxy)

	return r
}
