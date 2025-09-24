package asset

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {
	AssetIds []uuid.UUID `json:"assetIds,omitempty"`

	CameraMake      *string `json:"cameraMake,omitempty"`
	CameraModel     *string `json:"cameraModel,omitempty"`
	IsCamera        *bool
	IsFavorite      *bool
	IsScreenshot    *bool
	IsHidden        *bool
	NotInOnePHAsset *bool

	Albums       *[]string `json:"albums,omitempty"`       // Full album replacement
	AddAlbums    []string  `json:"addAlbums,omitempty"`    // PHAssets to add
	RemoveAlbums []string  `json:"removeAlbums,omitempty"` // PHAssets to remove

	Trips       *[]string `json:"trips,omitempty"`       // Full trip replacement
	AddTrips    []string  `json:"addTrips,omitempty"`    // Trips to add
	RemoveTrips []string  `json:"removeTrips,omitempty"` // Trips to remove

	Persons       *[]string `json:"persons,omitempty"`       // Full Person replacement
	AddPersons    []string  `json:"addPersons,omitempty"`    // Persons to add
	RemovePersons []string  `json:"removePersons,omitempty"` // Persons to remove
}

// Initialize updater
var metadataUpdater = update.NewUpdater[Asset, UpdateOptions]()

func init() {

	metadataUpdater.AddScalarUpdater(func(a *Asset, u UpdateOptions) {
		if u.IsFavorite != nil {
			a.IsFavorite = *u.IsFavorite
		}
		if u.IsCamera != nil {
			a.IsCamera = *u.IsCamera
		}
		if u.IsHidden != nil {
			a.IsHidden = *u.IsHidden
		}
		if u.IsScreenshot != nil {
			a.IsScreenshot = *u.IsScreenshot
		}
	})

	// Configure Collection operations
	metadataUpdater.AddCollectionUpdater(func(a *Asset, u UpdateOptions) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Albums,
			Add:         u.AddAlbums,
			Remove:      u.RemoveAlbums,
		}
		a.Albums = update.ApplyCollectionUpdate(a.Albums, op)
	})

	metadataUpdater.AddCollectionUpdater(func(a *Asset, u UpdateOptions) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Trips,
			Add:         u.AddTrips,
			Remove:      u.RemoveTrips,
		}
		a.Trips = update.ApplyCollectionUpdate(a.Trips, op)
	})

	metadataUpdater.AddCollectionUpdater(func(a *Asset, u UpdateOptions) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Persons,
			Add:         u.AddPersons,
			Remove:      u.RemovePersons,
		}
		a.Persons = update.ApplyCollectionUpdate(a.Persons, op)
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Asset) {
		a.UpdatedAt = time.Now()
	})
}

func Update(p *Asset, update UpdateOptions) *Asset {
	metadataUpdater.Apply(p, update)
	return p
}

// IsEmpty checks if the Place struct contains zero values for all its fields.
func (l Location) IsEmpty() bool {
	return l.Latitude == 0.0 &&
		l.Longitude == 0.0 &&
		l.City == "" &&
		l.Country == ""
}
