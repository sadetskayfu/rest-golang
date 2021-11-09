package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

func (r *PostgresRepository) GetAllCat() ([]*model.Cat, error) {
	res, err := r.DB.Query(context.Background(),
		"SELECT * FROM cats")
	if err != nil {
		return nil, err
	}
	allCat := []*model.Cat{}
	for res.Next() {
		u := &model.Cat{}
		err = res.Scan(&u.Id, &u.Name, &u.Age, &u.Color)
		if err != nil {
			return nil, err
		}
		allCat = append(allCat, u)
	}
	res.Close()
	return allCat, nil
}

func (r *PostgresRepository) GetByIdCat(id uuid.UUID) (*model.Cat, error) {
	res, err := r.DB.Query(context.Background(),
		"SELECT * FROM cats WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	u := &model.Cat{}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Name, &u.Age, &u.Color)
		if err != nil {
			return nil, err
		}
	}
	res.Close()
	return u, nil
}

func (r *PostgresRepository) CreateCat(u *model.Cat) (*model.Cat, error) {
	u.Id = uuid.New()
	res, err := r.DB.Query(context.Background(),
		"INSERT INTO cats (id,name,age,color) VALUES ($1,$2,$3,$4)",
		u.Id, u.Name, u.Age, u.Color)
	if err != nil {
		return nil, err
	}
	res.Close()
	return u, nil
}

func (r *PostgresRepository) DeleteByIdCat(id uuid.UUID) (string, error) {
	_, err := r.DB.Exec(context.Background(),
		"DELETE FROM cats WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	res := "Cat deleted"
	return res, nil
}

func (r *PostgresRepository) UpdateByIdCat(id uuid.UUID, u *model.Cat) (*model.Cat, error) {
	_, err := r.DB.Exec(context.Background(),
		"UPDATE cats SET name=$1, age=$2, color=$3 WHERE id=$4", u.Name, u.Age, u.Color, id)
	if err != nil {
		panic(err)
	}
	u.Id = id
	return u, nil
}
