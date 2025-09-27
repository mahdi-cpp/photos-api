package album

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
)

const userID = "018fe65d-8e4a-74b0-8001-c8a7c29367e1"
const albumID = "0199860a-7364-7de0-8ce5-f1a666df77a5"
const workDir = "/app/iris/com.iris.photos/users/018fe65d-8e4a-74b0-8001-c8a7c29367e1/metadata"

func TestCreatePhotos(t *testing.T) {
	photoManager, err := photo.NewManager(workDir)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		p := &photo.Photo{
			FileInfo: photo.FileInfo{
				MimeType:    "image/jpg",
				FileSize:    12000,
				OriginalURL: fmt.Sprintf("http://example.com/%d", i),
			},
			ImageInfo: photo.ImageInfo{
				Width:       1000,
				Height:      5000,
				Orientation: "6",
			},
			IsCamera:   true,
			IsFavorite: true,
		}

		_, err := photoManager.Create(p)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCreateAlbum(t *testing.T) {
	photoManager, err := photo.NewManager(workDir)
	if err != nil {
		t.Fatal(err)
	}

	albumManager, err := NewManager(photoManager, workDir)
	if err != nil {
		t.Fatal(err)
	}

	a := &Album{
		Title:    "Album 3",
		IsHidden: true,
		Version:  "1.0.0",
	}

	create, err := albumManager.Create(a)
	if err != nil {
		return
	}
	fmt.Println(create.ID)
}

func TestAddPhotoToAlbum(t *testing.T) {

	photoManager, err := photo.NewManager(workDir)
	if err != nil {
		t.Fatal(err)
	}

	albumManager, err := NewManager(photoManager, workDir)
	if err != nil {
		t.Fatal(err)
	}

	albumID, err := uuid.Parse(albumID)
	if err != nil {
		return
	}
	album, err := albumManager.Read(albumID)
	if err != nil {
		return
	}

	in := photoManager.ReadIndexes()
	for _, i := range in {
		fmt.Println(i.ID)
		err := albumManager.AddPhoto(album.ID, i.ID)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestAlbumPhotos(t *testing.T) {

	photoManager, err := photo.NewManager(workDir)
	if err != nil {
		t.Fatal(err)
	}

	albumManager, err := NewManager(photoManager, workDir)
	if err != nil {
		t.Fatal(err)
	}

	//all, err := albumManager.ReadAll()
	//if err != nil {
	//	return
	//}
	//for _, i := range all {
	//	fmt.Println(i.ID)
	//}

	a, err := uuid.Parse(albumID)
	if err != nil {
		t.Fatal(err)
	}

	with := &photo.SearchOptions{
		Sort:      "id",
		SortOrder: "desc",
	}

	all, err := albumManager.ReadCollections(a, with)
	if err != nil {
		t.Fatal(err)
	}

	for _, p := range all.Photos {
		fmt.Println(p.FileInfo.OriginalURL)
	}
}

func TestAlbumPhotosLimit(t *testing.T) {

	photoManager, err := photo.NewManager(workDir)
	if err != nil {
		t.Fatal(err)
	}

	albumManager, err := NewManager(photoManager, workDir)
	if err != nil {
		t.Fatal(err)
	}

	start := time.Now()

	a, err := uuid.Parse(albumID)
	if err != nil {
		t.Fatal(err)
	}

	with := &photo.SearchOptions{
		Sort:      "id",
		SortOrder: "desc",
		Page:      1,
		Size:      1,
	}

	all, err := albumManager.ReadCollections(a, with)
	if err != nil {
		t.Fatal(err)
	}

	for _, p := range all.Photos {
		fmt.Println(p.FileInfo.OriginalURL)
	}

	duration := time.Since(start)
	fmt.Println(duration)
}
