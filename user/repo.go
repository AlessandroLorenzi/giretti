package user

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var UsersRepo = []*User{}

func InitRepo() error {
	// read from users.yaml
	data, err := os.ReadFile("users.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &UsersRepo)
	if err != nil {
		return err
	}
	return nil
}

func SaveRepo() error {
	// write to users.yaml
	data, err := yaml.Marshal(UsersRepo)
	if err != nil {
		return err
	}

	err = os.WriteFile("users.yaml", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetAll() []*User {
	return UsersRepo
}

func GetByUsername(username string) *User {
	for _, u := range UsersRepo {
		if u.Username == username {
			return u
		}
	}
	return nil
}

func AddFollower(username, follower string) error {
	u := GetByUsername(username)
	if u == nil {
		return fmt.Errorf("User %s not found", username)
	}

	// check if follower already exists
	for _, f := range u.Followers {
		if f == follower {
			return nil
		}
	}

	u.Followers = append(u.Followers, follower)
	return SaveRepo()
}
