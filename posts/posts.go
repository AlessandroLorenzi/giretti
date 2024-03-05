package posts

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlessandroLorenzi/giretti/post"
)

var Posts = []*post.Post{}

func Init(postsPath string) error {
	// list files in postsPath
	files, err := os.ReadDir(postsPath)
	if err != nil {
		return fmt.Errorf("error reading posts directory: %w", err)
	}
	// for each file, read the post and append it to Posts
	for _, f := range files {
		if f.IsDir() {
			continue // Skip directories
		}
		if filepath.Ext(f.Name()) == ".md" {
			p, err := post.ReadPost(filepath.Join(postsPath, f.Name()))
			if err != nil {
				return fmt.Errorf("error reading post %s: %w", f.Name(), err)
			}
			Posts = append(Posts, p)
		}
	}
	return nil
}

func GetFromUrl(url string) *post.Post {
	for _, p := range Posts {
		if p.Url == url {
			return p
		}
	}
	return nil
}
