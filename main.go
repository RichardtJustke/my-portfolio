package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oi mundo"))
	})

	log.Println("rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
