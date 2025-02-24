package router

import (
	"fmt"
	"net/http"
	"rest_api_go/adapter/api"
	"rest_api_go/adapter/database"
	"rest_api_go/infrastructure/db"
	"rest_api_go/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	DBAdministrator := db.NewDBAdministrator(db.DB)
	userRepository := database.NewUserRepository(DBAdministrator)
	userUsacase := usecase.NewUserUsecase(userRepository)
	userApi := api.NewUserApi(userUsacase)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, world!")
	})
	r.Get("/users", userApi.Fetch)
	return r
}
