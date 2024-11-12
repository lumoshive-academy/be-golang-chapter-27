package router

import (
	"be-golang-chapter-27/implementation-zap-log/database"
	"be-golang-chapter-27/implementation-zap-log/handler"
	"be-golang-chapter-27/implementation-zap-log/library"
	"be-golang-chapter-27/implementation-zap-log/middleware"
	"be-golang-chapter-27/implementation-zap-log/repository"
	"be-golang-chapter-27/implementation-zap-log/service"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func RouteInit() (*sql.DB, *chi.Mux, *zap.Logger) {
	r := chi.NewRouter()

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("error :", err)
	}

	logger := library.InitLog()
	allRepository := repository.NewAllRepository(db, logger)
	bookService := service.NewBookService(allRepository)
	handlerBookHandler := handler.NewBookHandler(bookService, logger)
	middleware := middleware.NewMiddleware(logger)

	// db, err := database.InitDB()
	// if err != nil {
	// 	panic(err)
	// }
	// // defer db.Close()
	// repo := repository.NewBookRepository(db)
	// service := service.NewBookService(repo)
	// handler := handler.NewBookHandler(service)

	fs := http.FileServer(http.Dir("assests/cover_books"))
	r.Handle("/cover_books/*", http.StripPrefix("/cover_books/", fs))

	r.With(middleware.MinddlewareLogger).Post("/store_book", handlerBookHandler.CreateBook)

	return db, r, logger
}
