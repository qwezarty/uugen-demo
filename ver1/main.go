package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid" // importing from a domain
)

func main() { // main() is unique per package
	// router table
	http.HandleFunc("/uuid", handler)

	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	u, err := uuid.NewUUID() // multiple return values
	if err != nil {
		// error handling without try-catch
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(u.String())
}
