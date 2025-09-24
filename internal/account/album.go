package account

import (
	"fmt"

	"github.com/mahdi-cpp/photos-api/internal/collections/asset"
)

func (m *Manager) prepareAlbums() {

	items, err := m.Album.CollectionMemory.ReadAll()
	if err != nil {
	}

	for _, item := range items {

		with := &asset.SearchOptions{
			UserID:    m.userID,
			Albums:    []string{item.ID.String()},
			SortBy:    "createdAt",
			SortOrder: "start",
			Size:      6,
		}

		filterAssets, err := m.ReadAll(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		item.Count = len(filterAssets)
		m.Album.CoverAssetArray[item.ID] = filterAssets
	}
}
