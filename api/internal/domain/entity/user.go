package entity

import (
	"time"

	"github.com/google/uuid"
)

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

type UserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
