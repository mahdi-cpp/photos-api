package pinned

import (
	"time"

	"github.com/google/uuid"
)

func (p *Pinned) GetID() uuid.UUID         { return p.ID }
func (p *Pinned) SetID(id uuid.UUID)       { p.ID = id }
func (p *Pinned) SetCreatedAt(t time.Time) { p.CreatedAt = t }
func (p *Pinned) SetUpdatedAt(t time.Time) { p.UpdatedAt = t }
func (p *Pinned) GetRecordSize() int       { return 2048 }

type Pinned struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Type      string    `json:"type"`
	AlbumID   uuid.UUID `json:"albumID"`
	Icon      string    `json:"icon"`
	Count     int       `json:"count"`
	Index     int       `json:"index"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Version   string    `json:"version"`
}
