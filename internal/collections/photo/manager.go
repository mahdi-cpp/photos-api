package photo

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_index"
)

type Manager struct {
	mu               sync.RWMutex
	userID           uuid.UUID
	collection       *collection_manager_index.Manager[*Photo, *Index]
	onCreateCallback OnCreateCallback
}

func NewManager(userID uuid.UUID, onCreateCallback OnCreateCallback, path string) (*Manager, error) {
	manager := &Manager{
		userID:           userID,
		onCreateCallback: onCreateCallback,
	}

	var err error
	manager.collection, err = collection_manager_index.New[*Photo, *Index](path)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) Create(info *UploadInfo) (*Photo, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	info.Photo.Version = "2"
	info.Photo.IsFavorite = true

	photoCreated, err := m.collection.Create(&info.Photo)
	if err != nil {
		return nil, err
	}

	m.onCreateCallback(photoCreated)

	err = moveMedia(m.userID, info.Directory, &info.Photo)
	if err != nil {
		return nil, err
	}

	return photoCreated, nil
}

func (m *Manager) Read(id uuid.UUID) (*Photo, error) {
	item, err := m.collection.Read(id)
	if err != nil {
		fmt.Printf("Error read collection item: %v\n", err)
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadAll(with *SearchOptions) ([]*Photo, error) {

	all := m.collection.GetAllIndexes()
	var photos []*Photo

	filterIndexes := Search(all, with)

	for _, index := range filterIndexes {
		read, err := m.collection.Read(index.ID)
		if err != nil {
			return nil, fmt.Errorf("error reading message %s: %w", index.ID, err)
		}
		photos = append(photos, read)
	}

	return photos, nil
}

func (m *Manager) Update(with UpdateOptions) (string, error) {

	fmt.Println("photo manager Update")
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, id := range with.PhotosIds {

		item, err := m.collection.Read(id)
		if err != nil {
			continue
		}

		Update(item, with)

		_, err = m.collection.Update(item)
		if err != nil {
			return "", err
		}
	}

	// Merging strings with the integer ID
	merged := fmt.Sprintf(" %s, %d:", "with person_test count: ", len(with.PhotosIds))

	return merged, nil
}

func (m *Manager) Delete(id uuid.UUID) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	// Read person_test
	//person_test, err := m.GetAsset(id)
	//if err != nil {
	//	return err
	//}

	// Delete person_test file
	//assetPath := filepath.join(m.config.PhotosDir, person_test.Filename)
	//if err := os.Remove(assetPath); err != nil {
	//	return fmt.Errorf("failed to delete person_test file: %w", err)
	//}

	// Delete metadata
	err := m.collection.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete metadata: %w", err)
	}
	//if err := m.metadata.DeleteMetadata(id); err != nil {
	//	return fmt.Errorf("failed to delete metadata: %w", err)
	//}

	// Delete thumbnail (if exists)
	//m.thumbnail.DeleteThumbnails(id)

	// Remove from indexes
	//m.removeFromIndexes(id)

	// Remove from memory
	//m.memory.Remove(id)

	// UpdateOptions stats
	//m.statsMu.Lock()
	//m.stats.TotalPhotos--
	//m.statsMu.Unlock()

	return nil
}

func (m *Manager) ReadByIds(ids []uuid.UUID) ([]*Photo, error) {

	photos := make([]*Photo, 0, len(ids))

	for _, id := range ids {
		read, err := m.collection.Read(id)
		if err != nil {
			return nil, fmt.Errorf("error reading message %s: %w", id, err)
		}
		photos = append(photos, read)
	}

	return photos, nil
}

func (m *Manager) ReadJoinPhotos(joins []*Join, with *SearchOptions) ([]*Photo, error) {

	var photos []*Photo
	var indexes []*Index

	for _, join := range joins {
		index, err := m.collection.ReadIndex(join.PhotoID)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, index)
	}

	filterIndexes := Search(indexes, with)

	for _, index := range filterIndexes {
		read, err := m.collection.Read(index.ID)
		if err != nil {
			return nil, fmt.Errorf("error reading message %s: %w", index.ID, err)
		}
		photos = append(photos, read)
	}

	return photos, nil
}

func (m *Manager) ReadByIndexes(indexes []*Index, with *SearchOptions) ([]*Photo, error) {

	var photos []*Photo
	filterIndexes := Search(indexes, with)

	for _, index := range filterIndexes {
		read, err := m.collection.Read(index.ID)
		if err != nil {
			return nil, fmt.Errorf("error reading message %s: %w", index.ID, err)
		}
		photos = append(photos, read)
	}

	return photos, nil
}

func (m *Manager) ReadIndexes() []*Index {
	return m.collection.GetAllIndexes()
}
