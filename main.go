package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
	log.Println("---STARTING!!!!!---")

	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	})

	http.HandleFunc("/healthcheck", healthCheck)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WHAT?!!!!")
	})

	log.Println("Server is ready to handle requests at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm up and running!\n")
}
