package internal

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/golangtips/yuque/config"

	"github.com/golangtips/yuque/sdk/intf"
)

func TestYuQue_GetRepoDocList(t *testing.T) {
	c, _ := loadConfig()

	yq := NewYuQue(c.BaseURL, c.Token, c.UserAgent, c.Namespace)

	docs, err := yq.GetRepoDocList(context.TODO(), &intf.GetRepoDocListRequest{
		Namespace: c.Namespace,
		Offset:    0,
		Limit:     1,
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(docs)
}

func TestYuQue_GetRepoDocDetail(t *testing.T) {

	c, _ := loadConfig()
	yq := NewYuQue(c.BaseURL, c.Token, c.UserAgent, c.Namespace)

	detail, err := yq.GetRepoDocDetail(context.TODO(), &intf.GetRepoDocDetailRequest{
		Namespace: c.Namespace,
		Slug:      "ua17iz",
		Raw:       1,
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(detail)

}

func TestYuQue_Search(t *testing.T) {
	//
}

func TestLoadConfig(t *testing.T) {
	c, err := loadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}

func loadConfig() (*config.YuQue, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../")
	os.Chdir(dir)
	// 解析配置文件
	var config config.Toml
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		return nil, err
	}

	return &config.YuQue, nil
}
