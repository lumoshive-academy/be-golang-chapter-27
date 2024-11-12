package repository

import (
	"be-golang-chapter-27/implementation-zap-log/model"
	"database/sql"
	"fmt"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return BookRepository{
		DB: db,
	}
}

func (b *BookRepository) Create(book *model.Book) error {
	query := "INSERT INTO books (category_id, title, author, price, discount, cover_url, file_book) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id"
	err := b.DB.QueryRow(query, book.Category.ID, book.Title, book.Author, book.Price, book.Discount, book.CoverURL, book.FileBook).Scan(&book.ID)
	return err
}

func (b *BookRepository) GetAll(book model.Book) error {
	return nil
}

func (b *BookRepository) GetByID(id int) (*model.Book, error) {
	var book model.Book
	query := "SELECT categories.name as category, title, author, price, discount, cover_url, file_book FROM books INNER JOIN categories ON books.category_id = categories.id WHERE books.id=$1"
	err := b.DB.QueryRow(query, id).Scan(&book.Category.Name, &book.Title, &book.Author, &book.Price, &book.Discount, &book.CoverURL, &book.FileBook)
	if err != nil {
		return nil, fmt.Errorf("failed to get detail book: %w", err)
	}
	return &book, nil
}

func (b *BookRepository) Update(id int, book model.Book) error {
	return nil
}

func (b *BookRepository) Delete(id int) error {
	return nil
}
