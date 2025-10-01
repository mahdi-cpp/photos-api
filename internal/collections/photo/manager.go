package photo

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_index"
)

type Manager struct {
	mu         sync.RWMutex
	userID     uuid.UUID
	collection *collection_manager_index.Manager[*Photo, *Index]
	onCallback OnCallback
}

func NewManager(userID uuid.UUID, onCallback OnCallback, path string) (*Manager, error) {
	manager := &Manager{
		userID:     userID,
		onCallback: onCallback,
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

	f := strings.ToLower(info.FileName)
	if strings.Contains(f, "screenshot") {
		info.Photo.IsScreenshot = true
	}
	info.Photo.Version = "2"

	photoCreated, err := m.collection.Create(&info.Photo)
	if err != nil {
		return nil, err
	}

	m.onCallback("create", photoCreated.ID)

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
		photo, err := m.collection.Read(index.ID)
		if err != nil {
			return nil, fmt.Errorf("error reading message %s: %w", index.ID, err)
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func (m *Manager) Update(with *UpdateOptions) (string, error) {

	fmt.Println("photo manager Update")
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, id := range with.PhotosIDs {

		item, err := m.collection.Read(id)
		if err != nil {
			continue
		}

		Update(item, *with)

		_, err = m.collection.Update(item)
		if err != nil {
			return "", err
		}
	}

	// Merging strings with the integer ID
	merged := fmt.Sprintf(" %s, %d:", "with person_test count: ", len(with.PhotosIDs))

	return merged, nil
}

func (m *Manager) Delete(id uuid.UUID) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	p, err := m.collection.Read(id)
	if err != nil {
		return err
	}

	err = m.collection.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete metadata: %w", err)
	}

	err = safeRemove(p.FileInfo.OriginalURL)
	if err != nil {
		return fmt.Errorf("failed to remove originalURL: %w", err)
	}
	err = safeRemove(p.FileInfo.ThumbnailURL + "_270.jpg")
	if err != nil {
		return fmt.Errorf("failed to remove thumbnailURL: %w", err)
	}

	m.onCallback("delete", id)

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

func safeRemove(path string) error {
	// ۱. بررسی نام دایرکتوری
	// اگر نام دایرکتوری دقیقا "protected_dir" باشد، عملیات حذف را متوقف کن
	if path == "protected_dir" {
		return fmt.Errorf("Error: Deletion of the protected directory '%s' is not allowed", path)
	}

	// ۲. ادامه عملیات حذف (اگر نام مجاز باشد)
	// از os.RemoveAll برای اطمینان از حذف فایل/دایرکتوری استفاده می‌شود
	err := os.RemoveAll(path)
	if err != nil {
		// مدیریت خطاهای دیگر (مثل عدم وجود)
		return fmt.Errorf("failed to remove %s: %w", path, err)
	}
	return nil
}
