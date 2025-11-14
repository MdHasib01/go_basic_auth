package router

import (
	"net/http"
	"userStory/controller"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", controller.GetAlluser)
	mux.HandleFunc("POST /register", controller.RegisterUser)
	mux.HandleFunc("POST /login", controller.LoginUser)
	mux.HandleFunc("GET /users/{id}", controller.GetUserById)
	mux.HandleFunc("DELETE /users/{id}", controller.DeleteUserById)

	return mux
}
