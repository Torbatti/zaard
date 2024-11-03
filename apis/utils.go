package apis

import (
	"fmt"
	"net/http"
)

func health_check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
