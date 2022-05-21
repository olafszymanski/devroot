package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/olafszymanski/devroot/api/internal/domain/entity"

	"github.com/olafszymanski/devroot/api/internal/service"
	"github.com/olafszymanski/devroot/api/pkg/response"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{us}
}

func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req entity.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}

	res, err := uh.userService.Create(&req)
	if err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req entity.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}

	strId := chi.URLParam(r, "id")
	id, err := uuid.Parse(strId)
	if err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
	res, err := uh.userService.Update(id, &req)
	if err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
}

func (uh *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	strId := chi.URLParam(r, "id")
	id, err := uuid.Parse(strId)
	if err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
	res, err := uh.userService.GetByID(id)
	if err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
}

func (uh *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	strId := chi.URLParam(r, "id")
	id, err := uuid.Parse(strId)
	if err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
	err = uh.userService.Delete(id)
	if err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.NewSuccessResponse(http.StatusOK, "User deleted successfully.")); err != nil {
		http.Error(w, response.NewErrorResponse(http.StatusBadRequest, err).ToJSON(), http.StatusBadRequest)
		return
	}
}
