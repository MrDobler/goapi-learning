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
		log.Fatalln(err.Error())
	}
	defer rows.Close()

	var albums []domain.Album
	for rows.Next() {
		var album domain.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			log.Fatalln("Erro ao ler row do banco de dados")
			log.Fatalln(err)
			return nil, err
		}

		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		log.Fatalln("Erro ao iterar sobre o banco de dados")
		return nil, err
	}

	return albums, nil
}

func (repo *AlbumRepository) InsertAlbum(newAlbum *domain.Album) (*domain.Album, error) {
	_, err := repo.db.Exec("INSERT INTO albums (title, artist, price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		log.Fatalln("Erro ao inserir album no banco de dados")
		log.Fatalln(err)
		return nil, err
	}

	return newAlbum, nil
}
