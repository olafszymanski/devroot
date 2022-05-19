package api

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/olafszymanski/devroot/api/internal/domain/repository"
	"github.com/olafszymanski/devroot/api/internal/handler"
	"github.com/olafszymanski/devroot/api/internal/service"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	var db *sqlx.DB
	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)
	uh := handler.NewUserHandler(us)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", uh.Create)
		r.Get("/{id}", uh.GetByID)
		r.Put("/{id}", uh.Update)
		r.Delete("/{id}", uh.Delete)
	})

	return r
}
