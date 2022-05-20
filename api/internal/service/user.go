package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/olafszymanski/devroot/api/internal/domain/entity"
	"github.com/olafszymanski/devroot/api/internal/domain/repository"
	"github.com/olafszymanski/devroot/api/pkg/hash"
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
	pwd, err := hash.Hash(req.Password)
	if err != nil {
		return nil, err
	}
	user, err := u.repo.Create(&entity.User{
		ID:        uuid.New(),
		Name:      req.Name,
		Username:  req.Username,
		Email:     req.Email,
		Password:  pwd,
		Admin:     false,
		Image:     "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	})
	if err != nil {
		return nil, err
	}

	return &entity.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Admin:     user.Admin,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u *UserService) Update(id uuid.UUID, req *entity.UserUpdateRequest) (*entity.UserResponse, error) {
	user, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if len(req.Name) > 0 {
		user.Name = req.Name
	}
	if len(req.Username) > 0 {
		user.Username = req.Username
	}
	if len(req.Email) > 0 {
		user.Email = req.Email
	}
	if len(req.Password) > 0 {
		pwd, err := hash.Hash(req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = pwd
	}
	if len(req.Image) > 0 {
		user.Image = req.Image
	}
	user.UpdatedAt = time.Now()

	updated, err := u.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return &entity.UserResponse{
		ID:        updated.ID,
		Name:      updated.Name,
		Username:  updated.Username,
		Email:     updated.Email,
		Password:  updated.Password,
		Admin:     updated.Admin,
		Image:     updated.Image,
		CreatedAt: updated.CreatedAt,
	}, nil
}

func (u *UserService) GetByID(id uuid.UUID) (*entity.UserResponse, error) {
	user, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &entity.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Admin:     user.Admin,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u *UserService) Delete(id uuid.UUID) error {
	if err := u.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
