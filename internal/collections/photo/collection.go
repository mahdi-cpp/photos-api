package photo

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
)

type Collection[T collection_manager_memory.CollectionItem] struct {
	CollectionMemory *collection_manager_memory.Manager[T]
	CoverPhotoArray  map[uuid.UUID][]*Photo
}

func NewCollection[T collection_manager_memory.CollectionItem](path string, fileName string) *Collection[T] {

	c, err := collection_manager_memory.New[T](path, fileName)
	if err != nil {
		panic(err)
	}

	a := &Collection[T]{
		CollectionMemory: c,
		CoverPhotoArray:  make(map[uuid.UUID][]*Photo),
	}

	return a
}

type PHCollectionList[T any] struct {
	Status      string             `json:"status"` // "success" or "error"
	Collections []*PHCollection[T] `json:"collections"`
}

type PHCollection[T any] struct {
	Item   T        `json:"item"`       // Generic items
	Photos []*Photo `json:"collection"` // Specific collection
}

// https://chat.deepseek.com/a/chat/s/9b010f32-b23d-4f9b-ae0c-31a9b2c9408c

type SortableCollectionItem interface {
	GetID() int
	GetCreationDate() time.Time
	GetModificationDate() time.Time
}

type CollectionRequest struct {
	ID           string   `json:"id"`
	PhotoIds     []string `json:"photoIds,omitempty"` // Photo Ids
	Title        string   `json:"title,omitempty"`
	Subtitle     string   `json:"subtitle,omitempty"`
	TripType     string   `json:"trip,omitempty"`
	IsCollection *bool    `json:"isCollection,omitempty"`
}

type CollectionResponse struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
}
