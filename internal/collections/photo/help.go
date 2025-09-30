package photo

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/config"
)

const uploadsDir = "/app/iris/services/uploads"

func moveMedia(userID uuid.UUID, workDir string, photo *Photo) error {

	uID := userID.String()

	var id = strings.TrimSuffix(photo.FileInfo.FileName, ".jpg")
	if photo.FileInfo.MimeType == "image/jpeg" {
		var name = id + ".jpg"
		src := filepath.Join(uploadsDir, workDir, name)
		des := filepath.Join(config.RootDir, "users", uID, "assets", name)
		err := os.Rename(src, des)
		if err != nil {
			return err
		}

		var thumbnail200 = id + "_270.jpg"
		srcThumb := filepath.Join(uploadsDir, workDir, thumbnail200)
		desThumb := filepath.Join(config.RootDir, "users", uID, "thumbnails", thumbnail200)
		err = os.Rename(srcThumb, desThumb)
		if err != nil {
			return err
		}
	} else {
		var name = id + ".mp4"
		src := filepath.Join(uploadsDir, workDir, name)
		des := filepath.Join(config.RootDir, "users", uID, "assets", name)
		err := os.Rename(src, des)
		if err != nil {
			return err
		}

		var thumbnail200 = id + "_270.jpg"
		srcThumb := filepath.Join(uploadsDir, workDir, thumbnail200)
		desThumb := filepath.Join(config.RootDir, "users", uID, "thumbnails", thumbnail200)
		err = os.Rename(srcThumb, desThumb)
		if err != nil {
			return err
		}

		var thumbnail400 = id + "-400.jpg"
		srcThumb = filepath.Join(uploadsDir, workDir, thumbnail400)
		desThumb = filepath.Join(config.RootDir, "users", uID, "thumbnails", thumbnail400)
		err = os.Rename(srcThumb, desThumb)
		if err != nil {
			return err
		}
	}
	return nil
}
