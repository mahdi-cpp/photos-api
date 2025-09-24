package account

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/collections/asset"
	"github.com/mahdi-cpp/photos-api/internal/help"
)

func (m *Manager) preparePinned() {

	items, err := m.Pinned.CollectionMemory.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, item := range items {

		var with *asset.SearchOptions

		switch item.Type {
		case "camera":
			with = &asset.SearchOptions{
				IsCamera:  help.BoolPtr(true),
				SortBy:    "createdAt",
				SortOrder: "start",
				Size:      1,
			}
			break
		case "screenshot":
			with = &asset.SearchOptions{
				IsScreenshot: help.BoolPtr(true),
				SortBy:       "createdAt",
				SortOrder:    "start",
				Size:         1,
			}
			break
		case "favorite":
			with = &asset.SearchOptions{
				IsFavorite: help.BoolPtr(true),
				SortBy:     "createdAt",
				SortOrder:  "start",
				Size:       1,
			}
			break
		case "video":
			with = &asset.SearchOptions{
				MimeType:  help.StrPtr("video/mp4"),
				SortBy:    "createdAt",
				SortOrder: "start",
				Size:      1,
			}
			break
		case "map":
			var assets []*asset.Asset
			asset1 := asset.Asset{
				ID: uuid.Nil,
				FileInfo: asset.FileInfo{
					BaseURL:  "map",
					FileType: "map",
				},
			}
			assets = append(assets, &asset1)
			m.Pinned.CoverAssetArray[item.ID] = assets
			break
		case "album":
			selectedAlbum, err := m.Album.CollectionMemory.Read(item.AlbumID)
			if err != nil {
				continue
			}
			item.Title = selectedAlbum.Title
			with = &asset.SearchOptions{
				Albums:    []string{selectedAlbum.ID.String()},
				SortBy:    "createdAt",
				SortOrder: "start",
				Size:      1,
			}
			break
		}

		if with == nil || item.Type == "map" {
			continue
		}

		filterAssets, err := m.ReadAll(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		item.Count = len(filterAssets)
		m.Pinned.CoverAssetArray[item.ID] = filterAssets
	}
}
