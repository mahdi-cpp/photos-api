package album

import (
	"fmt"
	"testing"

	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
)

func TestNewManager(t *testing.T) {

	photoManager, err := photo.NewManager("app/tmp/photos/")
	if err != nil {
		t.Fatal(err)
	}

	m, err := NewManager(photoManager, "/app/tmp/photos/albums")
	if err != nil {
		t.Fatal(err)
	}

	a := &Album{
		Title:    "Test Album",
		IsHidden: true,
		Version:  "1.0.0",
	}

	create, err := m.Create(a)
	if err != nil {
		return
	}
	fmt.Println(create.ID)
}
