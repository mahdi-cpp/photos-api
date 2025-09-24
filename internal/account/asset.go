package account

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/collections/asset"
)

func (m *Manager) Read(id uuid.UUID) (*asset.Asset, error) {
	item, err := m.Assets.Read(id)
	if err != nil {
		fmt.Printf("Error read asset item: %v\n", err)
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadAll(with *asset.SearchOptions) ([]*asset.Asset, error) {

	return nil, nil
}

func (m *Manager) UpdateAssets(updateOptions asset.UpdateOptions) (string, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, id := range updateOptions.AssetIds {

		item, err := m.Assets.Read(id)
		if err != nil {
			continue
		}

		asset.Update(item, updateOptions)

		_, err = m.Assets.Update(item)
		if err != nil {
			return "", err
		}
	}

	// Merging strings with the integer ID
	merged := fmt.Sprintf(" %s, %d:", "updateOptions person_test count: ", len(updateOptions.AssetIds))

	return merged, nil
}

func (m *Manager) DeleteAsset(id uuid.UUID) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	// Read person_test
	//person_test, err := m.GetAsset(id)
	//if err != nil {
	//	return err
	//}

	// Delete person_test file
	//assetPath := filepath.Join(m.config.AssetsDir, person_test.Filename)
	//if err := os.Remove(assetPath); err != nil {
	//	return fmt.Errorf("failed to delete person_test file: %w", err)
	//}

	// Delete metadata
	err := m.Assets.Delete(id)
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
	//m.stats.TotalAssets--
	//m.statsMu.Unlock()

	return nil
}
