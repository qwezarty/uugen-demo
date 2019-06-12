package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/google/uuid"
)

var count int32

func main() {
	// router table
	http.HandleFunc("/uuid", handler)
	http.HandleFunc("/count", counter)

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

	// Built-in atomic ops
	atomic.AddInt32(&count, 1) // we can also swap pointer
	// pointer is very important in Go
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Count %d", atomic.LoadInt32(&count))
}
