package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/mygin"

	"github.com/mahdi-cpp/photos-api/internal/application"
	"github.com/mahdi-cpp/photos-api/internal/help"
)

type AssetHandler struct {
	appManager *application.AppManager
}

func NewAssetHandler(manager *application.AppManager) *AssetHandler {
	return &AssetHandler{appManager: manager}
}

func (h *AssetHandler) Create(c *mygin.Context) {
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

	assetID := c.Param("id")
	id, err := uuid.Parse(assetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	item, err := accountManager.Read(id)
	if err != nil {
		c.JSON(http.StatusNotFound, mygin.H{"error": "Asset not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *AssetHandler) Read(c *mygin.Context) {

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

	assetID := c.Param("id")
	id, err := uuid.Parse(assetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	item, err := accountManager.Read(id)
	if err != nil {
		c.JSON(http.StatusNotFound, mygin.H{"error": "Asset not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

//func (h *AssetHandler) ReadAll(c *gin.Context) {
//
//	userID, ok := help.GetUserID(c)
//	if !ok {
//		help.AbortWithUserIDInvalid(c)
//		return
//	}
//
//	fmt.Println("ReadAll userID: ", userID)
//
//	var with *asset.SearchOptions
//	if err := c.ShouldBindJSON(&with); err != nil {
//		help.AbortWithRequestInvalid(c)
//		return
//	}
//
//	userManager, err := h.appManager.GetAccountManager(c, userID)
//	if err != nil {
//		fmt.Println(err.Error())
//		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
//		return
//	}
//
//	items, err := userManager.ReadAll(with)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, mygin.H{"error": "failed account Read"})
//		return
//	}
//
//	fmt.Println("ReadAll count", len(items))
//
//	//result := asset.PHFetchResult[*asset.Assets]{
//	//	Items:  items,
//	//	Total:  total,
//	//	Size:  100,
//	//	Page: 100,
//	//}
//	//c.JSON(http.StatusOK, items)
//	c.JSON(http.StatusOK, mygin.H{"data": items})
//}
//
//func (h *AssetHandler) Update(c *gin.Context) {
//
//	userID, ok := help.GetUserID(c)
//	if !ok {
//		help.AbortWithUserIDInvalid(c)
//		return
//	}
//
//	var updateOptions asset.UpdateOptions
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
//	asset, err := userManager.UpdateAssets(updateOptions)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
//		return
//	}
//
//	userManager.UpdateCollections()
//
//	c.JSON(http.StatusCreated, asset)
//}
//
//func (h *AssetHandler) BuckUpdate(c *gin.Context) {
//	userID, ok := help.GetUserID(c)
//	if !ok {
//		help.AbortWithUserIDInvalid(c)
//		return
//	}
//
//	var updateOptions asset.UpdateOptions
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
//	for _, asset := range allAssets {
//		updateOptions.AssetIds = append(updateOptions.AssetIds, asset.ID)
//	}
//
//	asset, err := userManager.UpdateAssets(updateOptions)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
//		return
//	}
//
//	userManager.UpdateCollections()
//
//	c.JSON(http.StatusCreated, asset)
//}
//
//func (h *AssetHandler) Delete(c *gin.Context) {
//
//	//userID := c.Query("userID")
//	//userID, err := strconv.Atoi(userID)
//	//if err != nil {
//	//	c.JSON(400, mygin.H{"error": "userID must be an integer"})
//	//	return
//	//}
//
//	//var request asset.Delete
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
//	//err = userManager.DeleteAsset(request.AssetID)
//	//if err != nil {
//	//	c.JSON(http.StatusNotFound, mygin.H{"error": err.Error()})
//	//	return
//	//}
//	//
//	//c.JSON(http.StatusOK, "successful delete person_test with id: "+request.AssetID)
//}
//
//func (h *AssetHandler) Search(c *gin.Context) {
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
//	filters := asset.SearchOptions{
//		UserID:    userID,
//		TextQuery: query,
//		//MediaType: asset.MediaType(mediaType),
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
