package views

import (
	"giretti/config"
	"giretti/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Info struct {
	Config *config.Config
	Post   *post.Post
}

func GetPost(c *gin.Context) {
	post, err := post.ReadPost("./post/example.md")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	config, err := config.ParseConfig("./config/config.yaml")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	i := Info{
		Config: config,
		Post:   post,
	}
	c.HTML(http.StatusOK, "post.tmpl", i)
}
