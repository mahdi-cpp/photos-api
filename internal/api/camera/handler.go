package camera_handler

import (
	"fmt"
	"net/http"

	"github.com/mahdi-cpp/photos-api/internal/application"
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

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("account error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	collections := accountManager.CameraManager.ReadCollections()

	fmt.Println("Read Camera Collections count", len(collections))

	c.JSON(http.StatusOK, mygin.H{"collections": collections})
}
