package post_test

import (
	"os"
	"testing"

	"github.com/AlessandroLorenzi/giretti/post"
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

	err := post.InitRepo(testDir)
	posts := post.GetAll()
	a.Nil(err)

	a.Equal(2, len(posts))

	// Check that the posts are sorted by date
	a.Equal("My first post", posts[1].Title)
	a.Equal("Second post", posts[0].Title)
}
