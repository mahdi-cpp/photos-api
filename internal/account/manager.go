package account

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_index"
	"github.com/mahdi-cpp/photos-api/internal/collections/album"
	"github.com/mahdi-cpp/photos-api/internal/collections/ali"
	"github.com/mahdi-cpp/photos-api/internal/collections/asset"
	"github.com/mahdi-cpp/photos-api/internal/collections/camera"
	"github.com/mahdi-cpp/photos-api/internal/collections/person"
	"github.com/mahdi-cpp/photos-api/internal/collections/pinned"
	"github.com/mahdi-cpp/photos-api/internal/collections/shared_album"
	"github.com/mahdi-cpp/photos-api/internal/collections/trip"
	"github.com/mahdi-cpp/photos-api/internal/config"
)

type Manager struct {
	mu           sync.RWMutex
	userID       uuid.UUID
	Assets       *collection_manager_index.Manager[*asset.Asset, *asset.Index]
	Album        *asset.Collection[*album.Album]
	Trips        *asset.Collection[*trip.Trip]
	Persons      *asset.Collection[*person.Person]
	Pinned       *asset.Collection[*pinned.Pinned]
	SharedAlbums *asset.Collection[*shared_album.SharedAlbum]
	cameras      map[string]*ali.PHCollection[camera.Camera]
	statsMu      sync.Mutex
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

	a := config.GetUserMetadataPath(userID.String(), "assets")
	var err error
	m.Assets, err = collection_manager_index.New[*asset.Asset, *asset.Index](a)
	if err != nil {
		panic(err)
	}

	all := m.Assets.GetAllIndexes()
	fmt.Println(len(all))

	path := config.GetUserMetadataPath(userID.String(), "")
	m.Album = asset.NewCollection[*album.Album](path + "/albums")
	m.SharedAlbums = asset.NewCollection[*shared_album.SharedAlbum](path + "/shared_albums")
	m.Trips = asset.NewCollection[*trip.Trip](path + "/trips")
	m.Persons = asset.NewCollection[*person.Person](path + "/persons")
	m.Pinned = asset.NewCollection[*pinned.Pinned](path + "/pins")

	m.prepareAlbums()
	m.prepareTrips()
	m.preparePersons()
	//m.prepareCameras()
	m.preparePinned()

	return m, nil
}

func (m *Manager) UpdateCollections() {
	m.prepareAlbums()
	//m.prepareCameras()
	m.prepareTrips()
	m.preparePersons()
	m.preparePinned()
}
