package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/olafszymanski/devroot/api/internal/domain/entity"
	"github.com/olafszymanski/devroot/api/internal/domain/repository"
)

type IUserService interface {
	Create(req *entity.UserCreateRequest) (*entity.UserResponse, error)
	Update(id uuid.UUID, req *entity.UserUpdateRequest) (*entity.UserResponse, error)
	GetByID(id uuid.UUID) (*entity.UserResponse, error)
	Delete(id uuid.UUID) error
}

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Create(req *entity.UserCreateRequest) (*entity.UserResponse, error) {
	return &entity.UserResponse{
		ID:        uuid.New(),
		Name:      "test",
		Username:  "test",
		Email:     "test",
		Admin:     false,
		Image:     "test",
		Password:  "test",
		CreatedAt: time.Now(),
	}, nil
}

func (u *UserService) Update(id uuid.UUID, req *entity.UserUpdateRequest) (*entity.UserResponse, error) {
	return nil, nil
}

func (u *UserService) GetByID(id uuid.UUID) (*entity.UserResponse, error) {
	return nil, nil
}

func (u *UserService) Delete(id uuid.UUID) error {
	return nil
}
