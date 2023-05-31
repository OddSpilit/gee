package gee

import (
	"net/http"
)

/*
*
注册路由
*/
func InitRouter() {
	r := New()
	r.Static("/assets", "./gee/static")
	r.LoadHtmlGlob("./gee/templates/*")
	v1 := r.Group("v1")
	v1.Use(Logger())
	{
		v1.GET("/hello", func(c *Context) {
			c.HTML(http.StatusOK, "test.tmpl", H{
				"title":   "gee",
				"content": "hello v1",
			})
		})
	}

	v2 := r.Group("v2")
	v2.Use(Logger())
	{
		v2.GET("/hello", func(c *Context) {
			c.HTML(http.StatusOK, "test.tmpl", H{
				"title":   "gee",
				"content": "hello v2",
			})
		})
	}

	r.Run(":9999")
}
