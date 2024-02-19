package services

import (
	"context"
	"database/sql"
	"rest-api-golang/src/entity"
	"rest-api-golang/src/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
	db             *sql.DB
}

func NewUserService(userRepository *repositories.UserRepository, db *sql.DB) *UserService {
	return &UserService{UserRepository: userRepository, db: db}
}

func (s UserService) Get(ctx context.Context) ([]entity.Users, error) {
	users, err := s.UserRepository.Get(ctx, s.db)
	if err != nil {
		return []entity.Users{}, nil
	}
	return users, nil
}
