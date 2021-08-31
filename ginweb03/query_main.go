package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//name:=c.Query("query")
		//name := c.DefaultQuery("query", "somebody")
		name, ok := c.GetQuery("query")
		if !ok {
			name = "somebody"
		}

		age := c.Query("age")
		fmt.Println("================>>>", age)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":9090")
}
