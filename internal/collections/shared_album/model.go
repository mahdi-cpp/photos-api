package shared_album

import (
	"time"

	"github.com/google/uuid"
)

func (s *SharedAlbum) GetID() uuid.UUID         { return s.ID }
func (s *SharedAlbum) SetID(id uuid.UUID)       { s.ID = id }
func (s *SharedAlbum) SetCreatedAt(t time.Time) { s.CreatedAt = t }
func (s *SharedAlbum) SetUpdatedAt(t time.Time) { s.UpdatedAt = t }
func (s *SharedAlbum) GetRecordSize() int       { return 2048 }

type SharedAlbum struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Subtitle     string    `json:"subtitle"`
	Type         string    `json:"type"`
	Count        int       `json:"count"`
	IsCollection bool      `json:"isCollection"`
	IsHidden     bool      `json:"isHidden"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
	Version      string    `json:"version"`
}
