package post

import (
	"fmt"
	"os"
	"path/filepath"
)

var PostsRepo = []*Post{}

func InitRepo(postsPath string) error {
	// list files in postsPath
	if err := loadPosts(postsPath); err != nil {
		return err
	}
	sortPostsBydate()

	return nil
}

func GetAll() []*Post {
	return PostsRepo
}

func GetFromUrl(url string) *Post {
	for _, p := range PostsRepo {
		if p.Url() == url {
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
			p, err := ReadPost(filepath.Join(postsPath, f.Name()))
			if err != nil {
				return fmt.Errorf("error reading post %s: %w", f.Name(), err)
			}
			PostsRepo = append(PostsRepo, p)
		}
	}
	return nil
}

func sortPostsBydate() {
	for i := 0; i < len(PostsRepo); i++ {
		for j := i + 1; j < len(PostsRepo); j++ {
			if PostsRepo[i].Date().Before(PostsRepo[j].Date()) {
				PostsRepo[i], PostsRepo[j] = PostsRepo[j], PostsRepo[i]
			}
		}
	}
}
