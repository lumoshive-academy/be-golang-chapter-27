package repository

import (
	"be-golang-chapter-27-implem/model"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type OrderRepository struct {
	DB  *sql.DB
	Log *zap.Logger
}

func NewOrderRepository(db *sql.DB, log *zap.Logger) OrderRepository {
	return OrderRepository{
		DB:  db,
		Log: log,
	}
}

func (orderRepo *OrderRepository) Create(order *model.Order) error {
	query := "INSERT INTO public.orders(customer_id, street, city, postal_code, status, payment_method, total_amount, discount_amount) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := orderRepo.DB.QueryRow(query, order.Customer.ID, order.Address.Street, order.Address.City, order.Address.PostalCode, order.Status, order.PaymentMethod, order.TotalAmount, order.DiscountAmount).Scan(&order.ID)
	if err != nil {
		orderRepo.Log.Error("failed to store data order", zap.Error(err))
		return fmt.Errorf("failed to store data order: %w", err)
	}
	return nil
}

func (orderRepo *OrderRepository) GetAll(order model.Order) error {
	return nil
}

func (orderRepo *OrderRepository) GetByID(id int) error {
	return nil
}

func (orderRepo *OrderRepository) Update(id int, order model.Order) error {
	return nil
}

func (orderRepo *OrderRepository) Delete(id int) error {
	return nil
}

func (orderRepo *OrderRepository) OrderItem(order *model.Order) error {
	query := "INSERT INTO public.orders(customer_id, street, city, postal_code, status, payment_method, total_amount, discount_amount) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := orderRepo.DB.QueryRow(query, order.Customer.ID, order.Address.Street, order.Address.City, order.Address.PostalCode, order.Status, order.PaymentMethod, order.TotalAmount, order.DiscountAmount).Scan(&order.ID)
	if err != nil {
		return fmt.Errorf("failed to store data order: %w", err)
	}
	return nil
}
