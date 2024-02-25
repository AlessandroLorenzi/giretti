package views

import (
	"giretti/config"
	"giretti/post"
	"giretti/posts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetIndexInfo struct {
	Config *config.ConfigStruct
	Posts  []*post.Post
}

func GetIndex(c *gin.Context) {
	info := GetIndexInfo{
		Config: config.Config,
		Posts:  posts.Posts,
	}
	c.HTML(http.StatusOK, "index.tmpl", info)
}
