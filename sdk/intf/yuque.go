package intf

import (
	"context"
	"time"
)

type IYuQue interface {
	// GetRepoDocList 获取一个仓库的文档列表
	//  文档 https://www.yuque.com/yuque/developer/doc
	GetRepoDocList(ctx context.Context, request *GetRepoDocListRequest) (*GetRepoDocListResponse, error)

	// GetRepoDocDetail 获取单篇文档的详细信息
	//  文档 https://www.yuque.com/yuque/developer/doc
	GetRepoDocDetail(ctx context.Context, request *GetRepoDocDetailRequest) (*GetRepoDocDetailResponse, error)

	// Search 搜索
	//  文档 https://www.yuque.com/yuque/developer/high_level_api
	Search(ctx context.Context, request *SearchRequest) (*SearchResponse, error)
}

// GetRepoDocListRequest 获取一个仓库的文档列表
type GetRepoDocListRequest struct {
	Namespace          string //
	Offset             int    //
	Limit              int    //
	OptionalProperties int    // 获取文档浏览数
}

// GetRepoDocListResponse 获取一个仓库的文档列表
type GetRepoDocListResponse struct {
	Data []Doc `json:"data"`
}

// GetRepoDocDetailRequest 获取单篇文档的详细信息
type GetRepoDocDetailRequest struct {
	Namespace string
	Slug      string
	Raw       int // raw=1 返回文档最原始的格式
}

// GetRepoDocDetailResponse 获取单篇文档的详细信息
type GetRepoDocDetailResponse struct {
	Abilities struct {
		Update  bool `json:"update"`
		Destroy bool `json:"destroy"`
	} `json:"abilities"`
	Data DocDetail `json:"data"`
}

// SearchRequest 搜索请求
type SearchRequest struct {
	Type    string // 资源类型
	Offset  int    // 分页，1、2...
	Scope   int    // 搜索路径
	Related bool   // 搜索与我相关的传递 true
}

// SearchResponse 搜索结果
type SearchResponse struct {
	// ...
}

// Doc 文档基本信息，一般用在列表场景
// https://www.yuque.com/yuque/developer/docserializer
type Doc struct {
	Id                int         `json:"id"`
	Slug              string      `json:"slug"`
	Title             string      `json:"title"`
	Description       string      `json:"description"`
	UserId            int         `json:"user_id"`
	BookId            int         `json:"book_id"`
	Format            string      `json:"format"`
	Public            int         `json:"public"`
	Status            int         `json:"status"`
	ViewStatus        int         `json:"view_status"`
	ReadStatus        int         `json:"read_status"`
	LikesCount        int         `json:"likes_count"`
	ReadCount         int         `json:"read_count"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  time.Time   `json:"content_updated_at"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	FirstPublishedAt  time.Time   `json:"first_published_at"`
	DraftVersion      int         `json:"draft_version"`
	LastEditorId      int         `json:"last_editor_id"`
	WordCount         int         `json:"word_count"`
	Cover             interface{} `json:"cover"`
	CustomDescription string      `json:"custom_description"`
	LastEditor        struct {
		Id             int       `json:"id"`
		Type           string    `json:"type"`
		Login          string    `json:"login"`
		Name           string    `json:"name"`
		Description    string    `json:"description"`
		AvatarUrl      string    `json:"avatar_url"`
		FollowersCount int       `json:"followers_count"`
		FollowingCount int       `json:"following_count"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		Serializer     string    `json:"_serializer"`
	} `json:"last_editor"`
	Book       interface{} `json:"book"`
	Serializer string      `json:"_serializer"`
}

// DocDetail 文档详细信息
// https://www.yuque.com/yuque/developer/docdetailserializer
type DocDetail struct {
	Id     int    `json:"id"`
	Slug   string `json:"slug"`
	Title  string `json:"title"`
	BookId int    `json:"book_id"`
	Book   struct {
		Id               int       `json:"id"`
		Type             string    `json:"type"`
		Slug             string    `json:"slug"`
		Name             string    `json:"name"`
		UserId           int       `json:"user_id"`
		Description      string    `json:"description"`
		CreatorId        int       `json:"creator_id"`
		Public           int       `json:"public"`
		ItemsCount       int       `json:"items_count"`
		LikesCount       int       `json:"likes_count"`
		WatchesCount     int       `json:"watches_count"`
		ContentUpdatedAt time.Time `json:"content_updated_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		CreatedAt        time.Time `json:"created_at"`
		Namespace        string    `json:"namespace"`
		User             struct {
			Id               int         `json:"id"`
			Type             string      `json:"type"`
			Login            string      `json:"login"`
			Name             string      `json:"name"`
			Description      interface{} `json:"description"`
			AvatarUrl        string      `json:"avatar_url"`
			BooksCount       int         `json:"books_count"`
			PublicBooksCount int         `json:"public_books_count"`
			FollowersCount   int         `json:"followers_count"`
			FollowingCount   int         `json:"following_count"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			Serializer       string      `json:"_serializer"`
		} `json:"user"`
		Serializer string `json:"_serializer"`
	} `json:"book"`
	UserId  int `json:"user_id"`
	Creator struct {
		Id               int         `json:"id"`
		Type             string      `json:"type"`
		Login            string      `json:"login"`
		Name             string      `json:"name"`
		Description      interface{} `json:"description"`
		AvatarUrl        string      `json:"avatar_url"`
		BooksCount       int         `json:"books_count"`
		PublicBooksCount int         `json:"public_books_count"`
		FollowersCount   int         `json:"followers_count"`
		FollowingCount   int         `json:"following_count"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		Serializer       string      `json:"_serializer"`
	} `json:"creator"`
	Format            string      `json:"format"`
	Body              string      `json:"body"`
	BodyDraft         string      `json:"body_draft"`
	BodyHtml          string      `json:"body_html"`
	BodyLake          string      `json:"body_lake"`
	BodyDraftLake     string      `json:"body_draft_lake"`
	Public            int         `json:"public"`
	Status            int         `json:"status"`
	ViewStatus        int         `json:"view_status"`
	ReadStatus        int         `json:"read_status"`
	LikesCount        int         `json:"likes_count"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  time.Time   `json:"content_updated_at"`
	DeletedAt         time.Time   `json:"deleted_at"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	FirstPublishedAt  time.Time   `json:"first_published_at"`
	WordCount         int         `json:"word_count"`
	Cover             interface{} `json:"cover"`
	Description       string      `json:"description"`
	CustomDescription interface{} `json:"custom_description"`
	Hits              int         `json:"hits"`
	Serializer        string      `json:"_serializer"`
}
