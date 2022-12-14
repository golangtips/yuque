// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/service"
	"github.com/golangtips/yuque/service/set"
	"github.com/golangtips/yuque/transport"
	"github.com/google/wire"
)

// Injectors from wire.go:

func injectApp() (*App, error) {
	toml, err := config.NewTom()
	if err != nil {
		return nil, err
	}
	template, err := transport.NewTemplate(toml)
	if err != nil {
		return nil, err
	}
	iYuQue := service.NewYuQue(toml)
	segmenter := service.NewSegmenter()
	iArticle, err := set.NewArticle(iYuQue, segmenter, toml)
	if err != nil {
		return nil, err
	}
	homePage := &transport.HomePage{
		Template: template,
		Service:  iArticle,
		Config:   toml,
	}
	articleDetail := &transport.ArticleDetail{
		Template: template,
		Service:  iArticle,
		Config:   toml,
	}
	search := &transport.Search{
		Template: template,
		Service:  iArticle,
		Config:   toml,
	}
	about := &transport.About{
		Template: template,
		Service:  iArticle,
		Config:   toml,
	}
	cdnProxy := &transport.CDNProxy{
		Template: template,
		Service:  iArticle,
		Config:   toml,
	}
	options := &transport.Options{
		HomePage:      homePage,
		ArticleDetail: articleDetail,
		Search:        search,
		About:         about,
		CDNProxy:      cdnProxy,
		Config:        toml,
	}
	handler := transport.NewHandler(options)
	server := transport.NewHTTPServer(toml, handler)
	app := &App{
		Config:         toml,
		HTTPServer:     server,
		ArticleService: iArticle,
	}
	return app, nil
}

// wire.go:

// 公用配置
var CommonSet = wire.NewSet(config.NewTom)

// HTTP 传输层
var HTTPSet = wire.NewSet(transport.NewHTTPServer, transport.NewTemplate, transport.NewHandler, wire.Struct(new(transport.Options), "*"), wire.Struct(new(transport.ArticleDetail), "*"), wire.Struct(new(transport.About), "*"), wire.Struct(new(transport.CDNProxy), "*"), wire.Struct(new(transport.HomePage), "*"), wire.Struct(new(transport.Search), "*"))

// 服务层
var ServiceSet = wire.NewSet(set.NewArticle, service.NewYuQue, service.NewSegmenter)
