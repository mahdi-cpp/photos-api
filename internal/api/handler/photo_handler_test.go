package handler

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
	"github.com/mahdi-cpp/photos-api/internal/help"
)

const baseURL = "http://localhost:50000"

func TestAssetHandler_Create(t *testing.T) {

	currentURL := baseURL + "/photos/api/photos"

	body := &photo.Photo{
		FileInfo: photo.FileInfo{
			BaseURL:  "",
			FileSize: "1000",
			MimeType: "voice/mp3",
		},
		ImageInfo: photo.ImageInfo{
			Width:       1000,
			Height:      1000,
			Orientation: "portrait",
		},
		IsCamera: true,
	}

	resp, err := help.MakeRequestBody("POST", currentURL, body)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response: %v", err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		var r Error
		if err := json.Unmarshal(respBody, &r); err != nil {
			t.Fatalf("unmarshaling response: %v", err)
		}
		t.Fatalf("error %s", r.Message)
	}

	var a photo.Photo
	if err := json.Unmarshal(respBody, &a); err != nil {
		t.Fatalf("unmarshaling response: %v", err)
	}

	fmt.Println("new photo id: ", a.ID)
}
