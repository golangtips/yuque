//go:build wireinject
// +build wireinject

package main

import (
	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/service"
	"github.com/golangtips/yuque/service/set"
	"github.com/golangtips/yuque/transport"
	"github.com/google/wire"
)

// 公用配置
var CommonSet = wire.NewSet(
	config.NewTom,
)

// HTTP 传输层
var HTTPSet = wire.NewSet(
	transport.NewHTTPServer,
	transport.NewTemplate,
	transport.NewHandler,
	wire.Struct(new(transport.Options), "*"),
	wire.Struct(new(transport.ArticleDetail), "*"),
	wire.Struct(new(transport.About), "*"),
	wire.Struct(new(transport.CDNProxy), "*"),
	wire.Struct(new(transport.HomePage), "*"),
	wire.Struct(new(transport.Search), "*"),
)

// 服务层
var ServiceSet = wire.NewSet(
	set.NewArticle,
	service.NewYuQue,
	service.NewSegmenter,
)

// 仓储层
// var RepoSet = wire.NewSet()

func injectApp() (*App, error) {
	wire.Build(
		wire.Struct(new(App), "*"),
		CommonSet,
		HTTPSet,
		ServiceSet,
	)
	return new(App), nil
}
