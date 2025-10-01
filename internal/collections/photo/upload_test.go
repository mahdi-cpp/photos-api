package photo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/collections/metadata"
	"github.com/mahdi-cpp/photos-api/internal/config"
	"github.com/mahdi-cpp/photos-api/internal/help"
)

const uploadURL = "http://localhost:50103/api/v1/upload/"
const photoApi = "http://localhost:50000/photos/api/photos"

// httpClient is a shared instance of the HTTP client for efficiency.
var httpClient = &http.Client{Timeout: 30 * time.Second}

func TestMessageCreate(t *testing.T) {

	workDir := "/app/tmp/all/"
	entries, err := os.ReadDir(workDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			singleUpload(t, workDir, entry.Name())
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func singleUpload(t *testing.T, uploadFilePath string, fileName string) {

	var err error

	// 1 create upload directory
	var apiURL = uploadURL + "create"
	respBody, err := help.MakeRequestParam("POST", apiURL, nil)
	if err != nil {
		t.Errorf("create request failed: %v", err)
	}

	var workDir DirectoryRequest
	if err := json.Unmarshal(respBody, &workDir); err != nil {
		t.Errorf("unmarshaling response: %v", err)
	}

	// 2 upload image
	fmt.Println(workDir.ID)
	apiURL = uploadURL + "media"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	photo, err := upload(ctx, httpClient, apiURL, workDir.ID, uploadFilePath+fileName)
	if err != nil {
		t.Errorf("%v", err)
	}

	if photo == nil {
		t.Fatal("Expected a non-nil response, but got nil")
	}

	body := &UploadInfo{
		Directory: workDir.ID.String(),
		FileName:  fileName,
		Photo:     *photo,
	}

	// 3 send to photos-api
	response, err := help.MakeRequestBody("POST", photoApi, body)
	if err != nil {
		t.Errorf("create request body failed: %v", err)
	}

	if response.StatusCode == http.StatusBadRequest {
		var e help.Error
		err = json.NewDecoder(response.Body).Decode(&e)
		if err != nil {
			t.Fatal(err)
		}
		t.Fatal(e.Message)
	}
}

func upload(ctx context.Context, client *http.Client, apiURL string, directoryID uuid.UUID, filePath string) (*Photo, error) {

	var mimeType = ""
	if IsJPEG(filePath) {
		mimeType = "image/jpeg"
	} else {
		mimeType = "video/mp4"
	}

	// Create the Request struct with the necessary metadata.
	var uploadReq = &Request{}
	if mimeType == "image/jpeg" {
		uploadReq = &Request{
			Directory: directoryID,
			IsVideo:   false,
		}
	} else {
		uploadReq = &Request{
			Directory: directoryID,
			IsVideo:   true,
		}
	}

	// Call the mediaUpload function.
	resp, err := MediaUpload(ctx, client, apiURL, filePath, uploadReq)
	if err != nil {
		return nil, fmt.Errorf("media upload failed: %w", err)
	}
	var photo = &Photo{}

	fileName := strings.TrimSuffix(resp.FileInfo.BaseURL, ".jpg")

	photo.CameraMake = resp.Camera.Make
	photo.CameraModel = resp.Camera.Model

	photo.FileInfo = FileInfo{
		OriginalURL:  filepath.Join(config.RootDir, "users", config.TestUserID, "assets", fileName+".jpg"),
		ThumbnailURL: filepath.Join(config.RootDir, "users", config.TestUserID, "thumbnails", fileName),
		FileName:     fileName,
		FileSize:     resp.FileInfo.FileSize,
		MimeType:     resp.FileInfo.MimeType,
	}

	if mimeType == "image/jpeg" {
		photo.ImageInfo = ImageInfo{
			Width:       resp.Image.Width,
			Height:      resp.Image.Height,
			Orientation: resp.Image.Orientation,
		}
	} else {
		photo.VideoInfo = VideoInfo{
			Width:  resp.Video.Width,
			Height: resp.Video.Height,
		}
	}

	return photo, nil
}

func MediaUpload(ctx context.Context, client *http.Client, apiURL, filePath string, uploadRequest *Request) (*metadata.Metadata, error) {

	// Open the file to be uploaded.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	// Create a new multipart writer.
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Marshal the struct into a JSON string.
	jsonData, err := json.Marshal(uploadRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON data: %w", err)
	}

	// Create a new form field for the JSON data.
	if err := writer.WriteField("metadata", string(jsonData)); err != nil {
		return nil, fmt.Errorf("failed to write JSON data to form field: %w", err)
	}

	// Create a form file part for the media.
	filePart, err := writer.CreateFormFile("media", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	// Copy the file content into the form file part.
	if _, err := io.Copy(filePart, file); err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}

	// Close the multipart writer to finalize the body.
	writer.Close()

	// Create the HTTP request.
	req, err := http.NewRequestWithContext(ctx, "POST", apiURL, &requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Caption-Type header with the boundary from the writer.
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request.
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check for a successful response status code.
	if resp.StatusCode != http.StatusOK {
		// Read and log the server's error message.
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("server responded with status code %d, but could not read error body: %w", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("server responded with status code %d and body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Decode the JSON response body.
	var serverResponse metadata.Metadata
	if err := json.NewDecoder(resp.Body).Decode(&serverResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	fmt.Printf("Successfully uploaded image from %s\n", filePath)
	return &serverResponse, nil
}

func IsJPEG(filePath string) bool {
	// Convert the file path to lowercase for a case-insensitive check.
	lowerCasePath := strings.ToLower(filePath)

	return strings.HasSuffix(lowerCasePath, ".jpg") || strings.HasSuffix(lowerCasePath, ".jpeg")
}
