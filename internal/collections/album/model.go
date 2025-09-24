package album

import (
	"time"

	"github.com/google/uuid"
)

func (a *Album) GetID() uuid.UUID         { return a.ID }
func (a *Album) SetID(id uuid.UUID)       { a.ID = id }
func (a *Album) SetCreatedAt(t time.Time) { a.CreatedAt = t }
func (a *Album) SetUpdatedAt(t time.Time) { a.UpdatedAt = t }
func (a *Album) GetRecordSize() int       { return 2048 }

type Album struct {
	ID                uuid.UUID `json:"id"` // unique: true
	Title             string    `json:"title"`
	Subtitle          string    `json:"subtitle"`
	Type              string    `json:"type"`
	Number            int       `json:"number"`
	IsCollectionValid bool      `json:"isCollectionValid"`
	IsHidden          bool      `json:"isHidden"`
	Count             int       `json:"count"`
	LastSeen          time.Time `json:"lastSeen"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         time.Time `json:"deletedAt"`
	Version           string    `json:"version"`
}
