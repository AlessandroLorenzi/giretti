package views

import (
	"net/http"

	"github.com/AlessandroLorenzi/giretti/config"
	"github.com/AlessandroLorenzi/giretti/post"
	"github.com/AlessandroLorenzi/giretti/posts"
	"github.com/gin-gonic/gin"
)

type Info struct {
	Config *config.ConfigStruct
	Post   *post.Post
}

func GetPost(c *gin.Context) {
	post := posts.GetFromUrl(c.Request.URL.Path)

	c.HTML(http.StatusOK, "post.tmpl", Info{
		Config: config.Config,
		Post:   post,
	})
}
