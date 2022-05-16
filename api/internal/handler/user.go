package handler

import (
	"net/http"

	"github.com/olafszymanski/devroot/api/internal/service"
)

type UserHandler struct {
	*service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{us}
}

func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (uh *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (uh *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
