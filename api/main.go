package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hi 69")
	})

	http.ListenAndServe(":8000", r)
}
