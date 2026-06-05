package repository

import (
	"database/sql"
	"goapi/vinyl-store/internal/domain"
	"log"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (repo *AlbumRepository) GetAll() ([]domain.Album, error) {
	rows, err := repo.db.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	var albums []domain.Album
	for rows.Next() {
		var album domain.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			log.Println("Erro ao ler row do banco de dados")
			log.Println(err)
			return nil, err
		}

		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		log.Println("Erro ao iterar sobre o banco de dados")
		return nil, err
	}

	return albums, nil
}

func (repo *AlbumRepository) InsertAlbum(newAlbum *domain.Album) (*domain.Album, error) {
	query := "INSERT INTO albums (title, artist, price) VALUES (?, ?, ?) RETURNING id"
	err := repo.db.QueryRow(query, newAlbum.Title, newAlbum.Artist, newAlbum.Price).Scan(&newAlbum.ID)
	if err != nil {
		log.Println("Erro ao inserir album no banco de dados")
		log.Println(err)
		return nil, err
	}

	return newAlbum, nil
}

func (repo *AlbumRepository) DeleteAlbum(albumId int64) error {
	_, err := repo.db.Exec("DELETE FROM albums WHERE id = ?", albumId)
	if err != nil {
		log.Println("Erro ao deletar album no banco de dados")
		log.Println(err)
		return err
	}
	return nil
}

func (repo *AlbumRepository) UdateAlbum(newAlbum *domain.Album) (*domain.Album, error) {
	query := "UPDATE albums SET title = ?, artist = ?, price = ? WHERE id = ?"
	_, err := repo.db.Exec(query, newAlbum.Title, newAlbum.Artist, newAlbum.Price, newAlbum.ID)
	if err != nil {
		log.Println("Erro ao atualizar o album no banco de dados")
		log.Println(err)
		return nil, err
	}
	return newAlbum, nil
}
