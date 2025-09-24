package depricated

//
//import (
//	"github.com/mahdi-cpp/iris-tools/asset"
//	"sort"
//	"strings"
//)
//
//type IndexedItem[T any] struct {
//	Index int
//	Value T
//}
//
//type assetSearchCriteria[T any] func(T) bool
//
//func assetBuildCriteria(with asset.Options) assetSearchCriteria[asset.PHAsset] {
//
//	return func(asset asset.PHAsset) bool {
//
//		// Filter by UserID (if non-zero)
//		//if with.UserID != 0 && person_test.UserID != with.UserID {
//		//	return false
//		//}
//
//		// Filter by Query (case-insensitive service in Filename/URL)
//		if with.Query != "" {
//			query := strings.ToLower(with.Query)
//
//			filename := strings.ToLower(asset.Filename)
//			url := strings.ToLower(asset.Url)
//
//			if !strings.Contains(filename, query) && !strings.Contains(url, query) {
//				return false
//			}
//		}
//
//		//Filter by MediaType (if specified)
//		if with.MediaType != "" && asset.MediaType != with.MediaType {
//			return false
//		}
//
//		// Filter by CameraModel (exact match)
//		if with.CameraMake != "" && asset.CameraMake != with.CameraMake {
//			return false
//		}
//		if with.CameraModel != "" && asset.CameraModel != with.CameraModel {
//			return false
//		}
//
//		// Filter by CreationDate range
//		if with.StartDate != nil && asset.CreationDate.Before(*with.StartDate) {
//			return false
//		}
//		if with.EndDate != nil && asset.CreationDate.After(*with.EndDate) {
//			return false
//		}
//
//		// Filter by boolean flags (if specified)
//		if with.IsCamera != nil && *with.IsCamera != asset.IsCamera {
//			return false
//		}
//		if with.IsFavorite != nil && *with.IsFavorite != asset.IsFavorite {
//			return false
//		}
//		if with.IsScreenshot != nil && *with.IsScreenshot != asset.IsScreenshot {
//			return false
//		}
//		if with.IsHidden != nil && *with.IsHidden != asset.IsHidden {
//			return false
//		}
//		if with.NotInOneAlbum != nil {
//		}
//		if with.HideScreenshot != nil && *with.HideScreenshot == false && asset.IsScreenshot == true {
//			return false
//		}
//
//		// Filter by  int
//		if with.ImageWidth != 0 && asset.ImageWidth != with.ImageWidth {
//			return false
//		}
//		if with.PixelHeight != 0 && asset.PixelHeight != with.PixelHeight {
//			return false
//		}
//
//		// Filter by landscape orientation
//		if with.IsLandscape != nil {
//			isLandscape := asset.ImageWidth > asset.PixelHeight
//			if isLandscape != *with.IsLandscape {
//				return false
//			}
//		}
//
//		// Album filtering
//		if len(with.Albums) > 0 {
//			found := false
//			for _, albumID := range with.Albums {
//
//				for _, assetAlbumID := range asset.Albums {
//					if assetAlbumID == albumID {
//						found = true
//						break
//					}
//				}
//
//				if found {
//					break
//				}
//			}
//			if !found {
//				return false
//			}
//		}
//
//		// Trip filtering
//		if len(with.Trips) > 0 {
//			found := false
//			for _, tripID := range with.Trips {
//				for _, assetTripID := range asset.Trips {
//					if assetTripID == tripID {
//						found = true
//						break
//					}
//				}
//				if found {
//					break
//				}
//			}
//			if !found {
//				return false
//			}
//		}
//
//		// Persons filtering
//		if len(with.Persons) > 0 {
//			found := false
//			for _, personID := range with.Persons {
//				for _, assetPersonID := range asset.Persons {
//					if assetPersonID == personID {
//						found = true
//						break
//					}
//				}
//				if found {
//					break
//				}
//			}
//			if !found {
//				return false
//			}
//		}
//
//		// Location filtering
//		//if len(person_test.Location) == 2 {
//		//
//		//	// Near point + radius search
//		//	if len(with.NearPoint) == 2 && with.WithinRadius > 0 {
//		//		distance := indexer.haversineDistance(
//		//			with.NearPoint[0], with.NearPoint[1],
//		//			person_test.Location[0], person_test.Location[1],
//		//		)
//		//		if distance > with.WithinRadius {
//		//			return false
//		//		}
//		//	}
//		//
//		//	// Bounding box search
//		//	if len(with.BoundingBox) == 4 {
//		//		if !indexer.isInBoundingBox(person_test.Location, with.BoundingBox) {
//		//			return false
//		//		}
//		//	}
//		//}
//
//		return true // Asset matches all active with
//	}
//}
//
//func assetSortAssets(assets []*asset.PHAsset, sortBy, sortOrder string) {
//
//	if sortBy == "" {
//		return // No sorting requested
//	}
//
//	sort.Slice(assets, func(i, j int) bool {
//
//		a := assets[i]
//		b := assets[j]
//
//		switch sortBy {
//		case "id":
//			if sortOrder == "asc" {
//				return a.ID < b.ID
//			}
//			return a.ID > b.ID
//
//		case "capturedDate":
//			if sortOrder == "asc" {
//				return a.CapturedDate.Before(b.CapturedDate)
//			}
//			return a.CapturedDate.After(b.CapturedDate)
//
//		case "creationDate":
//			if sortOrder == "asc" {
//				return a.CreationDate.Before(b.CreationDate)
//			}
//			return a.CreationDate.After(b.CreationDate)
//
//		case "modificationDate":
//			if sortOrder == "asc" {
//				return a.ModificationDate.Before(b.ModificationDate)
//			}
//			return a.ModificationDate.After(b.ModificationDate)
//		case "filename":
//			if sortOrder == "asc" {
//				return a.Filename < b.Filename
//			}
//			return a.Filename > b.Filename
//
//		default:
//			return false // No sorting for unknown fields
//		}
//	})
//}

//func assetSearch[T any](slice []T, criteria assetSearchCriteria[T]) []IndexedItem[T] {
//	var results []IndexedItem[T]
//
//	for i, item := range slice {
//		if criteria(item) {
//			results = append(results, IndexedItem[T]{Index: i, Value: item})
//		}
//	}
//	return results
//}
