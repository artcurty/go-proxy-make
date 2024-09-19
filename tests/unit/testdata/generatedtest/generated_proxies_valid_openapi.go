package generated

import (
	"github.com/artcurty/go-proxy-make/internal"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	internal.AddRouteRegistration(registerUserRoutes)
}

func registerUserRoutes(r *mux.Router) {
	r.HandleFunc("/test", ProxyGETtest).Methods("GET")
}

func ProxyGETtest(w http.ResponseWriter, r *http.Request) {
	internal.ProxyRequest(w, r, "http://localhost:8080/proxy-test", "GET", map[string]string{
		"orderId": "id",
		"name":    "product",
		"count":   "quantity",
	})
}
