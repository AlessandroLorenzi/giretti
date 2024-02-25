package views

import (
	"giretti/config"
	"giretti/post"
	"giretti/posts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Info struct {
	Config *config.ConfigStruct
	Post   *post.Post
}

func GetPost(c *gin.Context) {
	post := posts.GetFromUrl(c.Request.URL.Path)

	i := Info{
		Config: config.Config,
		Post:   post,
	}
	c.HTML(http.StatusOK, "post.tmpl", i)
}
