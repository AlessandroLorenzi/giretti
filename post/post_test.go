package post_test

import (
	"html/template"
	"os"
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
func init() {
	err := os.Chdir("../example_site")
	if err != nil {
		panic(err)
	}
}

func TestReadPost(t *testing.T) {
	a := assert.New(t)

	p, err := post.ReadPost("posts/2024-02-25-my-first-post.md")
	a.NoError(err)

	a.Equal("My first post", p.Title)
	a.Equal([]string{"first", "post"}, p.Tags)
	a.Equal([]string{"media/posts/2024-02-25-my-first-post/track.gpx"}, p.Gpx)
	a.Equal("media/posts/2024-02-25-my-first-post/gallery/DSC07957.JPG", *p.OpenGraph.Image)
	a.Equal("This is the opengraph description", *p.OpenGraph.Description)
	a.Equal("media/posts/2024-02-25-my-first-post/gallery/DSC07957.JPG", p.Gallery()[0].Image)
	a.Equal("media/posts/2024-02-25-my-first-post/gallery/DSC07957_thumb.JPG", p.Gallery()[0].Thumbnail)

	a.Equal(2024, p.Date().Year())
	a.Equal(time.Month(2), p.Date().Month())
	a.Equal(25, p.Date().Day())

	a.Equal("/2024/02/25/my-first-post.html", p.Url())

	a.Equal(template.HTML("<p>This is my first post</p>\n"), p.HTML)

	a.Equal("2024-02-25-my-first-post", p.Id())

	a.NotNil(p.Gallery()[0].Position)
	a.Equal(45.880394, p.Gallery()[0].Position.Lat)
	a.Equal(8.903013, p.Gallery()[0].Position.Lon)

	a.Equal(45.880394, p.StartingPosition().Lat)
	a.Equal(8.903013, p.StartingPosition().Lon)
}
