package person

import (
	"time"

	"github.com/google/uuid"
)

func (p *Person) GetID() uuid.UUID         { return p.ID }
func (p *Person) SetID(id uuid.UUID)       { p.ID = id }
func (p *Person) SetCreatedAt(t time.Time) { p.CreatedAt = t }
func (p *Person) SetUpdatedAt(t time.Time) { p.UpdatedAt = t }
func (p *Person) GetRecordSize() int       { return 2048 }

type Person struct {
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
