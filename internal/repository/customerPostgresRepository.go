package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// GetAllCustomer in postgres db...

func (r *PostgresRepository) GetAllCustomer() ([]*model.Customer, error) {
	res, err := r.DB.Query(context.Background(), "SELECT * FROM Customers")
	if err != nil {
		return nil, err
	}
	allCustomers := []*model.Customer{}
	for res.Next() {
		u := &model.Customer{}
		err := res.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.Married)
		if err != nil {
			return nil, err
		}
		allCustomers = append(allCustomers, u)
	}
	res.Close()
	return allCustomers, nil
}

// GetByIdCustomer in postgres db...

func (r *PostgresRepository) GetByIDCustomer(ID uuid.UUID) (*model.Customer, error) {
	res, err := r.DB.Query(context.Background(), "SELECT * FROM Customers WHERE id=$1", ID)
	if err != nil {
		panic(err)
	}
	u := &model.Customer{}
	for res.Next() {
		err = res.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.Married)
		if err != nil {
			panic(err)
		}
	}
	res.Close()
	return u, nil
}

// DeleteByIdCustomer in postgres db...

func (r *PostgresRepository) DeleteByIDCustomer(ID uuid.UUID) (string, error) {
	_, err := r.DB.Exec(context.Background(), "DELETE FROM Customers WHERE id = $1", ID)
	if err != nil {
		panic(err)
	}
	const res = "Customer deleted"
	return res, nil
}

// UpdateByIdCustomer in postgres db...

func (r *PostgresRepository) UpdateByIDCustomer(ID uuid.UUID, u *model.Customer) (*model.Customer, error) {
	_, err := r.DB.Exec(context.Background(), "UPDATE Customers SET name = $1, email = $2, age = $3, married = $4 WHERE id = $5",
		u.Name, u.Email, u.Age, u.Married, ID)
	if err != nil {
		return nil, err
	}
	u.ID = ID
	return u, nil
}

// CreateCustomer in postgres db...

func (r *PostgresRepository) CreateCustomer(u *model.Customer) (*model.Customer, error) {
	uuid := uuid.New()
	u.ID = uuid
	res, err := r.DB.Query(context.Background(), "INSERT INTO Customers (id,name, email, age, married) VALUES ($1,$2,$3,$4,$5) ", u.ID, u.Name, u.Email, u.Age, u.Married)
	if err != nil {
		panic(err)
	}
	res.Close()
	return u, nil
}
