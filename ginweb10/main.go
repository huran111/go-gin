package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	get := c.MustGet("name")
	fmt.Println("get....", get)
	time.Sleep(time.Second * 2)
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

func m1(c *gin.Context) {
	fmt.Println("m1 in ....")
	start := time.Now()
	c.Next() //调用 后续的处理函数
	//c.Abort()//阻止后续处理函数
	cost := time.Since(start)
	fmt.Println("耗时:", cost)
	fmt.Println("m1 out ....")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "huran")
	//	c.Abort() //阻止后续处理函数
	//return
	fmt.Println("m2 out ....")
	// 中间件中使用goroute的时候  不能 直接使用c 而是使用c.Copy()

	go func() {

	}()
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		if doCheck {
			//数据库连接
		} else {
			c.Next()
		}
	}
}

//中间件
func main() {
	r := gin.New() //默认使用不包含任何得中间件
	// r := gin.Default()//Logger  Recovery
	r.Use(m1, m2, authMiddleware(true)) //注册全局中间件
	r.GET("/index", indexHandler)
	r.GET("shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "stop",
		})
	})
	//路由组定义中间件
	xxxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxxGroup.GET("/index", func(context *gin.Context) {

		})
	}
	r.Run(":9090")
}
