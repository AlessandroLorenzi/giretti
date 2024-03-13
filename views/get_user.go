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

	c.HTML(http.StatusOK, "user.tmpl", map[string]interface{}{
		"Config": config.Config,
		"Posts":  post.GetAll(),
		"User":   u,
	})
}
