package album

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_join"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
	"github.com/mahdi-cpp/photos-api/internal/config"
)

func TestAlbum_Create(t *testing.T) {

	path := config.GetUserMetadataPath("01997cba-6dab-7636-a1f8-2c03174c7b6e", "")
	alumCollection, _ := collection_manager_memory.New[*Album](path+"/albums", "albums")
	photoAlbumsCollection, _ := collection_manager_join.New[*photo.Join](path+"/albums", "albums_join")

	a := &Album{
		Title:    "collection 1",
		Type:     "collection",
		IsHidden: false,
	}
	_, err := alumCollection.Create(a)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		photoID, err := uuid.NewV7()
		if err != nil {
			t.Fatal(err)
		}
		photoAlbum := &photo.Join{
			ParentID: a.ID,
			PhotoID:  photoID,
		}
		_, err = photoAlbumsCollection.Create(photoAlbum)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestAlbum_AddPhoto(t *testing.T) {

	path := config.GetUserMetadataPath("01997cba-6dab-7636-a1f8-2c03174c7b6e", "")

	photoAlbumsCollection, _ := collection_manager_join.New[*photo.Join](path+"/albums", "albums_join")

	albumID, err := uuid.Parse("01997cba-6dab-7636-a1f8-2c03174c7b6e")
	if err != nil {
		t.Fatal(err)
	}

	p := &photo.Join{
		ParentID: albumID,
		PhotoID:  uuid.New(),
	}

	_, err = photoAlbumsCollection.Create(p)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlbum_ReadAlbumPhotos(t *testing.T) {
	path := config.GetUserMetadataPath("01997cba-6dab-7636-a1f8-2c03174c7b6e", "")
	photoAlbumsCollection, _ := collection_manager_join.New[*photo.Join](path+"/albums", "albums_join")

	albumID, err := uuid.Parse("01997cba-6dab-7636-a1f8-2c03174c7b6e")
	if err != nil {
		t.Fatal(err)
	}

	ids, err := photoAlbumsCollection.GetByParentID(albumID)
	if err != nil {
		return
	}
	for _, id := range ids {
		fmt.Println("-------:", id.PhotoID)
	}
}

func TestAlbum_RemovePhoto(t *testing.T) {

}
