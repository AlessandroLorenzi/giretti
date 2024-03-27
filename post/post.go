package post

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/AlessandroLorenzi/giretti/position"
	"github.com/adrg/frontmatter"
	"github.com/disintegration/imaging"
	"github.com/gomarkdown/markdown"
	"github.com/lib/pq"
)

type PostHeaders struct {
	Title     string   `yaml:"title"`
	Tags      []string `yaml:"tags"`
	Gpx       []string `yaml:"gpx"`
	OpenGraph struct {
		Image       *string
		Description *string
	} `yaml:"open_graph"`
}

type Post struct {
	Id                   string `gorm:"primaryKey"`
	Title                string
	Date                 time.Time
	Tags                 pq.StringArray `gorm:"type:text[]"`
	Gpx                  pq.StringArray `gorm:"type:text[]"`
	OpenGraphImage       *string
	OpenGraphDescription *string
	HTML                 template.HTML
}

type GalleryItem struct {
	Image     string
	Thumbnail string
	Position  *position.Position
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
		Id:                   id,
		Title:                headers.Title,
		Date:                 postDate(id),
		Tags:                 headers.Tags,
		Gpx:                  headers.Gpx,
		OpenGraphImage:       headers.OpenGraph.Image,
		OpenGraphDescription: headers.OpenGraph.Description,
		HTML:                 template.HTML(string(html)),
	}

	return p, nil
}

func (p *Post) Url() string {
	// Split the input string by underscores
	components := strings.Split(p.Id, "-")

	// Extract year, month, and day
	year := components[0]
	month := components[1]
	day := components[2]

	htmlName := strings.Join(components[3:], "-")
	htmlName = strings.TrimSuffix(htmlName, ".md")

	// Construct the output path
	return fmt.Sprintf("/%s/%s/%s/%s.html", year, month, day, htmlName)
}

func postDate(id string) time.Time {
	components := strings.Split(id, "-")

	// Convert string components to integers
	year, err := strconv.Atoi(components[0])
	if err != nil {
		fmt.Println("Error converting year:", err)
		return time.Time{}
	}

	month, err := strconv.Atoi(components[1])
	if err != nil {
		fmt.Println("Error converting month:", err)
		return time.Time{}
	}

	day, err := strconv.Atoi(components[2])
	if err != nil {
		fmt.Println("Error converting day:", err)
		return time.Time{}
	}

	// Construct the time.Time value
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (p *Post) MediaDir() string {
	return "media/posts/" + p.Id + "/"
}

func (p *Post) Gallery() []*GalleryItem {
	gallery := []*GalleryItem{}
	galleryPath := fmt.Sprintf("media/post/%s/gallery", p.Id)
	files, err := os.ReadDir(galleryPath)
	if err != nil {
		fmt.Printf("Error reading media dir: %v\n", err)
		return gallery
	}
	for _, file := range files {

		if strings.Contains(file.Name(), "_thumb.JPG") {
			continue
		}

		if filepath.Ext(file.Name()) == ".JPG" {
			fullFileName := fmt.Sprintf("%s/%s", galleryPath, file.Name())
			thumbnail := strings.TrimSuffix(fullFileName, ".JPG") + "_thumb.JPG"

			genThumbIfNeeded(fullFileName, thumbnail)

			pos := &position.Position{}

			if len(p.Gpx) > 0 {
				pos, err = position.ImagePosition(fullFileName, p.Gpx[0])
				if err != nil {
					fmt.Printf("Error getting image position: %s %v", fullFileName, err)
					continue
				}
			}

			gallery = append(gallery, &GalleryItem{
				Image:     fullFileName,
				Thumbnail: thumbnail,
				Position:  pos,
			})
		}
	}
	return gallery
}

func genThumbIfNeeded(fileName, thumbnail string) {
	_, err := os.Stat(thumbnail)
	if !os.IsNotExist(err) {
		return
	}
	src, _ := imaging.Open(fileName, imaging.AutoOrientation(true))
	dst := imaging.Fit(src, 800, 600, imaging.Lanczos)
	imaging.Save(dst, thumbnail)
}

func (p *Post) StartingPosition() *position.Position {
	if len(p.Gpx) == 0 {
		return nil
	}
	pat, err := position.GetPositionsFromGpx(p.Gpx[0])
	if err != nil {
		fmt.Printf("Error getting positions from gpx: %v", err)
		return nil
	}
	return &pat[0].Position
}
