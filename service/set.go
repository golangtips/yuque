package service

import (
	sdk "github.com/golangtips/yuque/sdk/intf"
	"github.com/golangtips/yuque/service/cache"
	"github.com/golangtips/yuque/service/internal"
	"github.com/golangtips/yuque/service/intf"
)

type Set struct {
	Article intf.IArticle
}

func NewSet(yq sdk.IYuQue) (set *Set, err error) {

	var article intf.IArticle
	{
		article, err = internal.NewArticle(yq)
		if err != nil {
			return nil, err
		}

		article = cache.NewArticle(article)
	}

	return &Set{
		Article: article,
	}, nil
}
