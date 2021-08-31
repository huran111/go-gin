package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {

	r := gin.Default()
	//进入静态文件
	r.Static("/xxx", "./statics")
	//添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "huran.com",
		})
	})
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts.com",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.Copy()
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='www.baidu.com'>百度",
		})
	})
	//启动server
	r.Run(":9090")
}
