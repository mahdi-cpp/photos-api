package album

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/search"
)

type SearchOptions struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Title     *string   `json:"title,omitempty"`
	Subtitle  *string   `json:"subtitle,omitempty"`
	Type      *string   `json:"type,omitempty"`
	TextQuery *string   `json:"textQuery,omitempty"`

	// Date filters
	CreatedAfter  *time.Time `json:"createdAfter,omitempty"`
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
	ActiveAfter   *time.Time `json:"activeAfter,omitempty"`

	// Pagination
	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`

	// Sorting
	Sort      string `json:"sort,omitempty"`      // "title", "created", "members", "lastActivity"
	SortOrder string `json:"sortOrder,omitempty"` // "asc" or "desc"
}

const MaxLimit = 1000

var LessFunks = map[string]search.LessFunction[*Album]{
	"id":        func(a, b *Album) bool { return a.ID.String() < b.ID.String() },
	"createdAt": func(a, b *Album) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt": func(a, b *Album) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*Album] {

	fn, exists := LessFunks[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "desc" {
		return func(a, b *Album) bool { return !fn(a, b) }
	}
	return fn
}

func BuildAlbumSearch(with SearchOptions) search.Criteria[*Album] {

	return func(a *Album) bool {

		// ID filter
		if with.ID != uuid.Nil && a.ID != with.ID {
			return false
		}

		// Title search_manager (case-insensitive)
		if with.TextQuery != nil {
			query := strings.ToLower(*with.Title)
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

func Search(chats []*Album, with SearchOptions) []*Album {

	// Build criteria
	criteria := BuildAlbumSearch(with)

	// Execute search_manager
	results := search.Find(chats, criteria)

	// Sort results if needed
	if with.Sort != "" {
		lessFn := GetLessFunc(with.Sort, with.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final assets
	final := make([]*Album, len(results))
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
		return []*Album{}
	}

	end := start + with.Size
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}
