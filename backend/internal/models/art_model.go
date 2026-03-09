package internal

type ArtModel struct {
    Id int
    Title string
    URL string
    Description string
}

func CreateArtModel(id: int, title: string, url: string, description: string) *ArtModel {
    return &ArtModel{
        Id: id
        Title: title
        URL: url
        Description: description
    }
}

func (m *ArtModel) EditRecord(title: string, url: string, description: string) {
	//m.Title: title
	//m.URL: url
	//m.Description: description
}
