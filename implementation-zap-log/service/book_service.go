package service

import (
	"be-golang-chapter-27/implementation-zap-log/model"
	"be-golang-chapter-27/implementation-zap-log/repository"
	"fmt"
)

type BookService struct {
	Repo repository.AllRepository
}

func NewBookService(repo repository.AllRepository) BookService {
	return BookService{
		Repo: repo,
	}
}

func (bookService *BookService) Create(book *model.Book) error {
	err := bookService.Repo.Repobook.Create(book)
	if err != nil {
		return fmt.Errorf("failed to store data book : %w", err)
	}
	return nil
}
