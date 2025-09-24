package account

import (
	"fmt"

	"github.com/mahdi-cpp/photos-api/internal/collections/asset"
)

func (m *Manager) preparePersons() {

	items, err := m.Persons.CollectionMemory.ReadAll()
	if err != nil {
	}

	for _, item := range items {
		with := &asset.SearchOptions{
			UserID:    m.userID,
			Persons:   []string{item.ID.String()},
			SortBy:    "createdAt",
			SortOrder: "start",
			Size:      1,
		}
		filterAssets, err := m.ReadAll(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		item.Count = len(filterAssets)
		m.Persons.CoverAssetArray[item.ID] = filterAssets
	}
}
