package camera

import (
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
	"github.com/mahdi-cpp/photos-api/internal/help"
)

type Manager struct {
	photoManager *photo.Manager
	collection   *collection_manager_memory.Manager[*Camera]
	cameras      []*photo.Collection[*Camera]
}

func NewManager(photoManager *photo.Manager, path string) (*Manager, error) {

	manager := &Manager{
		photoManager: photoManager,
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

	m.cameras = []*photo.Collection[*Camera]{}

	for _, item := range all {

		photoOptions := &photo.SearchOptions{
			CameraMake: help.StrPtr(item.CameraMake),
			Sort:       "id",
			SortOrder:  "desc",
			Page:       1,
			Size:       6,
		}

		photos, err := m.photoManager.ReadByIndexes(indexes, photoOptions)
		if err != nil {
			return err
		}

		a := &photo.Collection[*Camera]{
			Item:   item,
			Photos: photos,
		}

		m.cameras = append(m.cameras, a)
	}

	return nil
}

func (m *Manager) ReadCollections() []*photo.Collection[*Camera] {
	return m.cameras
}

//--- events

func (m *Manager) OnCreated(photo *photo.Photo) {

	if photo.CameraMake != "" {

		isExist := false
		for _, camera := range m.cameras {
			if camera.Item.CameraMake == photo.CameraMake {
				isExist = true
				break
			}
		}

		if isExist == false { // if is new. create and reload cameras
			camera := &Camera{
				Title:      photo.CameraMake + " " + photo.CameraModel,
				CameraMake: photo.CameraMake,
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
}
