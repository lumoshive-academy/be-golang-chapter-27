package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type AllRepository struct {
	Repobook     BookRepository
	RepoCustomer CustomerRepository
	OrderRepo    OrderRepository
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		Repobook:     NewBookRepository(db),
		RepoCustomer: NewCustomerRepository(db),
		OrderRepo:    NewOrderRepository(db, log),
	}
}
