package post

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/gomarkdown/markdown"
)

type PostHeaders struct {
	Title     string   `yaml:"title"`
	Tags      []string `yaml:"tags"`
	Gpx       []string `yaml:"gpx"`
	OpenGraph struct {
		Image       *string `yaml:"image"`
		Description *string `yaml:"description"`
	} `yaml:"open_graph"`
	Gallery []struct {
		Image       string `yaml:"image"`
		Thumbnail   string `yaml:"thumbnail"`
		Description string `yaml:"description"`
	} `yaml:"gallery"`
}

type Post struct {
	ID      string
	Headers *PostHeaders
	HTML    template.HTML
}

// Open a file and return rendered html as a string
func ReadPost(path string) (*Post, error) {
	headers := PostHeaders{}

	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	rest, err := frontmatter.Parse(input, &headers)
	if err != nil {
		return nil, err
	}

	html := markdown.ToHTML(rest, nil, nil)

	id := strings.TrimSuffix(
		filepath.Base(path),
		filepath.Ext(path),
	)

	p := &Post{
		ID:      id,
		Headers: &headers,
		HTML:    template.HTML(string(html)),
	}

	return p, nil
}
