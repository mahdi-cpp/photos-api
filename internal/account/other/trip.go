package other

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/account"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
)

func (m *account.Manager) prepareTrips() {

	items, err := m.Trips.CollectionMemory.ReadAll()
	if err != nil {
	}

	for _, item := range items {
		with := &photo.SearchOptions{
			UserID:    uuid.Nil,
			Trips:     []string{item.ID.String()},
			Sort:      "createdAt",
			SortOrder: "start",
			Size:      2,
		}

		filterAssets, err := m.ReadAll(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		item.Count = len(filterAssets)
		m.Trips.CoverPhotoArray[item.ID] = filterAssets
	}
}
