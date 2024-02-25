package main

import (
	"fmt"
	"giretti/views"

	"github.com/gin-gonic/gin"
)

var baseDir = "../giretti.alorenzi.eu" // TODO: from cli

func main() {
	fmt.Println("Giretti")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.StaticFS("/media", gin.Dir(baseDir+"/media", false))

	r.GET("/:year/:month/:day/:title", views.GetPost)
	r.Run()
}
