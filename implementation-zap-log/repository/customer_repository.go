package repository

import (
	"be-golang-chapter-27/implementation-zap-log/model"
	"database/sql"
	"fmt"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return CustomerRepository{
		DB: db,
	}
}

func (c *CustomerRepository) Create(customer *model.Customer) error {
	query := "INSERT INTO customers (name, phone) VALUES ($1, $2) RETURNING id"
	err := c.DB.QueryRow(query, customer.Name, customer.Phone).Scan(&customer.ID)
	if err != nil {
		return fmt.Errorf("failed to store data customer: %w", err)
	}
	return nil
}

func (c *CustomerRepository) GetAll(customer *[]model.Customer) error {
	return nil
}

func (c *CustomerRepository) GetByID(id int) error {
	return nil
}

func (c *CustomerRepository) GetByPhone(phone string) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM customers WHERE phone=$1"
	err := c.DB.QueryRow(query, phone).Scan(&count)
	return count, err
}

func (c *CustomerRepository) Update(id int, customer model.Customer) error {
	return nil
}

func (c *CustomerRepository) Delete(id int) error {
	return nil
}
