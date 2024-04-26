package views

import (
	"net/http"

	"github.com/AlessandroLorenzi/giretti/config"
	"github.com/AlessandroLorenzi/giretti/post"
	"github.com/AlessandroLorenzi/giretti/user"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	username := c.Param("username")
	u := user.GetByUsername(username)
	if u == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if c.GetHeader("Accept") == "application/activity+json" {
		c.JSON(http.StatusOK, u.ToActivityPub())
		return
	}

	c.HTML(http.StatusOK, "user.tmpl", map[string]interface{}{
		"Config": config.Config,
		"Posts":  post.GetAll(),
		"User":   u,
	})
}
