package handlers

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type HttpResponse struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func AddAlbum(context *gin.Context) {
	var newAlbum Album

	if err := context.BindJSON(&newAlbum); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func RemoveAlbum(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.IndentedJSON(http.StatusBadRequest, HttpResponse{Message: "Bad Request", Status: 400})
	}

	var originalAlbums = len(albums)

	albums = slices.DeleteFunc(albums, func(album Album) bool {
		return album.ID == id
	})

	if originalAlbums != len(albums) {
		context.IndentedJSON(http.StatusOK, HttpResponse{Message: "Success", Status: 200})
	} else {
		context.JSON(http.StatusBadRequest, HttpResponse{Message: "Not Found", Status: 404})
	}
}

func EditAlbum(context *gin.Context) {
	var albumToEdit Album

	if err := context.BindJSON(&albumToEdit); err != nil {
		context.IndentedJSON(http.StatusBadRequest, HttpResponse{Message: "Bad Request", Status: 400})
		return
	}

	for i := range albums {
		if albums[i].ID == albumToEdit.ID {
			albums[i] = albumToEdit
			context.IndentedJSON(http.StatusOK, albumToEdit)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, HttpResponse{Message: "Not Found", Status: 404})
}

func GetAlbumByID(context *gin.Context) {
	id := context.Param("id")

	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, HttpResponse{Message: "Not Found", Status: 404})
}
