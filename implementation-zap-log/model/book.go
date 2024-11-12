package model

import "database/sql"

type Book struct {
	ID       uint           `json:"id"`
	Category Category       `json:"category"`
	Title    string         `json:"title"`
	Author   string         `json:"author"`
	Price    float32        `json:"price"`
	Discount float32        `json:"dicount"`
	CoverURL sql.NullString `json:"cover_url"`
	FileBook sql.NullString `json:"file_book"`
	TimeStamp
}
