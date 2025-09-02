package asset

import (
	"time"

	"github.com/mahdi-cpp/photos-api/internal/collections/phasset"
)

// https://chat.deepseek.com/a/chat/s/9b010f32-b23d-4f9b-ae0c-31a9b2c9408c

type PHCollectionList[T any] struct {
	Status      string             `json:"status"` // "success" or "error"
	Collections []*PHCollection[T] `json:"collections"`
}

type PHCollection[T any] struct {
	Item   T                  `json:"item"`   // Generic items
	Assets []*phasset.PHAsset `json:"assets"` // Specific assets
}

type SortableCollectionItem interface {
	GetID() int
	GetCreationDate() time.Time
	GetModificationDate() time.Time
}

type CollectionRequest struct {
	ID           string   `json:"id"`
	Title        string   `json:"title,omitempty"`
	Subtitle     string   `json:"subtitle,omitempty"`
	TripType     string   `json:"trip,omitempty"`
	IsCollection *bool    `json:"isCollection,omitempty"`
	AssetIds     []string `json:"assetIds,omitempty"` // Asset Ids
}

type CollectionResponse struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
}
