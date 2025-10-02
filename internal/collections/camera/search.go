package camera

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/search"
)

type SearchOptions struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Subtitle string    `json:"subtitle,omitempty"`
	Type     string    `json:"type,omitempty"`

	TextQuery string `json:"textQuery,omitempty"`

	// Date filters
	CreatedAfter  *time.Time `json:"createdAfter,omitempty"`
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
	ActiveAfter   *time.Time `json:"activeAfter,omitempty"`

	// Pagination
	Page int `form:"page,omitempty"`
	Size int `form:"size,omitempty"`

	// Sorting
	Sort      string `form:"sort,omitempty"`      // "title", "created", "members", "lastActivity"
	SortOrder string `form:"sortOrder,omitempty"` // "asc" or "desc"
}

const MaxLimit = 1000

var LessFunks = map[string]search.LessFunction[*Camera]{
	"id":        func(a, b *Camera) bool { return a.ID.String() < b.ID.String() },
	"createdAt": func(a, b *Camera) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt": func(a, b *Camera) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*Camera] {

	fn, exists := LessFunks[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "desc" {
		return func(a, b *Camera) bool { return !fn(a, b) }
	}
	return fn
}

func BuildCameraSearch(with *SearchOptions) search.Criteria[*Camera] {

	return func(a *Camera) bool {

		// ID filter
		if with.ID != uuid.Nil && a.ID != with.ID {
			return false
		}

		// Title search_manager (case-insensitive)
		if with.TextQuery != "" {
			query := strings.ToLower(with.Title)
			title := strings.ToLower(a.Subtitle)
			if !strings.Contains(title, query) {
				return false
			}
		}

		// Date filters
		if with.CreatedAfter != nil && a.CreatedAt.Before(*with.CreatedAfter) {
			return false
		}
		if with.CreatedBefore != nil && a.CreatedAt.After(*with.CreatedBefore) {
			return false
		}

		return true
	}
}

func Search(chats []*Camera, with *SearchOptions) []*Camera {

	// Build criteria
	criteria := BuildCameraSearch(with)

	// Execute search_manager
	results := search.Find(chats, criteria)

	// Sort results if needed
	if with.Sort != "" {
		lessFn := GetLessFunc(with.Sort, with.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final photos
	final := make([]*Camera, len(results))
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
		return []*Camera{}
	}

	end := start + with.Size
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}
