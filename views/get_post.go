package views

import (
	"net/http"
	"strings"

	"github.com/AlessandroLorenzi/giretti/config"
	"github.com/AlessandroLorenzi/giretti/post"
	"github.com/gin-gonic/gin"
)

type Info struct {
	Config *config.ConfigStruct
	Post   *post.Post
}

func GetPost(c *gin.Context) {
	post := post.Get(urlToPostId(c.Request.URL.Path))

	c.HTML(http.StatusOK, "post.tmpl", Info{
		Config: config.Config,
		Post:   post,
	})
}

func urlToPostId(url string) string {
	// Remove leading slashes
	url = strings.TrimLeft(url, "/")

	// Split the URL path by slashes
	parts := strings.Split(url, "/")

	// Extract the date and post title
	title := strings.TrimSuffix(parts[len(parts)-1], ".html")

	// Replace slashes with hyphens
	postID := strings.Join(parts[:len(parts)-1], "-") + "-" + title

	return postID
}
