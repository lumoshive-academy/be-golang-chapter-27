package service

import (
	"be-golang-chapter-27/implementation-zap-log/model"
	"be-golang-chapter-27/implementation-zap-log/repository"
	"fmt"
)

type OrderService struct {
	Repo repository.AllRepository
}

func NewCustomerService(repo repository.AllRepository) OrderService {
	return OrderService{
		Repo: repo,
	}
}

func (OrderService *OrderService) Create(order *model.Order) error {
	// check account customer
	count, err := OrderService.Repo.RepoCustomer.GetByPhone(order.Customer.Phone)
	if err != nil {
		return fmt.Errorf("failed to get customer count: %w", err)
	}

	// create new customer
	if count < 0 {
		OrderService.Repo.RepoCustomer.Create(&order.Customer)
	}

	// store order
	err = OrderService.Repo.OrderRepo.Create(order)
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	// // strore item order
	// err :=

	return nil
}
