package camera_handler

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/mahdi-cpp/photos-api/internal/application"
	"github.com/mahdi-cpp/photos-api/internal/collections/camera"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
	"github.com/mahdi-cpp/photos-api/internal/help"
	"github.com/mahdi-cpp/photos-api/mygin"
)

type CameraHandler struct {
	appManager *application.AppManager
}

func New(manager *application.AppManager) *CameraHandler {
	return &CameraHandler{appManager: manager}
}

func (h *CameraHandler) ReadCollections(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		fmt.Println("user id invalid")
		return
	}

	var with *camera.SearchOptions
	err := json.NewDecoder(c.Req.Body).Decode(&with)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("account error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	collections := accountManager.CameraManager.ReadCollections(with)

	fmt.Println("Read Camera Collections count", len(collections))

	c.JSON(http.StatusOK, mygin.H{"collections": collections})
}

func (h *CameraHandler) ReadCollectionPhotos(c *mygin.Context) {

	fmt.Println("Read cameras Photos")

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		fmt.Println("user id invalid")
		return
	}

	var with *photo.SearchOptions
	err := json.NewDecoder(c.Req.Body).Decode(&with)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	photos, err := accountManager.CameraManager.ReadCameraPhotos(with.AlbumID, with)

	fmt.Println("Read Camera Photos count", len(photos))

	c.JSON(http.StatusOK, photos)
}
