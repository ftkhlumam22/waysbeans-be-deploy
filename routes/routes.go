package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	UserRoute(r)
	AuthRoute(r)
	ProductRoute(r)
	CartRoutes(r)
	TransactionRoutes(r)
}
