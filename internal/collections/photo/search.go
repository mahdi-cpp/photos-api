package photo

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/search"
)

type SearchOptions struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	TextQuery      *string `json:"textQuery,omitempty"`
	FileSize       *string `json:"fileSize,omitempty"`
	MimeType       *string `json:"mimeType,omitempty"`
	CameraMake     *string `json:"cameraMake,omitempty"`
	CameraModel    *string `json:"cameraModel,omitempty"`
	IsCamera       *bool   `json:"isCamera,omitempty"`
	IsFavorite     *bool   `json:"isFavorite,omitempty"`
	IsScreenshot   *bool   `json:"isScreenshot,omitempty"`
	IsHidden       *bool   `json:"isHidden,omitempty"`
	IsLandscape    *bool   `json:"isLandscape,omitempty"`
	NotInOnePhoto  *bool   `json:"notInOnePhoto,omitempty"`
	HideScreenshot *bool   `json:"hideScreenshot"`

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
	"id":        func(a, b *Index) bool { return a.ID.String() < b.ID.String() },
	"createdAt": func(a, b *Index) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt": func(a, b *Index) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
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

		// Boolean flags
		if with.IsCamera != nil && c.IsCamera != *with.IsCamera {
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

		// Collection_old membership filters
		//if len(with.AlbumsManager) > 0 {
		//	found := false
		//	for _, memberID := range with.AlbumsManager {
		//		if search.StringInSlice(memberID, c.AlbumsManager) {
		//			found = true
		//			break
		//		}
		//	}
		//	if !found {
		//		return false
		//	}
		//}
		//
		//if len(with.Trips) > 0 {
		//	found := false
		//	for _, adminID := range with.Trips {
		//		if search.StringInSlice(adminID, c.Trips) {
		//			found = true
		//			break
		//		}
		//	}
		//	if !found {
		//		return false
		//	}
		//}
		//
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
