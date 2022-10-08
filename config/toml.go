package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Toml struct {
	HTTPPort int
	Theme    string // 主题
	Site     Site   `toml:"site"`
	YuQue    YuQue  `toml:"yuque"`
}

func NewTom() (*Toml, error) {

	envFile := "config.toml"
	// 读取配置文件, 解决跑测试的时候找不到配置文件的问题，最多往上找10层目录
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(envFile); err == nil {
			break
		} else {
			envFile = "../" + envFile
		}
	}

	var config Toml
	if _, err := toml.DecodeFile(envFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

type Site struct {
	Title         string
	Copyright     string
	FeatureSlugs  []string // 功能性文章
	Tags          []string // 标签云
	Archives      []Archive
	Categories    []Category
	Newests       []Newest
	CommentSwitch bool
}

type YuQue struct {
	BaseURL   string
	Token     string
	UserAgent string
	Namespace string
}

type Archive struct {
	Title string
	Link  string
}

type Category struct {
	Title string
	Link  string
	Count int
}

type Newest struct {
	Title string
	Link  string
}
