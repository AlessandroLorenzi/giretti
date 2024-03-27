package post

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitRepo(postsPath string) error {
	var err error
	db, err = gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&Post{})

	return loadPostsFromDb(postsPath)
}

func GetAll() []*Post {
	var posts []*Post
	db.Order("date desc").Find(&posts)
	return posts
}

func Get(id string) *Post {
	var post Post
	db.Where("id = ?", id).First(&post)
	return &post
}

func loadPostsFromDb(postsPath string) error {
	files, err := os.ReadDir(postsPath)
	if err != nil {
		return fmt.Errorf("error reading posts directory: %w", err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue // Skip directories
		}
		if filepath.Ext(f.Name()) == ".md" {
			p, err := ReadPost(filepath.Join(postsPath, f.Name()))
			if err != nil {
				return fmt.Errorf("error reading post %s: %w", f.Name(), err)
			}

			if db.Where("id = ?", p.Id).First(&Post{}).RowsAffected == 0 {
				db.Create(p)
			}
		}
	}
	return nil
}
