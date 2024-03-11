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
	Position Position
	time     time.Time
}

func GetImagePosition(path string, gpxPath string) (*Position, error) {
	shootingTime, err := getShootingTime(path)
	if err != nil {
		return nil, err
	}

	positions, err := GetPositionsFromGpx(gpxPath)
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

func GetPositionsFromGpx(gpxPath string) ([]positionAtTime, error) {
	gpxFile, err := gpx.ParseFile(gpxPath)
	if err != nil {
		return nil, err
	}
	positions := make([]positionAtTime, 0)
	for _, track := range gpxFile.Tracks {
		for _, segment := range track.Segments {
			for _, point := range segment.Points {
				positions = append(positions, positionAtTime{
					Position: Position{
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

func getNearestPointInTime(positions []positionAtTime, shootingTime time.Time) *Position {
	p := positions[0].Position
	shootingTimeDiff := shootingTime.Sub(positions[0].time).Abs()
	for i := 0; i < len(positions)-1; i++ {
		if shootingTimeDiff > shootingTime.Sub(positions[i].time).Abs() {
			p = positions[i].Position
			shootingTimeDiff = shootingTime.Sub(positions[i].time).Abs()
		}
	}
	return &p
}
