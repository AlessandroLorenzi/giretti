package posts_test

import (
	"os"
	"testing"

	"github.com/AlessandroLorenzi/giretti/posts"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := os.Chdir("../example_site")
	if err != nil {
		panic(err)
	}
}
func TestPostsInit(t *testing.T) {
	testDir := "posts"

	a := assert.New(t)

	err := posts.Init(testDir)
	a.Nil(err)

	a.Equal(2, len(posts.Posts))

	// Check that the posts are sorted by date
	a.Equal("My first post", posts.Posts[1].Headers.Title)
	a.Equal("Second post", posts.Posts[0].Headers.Title)
}
