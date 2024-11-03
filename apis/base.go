package apis

import (
	"net/http"
)

func init_mux() *http.ServeMux {

	r := http.NewServeMux()

	// returns ok
	r.HandleFunc("GET /health", health_check)

	// BINDINGS

	// MIDDLE WARES

	return r

}
