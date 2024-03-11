package position_test

import (
	"testing"

	"github.com/AlessandroLorenzi/giretti/position"
	"github.com/stretchr/testify/assert"
)

func TestGetImagePosition(t *testing.T) {
	gpxPath := "../example_site/media/posts/2024-02-25-my-first-post/track.gpx"
	path := "../example_site/media/posts/2024-02-25-my-first-post/DSC07957.JPG"

	a := assert.New(t)
	p, err := position.GetImagePosition(path, gpxPath)

	a.Nil(err)

	assert.Equal(t, 45.880394, p.Lat)
	assert.Equal(t, 8.903013, p.Lon)
}
