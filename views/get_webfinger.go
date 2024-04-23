package views

import "github.com/gin-gonic/gin"

type Link struct {
	Rel      string `json:"rel"`
	Type     string `json:"type"`
	Href     string `json:"href,omitempty"`
	Template string `json:"template,omitempty"`
}

type WebFinger struct {
	Subject string   `json:"subject"`
	Aliases []string `json:"aliases"`
	Links   []Link   `json:"links"`
}

func GetWebFinger(c *gin.Context) {
	wf := WebFinger{
		Subject: "acct:alorenzi@giretti.alorenzi.eu",
		Aliases: []string{
			"http://localhost:8080/@alorenzi",
		},
		Links: []Link{
			{
				Href: "https://cdn.masto.host/livellosegretoit/accounts/avatars/108/752/037/243/682/607/original/c14849082f810124.jpeg",
				Rel:  "http://webfinger.net/rel/avatar",
				Type: "image/jpeg",
			},
		},
	}
	c.JSON(200, wf)
}
