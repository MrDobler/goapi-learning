package handlers

import (
	"goapi/vinyl-store/internal/domain"
	"goapi/vinyl-store/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

type AlbumHandler struct {
	albums     []domain.Album
	repository *repository.AlbumRepository
}

func NewAlbumHandler(repo *repository.AlbumRepository) *AlbumHandler {
	return &AlbumHandler{
		albums:     []domain.Album{},
		repository: repo,
	}
}

func (h *AlbumHandler) GetAlbums(context *gin.Context) {
	albums, err := h.repository.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, albums)
}

func (h *AlbumHandler) AddAlbum(context *gin.Context) {
	var newAlbum domain.Album

	if err := context.BindJSON(&newAlbum); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	var _, err = h.repository.InsertAlbum(&newAlbum)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func (h *AlbumHandler) RemoveAlbum(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.IndentedJSON(http.StatusBadRequest, HttpResponse{Message: "Bad Request", Status: 400})
	}

	albumId, conversionError := strconv.Atoi(id)
	if conversionError != nil {
		log.Println(conversionError.Error())
		return
	}

	if err := h.repository.DeleteAlbum(int64(albumId)); err != nil {
		context.JSON(http.StatusBadRequest, HttpResponse{Message: err.Error(), Status: 404})
	} else {
		context.IndentedJSON(http.StatusOK, HttpResponse{Message: "Success", Status: 200})
	}
}

func (h *AlbumHandler) EditAlbum(context *gin.Context) {
	var albumToEdit domain.Album

	if err := context.BindJSON(&albumToEdit); err != nil {
		context.IndentedJSON(http.StatusBadRequest, HttpResponse{Message: err.Error(), Status: 400})
		return
	}

	updatedAlbum, err := h.repository.UdateAlbum(&albumToEdit)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, HttpResponse{Message: err.Error(), Status: 500})
	}

	context.IndentedJSON(http.StatusOK, updatedAlbum)
}

func (h *AlbumHandler) GetAlbumByID(context *gin.Context) {
	id := context.Param("id")
	albumId, conversionError := strconv.Atoi(id)
	if conversionError != nil {
		log.Println(conversionError.Error())
		return
	}

	for _, a := range h.albums {
		if a.ID == int64(albumId) {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, HttpResponse{Message: "Not Found", Status: 404})
}
