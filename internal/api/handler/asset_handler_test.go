package handler

import (
	"io"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
	"github.com/mahdi-cpp/photos-api/internal/collections/asset"
	"github.com/mahdi-cpp/photos-api/internal/help"
)

const baseURL = "http://localhost:50000"

func TestAssetHandler_Create(t *testing.T) {

	currentURL := baseURL + "/photos/api/assets"

	body := &asset.Asset{
		FileInfo: asset.FileInfo{
			BaseURL:  "",
			FileSize: "1000",
			MimeType: "video/mp4",
		},
		ImageInfo: asset.ImageInfo{
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

	if resp.StatusCode != http.StatusCreated {
		var r Error
		if err := json.Unmarshal(respBody, &r); err != nil {
			t.Fatalf("unmarshaling response: %v", err)
		}
		t.Fatalf("error %s", r.Message)
	}

	var asset asset.Asset
	if err := json.Unmarshal(respBody, &asset); err != nil {
		t.Fatalf("unmarshaling response: %v", err)
	}
}
