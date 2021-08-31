package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":    "胡冉",
			"message": "hello",
			"age":     23,
		}
		data2 := gin.H{"name": "阿斯蒂芬", "age": 232}
		c.JSON(http.StatusOK, map[string]interface{}{
			"data":  data,
			"data2": data2,
		})
	})
	type msg struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	r.GET("/anther_json", func(context *gin.Context) {
		data := msg{
			Name: "sdf",
			Age:  23,
		}
		context.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
