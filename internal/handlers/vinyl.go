package handlers

import (
	"goapi/vinyl-store/internal/domain"
	"goapi/vinyl-store/internal/mock"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

type AlbumHandler struct {
	albums []domain.Album
}

func NewAlbumHandler() *AlbumHandler {
	return &AlbumHandler{
		albums: mock.Albums,
	}
}

func (h *AlbumHandler) GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, h.albums)
}

func (h *AlbumHandler) AddAlbum(context *gin.Context) {
	var newAlbum domain.Album

	if err := context.BindJSON(&newAlbum); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	h.albums = append(h.albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func (h *AlbumHandler) RemoveAlbum(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.IndentedJSON(http.StatusBadRequest, HttpResponse{Message: "Bad Request", Status: 400})
	}

	var originalAlbums = len(h.albums)

	h.albums = slices.DeleteFunc(h.albums, func(album domain.Album) bool {
		return album.ID == id
	})

	if originalAlbums != len(h.albums) {
		context.IndentedJSON(http.StatusOK, HttpResponse{Message: "Success", Status: 200})
	} else {
		context.JSON(http.StatusBadRequest, HttpResponse{Message: "Not Found", Status: 404})
	}
}

func (h *AlbumHandler) EditAlbum(context *gin.Context) {
	var albumToEdit domain.Album

	if err := context.BindJSON(&albumToEdit); err != nil {
		context.IndentedJSON(http.StatusBadRequest, HttpResponse{Message: "Bad Request", Status: 400})
		return
	}

	for i := range h.albums {
		if h.albums[i].ID == albumToEdit.ID {
			h.albums[i] = albumToEdit
			context.IndentedJSON(http.StatusOK, albumToEdit)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, HttpResponse{Message: "Not Found", Status: 404})
}

func (h *AlbumHandler) GetAlbumByID(context *gin.Context) {
	id := context.Param("id")

	for _, a := range h.albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, HttpResponse{Message: "Not Found", Status: 404})
}
