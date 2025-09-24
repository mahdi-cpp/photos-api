package shared_album

import (
	"time"

	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {
	ID       *string `json:"id"`
	Title    *string `json:"title,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// Initialize updater
var metadataUpdater = update.NewUpdater[SharedAlbum, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *SharedAlbum, u UpdateOptions) {
		if u.Title != nil {
			a.Title = *u.Title
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *SharedAlbum, u UpdateOptions) {
		if u.Subtitle != nil {
			a.Subtitle = *u.Subtitle
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *SharedAlbum, u UpdateOptions) {
		if u.Type != nil {
			a.Type = *u.Type
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *SharedAlbum) {
		a.UpdatedAt = time.Now()
	})
}

func Update(item *SharedAlbum, update UpdateOptions) *SharedAlbum {
	metadataUpdater.Apply(item, update)
	return item
}
