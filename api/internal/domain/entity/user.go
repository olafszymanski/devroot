package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type User struct {
	ID        uuid.UUID  `db:"id"`
	Name      string     `db:"name"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	Admin     bool       `db:"admin"`
	Image     string     `db:"image"`
	CreatedAt time.Time  `db:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt"`
	DeletedAt *time.Time `db:"deletedAt"`
}

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,max=32"`
	Username string `json:"username" validate:"required,max=32"`
	Email    string `json:"email" validate:"required,max=64,email"`
	Password string `json:"password" validate:"required,max=64"`
}

func (u *UserCreateRequest) Validate() error {
	return validate.Struct(*u)
}

type UserUpdateRequest struct {
	Name     string `json:"name" validate:"max=32"`
	Username string `json:"username" validate:"max=32"`
	Email    string `json:"email" validate:"max=64,email"`
	Password string `json:"password" validate:"max=64"`
}

func (u *UserUpdateRequest) Validate() error {
	return validate.Struct(*u)
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Admin     bool      `json:"admin"`
	Image     string    `json:"image"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
