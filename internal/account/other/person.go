package other

import (
	"fmt"

	"github.com/mahdi-cpp/photos-api/internal/account"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
)

func (m *account.Manager) preparePersons() {

	items, err := m.Persons.CollectionMemory.ReadAll()
	if err != nil {
	}

	for _, item := range items {
		with := &photo.SearchOptions{
			UserID:    m.userID,
			Persons:   []string{item.ID.String()},
			Sort:      "createdAt",
			SortOrder: "start",
			Size:      1,
		}
		filterAssets, err := m.ReadAll(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		item.Count = len(filterAssets)
		m.Persons.CoverPhotoArray[item.ID] = filterAssets
	}
}
