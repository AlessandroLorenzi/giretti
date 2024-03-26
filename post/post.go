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
)

type PostHeaders struct {
	Title     string   `yaml:"title"`
	Tags      []string `yaml:"tags"`
	Gpx       []string `yaml:"gpx"`
	OpenGraph struct {
		Image       *string `yaml:"image"`
		Description *string `yaml:"description"`
	} `yaml:"open_graph"`
}

type Post struct {
	Title     string
	Path      string
	Tags      []string
	Gpx       []string
	OpenGraph struct {
		Image       *string `yaml:"image"`
		Description *string `yaml:"description"`
	}
	HTML template.HTML
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

	p := &Post{
		Title: headers.Title,
		Path:  path,
		Tags:  headers.Tags,
		Gpx:   headers.Gpx,
		OpenGraph: struct {
			Image       *string `yaml:"image"`
			Description *string `yaml:"description"`
		}{
			Image:       headers.OpenGraph.Image,
			Description: headers.OpenGraph.Description,
		},
		HTML: template.HTML(string(html)),
	}

	return p, nil
}

func (p *Post) Id() string {
	return strings.TrimSuffix(
		filepath.Base(p.Path),
		filepath.Ext(p.Path),
	)
}

func (p *Post) Url() string {
	base := filepath.Base(p.Path)
	// Split the input string by underscores
	components := strings.Split(base, "-")

	// Extract year, month, and day
	year := components[0]
	month := components[1]
	day := components[2]

	htmlName := strings.Join(components[3:], "-")
	htmlName = strings.TrimSuffix(htmlName, ".md")

	// Construct the output path
	return fmt.Sprintf("/%s/%s/%s/%s.html", year, month, day, htmlName)
}

func (p *Post) Date() time.Time {
	base := filepath.Base(p.Path)
	// Split the input string by underscores
	components := strings.Split(base, "-")

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
	return "media/posts/" + p.Id() + "/"
}

func (p *Post) Gallery() []*GalleryItem {
	gallery := []*GalleryItem{}
	galleryPath := fmt.Sprintf("media/post/%s/gallery", p.Id())
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
