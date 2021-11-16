package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/sadetskayfu/rest-golang/internal/model"
	"time"
)

// Registration user...

func (s *Service) Registration(u *model.User) (*model.User, error) {
	res, err := s.repo.Registration(u)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LogIn user...

func (s *Service) LogIn(u *model.AuthUser) (string, error) {
	res, err := s.repo.LogIn(u)
	if err != nil {
		panic(err)
	}
    // Claims ...

    Claims := &model.JwtCustomClaims{
		ID : res.ID,
		StandardClaims: jwt.StandardClaims{
        ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt: time.Now().Unix(),
		},
	}
	// Create token with claims...
	
	atToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	token, err := atToken.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return token, nil
}
