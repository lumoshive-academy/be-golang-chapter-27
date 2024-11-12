//go:build wireinject
// +build wireinject

package wire

import (
	"be-golang-chapter-27-implem/database"
	"be-golang-chapter-27-implem/handler"
	"be-golang-chapter-27-implem/repository"
	"be-golang-chapter-27-implem/service"

	"github.com/google/wire"
)

var bookHandler = wire.NewSet(
	database.InitDB,
	log.InitLog
	repository.NewAllRepository,
	service.NewBookService,
	handler.NewBookHandler,
)

func InitHandlerBook() (handler.BookHandler, error) {
	wire.Build(bookHandler)
	return handler.BookHandler{}, nil
}


var middlewareSet = wire.NewSet(
	log.InitLog
	middleware.NewMiddleware
)


func InitMiddleware() middleware.Middleware {
	wire.Build(middlewareSet)
	return middleware.Middleware{}
}
