package api

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/fetch/data", getData)
	http.HandleFunc("/search", search)


	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}


}

