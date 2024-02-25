package main

import (
	"fmt"
	"giretti/config"
	"giretti/posts"
	"giretti/views"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Giretti")

	var baseDir = getBaseDirFromArgs()

	if err := posts.Init(baseDir + "/posts"); err != nil {
		fmt.Println("Error initializing posts", err)
		return
	}

	if err := config.Init(baseDir + "/config.yaml"); err != nil {
		fmt.Println("Error initializing config", err)
		return
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.StaticFS("/media", gin.Dir(baseDir+"/media", false))

	r.GET("/", views.GetIndex)
	r.GET("/:year/:month/:day/:title", views.GetPost)
	r.Run()
}

func getBaseDirFromArgs() string {
	args := os.Args
	if len(args) > 1 {
		return args[1]
	}
	return "."
}
