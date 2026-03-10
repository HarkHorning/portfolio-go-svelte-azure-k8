package models

type ArtModel struct {
	Id          int
	Title       string
	URL         string
	Description string
	Portrait    bool
}

func CreateArtModel(id int, title string, url string, description string, portrait bool) *ArtModel {
	return &ArtModel{
		Id:          id,
		Title:       title,
		URL:         url,
		Description: description,
		Portrait:    portrait,
	}
}
