package api

import (
	"fmt"
	"goapi/vinyl-store/internal/handlers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.AddAlbum)
	router.PUT("/albums", handlers.EditAlbum)
	router.DELETE("/albums/:id", handlers.RemoveAlbum)

	return router
}

func StartServer() {
	router := setupRouter()
	fmt.Println("Server Starting on Localhost:8000")

	if err := router.Run("localhost:8000"); err != nil {
		fmt.Println(err)
	}
}
