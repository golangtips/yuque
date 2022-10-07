package transport

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/util/xtemplate"

	"github.com/gorilla/mux"
)

type Options struct {
	HomePage      *HomePage
	ArticleDetail *ArticleDetail
	Search        *Search
	About         *About
	CDNProxy      *CDNProxy
}

func NewHandler(o *Options) http.Handler {
	r := mux.NewRouter()

	// 静态文件
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// 首页
	r.Methods("GET").Path("/").Handler(o.HomePage)

	// 文章详情页
	r.Methods("GET").Path("/articles/{slug:[a-zA-Z0-9.]+}").Handler(o.ArticleDetail)

	// 文章搜索页
	r.Methods("GET").Path("/search").Handler(o.Search)

	// 关于我
	r.Methods("GET").Path("/about").Handler(o.About)

	// 反向代理，解决跨域图片加载问题
	r.PathPrefix("/yuque/").Handler(o.CDNProxy)

	return r
}

func NewHTTPServer(config *config.Toml, handler http.Handler) *http.Server {
	server := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%d", config.HTTPPort),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	return server
}

// NewTemplate 加载主题模板文件
func NewTemplate(config *config.Toml) (*template.Template, error) {
	return xtemplate.New("htmlTplEngine", fmt.Sprintf("theme/%s/*.html", config.Theme))
}
