package entity

type SearchResult struct {
	Id             int    `json:"id"`
	Url            string `json:"url"`
	Name           string `json:"name"`
	Content        string `json:"content"`
	RelationPatent string `json:"relationPatent"`
	Owner          string `json:"owner"`
	RegisterDate   string `json:"registerDate"`
	Type           string `json:"type"`
	Score          string `json:"score"`
	Score1         string `json:"score1"`
	Score2         string `json:"score2"`
	Score3         string `json:"score3"`
}
