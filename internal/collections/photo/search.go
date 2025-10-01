package photo

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/search"
)

type SearchOptions struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	TextQuery   *string `json:"textQuery,omitempty"`
	FileSize    *string `json:"fileSize,omitempty"`
	MimeType    *string `json:"mimeType,omitempty"`
	CameraMake  *string `json:"cameraMake,omitempty"`
	CameraModel *string `json:"cameraModel,omitempty"`

	IsVideo        *bool `json:"isVideo,omitempty"`
	IsFavorite     *bool `json:"isFavorite,omitempty"`
	IsScreenshot   *bool `json:"isScreenshot,omitempty"`
	IsHidden       *bool `json:"isHidden,omitempty"`
	NotInOneAlbum  *bool `json:"notInOneAlbum,omitempty"`
	HideScreenshot *bool `json:"hideScreenshot"`

	// Date filters
	CreatedAfter  *time.Time `json:"createdAfter,omitempty"`
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
	ActiveAfter   *time.Time `json:"activeAfter,omitempty"`

	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`

	Sort      string `json:"sort,omitempty"`      // "title", "created", "members", "lastActivity"
	SortOrder string `json:"sortOrder,omitempty"` // "asc" or "desc"
}

const MaxLimit = 1000

var LessFunks = map[string]search.LessFunction[*Index]{
	"id":               func(a, b *Index) bool { return a.ID.String() < b.ID.String() },
	"createdAt":        func(a, b *Index) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt":        func(a, b *Index) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
	"dateTimeOriginal": func(a, b *Index) bool { return a.DateTimeOriginal.Before(b.DateTimeOriginal) },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*Index] {

	fn, exists := LessFunks[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "desc" {
		return func(a, b *Index) bool { return !fn(a, b) }
	}
	return fn
}

func BuildPhotoCriteria(with *SearchOptions) search.Criteria[*Index] {

	return func(c *Index) bool {

		// ID filter
		if with.ID != uuid.Nil && c.ID != with.ID {
			return false
		}

		if with.CameraMake != nil && c.CameraMake != *with.CameraMake {
			return false
		}
		if with.CameraModel != nil && c.CameraModel != *with.CameraModel {
			return false
		}

		// Boolean flags
		if with.IsVideo != nil && c.IsVideo != *with.IsVideo {
			return false
		}
		if with.IsFavorite != nil && c.IsFavorite != *with.IsFavorite {
			return false
		}
		if with.IsScreenshot != nil && c.IsScreenshot != *with.IsScreenshot {
			return false
		}
		if with.IsHidden != nil && c.IsHidden != *with.IsHidden {
			return false
		}
		if with.NotInOneAlbum != nil && c.NotInOneAlbum != *with.NotInOneAlbum {
			return false
		}

		// Collection_old membership filters
		//if len(with.Persons) > 0 {
		//	found := false
		//	for _, blockID := range with.Persons {
		//		if search.StringInSlice(blockID, c.Persons) {
		//			found = true
		//			break
		//		}
		//	}
		//	if !found {
		//		return false
		//	}
		//}

		// Date filters
		if with.CreatedAfter != nil && c.CreatedAt.Before(*with.CreatedAfter) {
			return false
		}
		if with.CreatedBefore != nil && c.CreatedAt.After(*with.CreatedBefore) {
			return false
		}

		return true
	}
}

func Search(index []*Index, with *SearchOptions) []*Index {

	// Build criteria
	criteria := BuildPhotoCriteria(with)

	// Execute search_manager
	results := search.Find(index, criteria)

	// Sort results if needed
	if with.Sort != "" {
		lessFn := GetLessFunc(with.Sort, with.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final collection
	final := make([]*Index, len(results))
	for i, item := range results {
		final[i] = item.Value
	}

	if with.Size == 0 { // if not set default is MAX_LIMIT
		with.Size = MaxLimit
	}

	// Apply pagination
	start := (with.Page - 1) * with.Size // Corrected pagination logic
	if start < 0 {
		start = 0
	}

	// Check if the start index is out of bounds. If so, return an empty slice.
	if start >= len(final) {
		return []*Index{}
	}
	end := start + with.Size
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}
