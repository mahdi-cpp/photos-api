package album

import (
	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_join"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
)

type Manager struct {
	photoManager *photo.Manager
	album        *collection_manager_memory.Manager[*Album]
	Join         *collection_manager_join.Manager[*photo.Join]
}

func NewManager(photoManager *photo.Manager, path string) (*Manager, error) {

	manager := &Manager{
		photoManager: photoManager,
	}

	var err error
	manager.album, err = collection_manager_memory.New[*Album](path, "album")
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) Create(album *Album) (*Album, error) {
	item, err := m.album.Create(album)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (m *Manager) Read(id uuid.UUID) (*Album, error) {
	item, err := m.album.Read(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadAll() ([]*Album, error) {
	items, err := m.album.ReadAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (m *Manager) Update(with UpdateOptions) (*Album, error) {

	item, err := m.album.Read(with.ID)
	if err != nil {
		return nil, err
	}

	Update(item, with)

	create, err := m.album.Update(item)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (m *Manager) Delete(id uuid.UUID) error {
	err := m.album.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

//--- photo

func (m *Manager) AddItem(albumID, photoID uuid.UUID) error {

	j := &photo.Join{
		ParentID: albumID,
		PhotoID:  photoID,
	}

	_, err := m.Join.Create(j)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) ReadCollection(albumID uuid.UUID, with *photo.SearchOptions) (*photo.PHCollection[*Album], error) {

	item, err := m.album.Read(albumID)
	if err != nil {
		return nil, err
	}

	all, err := m.Join.GetByParentID(albumID)
	if err != nil {
		return nil, err
	}

	photos, err := m.photoManager.ReadAlbumPhotos(all, with)
	if err != nil {
		return nil, err
	}

	a := &photo.PHCollection[*Album]{
		Item:   item,
		Photos: photos,
	}

	return a, nil
}

func (m *Manager) ReadCollections(id uuid.UUID) (*Album, error) {
	item, err := m.album.Read(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

//func (m *Manager) PrepareAlbums() {
//
//	items, err := m.album.ReadAll()
//	if err != nil {
//	}
//
//	for _, item := range items {
//
//		with := &photo.SearchOptions{
//			UserID:    m.userID,
//			Albums:    []string{item.ID.String()},
//			Sort:      "createdAt",
//			SortOrder: "start",
//			Size:      6,
//		}
//
//		filterPhotos, err := m.ReadAll(with)
//		if err != nil {
//			fmt.Printf("Error getting all person_test: %v\n", err)
//			return
//		}
//		item.Count = len(filterPhotos)
//		m.album.CoverPhotoArray[item.ID] = filterPhotos
//	}
//}
