package models

type Post struct {
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Summary   string `json:"summary"`
	ReadTime  string `json:"read_time"`
	Tags      string `json:"tags"`
	Category  string `json:"category"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
	Published string `json:"published"`
	Featured  string `json:"featured"`
}
