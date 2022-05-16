package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/olafszymanski/devroot/api/internal/domain/entity"
)

type IUserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	GetByID(id uuid.UUID) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
	Delete(id uuid.UUID) error
}

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Create(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepository) Update(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepository) GetByID(id uuid.UUID) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepository) GetByUsername(username string) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepository) Delete(id uuid.UUID) error {
	return nil
}
