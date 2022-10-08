package internal

import (
	"context"
	"testing"

	"github.com/golangtips/yuque/service"

	"github.com/golangtips/yuque/config"
)

func TestArticle_GetDetail(t *testing.T) {
	c, err := config.NewTom()
	if err != nil {
		t.Fatal(err)
	}

	yq := service.NewYuQue(c)
	seg := service.NewSegmenter()
	article, err := NewArticle(yq, seg, c)

	if err != nil {
		t.Fatal(err)
	}

	detail, err := article.GetDetail(context.TODO(), &service.GetDetailRequest{
		Slug: "yleyb8",
	})

	if err != nil {
		t.Fatal(err)
	}
	// ...
	h := article.GetToc(context.TODO(), detail.Data.Content)
	t.Log(h)
}
