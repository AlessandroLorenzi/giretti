package post

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
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
	Gallery []struct {
		Image       string `yaml:"image"`
		Thumbnail   string `yaml:"thumbnail"`
		Description string `yaml:"description"`
	} `yaml:"gallery"`
}

type Post struct {
	ID      string
	Date    time.Time `yaml:"date"`
	Url     string
	Headers *PostHeaders
	HTML    template.HTML
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
		ID:      id,
		Headers: &headers,
		Url:     getUrlFromPath(path),
		Date:    getDateFromPath(path),
		HTML:    template.HTML(string(html)),
	}

	return p, nil
}

func getUrlFromPath(path string) string {
	base := filepath.Base(path)
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

func getDateFromPath(path string) time.Time {
	base := filepath.Base(path)
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
