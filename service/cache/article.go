package cache

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/golangtips/yuque/service/intf"
)

var _ intf.IArticle = (*Article)(nil)

type Article struct {
	intf.IArticle
	cache *cache.Cache
}

func NewArticle(next intf.IArticle) *Article {
	// 默认5分钟缓存，每1小时清除过期项目
	c := cache.New(30*time.Minute, 1*time.Hour)
	return &Article{
		IArticle: next,
		cache:    c,
	}
}

func (s *Article) GetList(ctx context.Context, request *intf.GetListRequest) (*intf.GetListResponse, error) {

	key := s.buildCacheKey("GetList", request)
	if value, ok := s.cache.Get(key); ok {
		foo := value.(*intf.GetListResponse)
		return foo, nil
	}

	value, err := s.IArticle.GetList(ctx, request)
	if err != nil {
		return nil, err
	}

	//fmt.Println("cache === ", value)
	s.cache.Set(key, value, cache.DefaultExpiration)

	return value, nil
}

func (s *Article) GetDetail(ctx context.Context, request *intf.GetDetailRequest) (*intf.GetDetailResponse, error) {

	key := s.buildCacheKey("GetDetail", request)
	if value, ok := s.cache.Get(key); ok {
		foo := value.(*intf.GetDetailResponse)
		return foo, nil
	}

	value, err := s.IArticle.GetDetail(ctx, request)
	if err != nil {
		return nil, err
	}

	s.cache.Set(key, value, cache.DefaultExpiration)

	return value, nil
}

func (s *Article) GetTotal(ctx context.Context) (int, error) {

	key := "GetTotal"
	if value, ok := s.cache.Get(key); ok {
		foo := value.(int)
		return foo, nil
	}

	value, err := s.IArticle.GetTotal(ctx)
	if err != nil {
		return 0, err
	}

	s.cache.Set(key, value, cache.DefaultExpiration)

	return value, nil
}

func (s *Article) buildCacheKey(prefix string, request interface{}) string {
	x, _ := json.Marshal(request)
	return fmt.Sprintf("%s:%x", prefix, md5.Sum(x))
}
