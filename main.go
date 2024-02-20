package main

import (
	"fmt"
	"giretti/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Giretti")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		p, err := post.ReadPost("./post/example.md")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.HTML(http.StatusOK, "post.tmpl", p)
	})
	r.Run()
}
