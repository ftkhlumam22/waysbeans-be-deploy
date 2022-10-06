package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	UserRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(UserRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
}
