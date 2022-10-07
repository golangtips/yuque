package set

import (
	sdk "github.com/golangtips/yuque/sdk/intf"
	"github.com/golangtips/yuque/service"
	"github.com/golangtips/yuque/service/cache"
	"github.com/golangtips/yuque/service/internal"
	"github.com/wangbin/jiebago"
)

func NewArticle(yq sdk.IYuQue, jieba *jiebago.Segmenter) (service.IArticle, error) {
	var article service.IArticle
	article, err := internal.NewArticle(yq, jieba)
	if err != nil {
		return nil, err
	}

	article = cache.NewArticle(article)
	return article, nil
}
