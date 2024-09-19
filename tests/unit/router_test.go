package unit

import (
	"github.com/artcurty/go-proxy-make/internal"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func TestRouteRegistrations(t *testing.T) {
	tests := []struct {
		name     string
		register func(*mux.Router)
		expected []string
	}{
		{
			name: "Adds a single route registration",
			register: func(r *mux.Router) {
				r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
			},
			expected: []string{"/test"},
		},
		{
			name: "Adds multiple route registrations",
			register: func(r *mux.Router) {
				r.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
				r.HandleFunc("/test2", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
			},
			expected: []string{"/test1", "/test2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			internal.RouteRegistrations = nil
			internal.AddRouteRegistration(tt.register)

			if len(internal.RouteRegistrations) != 1 {
				t.Errorf("Expected 1 route registration function, got %d", len(internal.RouteRegistrations))
			}

			r := mux.NewRouter()
			internal.RegisterRoutes(r)

			for _, route := range tt.expected {
				found := false
				r.Walk(func(muxRoute *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
					path, err := muxRoute.GetPathTemplate()
					if err == nil && path == route {
						found = true
						t.Logf("path %s found with route %s", path, route)
					}
					return nil
				})
				if !found {
					t.Errorf("Expected route %s to be registered, but it was not", route)
				}
			}
		})
	}
}
