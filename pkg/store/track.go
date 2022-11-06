package store

type Track struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}
