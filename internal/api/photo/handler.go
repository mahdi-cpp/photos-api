package photo_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/mygin"
	"github.com/mahdi-cpp/photos-api/internal/application"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
	"github.com/mahdi-cpp/photos-api/internal/help"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type PhotoHandler struct {
	appManager *application.AppManager
}

func New(manager *application.AppManager) *PhotoHandler {
	return &PhotoHandler{appManager: manager}
}

func SendError(c *mygin.Context, message string, code int) {
	c.JSON(http.StatusBadRequest, mygin.H{"message": message, "code": code})
}

func (h *PhotoHandler) Create(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		SendError(c, "user id invalid", http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var request *photo.Photo
	err = json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	create, err := accountManager.PhotosManager.Create(request)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, create)
}

func (h *PhotoHandler) Read(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	photoID := c.Param("id")
	id, err := uuid.Parse(photoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	item, err := accountManager.PhotosManager.Read(id)
	if err != nil {
		c.JSON(http.StatusNotFound, mygin.H{"error": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *PhotoHandler) ReadAll(c *mygin.Context) {

	fmt.Println("1")

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		fmt.Println("user id invalid")
		return
	}

	fmt.Println("2")

	page, err := c.GetQueryInt("page")
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
	}
	size, err := c.GetQueryInt("size")
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
	}

	fmt.Println("page:", page)
	fmt.Println("size:", size)

	with := &photo.SearchOptions{
		Page: page,
		Size: size,
	}

	//err := json.NewDecoder(c.Request.Body).Decode(&with)
	//if err != nil {
	//	fmt.Println("%w", err)
	//	help.AbortWithRequestInvalid(c)
	//	return
	//}

	fmt.Println("3")

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	items, err := accountManager.PhotosManager.ReadAll(with)
	if err != nil {
		c.JSON(http.StatusInternalServerError, mygin.H{"error": "failed account Read"})
		return
	}

	for _, item := range items {
		fmt.Println(item.FileInfo.OriginalURL)
	}

	fmt.Println("ReadAll count", len(items))

	c.JSON(http.StatusOK, items)
}

//
//func (h *PhotoHandler) Update(c *gin.Context) {
//
//	userID, ok := help.GetUserID(c)
//	if !ok {
//		help.AbortWithUserIDInvalid(c)
//		return
//	}
//
//	var updateOptions photo.UpdateOptions
//	if err := c.ShouldBindJSON(&updateOptions); err != nil {
//		help.AbortWithRequestInvalid(c)
//		return
//	}
//
//	userManager, err := h.appManager.GetAccountManager(c, userID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
//		return
//	}
//
//	photo, err := userManager.Update(updateOptions)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
//		return
//	}
//
//	userManager.UpdateCollections()
//
//	c.JSON(http.StatusCreated, photo)
//}
//
//func (h *PhotoHandler) BuckUpdate(c *gin.Context) {
//	userID, ok := help.GetUserID(c)
//	if !ok {
//		help.AbortWithUserIDInvalid(c)
//		return
//	}
//
//	var updateOptions photo.UpdateOptions
//	if err := c.ShouldBindJSON(&updateOptions); err != nil {
//		help.AbortWithRequestInvalid(c)
//		return
//	}
//
//	userManager, err := h.appManager.GetAccountManager(c, userID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
//		return
//	}
//
//	allAssets, err := userManager.ReadAll()
//	if err != nil {
//		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
//		return
//	}
//
//	for _, photo := range allAssets {
//		updateOptions.PhotosIds = append(updateOptions.PhotosIds, photo.ID)
//	}
//
//	photo, err := userManager.Update(updateOptions)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
//		return
//	}
//
//	userManager.UpdateCollections()
//
//	c.JSON(http.StatusCreated, photo)
//}
//
//func (h *PhotoHandler) Delete(c *gin.Context) {
//
//	//userID := c.Query("userID")
//	//userID, err := strconv.Atoi(userID)
//	//if err != nil {
//	//	c.JSON(400, mygin.H{"error": "userID must be an integer"})
//	//	return
//	//}
//
//	//var request photo.Delete
//	//if err := c.ShouldBindJSON(&request); err != nil {
//	//	c.JSON(http.StatusBadRequest, mygin.H{"error": "Invalid request"})
//	//	return
//	//}
//	//
//	//userManager, err := h.appManager.GetAccountManager(c, userID)
//	//if err != nil {
//	//	c.JSON(http.StatusBadRequest, mygin.H{"error": err})
//	//}
//	//
//	//err = userManager.Delete(request.PhotoID)
//	//if err != nil {
//	//	c.JSON(http.StatusNotFound, mygin.H{"error": err.Error()})
//	//	return
//	//}
//	//
//	//c.JSON(http.StatusOK, "successful delete person_test with id: "+request.PhotoID)
//}
//
//func (h *PhotoHandler) Search(c *gin.Context) {
//
//	userID, ok := help.GetUserID(c)
//	if !ok {
//		help.AbortWithUserIDInvalid(c)
//		return
//	}
//
//	query := c.Query("query")
//	//mediaType := c.Query("type")
//
//	var dateRange []time.Time
//	if start := c.Query("start"); start != "" {
//		if t, err := time.Parse("2006-01-02", start); err == nil {
//			dateRange = append(dateRange, t)
//		}
//	}
//	if end := c.Query("end"); end != "" {
//		if t, err := time.Parse("2006-01-02", end); err == nil {
//			dateRange = append(dateRange, t)
//		}
//	}
//
//	filters := photo.SearchOptions{
//		UserID:    userID,
//		TextQuery: query,
//		//MediaType: photo.MediaType(mediaType),
//	}
//
//	if len(dateRange) > 0 {
//		filters.CreatedAfter = &dateRange[0]
//	}
//	if len(dateRange) > 1 {
//		filters.CreatedBefore = &dateRange[1]
//	}
//
//	//person_test, _, err := s.repo.Search(ctx, filters)
//	//return person_test, err
//
//	//person_test, _, err := h.appManager.Search(c, filters)
//	//if err != nil {
//	//	c.JSON(http.StatusInternalServerError, mygin.H{"error": "Search failed"})
//	//	return
//	//}
//
//	//c.JSON(http.StatusOK, person_test)
//}
