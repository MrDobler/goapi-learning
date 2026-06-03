package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"goapi/vinyl-store/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.AddAlbum)
	router.PUT("/albums", handlers.EditAlbum)
	router.DELETE("/albums/:id", handlers.RemoveAlbum)

	fmt.Println("Server Starting")

	err := router.Run("localhost:8000")
	if err != nil {
		fmt.Println(err)
	}
}
