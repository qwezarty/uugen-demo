package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

// all basic data types are not thread-safe in Go
// behave like a anonymous class, def and init.
var count struct {
	num int // zero value of int, i.e. 0
	mux sync.Mutex
}

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

	count.mux.Lock()
	defer count.mux.Unlock() // defer doing after leaving this block
	count.num++
}

func counter(w http.ResponseWriter, r *http.Request) {
	count.mux.Lock()
	defer count.mux.Unlock()

	fmt.Fprintf(w, "Count %d", count.num)
}
