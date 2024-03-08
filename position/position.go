package position

import (
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/tkrajina/gpxgo/gpx"
)

type Position struct {
	Lat float64 `yaml:"lat"`
	Lon float64 `yaml:"lon"`
}

type positionAtTime struct {
	position Position
	time     time.Time
}

func GetImagePosition(path string, gpxPath string) (*Position, error) {
	shootingTime, err := getShootingTime(path)
	if err != nil {
		return nil, err
	}

	positions, err := getPositionsFromGpx(gpxPath)
	if err != nil {
		return nil, err
	}

	return getNearestPointInTime(positions, *shootingTime), nil
}

func getShootingTime(path string) (*time.Time, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(f)
	if err != nil {
		return nil, err
	}
	shootingTime, err := x.DateTime()
	if err != nil {
		return nil, err
	}
	return &shootingTime, nil
}

func getPositionsFromGpx(gpxPath string) ([]positionAtTime, error) {
	gpxFile, err := gpx.ParseFile(gpxPath)
	if err != nil {
		return nil, err
	}
	positions := make([]positionAtTime, 0)
	for _, track := range gpxFile.Tracks {
		for _, segment := range track.Segments {
			for _, point := range segment.Points {
				positions = append(positions, positionAtTime{
					position: Position{
						Lat: point.Latitude,
						Lon: point.Longitude,
					},
					time: point.Timestamp,
				})
			}
		}
	}
	return positions, nil
}

func getNearestPointInTime(positions []positionAtTime, time time.Time) *Position {
	p := positions[0].position
	shootingTimeDiff := time.Sub(positions[0].time)

	for i := 0; i < len(positions)-1; i++ {
		if shootingTimeDiff < positions[i+1].time.Sub(positions[i].time) {
			p = positions[i].position
			shootingTimeDiff = time.Sub(positions[i].time)
		}
	}

	return &p
}
