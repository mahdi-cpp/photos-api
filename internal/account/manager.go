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

	//Trips        *photo.Collection[*trip.Trip]
	//Persons      *photo.Collection[*person.Person]
	//Pinned       *photo.Collection[*pinned.Pinned]
	//SharedAlbums *photo.Collection[*shared_album.SharedAlbum]
	//cameras      map[string]*ali.PHCollection[camera.Camera]
	statsMu sync.Mutex
}

func New(userID uuid.UUID) (*Manager, error) {

	m := &Manager{
		userID: userID,
	}

	// Ensure account directories exist
	userDirectory := filepath.Join(config.GetUserPath(userID.String()))
	userMetadata := filepath.Join(userDirectory, "metadata")
	userDirs := []string{userDirectory, userMetadata}

	for _, dir := range userDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create account directory %s: %w", dir, err)
		}
	}

	//a := config.GetUserMetadataPath(userID.String(), "")
	path := config.GetUserMetadataPath(userID.String(), "")

	var err error
	m.PhotosManager, err = photo.NewManager(path)
	if err != nil {
		panic(err)
	}

	m.AlbumsManager, err = album.NewManager(m.PhotosManager, path)
	if err != nil {
		panic(err)
	}

	//m.SharedAlbums = photo.NewCollection[*shared_album.SharedAlbum](path, "shared_album")
	//m.Trips = photo.NewCollection[*trip.Trip](path, "trip")
	//m.Persons = photo.NewCollection[*person.Person](path, "person")
	//m.Pinned = photo.NewCollection[*pinned.Pinned](path, "pinned")

	//m.prepareAlbums()
	//m.prepareTrips()
	//m.preparePersons()
	//m.preparePinned()

	return m, nil
}

func (m *Manager) UpdateCollections() {
	//m.prepareAlbums()
	//m.prepareTrips()
	//m.preparePersons()
	//m.preparePinned()
}
