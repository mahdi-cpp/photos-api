package album

import (
	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_join"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
)

type Manager struct {
	photoManager *photo.Manager
	collection   *collection_manager_memory.Manager[*Album]
	join         *collection_manager_join.Manager[*photo.Join]
}

func NewManager(photoManager *photo.Manager, path string) (*Manager, error) {

	manager := &Manager{
		photoManager: photoManager,
	}

	var err error
	manager.collection, err = collection_manager_memory.New[*Album](path, "collection")
	if err != nil {
		return nil, err
	}

	manager.join, err = collection_manager_join.New[*photo.Join](path, "join")
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) Create(album *Album) (*Album, error) {
	item, err := m.collection.Create(album)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (m *Manager) Read(id uuid.UUID) (*Album, error) {
	item, err := m.collection.Read(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadAll() ([]*Album, error) {
	items, err := m.collection.ReadAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (m *Manager) Update(with UpdateOptions) (*Album, error) {

	item, err := m.collection.Read(with.ID)
	if err != nil {
		return nil, err
	}

	Update(item, with)

	create, err := m.collection.Update(item)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (m *Manager) Delete(id uuid.UUID) error {
	err := m.collection.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

//--- photo

func (m *Manager) AddPhoto(albumID, photoID uuid.UUID) error {

	j := &photo.Join{
		ParentID: albumID,
		PhotoID:  photoID,
	}

	_, err := m.join.Create(j)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) DeletePhoto(albumID, photoID uuid.UUID) error {

	j := &photo.Join{
		ParentID: albumID,
		PhotoID:  photoID,
	}

	err := m.join.Delete(j.GetCompositeKey())
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) ReadCollection(id uuid.UUID) (*Album, error) {
	item, err := m.collection.Read(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadCollections(albumID uuid.UUID, with *photo.SearchOptions) (*photo.PHCollection[*Album], error) {

	item, err := m.collection.Read(albumID)
	if err != nil {
		return nil, err
	}

	all, err := m.join.GetByParentID(item.ID)
	if err != nil {
		return nil, err
	}

	photos, err := m.photoManager.ReadJoinPhotos(all, with)
	if err != nil {
		return nil, err
	}

	a := &photo.PHCollection[*Album]{
		Item:   item,
		Photos: photos,
	}

	return a, nil
}
