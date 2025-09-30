package config

import (
	"fmt"
	"path/filepath"
)

const (
	root        = "/app/iris/"
	application = "com.iris.photos"
	users       = "users"
	Metadata    = "metadata"
	Version     = "v3"
	RootDir     = "/app/iris/com.iris.photos/"
	TestUserID  = "018fe65d-8e4a-74b0-8001-c8a7c29367e1"
)

func GetPath(file string) string {
	return filepath.Join(root, application, file)
}

func GetUserPath(userID string) string {
	pp := filepath.Join(root, application, users, userID)
	fmt.Println(pp)
	return pp
}

func GetUserMetadataPath(id string, directory string) string {
	pp := filepath.Join(root, application, users, id, Metadata, directory)
	fmt.Println(pp)
	return pp
}
