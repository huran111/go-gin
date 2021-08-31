package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

//文件上传
func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 //8M
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//dst := fmt.Sprintf("./%s", file.Filename)
			dst := path.Join("./", file.Filename)
			_ = c.SaveUploadedFile(file, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	})
	r.Run(":9090")
}
