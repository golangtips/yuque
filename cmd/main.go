package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/golangtips/yuque/service"

	"github.com/golangtips/yuque/config"
)

func main() {
	// 注入 App && 运行
	app, _ := injectApp()
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

type App struct {
	Config         *config.Toml
	HTTPServer     *http.Server
	ArticleService service.IArticle
}

func (app *App) Run() error {

	// 构建所有文章索引
	err := app.ArticleService.BuildAllIndex(context.Background())
	if err != nil {
		return err
	}

	errs := make(chan error)

	// HTTP 服务
	go (func() {
		log.Printf("服务HTTP，端口:%d\n", app.Config.HTTPPort)
		errs <- app.HTTPServer.ListenAndServe()
	})()

	// 退出信号处理
	go (func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-ch)
	})()

	// 退出日志
	log.Println("exit", <-errs)

	return nil
}
