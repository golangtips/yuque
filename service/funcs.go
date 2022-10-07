package service

import (
	"github.com/golangtips/yuque/config"
	"github.com/golangtips/yuque/sdk"
	"github.com/wangbin/jiebago"
)

func NewYuQue(config *config.Toml) sdk.IYuQue {
	c := config.YuQue
	yq := sdk.NewYuQue(c.BaseURL, c.Token, c.UserAgent, c.Namespace)
	return yq
}

func NewSegmenter() *jiebago.Segmenter {
	var seg jiebago.Segmenter
	// 初始结巴分词
	seg.LoadDictionary("dict.txt")
	return &seg
}
