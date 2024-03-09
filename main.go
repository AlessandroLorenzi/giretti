package main

import (
	"embed"
	"fmt"
	"html/template"
	"os"

	"github.com/AlessandroLorenzi/giretti/config"
	"github.com/AlessandroLorenzi/giretti/posts"
	"github.com/AlessandroLorenzi/giretti/views"
	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

func main() {
	fmt.Println("Giretti")

	os.Chdir(getBaseDirFromArgs())

	if err := posts.Init("posts"); err != nil {
		fmt.Println("Error initializing posts", err)
		return
	}

	if err := config.Init("config.yaml"); err != nil {
		fmt.Println("Error initializing config", err)
		return
	}

	r := gin.Default()
	templ := template.Must(
		template.New("").ParseFS(f, "templates/*.tmpl"),
	)
	r.SetHTMLTemplate(templ)

	r.StaticFS("/media", gin.Dir("media", false))

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
