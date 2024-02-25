package main

import (
	"fmt"
	"giretti/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

var baseDir = "../giretti.alorenzi.eu" // TODO: from cli

func main() {
	fmt.Println("Giretti")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.StaticFS("/media", gin.Dir(baseDir+"/media", false))

	r.GET("/", func(c *gin.Context) {
		p, err := post.ReadPost("./post/example.md")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		Info := struct {
			Post *post.Post
		}{
			Post: p,
		}
		c.HTML(http.StatusOK, "post.tmpl", Info)
	})
	r.Run()
}
