package camera

import (
	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
	"github.com/mahdi-cpp/photos-api/internal/help"
)

type Manager struct {
	photoManager *photo.Manager
	collection   *collection_manager_memory.Manager[*Camera]
	cameras      []*Camera
	photos       map[uuid.UUID][]*photo.Photo //key is albumId
}

func NewManager(photoManager *photo.Manager, path string) (*Manager, error) {

	manager := &Manager{
		photoManager: photoManager,
		photos:       make(map[uuid.UUID][]*photo.Photo),
	}

	var err error
	manager.collection, err = collection_manager_memory.New[*Camera](path, "camera")
	if err != nil {
		return nil, err
	}

	err = manager.load()
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) load() error {

	all, err := m.collection.ReadAll()
	if err != nil {
		return err
	}

	indexes := m.photoManager.ReadIndexes()

	m.cameras = []*Camera{}

	for _, item := range all {

		with := &photo.SearchOptions{
			CameraMake:  help.StrPtr(item.CameraMake),
			CameraModel: help.StrPtr(item.CameraModel),
			Sort:        "id",
			SortOrder:   "desc",
			Page:        1,
			Size:        6,
		}

		//item.Count = len(all)

		photos, err := m.photoManager.ReadByIndexes(indexes, with)
		if err != nil {
			return err
		}

		//a := &photo.Collection[*Camera]{
		//	Item:   item,
		//	Photos: photos,
		//}

		m.cameras = append(m.cameras, item)
		m.photos[item.ID] = photos
	}

	return nil
}

func (m *Manager) ReadCollections(with *SearchOptions) []*photo.Collection[*Camera] {

	var results []*photo.Collection[*Camera]
	filterCameras := Search(m.cameras, with)

	for _, camera := range filterCameras {
		collection := &photo.Collection[*Camera]{
			Item:   camera,
			Photos: m.photos[camera.ID],
		}
		results = append(results, collection)
	}

	return results
}

func (m *Manager) ReadCameraPhotos(cameraID uuid.UUID, with *photo.SearchOptions) ([]*photo.Photo, error) {

	camera, err := m.collection.Read(cameraID)
	if err != nil {
		return nil, err
	}

	indexes := m.photoManager.ReadIndexes()

	with.CameraMake = help.StrPtr(camera.CameraMake)
	with.CameraModel = help.StrPtr(camera.CameraModel)

	photos, err := m.photoManager.ReadByIndexes(indexes, with)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (m *Manager) HandlePhotoCreate(id uuid.UUID) {

	p, err := m.photoManager.Read(id)
	if err != nil {
		return
	}

	if p.CameraMake != "" {

		isExist := false
		for _, camera := range m.cameras {
			if camera.CameraMake == p.CameraMake && camera.CameraModel == p.CameraModel {
				isExist = true
				break
			}
		}

		if isExist == false { // if is new. create and reload cameras
			camera := &Camera{
				Title:       p.CameraMake + " " + p.CameraModel,
				CameraMake:  p.CameraMake,
				CameraModel: p.CameraModel,
			}
			_, err := m.collection.Create(camera)
			if err != nil {
				return
			}
			err = m.load()
			if err != nil {
				return
			}
		}
	}
	return
}

func (m *Manager) HandlePhotoUpdate(id uuid.UUID) {

}

func (m *Manager) HandlePhotoDelete(id uuid.UUID) {

}
