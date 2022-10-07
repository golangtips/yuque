package config

import (
	"github.com/BurntSushi/toml"
)

type Toml struct {
	HTTPPort int
	Theme    string // 主题
	Site     Site   `toml:"site"`
	YuQue    YuQue  `toml:"yuque"`
}

func NewTom() (*Toml, error) {
	var config Toml
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		return nil, err
	}

	return &config, nil
}

type Site struct {
	Title     string
	Copyright string
}

type YuQue struct {
	BaseURL   string
	Token     string
	UserAgent string
	Namespace string
}
