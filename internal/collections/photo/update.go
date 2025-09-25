package photo

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {
	PhotosIds []uuid.UUID `json:"photosIds,omitempty"`

	CameraMake     *string `json:"cameraMake,omitempty"`
	CameraModel    *string `json:"cameraModel,omitempty"`
	IsCamera       *bool
	IsFavorite     *bool
	IsScreenshot   *bool
	IsHidden       *bool
	NotInOnePhotos *bool

	Trips       *[]string `json:"trips,omitempty"`       // Full trip replacement
	AddTrips    []string  `json:"addTrips,omitempty"`    // Trips to add
	RemoveTrips []string  `json:"removeTrips,omitempty"` // Trips to remove

	Persons       *[]string `json:"persons,omitempty"`       // Full Person replacement
	AddPersons    []string  `json:"addPersons,omitempty"`    // Persons to add
	RemovePersons []string  `json:"removePersons,omitempty"` // Persons to remove
}

// Initialize updater
var metadataUpdater = update.NewUpdater[Photo, UpdateOptions]()

func init() {

	metadataUpdater.AddScalarUpdater(func(a *Photo, u UpdateOptions) {
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

	//metadataUpdater.AddCollectionUpdater(func(a *Photo, u UpdateOptions) {
	//	op := update.CollectionUpdateOp[string]{
	//		FullReplace: u.Trips,
	//		Add:         u.AddTrips,
	//		Remove:      u.RemoveTrips,
	//	}
	//	a.Trips = update.ApplyCollectionUpdate(a.Trips, op)
	//})
	//
	//metadataUpdater.AddCollectionUpdater(func(a *Photo, u UpdateOptions) {
	//	op := update.CollectionUpdateOp[string]{
	//		FullReplace: u.Persons,
	//		Add:         u.AddPersons,
	//		Remove:      u.RemovePersons,
	//	}
	//	a.Persons = update.ApplyCollectionUpdate(a.Persons, op)
	//})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Photo) {
		a.UpdatedAt = time.Now()
	})
}

func Update(p *Photo, update UpdateOptions) *Photo {
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
