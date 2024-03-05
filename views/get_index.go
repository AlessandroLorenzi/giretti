package views

import (
	"net/http"

	"github.com/AlessandroLorenzi/giretti/config"
	"github.com/AlessandroLorenzi/giretti/post"
	"github.com/AlessandroLorenzi/giretti/posts"
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
