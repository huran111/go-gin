package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//url参数
func main() {
	r := gin.Default()
	r.GET("/users/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		fmt.Println(name, age)
	})
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	r.Run(":9090")

}
