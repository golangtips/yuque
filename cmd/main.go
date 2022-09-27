package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/sdk"
	"github.com/golangtips/yuque/service"
	"github.com/golangtips/yuque/transport"
	"github.com/golangtips/yuque/util/template"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	// 解析配置文件
	var config config.Toml
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		return err
	}

	// 初始 SDK
	var yq sdk.IYuQue
	{
		c := config.YuQue
		yq = sdk.NewYuQue(c.BaseURL, c.Token, c.UserAgent, c.Namespace)
	}

	// 初始化服务
	s, err := service.NewSet(yq)
	if err != nil {
		return err
	}

	// 加载主题模板文件
	htmlTplEngine, err := template.New("htmlTplEngine", fmt.Sprintf("theme/%s/*.html", config.Theme))
	if err != nil {
		return err
	}

	handler := transport.MakeHandler(htmlTplEngine, s, config)

	server := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%d", config.HTTPPort),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	errs := make(chan error)

	// HTTP 服务
	go (func() {
		log.Printf("服务HTTP，端口:%d\n", config.HTTPPort)
		errs <- server.ListenAndServe()
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
