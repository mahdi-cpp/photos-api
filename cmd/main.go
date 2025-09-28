package main

import (
	"fmt"
	"log"
	"net/http"

	album_handler "github.com/mahdi-cpp/photos-api/internal/api/album"
	"github.com/mahdi-cpp/photos-api/internal/api/photo"
	"github.com/mahdi-cpp/photos-api/internal/application"
	"github.com/mahdi-cpp/photos-api/mygin"
)

func main() {

	appManager, err := application.New()
	if err != nil {
		log.Fatal(err)
	}

	//if !utils.CheckVersionIsUpToDate(2) {
	//upgrade.Start(appManager.AccountManager)
	//upgrade_v3.Start(appManager.AccountManager)
	//}

	//ginInit()
	// Create a new engine with default middleware
	router := mygin.New()

	assetHandler := photo_handler.New(appManager)
	albumHandler := album_handler.New(appManager)

	assetRoute(assetHandler)

	router.POST("/api/photos", assetHandler.Create)
	//router.GET("/api/photos/photoId", assetHandler.Read)
	router.GET("/api/photos", assetHandler.ReadAll)

	router.POST("/api/albums/photos", albumHandler.AddPhotos)
	router.POST("/api/albums", albumHandler.Create)

	router.GET("/api/albums/", albumHandler.ReadAll)
	router.GET("/api/albums/collections", albumHandler.ReadCollections)

	// Start the server
	fmt.Println("Server is running on http://localhost:50151")
	log.Fatal(http.ListenAndServe(":50151", router))

	//albumHandler := handler.NewAlbumHandler(appManager)
	//RegisterAlbumRoutes(albumHandler)
	//
	//tripHandler := handler.NewTripHandler(appManager)
	//tripRoute(tripHandler)
	//
	//sharedAlbumHandler := handler.NewSharedAlbumHandler(appManager)
	//sharedAlbumRoute(sharedAlbumHandler)
	//
	//personHandler := handler.NewPersonsHandler(appManager)
	//personRoute(personHandler)
	//
	//villageHandler := handler.NewVillageHandler(appManager)
	//villageRoute(villageHandler)
	//
	//pinnedHandler := handler.NewPinnedHandler(appManager)
	//pinnedRoute(pinnedHandler)
	//
	//cameraHandler := handler.NewCameraHandler(appManager)
	//cameraRoute(cameraHandler)

	//startServer(router)
}

func assetRoute(h *photo_handler.PhotoHandler) {

	//api := router.Group("")

	//router.POST("/api/assets", h.Create)

	//router.GET("/api/assets/assetId", h.Read)
	//router.GET("/api/assets/", h.ReadAll)
	//
	//router.PATCH("/api/assets/assetId", h.Update)
	//router.PATCH("/api/assets/", h.BuckUpdate)
	//
	//router.DELETE("/api/assets", h.Delete)
}

//
//func RegisterAlbumRoutes(h *handler.AlbumHandler) {
//
//	api := router.Group("/api/v1/album")
//
//	api.POST("thumbnail", h.Create)
//	api.POST("update", h.Update)
//	api.POST("delete", h.Delete)
//	api.POST("list", h.ReadAll)
//	api.POST("search", h.GetBySearchOptions)
//}
//
//func pinnedRoute(h *handler.PinnedHandler) {
//
//	api := router.Group("/api/v1/pinned")
//
//	api.POST("thumbnail", h.Create)
//	api.POST("update", h.Update)
//	api.POST("delete", h.Delete)
//	api.POST("list", h.GetList)
//}
//
//func sharedAlbumRoute(h *handler.SharedAlbumHandler) {
//
//	api := router.Group("/api/v1/shared_album")
//
//	api.POST("thumbnail", h.Create)
//	api.POST("update", h.Update)
//	api.POST("delete", h.Delete)
//	api.POST("list", h.GetList)
//}
//
//func tripRoute(h *handler.TripHandler) {
//
//	api := router.Group("/api/v1/trip")
//
//	api.POST("thumbnail", h.Create)
//	api.POST("update", h.Update)
//	api.POST("delete", h.Delete)
//	api.POST("list", h.GetCollectionList)
//}
//
//func personRoute(h *handler.PersonHandler) {
//
//	api := router.Group("/api/v1/person")
//
//	api.POST("thumbnail", h.Create)
//	api.POST("update", h.Update)
//	api.POST("delete", h.Delete)
//	api.POST("list", h.GetCollectionList)
//}
//
//func cameraRoute(h *handler.CameraHandler) {
//
//	api := router.Group("/api/v1/camera")
//
//	api.POST("/list", h.GetList)
//}
//
//func villageRoute(h *handler.VillageHandler) {
//
//	api := router.Group("/api/v1/village")
//
//	api.POST("list", h.GetList)
//}
