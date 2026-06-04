package api

import (
	"fmt"
	"goapi/vinyl-store/internal/handlers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	handler := handlers.NewAlbumHandler()
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.AddAlbum)
	router.PUT("/albums", handler.EditAlbum)
	router.DELETE("/albums/:id", handler.RemoveAlbum)
	router.GET("/albums/:id", handler.GetAlbumByID)

	return router
}

func StartServer() {
	router := setupRouter()
	fmt.Println("Server Starting on Localhost:8000")

	if err := router.Run("localhost:8000"); err != nil {
		fmt.Println(err)
	}
}
