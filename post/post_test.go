package post_test

import (
	"giretti/post"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinleNotFound(t *testing.T) {
	a := assert.New(t)

	_, err := post.ReadPost("notfound.md")
	a.Error(err)
}

func TestRenderHTML(t *testing.T) {
	a := assert.New(t)

	p, err := post.ReadPost("../post/example.md")
	a.NoError(err)

	a.Equal("My first post", p.Headers.Title)
	a.Equal([]string{"first", "post"}, p.Headers.Tags)
	a.Equal([]string{"example.gpx"}, p.Headers.Gpx)
	a.Equal("example.jpg", *p.Headers.OpenGraph.Image)
	a.Equal("This is the opengraph description", *p.Headers.OpenGraph.Description)
	a.Equal("example.jpg", p.Headers.Gallery[0].Image)
	a.Equal("example-thumb.jpg", p.Headers.Gallery[0].Thumbnail)
	a.Equal("This is the image description", p.Headers.Gallery[0].Description)

	a.Equal(template.HTML("<p>This is my first post</p>\n"), p.HTML)

	a.Equal("example", p.ID)
}
