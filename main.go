package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	serveMux := http.NewServeMux()
	server := http.Server{
		Addr:    port,
		Handler: serveMux,
	}

	fmt.Printf("Server up and running on port %v\n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while listening and serving server, %v", err)
	}
}
