package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// GET ALL
func (r *PostgresRepository) GetAllCustomer() ([]*model.Customer, error) {
	res, err := r.DB.Query(context.Background(), "SELECT * FROM Customers")
	if err != nil {
		return nil, err
	}
	allCustomers := []*model.Customer{}
	for res.Next() {
		u := &model.Customer{}
		err := res.Scan(&u.Id, &u.Name, &u.Email, &u.Age, &u.Married)
		if err != nil {
			return nil, err
		}
		allCustomers = append(allCustomers, u)
	}
	res.Close()
	return allCustomers, nil
}

// GET BY ID
func (r *PostgresRepository) GetByIdCustomer(id uuid.UUID) (*model.Customer, error) {
	res, err := r.DB.Query(context.Background(), "SELECT * FROM Customers WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	u := &model.Customer{}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Name, &u.Email, &u.Age, &u.Married)
		if err != nil {
			panic(err)
		}
	}
	res.Close()
	return u, nil
}

// DELETE BY ID
func (r *PostgresRepository) DeleteByIdCustomer(id uuid.UUID) (string, error) {
	_, err := r.DB.Exec(context.Background(), "DELETE FROM Customers WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	res := "Customer deleted"
	return res, nil
}

// UPDATE BY ID
func (r *PostgresRepository) UpdateByIdCustomer(id uuid.UUID, u *model.Customer) (*model.Customer, error) {
	_, err := r.DB.Exec(context.Background(), "UPDATE Customers SET name = $1, email = $2, age = $3, married = $4 WHERE id = $5",
		u.Name, u.Email, u.Age, u.Married, id)
	if err != nil {
		return nil, err
	}
	u.Id = id
	return u, nil
}

// CREATE
func (r *PostgresRepository) CreateCustomer(u *model.Customer) (*model.Customer, error) {
	uuid := uuid.New()
	u.Id = uuid
	res, err := r.DB.Query(context.Background(), "INSERT INTO Customers (id,name, email, age, married) VALUES ($1,$2,$3,$4,$5) ", u.Id, u.Name, u.Email, u.Age, u.Married)
	if err != nil {
		panic(err)
	}
	res.Close()
	return u, nil
}
