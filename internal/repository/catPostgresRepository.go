package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// GetAllCat in postgres DB...

func (r *PostgresRepository) GetAllCat() ([]*model.Cat, error) {
	res, err := r.DB.Query(context.Background(),
		"SELECT * FROM cats")
	if err != nil {
		return nil, err
	}
	allCat := []*model.Cat{}
	for res.Next() {
		u := &model.Cat{}
		err = res.Scan(&u.ID, &u.Name, &u.Age, &u.Color)
		if err != nil {
			return nil, err
		}
		allCat = append(allCat, u)
	}
	res.Close()
	return allCat, nil
}

// GetByIdCat in postgres DB...

func (r *PostgresRepository) GetByIDCat(ID uuid.UUID) (*model.Cat, error) {
	res, err := r.DB.Query(context.Background(),
		"SELECT * FROM cats WHERE id=$1", ID)
	if err != nil {
		return nil, err
	}
	u := &model.Cat{}
	for res.Next() {
		err = res.Scan(&u.ID, &u.Name, &u.Age, &u.Color)
		if err != nil {
			return nil, err
		}
	}
	res.Close()
	return u, nil
}

//  CreateCat in postgres DB...

func (r *PostgresRepository) CreateCat(u *model.Cat) (*model.Cat, error) {
	u.ID = uuid.New()
	res, err := r.DB.Query(context.Background(),
		"INSERT INTO cats (id,name,age,color) VALUES ($1,$2,$3,$4)",
		u.ID, u.Name, u.Age, u.Color)
	if err != nil {
		return nil, err
	}
	res.Close()
	return u, nil
}

// DeleteByIdCat in postgres DB...

func (r *PostgresRepository) DeleteByIDCat(ID uuid.UUID) (string, error) {
	_, err := r.DB.Exec(context.Background(),
		"DELETE FROM cats WHERE id=$1", ID)
	if err != nil {
		panic(err)
	}
	const res = "Cat deleted"
	return res, nil
}

// UpdateByIdCat in postgres DB...

func (r *PostgresRepository) UpdateByIDCat(ID uuid.UUID, u *model.Cat) (*model.Cat, error) {
	_, err := r.DB.Exec(context.Background(),
		"UPDATE cats SET name=$1, age=$2, color=$3 WHERE id=$4", u.Name, u.Age, u.Color, ID)
	if err != nil {
		panic(err)
	}
	u.ID = ID
	return u, nil
}
