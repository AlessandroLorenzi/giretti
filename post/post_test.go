package post_test

import (
	"html/template"
	"testing"
	"time"

	"github.com/AlessandroLorenzi/giretti/post"
	"github.com/stretchr/testify/assert"
)

func TestFinleNotFound(t *testing.T) {
	a := assert.New(t)

	_, err := post.ReadPost("notfound.md")
	a.Error(err)
}

func TestRenderHTML(t *testing.T) {
	a := assert.New(t)

	p, err := post.ReadPost("../example_site/posts/2024-02-25-my-first-post.md")
	a.NoError(err)

	a.Equal("My first post", p.Headers.Title)
	a.Equal([]string{"first", "post"}, p.Headers.Tags)
	a.Equal([]string{"example.gpx"}, p.Headers.Gpx)
	a.Equal("example.jpg", *p.Headers.OpenGraph.Image)
	a.Equal("This is the opengraph description", *p.Headers.OpenGraph.Description)
	a.Equal("example.jpg", p.Headers.Gallery[0].Image)
	a.Equal("example-thumb.jpg", p.Headers.Gallery[0].Thumbnail)
	a.Equal("This is the image description", p.Headers.Gallery[0].Description)

	a.Equal(2024, p.Date.Year())
	a.Equal(time.Month(2), p.Date.Month())
	a.Equal(25, p.Date.Day())

	a.Equal("/2024/02/25/my-first-post.html", p.Url)

	a.Equal(template.HTML("<p>This is my first post</p>\n"), p.HTML)

	a.Equal("2024-02-25-my-first-post", p.ID)
}
