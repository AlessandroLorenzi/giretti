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
	if err := loadPosts(postsPath); err != nil {
		return err
	}
	sortPostsBydate()

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

func loadPosts(postsPath string) error {
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

func sortPostsBydate() {
	for i := 0; i < len(Posts); i++ {
		for j := i + 1; j < len(Posts); j++ {
			if Posts[i].Date.Before(Posts[j].Date) {
				Posts[i], Posts[j] = Posts[j], Posts[i]
			}
		}
	}
}
