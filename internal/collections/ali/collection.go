package ali

import (
	"time"
)

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
