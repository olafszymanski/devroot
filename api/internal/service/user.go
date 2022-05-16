package service

import (
	"github.com/google/uuid"
	"github.com/olafszymanski/devroot/api/internal/domain/entity"
	"github.com/olafszymanski/devroot/api/internal/domain/repository"
)

type IUserService interface {
	Create(req *entity.UserRequest) (*entity.UserResponse, error)
	Update(id uuid.UUID, req *entity.UserRequest) (*entity.UserResponse, error)
	GetByID(id uuid.UUID) (*entity.UserResponse, error)
	Delete(id uuid.UUID) error
}

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Create(req *entity.UserRequest) (*entity.UserResponse, error) {
	return nil, nil
}

func (u *UserService) Update(id uuid.UUID, req *entity.UserRequest) (*entity.UserResponse, error) {
	return nil, nil
}

func (u *UserService) GetByID(id uuid.UUID) (*entity.UserResponse, error) {
	return nil, nil
}

func (u *UserService) Delete(id uuid.UUID) error {
	return nil
}
