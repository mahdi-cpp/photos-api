package pinned

func preparePinned() {

	//
	//for _, item := range items {
	//
	//	var with *photo.SearchOptions
	//
	//	switch item.Type {
	//	case "camera":
	//		with = &photo.SearchOptions{
	//			IsCamera:  help.BoolPtr(true),
	//			Sort:      "createdAt",
	//			SortOrder: "start",
	//			Size:      1,
	//		}
	//		break
	//	case "screenshot":
	//		with = &photo.SearchOptions{
	//			IsScreenshot: help.BoolPtr(true),
	//			Sort:         "createdAt",
	//			SortOrder:    "start",
	//			Size:         1,
	//		}
	//		break
	//	case "favorite":
	//		with = &photo.SearchOptions{
	//			IsFavorite: help.BoolPtr(true),
	//			Sort:       "createdAt",
	//			SortOrder:  "start",
	//			Size:       1,
	//		}
	//		break
	//	case "video":
	//		with = &photo.SearchOptions{
	//			MimeType:  help.StrPtr("video/mp4"),
	//			Sort:      "createdAt",
	//			SortOrder: "start",
	//			Size:      1,
	//		}
	//		break
	//	case "map":
	//		var photos []*photo.Photo
	//		a := photo.Photo{
	//			ID: uuid.Nil,
	//			FileInfo: photo.FileInfo{
	//				URL:      "map",
	//				MimeType: "map",
	//			},
	//		}
	//		photos = append(photos, &a)
	//		m.Pinned.CoverPhotoArray[item.ID] = photos
	//		break
	//	case "album":
	//		selectedAlbum, err := m.Album.CollectionMemory.Read(item.AlbumID)
	//		if err != nil {
	//			continue
	//		}
	//		item.Title = selectedAlbum.Title
	//		with = &photo.SearchOptions{
	//			Albums:    []string{selectedAlbum.ID.String()},
	//			Sort:      "createdAt",
	//			SortOrder: "start",
	//			Size:      1,
	//		}
	//		break
	//	}
	//
	//	if with == nil || item.Type == "map" {
	//		continue
	//	}
	//
	//	filterAssets, err := m.ReadAll(with)
	//	if err != nil {
	//		fmt.Printf("Error getting all person_test: %v\n", err)
	//		return
	//	}
	//	item.Count = len(filterAssets)
	//	m.Pinned.CoverPhotoArray[item.ID] = filterAssets
	//}
}
