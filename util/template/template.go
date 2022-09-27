package template

import (
	"html/template"
	"time"
)

func New(name string, patterns ...string) (*template.Template, error) {
	// 初始化模板引擎 并加载模板文件
	htmlTplEngine := template.New(name).Funcs(template.FuncMap{
		"unescapeHTML": unescapeHTML,
		"timeFormat":   timeFormat,
	})

	for _, pattern := range patterns {
		// 模板根目录下的模板文件 一些公共文件
		_, err := htmlTplEngine.ParseGlob(pattern)
		if err != nil {
			return nil, err
		}
	}

	return htmlTplEngine, nil
}

// unescapeHTML 自定义模版函数
func unescapeHTML(s string) template.HTML {
	return template.HTML(s)
}

func timeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
