package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouting adds all the routes on the http server.
func SetupRouting(r *mux.Router) *mux.Router {
	r.HandleFunc("/acquire", acquire).Methods(http.MethodPost)
	r.HandleFunc("/checkAcquire", checkAcquire).Methods(http.MethodPost)
	r.HandleFunc("/release", release).Methods(http.MethodPost)
	r.HandleFunc("/checkRelease", checkRelease).Methods(http.MethodPost)
	return r
}
