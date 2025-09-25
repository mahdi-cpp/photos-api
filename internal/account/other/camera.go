package other

//func (m *Manager) prepareCameras() {
//
//	//items, err := m.CameraManager.ReadAll()
//	//if err != nil {
//	//}
//
//	if m.cameras == nil {
//		m.cameras = map[string]*ali.PHCollection[camera.Camera]{}
//	}
//
//	indexes := m.PhotosManager.GetAllIndexes()
//
//	for _, a := range indexes {
//		if a.Camera.Model == "" {
//			continue
//		}
//
//		cameraManager, exists := m.cameras[a.Camera.Model]
//		if exists {
//			cameraManager.Item.Count = cameraManager.Item.Count + 1
//			m.cameras[a.Camera.Model] = cameraManager
//		} else {
//			collection := &ali.PHCollection[camera.Camera]{
//				Item: camera.Camera{
//					ID:          "1",
//					CameraMake:  a.Camera.Make,
//					CameraModel: a.Camera.Model,
//					Count:       1},
//			}
//			//fmt.Println(collection)
//			m.cameras[a.Camera.Model] = collection
//		}
//	}
//
//	for _, collection := range m.cameras {
//
//		with := &photo.SearchOptions{
//			UserID:      m.userID,
//			CameraMake:  collection.Item.CameraMake,
//			CameraModel: collection.Item.CameraModel,
//			SortBy:      "createdAt",
//			SortOrder:   "start",
//			Size:        6,
//		}
//
//		filterAssets, err := m.ReadAll(with)
//		if err != nil {
//			fmt.Printf("Error getting all person_test: %v\n", err)
//			return
//		}
//		collection.PhotosManager = filterAssets
//	}
//}
//
//func (m *Manager) GetAllCameras() map[string]*ali.PHCollection[camera.Camera] {
//	return m.cameras
//}
