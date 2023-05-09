package gee

import "net/http"

/**
注册路由
*/
func InitRouter() {
	r := New()
	r.Static("/assets", "./gee/static")
	r.LoadHtmlGlob("./gee/templates/*")
	r.Group("v1")
	{
		r.GET("/", func(c *Context) {
			c.HTML(http.StatusOK, "test.tmpl", H{
				"title":   "gee",
				"content": "hello world",
			})
		})
	}

	r.Run(":9999")
}
