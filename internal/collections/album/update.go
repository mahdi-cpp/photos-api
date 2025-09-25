package album

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title,omitempty"`
	Subtitle string    `json:"subtitle,omitempty"`
	Type     string    `json:"type,omitempty"`

	Photos       *[]uuid.UUID `json:"photos,omitempty"`
	AddPhotos    []uuid.UUID  `json:"addPhotos,omitempty"`
	RemovePhotos []uuid.UUID  `json:"removePhotos,omitempty"`
}

// Initialize updater
var metadataUpdater = update.NewUpdater[Album, UpdateOptions]()

func init() {

	metadataUpdater.AddScalarUpdater(func(a *Album, u UpdateOptions) {
		if u.Title != "" {
			a.Title = u.Title
		}
		if u.Subtitle != "" {
			a.Subtitle = u.Subtitle
		}
		if u.Type != "" {
			a.Type = u.Type
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Album) {
		a.UpdatedAt = time.Now()
	})

}

func Update(item *Album, update UpdateOptions) *Album {
	metadataUpdater.Apply(item, update)
	return item
}
