package services

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
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

func (s UserService) Insert(ctx context.Context, user entity.UserRequest) (int, error) {
	// set data id to user from google uuid
	user.ID = uuid.New().String()
	result, err := s.UserRepository.Insert(ctx, s.db, user)
	if err != nil {
		return 0, err
	}
	return result, nil
}
