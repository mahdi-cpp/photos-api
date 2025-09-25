package photo

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
)

func (i *Index) GetID() uuid.UUID         { return i.ID }
func (i *Index) SetID(id uuid.UUID)       { i.ID = id }
func (i *Index) SetCreatedAt(t time.Time) { i.CreatedAt = t }
func (i *Index) SetUpdatedAt(t time.Time) { i.UpdatedAt = t }
func (i *Index) GetRecordSize() int       { return 400 }

func (p *Photo) GetID() uuid.UUID         { return p.ID }
func (p *Photo) SetID(id uuid.UUID)       { p.ID = id }
func (p *Photo) SetCreatedAt(t time.Time) { p.CreatedAt = t }
func (p *Photo) SetUpdatedAt(t time.Time) { p.UpdatedAt = t }
func (p *Photo) GetRecordSize() int       { return 2048 }

func (a *Join) GetRecordSize() int { return 100 }
func (a *Join) GetCompositeKey() string {
	return fmt.Sprintf("%s:%s", a.ParentID.String(), a.PhotoID.String())
}

type Join struct {
	ParentID uuid.UUID `json:"parentID"`
	PhotoID  uuid.UUID `json:"photoId"`
}

type Index struct {
	ID                  uuid.UUID `json:"id"`
	UserID              uuid.UUID `json:"userId"`
	CameraMake          string    `json:"cameraMake,omitempty"`
	CameraModel         string    `json:"cameraModel,omitempty"`
	IsCamera            bool      `json:"isCamera"`
	IsFavorite          bool      `json:"isFavorite"`
	IsScreenshot        bool      `json:"isScreenshot"`
	IsHidden            bool      `json:"isHidden"`
	CanDelete           bool      `json:"canDelete"`
	CanEditContent      bool      `json:"canEditContent"`
	CanAddToSharedAlbum bool      `json:"canAddToSharedAlbum"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type Photo struct {
	ID           uuid.UUID  `json:"id" index:"true"`
	UserID       uuid.UUID  `json:"userId" index:"true"`
	FileInfo     FileInfo   `json:"fileInfo"`
	ImageInfo    ImageInfo  `json:"imageInfo"`
	VideoInfo    VideoInfo  `json:"videoInfo"`
	CameraInfo   CameraInfo `json:"cameraInfo"`
	Location     Location   `json:"location"`
	IsCamera     bool       `json:"isCamera"`
	IsFavorite   bool       `json:"isFavorite"`
	IsScreenshot bool       `json:"isScreenshot"`
	IsHidden     bool       `json:"isHidden"`
	CameraMake   string     `json:"cameraMake,omitempty"`
	CameraModel  string     `json:"cameraModel,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    time.Time  `json:"deletedAt"`
	Version      string     `json:"version"`
}

type FileInfo struct {
	BaseURL  string `json:"baseURL"`
	FileSize string `json:"fileSize"`
	MimeType string `json:"mimeType"`
}

type ImageInfo struct {
	Width           int    `json:"width,omitempty"`
	Height          int    `json:"height,omitempty"`
	Megapixels      int    `json:"megapixels,omitempty"`
	Orientation     string `json:"orientation,omitempty"`
	ColorSpace      string `json:"colorSpace,omitempty"`
	EncodingProcess string `json:"encodingProcess,omitempty"`
}
type CameraInfo struct {
	Software         string    `json:"software,omitempty"`
	DateTimeOriginal time.Time `json:"dateTimeOriginal,omitempty"`
	ExposureTime     string    `json:"exposureTime,omitempty"`
	FNumber          float64   `json:"fNumber,omitempty"` // دیافراگم معمولاً float است
	ISO              int       `json:"iso,omitempty"`     // ISO معمولاً عدد صحیح است
	FocalLength      string    `json:"focalLength,omitempty"`
	FocalLength35mm  string    `json:"focalLength35mm,omitempty"`
	Flash            string    `json:"flash,omitempty"`
	LightSource      string    `json:"lightSource,omitempty"`
	ExposureMode     string    `json:"exposureMode,omitempty"`
	WhiteBalance     string    `json:"whiteBalance,omitempty"`
}

type VideoInfo struct {
	MediaDuration      string  `json:"mediaDuration,omitempty"`  // Video duration
	Width              int     `json:"width,omitempty"`          // Video frame width in pixels
	Height             int     `json:"height,omitempty"`         // Video frame height in pixels
	VideoFrameRate     float64 `json:"videoFrameRate,omitempty"` // Video frame rate
	AvgBitrate         string  `json:"avgBitrate,omitempty"`     // Average bitrate (quality and data volume)
	Encoder            string  `json:"encoder,omitempty"`        // Video encoding software
	Rotation           int     `json:"rotation,omitempty"`
	AudioFormat        string  `json:"audioFormat,omitempty"`
	AudioChannels      int     `json:"audioChannels,omitempty"`   // Number of audio channels (e.g., 2 for stereo)
	AudioSampleRate    int     `json:"audioSampleRate,omitempty"` // Audio sample rate
	AudioBitsPerSample int     `json:"audioBitsPerSample,omitempty"`
}

type Location struct {
	Latitude   float64 `json:"location,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
	Country    string  `json:"country,omitempty"`
	Province   string  `json:"province,omitempty"`
	County     string  `json:"county,omitempty"`
	City       string  `json:"city,omitempty"`
	Village    string  `json:"village,omitempty"`
	Electronic int     `json:"electronic,omitempty"`
}

type Collection[T collection_manager_memory.CollectionItem] struct {
	CollectionMemory *collection_manager_memory.Manager[T]
	CoverPhotoArray  map[uuid.UUID][]*Photo
}

func NewCollection[T collection_manager_memory.CollectionItem](path string, fileName string) *Collection[T] {

	c, err := collection_manager_memory.New[T](path, fileName)
	if err != nil {
		panic(err)
	}

	a := &Collection[T]{
		CollectionMemory: c,
		CoverPhotoArray:  make(map[uuid.UUID][]*Photo),
	}

	return a
}

type PHCollectionList[T any] struct {
	Status      string             `json:"status"` // "success" or "error"
	Collections []*PHCollection[T] `json:"collections"`
}

type PHCollection[T any] struct {
	Item   T        `json:"item"`   // Generic items
	Photos []*Photo `json:"photos"` // Specific photos
}
