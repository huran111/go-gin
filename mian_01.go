package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello HuRan",
	})

}
func main() {
	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default() //返回默认引擎
	//自定用户使用get请求访问	执行sayHello方法
	r.GET("/hello", sayHello)
	//启动服务
	r.Run()
}
