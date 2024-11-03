package apis

import (
	"log"
	"net/http"
)

func Serve() {
	r := init_mux()

	server := http.Server{
		Addr:    "127.0.0.1:8181",
		Handler: r,
	}
	log.Println("Starting server on :", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
