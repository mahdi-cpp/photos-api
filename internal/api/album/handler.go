package album_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/application"
	"github.com/mahdi-cpp/photos-api/internal/collections/album"
	"github.com/mahdi-cpp/photos-api/internal/collections/photo"
	"github.com/mahdi-cpp/photos-api/internal/help"
	"github.com/mahdi-cpp/photos-api/mygin"
)

type AlbumHandler struct {
	appManager *application.AppManager
}

func New(manager *application.AppManager) *AlbumHandler {
	return &AlbumHandler{appManager: manager}
}

func (h *AlbumHandler) Create(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.SendError(c, "user id invalid", http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var request *album.Album
	err = json.NewDecoder(c.Req.Body).Decode(&request)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	create, err := accountManager.AlbumsManager.Create(request)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("album is created:", request.Title)

	c.JSON(http.StatusOK, create)
}

func (h *AlbumHandler) Read(c *mygin.Context) {

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

	albumID := c.Param("id")
	id, err := uuid.Parse(albumID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	item, err := accountManager.AlbumsManager.Read(id)
	if err != nil {
		c.JSON(http.StatusNotFound, mygin.H{"error": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *AlbumHandler) ReadAll(c *mygin.Context) {

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
		help.SendError(c, err.Error(), http.StatusBadRequest)
	}
	size, err := c.GetQueryInt("size")
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
	}

	fmt.Println("page:", page)
	fmt.Println("size:", size)

	with := &album.SearchOptions{
		Page: page,
		Size: size,
	}

	fmt.Println("3")

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	items, err := accountManager.AlbumsManager.ReadAll(with)
	if err != nil {
		c.JSON(http.StatusInternalServerError, mygin.H{"error": "failed account Read"})
		return
	}

	fmt.Println("ReadAll count", len(items))

	c.JSON(http.StatusOK, items)
}

func (h *AlbumHandler) ReadCollections(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		fmt.Println("user id invalid")
		return
	}

	//page, err := c.GetQueryInt("page")
	//if err != nil {
	//	help.SendError(c, err.Error(), http.StatusBadRequest)
	//}
	//size, err := c.GetQueryInt("size")
	//if err != nil {
	//	help.SendError(c, err.Error(), http.StatusBadRequest)
	//}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	with := &album.SearchOptions{
		Sort:      "id",
		SortOrder: "desc",
		Page:      1,
		Size:      20,
	}
	items := accountManager.AlbumsManager.ReadCollections(with)

	fmt.Println("Read Album Collections count", len(items))
	c.JSON(http.StatusOK, items)
}

func (h *AlbumHandler) AddPhotos(c *mygin.Context) {

	fmt.Println("AddPhoto")

	userID, ok := help.GetUserID(c)
	if !ok {
		help.SendError(c, "user id invalid", http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Println("account error:", err)
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var request *photo.CollectionPhoto
	err = json.NewDecoder(c.Req.Body).Decode(&request)
	if err != nil {
		fmt.Println("Decode error:", err)
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = accountManager.AlbumsManager.IsExist(request.ParentID)
	if err != nil {
		fmt.Println("Album isExist error:", err)
		help.SendError(c, err.Error(), http.StatusBadRequest)
	}

	for _, photoID := range request.PhotoIDs {
		err := accountManager.AlbumsManager.AddPhoto(request.ParentID, photoID)
		if err != nil {
			fmt.Println("Album addPhoto error:", err)
			help.SendError(c, err.Error(), http.StatusBadRequest)
			return
		}
	}

	fmt.Println("photos add to album with id: ", request.ParentID.String())

	c.JSON(http.StatusOK, "create")
}

func (h *AlbumHandler) Delete(c *mygin.Context) {

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

	albumID := c.Param("id")
	id, err := uuid.Parse(albumID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	err = accountManager.AlbumsManager.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, mygin.H{"error": "delete album " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, "")
}
