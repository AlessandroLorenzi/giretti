package posts_test

import (
	"giretti/posts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostsInit(t *testing.T) {
	testDir := "../example_site/posts"

	a := assert.New(t)

	err := posts.Init(testDir)
	a.Nil(err)

	a.Equal(2, len(posts.Posts))

	a.Equal("My first post", posts.Posts[0].Headers.Title)
	a.Equal("Second post", posts.Posts[1].Headers.Title)
}
