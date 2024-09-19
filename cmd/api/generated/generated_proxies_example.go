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
	r.HandleFunc("/order", ProxyPOSTorder).Methods("POST")
	r.HandleFunc("/order", ProxyGETorder).Methods("GET")
	r.HandleFunc("/payment", ProxyPOSTpayment).Methods("POST")
}

func ProxyPOSTorder(w http.ResponseWriter, r *http.Request) {
	internal.ProxyRequest(w, r, "http://localhost:8080/proxy-order", "POST", map[string]string{
		"name":    "product",
		"count":   "quantity",
		"orderId": "id",
	})
}

func ProxyGETorder(w http.ResponseWriter, r *http.Request) {
	internal.ProxyRequest(w, r, "http://localhost:8080/proxy-order", "GET", map[string]string{
		"orderId": "id",
		"name":    "product",
		"count":   "quantity",
	})
}

func ProxyPOSTpayment(w http.ResponseWriter, r *http.Request) {
	internal.ProxyRequest(w, r, "http://localhost:8080/payment-order", "PUT", map[string]string{
		"id":     "orderId",
		"status": "orderStatus",
	})
}
