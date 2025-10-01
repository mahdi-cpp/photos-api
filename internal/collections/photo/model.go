package photo

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (a *Join) GetRecordSize() int { return 110 }
func (a *Join) GetCompositeKey() string {
	return fmt.Sprintf("%s:%s", a.ParentID.String(), a.PhotoID.String())
}

type Join struct {
	ParentID uuid.UUID `json:"parentID"`
	PhotoID  uuid.UUID `json:"photoId"`
}

func (i *Index) GetID() uuid.UUID         { return i.ID }
func (i *Index) SetID(id uuid.UUID)       { i.ID = id }
func (i *Index) SetCreatedAt(t time.Time) { i.CreatedAt = t }
func (i *Index) SetUpdatedAt(t time.Time) { i.UpdatedAt = t }
func (i *Index) GetRecordSize() int       { return 450 }

type Index struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"userId"`
	CameraMake   string    `json:"cameraMake"`
	CameraModel  string    `json:"cameraModel"`
	IsCamera     bool      `json:"isCamera"`
	IsFavorite   bool      `json:"isFavorite"`
	IsScreenshot bool      `json:"isScreenshot"`
	IsHidden     bool      `json:"isHidden"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (p *Photo) GetID() uuid.UUID         { return p.ID }
func (p *Photo) SetID(id uuid.UUID)       { p.ID = id }
func (p *Photo) SetCreatedAt(t time.Time) { p.CreatedAt = t }
func (p *Photo) SetUpdatedAt(t time.Time) { p.UpdatedAt = t }
func (p *Photo) GetRecordSize() int       { return 2048 }

type Photo struct {
	ID           uuid.UUID  `json:"id" index:"true"`
	UserID       uuid.UUID  `json:"userId" index:"true"`
	FileInfo     FileInfo   `json:"fileInfo"`
	ImageInfo    ImageInfo  `json:"imageInfo"`
	VideoInfo    VideoInfo  `json:"videoInfo"`
	CameraInfo   CameraInfo `json:"cameraInfo"`
	Location     Location   `json:"location"`
	CameraMake   string     `json:"cameraMake" index:"true"`
	CameraModel  string     `json:"cameraModel" index:"true"`
	IsCamera     bool       `json:"isCamera"`
	IsFavorite   bool       `json:"isFavorite"`
	IsScreenshot bool       `json:"isScreenshot"`
	IsHidden     bool       `json:"isHidden"`
	CreatedAt    time.Time  `json:"createdAt" index:"true"`
	UpdatedAt    time.Time  `json:"updatedAt" index:"true"`
	DeletedAt    time.Time  `json:"deletedAt"`
	Version      string     `json:"version"`
}

type FileInfo struct {
	OriginalURL  string `json:"originalURL"`
	ThumbnailURL string `json:"thumbnailURL"`
	FileName     string `json:"fileName"`
	FileSize     int    `json:"fileSize"`
	MimeType     string `json:"mimeType"`
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

// ---

type Request struct {
	Directory uuid.UUID `json:"directory"`
	IsVideo   bool      `json:"isVideo"`
	//Hash      string    `json:"hash"`
}

type Response struct {
	Message  string    `json:"message,omitempty"`
	Filename string    `json:"filename,omitempty"`
	ID       uuid.UUID `json:"id,omitempty"`
	Error    string    `json:"error,omitempty"`
}

type DirectoryRequest struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
	Errors  string    `json:"errors,omitempty"`
}

type UploadInfo struct {
	Directory string `json:"directory"`
	FileName  string `json:"fileName"`
	Photo     Photo  `json:"photo"`
}

type BulkPhoto struct {
	PhotoIDs []uuid.UUID `json:"photoIds"`
}
