package handler

import (
	"be-golang-chapter-27/implementation-zap-log/model"
	"be-golang-chapter-27/implementation-zap-log/service"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type BookHandler struct {
	BookService service.BookService
	Log         *zap.Logger
}

func NewBookHandler(bookService service.BookService, log *zap.Logger) BookHandler {
	return BookHandler{
		BookService: bookService,
		Log:         log,
	}
}

func (bookHandler *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var title, author, conver_url, file_book, category_id, price, discount string

	//set domain
	domain := strings.Join([]string{"http://", r.Host}, "")

	// get data from form
	title = r.FormValue("title")
	author = r.FormValue("author")
	category_id = r.FormValue("category_id")
	price = r.FormValue("price")
	discount = r.FormValue("discount")

	// convertion type data
	category_id_int, _ := strconv.Atoi(category_id)
	price_float, _ := strconv.ParseFloat(price, 32)
	discount_float, _ := strconv.ParseFloat(discount, 32)

	file, data, err := r.FormFile("conver_url")
	if err != nil {
		bookHandler.Log.Error("bookhandler_createBook_formfile", zap.Error(err))
		fmt.Println("Error : ", err)
	}
	defer file.Close()

	// store file
	dst, err := os.Create(filepath.Join("assests/cover_books", data.Filename))
	if err != nil {
		fmt.Println("Error : ", err)
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	// set url
	conver_url = strings.Join([]string{domain, "/cover_books/", data.Filename}, "")
	file_book = strings.Join([]string{domain, "/cover_books/", data.Filename}, "")

	// ini data
	book := model.Book{
		Category: model.Category{
			ID: uint(category_id_int),
		},
		Title:    title,
		Author:   author,
		Price:    float32(price_float),
		Discount: float32(discount_float),
		CoverURL: sql.NullString{String: conver_url, Valid: true},
		FileBook: sql.NullString{String: file_book, Valid: true},
	}

	// fmt.Println(book)

	// store data to database
	err = bookHandler.BookService.Create(&book)
	if err != nil {
		bookHandler.Log.Error("failed to store data book", zap.Error(err))
		// fmt.Println("Error : ", err)
	}

	// response json
}
