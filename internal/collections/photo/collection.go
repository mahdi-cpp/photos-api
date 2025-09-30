package photo

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
)

type CollectionOld[T collection_manager_memory.CollectionItem] struct {
	CollectionMemory *collection_manager_memory.Manager[T]
	CoverPhotoArray  map[uuid.UUID][]*Photo
}

func NewCollection[T collection_manager_memory.CollectionItem](path string, fileName string) *CollectionOld[T] {

	c, err := collection_manager_memory.New[T](path, fileName)
	if err != nil {
		panic(err)
	}

	a := &CollectionOld[T]{
		CollectionMemory: c,
		CoverPhotoArray:  make(map[uuid.UUID][]*Photo),
	}

	return a
}

type PHCollectionList[T any] struct {
	Status      string           `json:"status"` // "success" or "error"
	Collections []*Collection[T] `json:"collections"`
}

type Collection[T any] struct {
	Item   T        `json:"item"`   // Generic items
	Photos []*Photo `json:"photos"` // Specific collection
}

type CollectionPhoto struct {
	ParentID uuid.UUID   `json:"ParentId"`
	PhotoIDs []uuid.UUID `json:"photoIds"`
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

type OnCreateCallback func(*Photo)
