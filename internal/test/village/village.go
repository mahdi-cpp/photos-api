package village

import (
	"time"

	"github.com/google/uuid"
)

func (v *Village) GetID() uuid.UUID         { return v.ID }
func (v *Village) SetID(id uuid.UUID)       { v.ID = id }
func (v *Village) SetCreatedAt(t time.Time) { v.CreatedAt = t }
func (v *Village) SetUpdatedAt(t time.Time) { v.UpdatedAt = t }
func (v *Village) GetRecordSize() int       { return 2048 }

type Village struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Version   string    `json:"version"`
}

type Polygon struct {
	Name string        `json:"name"`
	Data [][][]float64 `json:"data"`
}
