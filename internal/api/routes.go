package api

import (
	"database/sql"
	"fmt"
	"goapi/vinyl-store/internal/handlers"
	"goapi/vinyl-store/internal/repository"

	"github.com/gin-gonic/gin"
)

func setupRouter(db *sql.DB) *gin.Engine {
	albumRepo := repository.NewAlbumRepository(db)
	handler := handlers.NewAlbumHandler(albumRepo)
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.AddAlbum)
	router.PUT("/albums", handler.EditAlbum)
	router.DELETE("/albums/:id", handler.RemoveAlbum)
	router.GET("/albums/:id", handler.GetAlbumByID)

	return router
}

func StartServer(db *sql.DB) {
	router := setupRouter(db)
	fmt.Println("Server Starting on Localhost:8000")

	if err := router.Run("localhost:8000"); err != nil {
		fmt.Println(err)
	}
}
