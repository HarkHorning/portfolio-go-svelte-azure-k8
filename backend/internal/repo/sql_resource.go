package repo

import (
	"fmt"
	"log"

	"github.com/HarkHorning/portfolio-go-svelte-azure-k8/internal/models"
	_ "github.com/HarkHorning/portfolio-go-svelte-azure-k8/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (repo *Repo) TopTiles() ([]models.ArtModel, error) {
	var artTiles []models.ArtModel
	query := "SELECT * FROM art_tiles"

	err := repo.db.Select(&artTiles, query)
	if err != nil {
		return nil, fmt.Errorf("Could not list art tiles: %w", err)
	}

	return artTiles, err
}

func (repo *Repo) ListMoreTiles() ([]models.ArtModel, error) {
	log.Println("Function ListMoreTiles is not finished yet.")
	return nil, nil
}

func (repo *Repo) GetHigherQualityImage(artId int) {
	log.Println("Function GetHigherQualityImage is not finished yet. Will need to return image file.")
}

func (repo *Repo) NewModel(id int, title string, url string, description string, portrait bool) (string, error) {
	log.Println("Function NewModel is not finished yet.")
	return "", nil
}

func (repo *Repo) CreateModelTable() (string, error) {
	//query := "CREATE TABLE art_tiles (" +
	//	"id SERIAL PRIMARY KEY," +
	//	")"
	log.Println("Function CreateModelTable is not finished yet.")
	return "", nil
}
