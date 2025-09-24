package camera

import (
	"time"

	"github.com/google/uuid"
)

func (c *Camera) GetID() uuid.UUID         { return c.ID }
func (c *Camera) SetID(id uuid.UUID)       { c.ID = id }
func (c *Camera) SetCreatedAt(t time.Time) { c.CreatedAt = t }
func (c *Camera) SetUpdatedAt(t time.Time) { c.UpdatedAt = t }
func (c *Camera) GetRecordSize() int       { return 2048 }

type Camera struct {
	ID          uuid.UUID `json:"id"` // unique: true
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Type        string    `json:"type"`
	CameraMake  string    `json:"cameraMake"`
	CameraModel string    `json:"cameraModel"`
	Count       int       `json:"count"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
	Version     string    `json:"version"`
}
