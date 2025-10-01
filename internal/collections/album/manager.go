package album

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_join"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
)

type Manager struct {
	photoManager *photo.Manager
	collection   *collection_manager_memory.Manager[*Album]
	join         *collection_manager_join.Manager[*photo.Join]
	albums       []*Album
	photos       map[uuid.UUID][]*photo.Photo //key is albumId
}

func NewManager(photoManager *photo.Manager, path string) (*Manager, error) {

	manager := &Manager{
		photoManager: photoManager,
		photos:       make(map[uuid.UUID][]*photo.Photo),
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

	err = manager.load()
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) load() error {

	items, err := m.collection.ReadAll()
	if err != nil {
		return err
	}

	m.albums = []*Album{}

	for _, item := range items {

		allPhotos, err := m.join.GetByParentID(item.ID)
		if err != nil {
			continue
		}

		item.Count = len(allPhotos)
		with := &photo.SearchOptions{
			Sort:      "id",
			SortOrder: "desc",
			Page:      0,
			Size:      5,
		}
		albumPhotos, err := m.photoManager.ReadJoinPhotos(allPhotos, with)
		if err != nil {
			continue
		}

		m.albums = append(m.albums, item)
		m.photos[item.ID] = albumPhotos
	}

	return nil
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

func (m *Manager) ReadAll(with *SearchOptions) ([]*Album, error) {
	items, err := m.collection.ReadAll()
	if err != nil {
		return nil, err
	}

	filterItems := Search(items, with)
	return filterItems, nil
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

	all, err := m.join.ReadAll()
	if err != nil {
		return err
	}

	for _, item := range all {
		if item.ParentID == id {
			err := m.join.Delete(item.GetCompositeKey())
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	return nil
}

func (m *Manager) IsExist(id uuid.UUID) error {
	_, err := m.collection.Read(id)
	if err != nil {
		return fmt.Errorf("album not found: %s", id)
	}

	return nil
}

//--- photo

func (m *Manager) AddPhoto(albumID, photoID uuid.UUID) error {

	if albumID == uuid.Nil {
		return fmt.Errorf("albumID must not be an empty string")
	}
	if photoID == uuid.Nil {
		return fmt.Errorf("photoID must not be an empty string")
	}

	j := &photo.Join{
		ParentID: albumID,
		PhotoID:  photoID,
	}

	_, err := m.join.Create(j)
	if err != nil {
		return err
	}

	err = m.load()
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

	err = m.load()
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

func (m *Manager) ReadCollections(with *SearchOptions) []*photo.Collection[*Album] {

	var results []*photo.Collection[*Album]

	filterAlbums := Search(m.albums, with)

	for _, album := range filterAlbums {
		collection := &photo.Collection[*Album]{
			Item:   album,
			Photos: m.photos[album.ID],
		}
		results = append(results, collection)
	}

	return results
}

//--- events

func (m *Manager) HandlePhotoCreate(id uuid.UUID) {

}

func (m *Manager) HandlePhotoUpdate(id uuid.UUID) {

}

func (m *Manager) HandlePhotoDelete(id uuid.UUID) {

}
