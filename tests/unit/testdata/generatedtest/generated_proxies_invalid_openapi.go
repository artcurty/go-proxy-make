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
	r.HandleFunc("/test", ProxyPROXY_MAPPINGtest).Methods("PROXY_MAPPING")
}

func ProxyPROXY_MAPPINGtest(w http.ResponseWriter, r *http.Request) {
	internal.ProxyRequest(w, r, "", "", map[string]string{})
}
