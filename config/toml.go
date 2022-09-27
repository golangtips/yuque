package config

type Toml struct {
	HTTPPort int
	Theme    string // 主题
	Site     Site   `toml:"site"`
	YuQue    YuQue  `toml:"yuque"`
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
