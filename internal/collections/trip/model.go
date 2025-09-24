package trip

import (
	"time"

	"github.com/google/uuid"
)

func (tr *Trip) GetID() uuid.UUID         { return tr.ID }
func (tr *Trip) SetID(id uuid.UUID)       { tr.ID = id }
func (tr *Trip) SetCreatedAt(t time.Time) { tr.CreatedAt = t }
func (tr *Trip) SetUpdatedAt(t time.Time) { tr.UpdatedAt = t }
func (tr *Trip) GetRecordSize() int       { return 2048 }

type Trip struct {
	ID           uuid.UUID `json:"id"` // unique: true
	Title        string    `json:"title"`
	Subtitle     string    `json:"subtitle"`
	Type         string    `json:"type"`
	Count        int       `json:"count"`
	IsCollection bool      `json:"isCollection"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
	Version      string    `json:"version"`
}
