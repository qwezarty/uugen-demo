package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/google/uuid"
)

var num int32

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

	// Built-in Atomic ops
	atomic.AddInt32(&num, 1)
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", atomic.LoadInt32(&num))
}
