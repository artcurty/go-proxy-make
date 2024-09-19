package internal

import (
	"github.com/gorilla/mux"
)

var RouteRegistrations []func(*mux.Router)

func RegisterRoutes(r *mux.Router) {
	for _, register := range RouteRegistrations {
		register(r)
	}
}

func AddRouteRegistration(register func(*mux.Router)) {
	RouteRegistrations = append(RouteRegistrations, register)
}
