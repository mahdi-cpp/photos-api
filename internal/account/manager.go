package account

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/collections/album"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
	"github.com/mahdi-cpp/photos-api/internal/config"
)

type Manager struct {
	mu     sync.RWMutex
	userID uuid.UUID

	PhotosManager *photo.Manager
	AlbumsManager *album.Manager

	statsMu sync.Mutex
}

func New(userID uuid.UUID) (*Manager, error) {

	m := &Manager{
		userID: userID,
	}

	err := directory(userID)
	if err != nil {
		return nil, err
	}

	accountDir := config.GetUserMetadataPath(userID.String(), "")

	m.PhotosManager, err = photo.NewManager(accountDir)
	if err != nil {
		panic(err)
	}

	m.AlbumsManager, err = album.NewManager(m.PhotosManager, accountDir)
	if err != nil {
		panic(err)
	}

	return m, nil
}

func directory(userID uuid.UUID) error {

	// Ensure account directories exist
	userDirectory := filepath.Join(config.GetUserPath(userID.String()))
	userMetadata := filepath.Join(userDirectory, "metadata")
	userDirs := []string{userDirectory, userMetadata}

	for _, dir := range userDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create account directory %s: %w", dir, err)
		}
	}
	return nil
}
