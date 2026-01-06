package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Cat struct {
	CatName string `json:"catName"`
	CatAge  uint8  `json:"catAge"`
}

var CatCache = make(map[int]Cat)
var cacheMutex sync.RWMutex

func main() {

	mux := http.NewServeMux()
	server := http.Server{}
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /cat", createCat)
	mux.HandleFunc("GET /cat/{id}", getCat)
	server.Handler = mux
	fmt.Println("Server Listening At localhost")
	server.ListenAndServe()

}
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome back master CCN !!!")
}
func createCat(w http.ResponseWriter, r *http.Request) {
	var cat Cat
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if cat.CatName == "" {
		http.Error(w, "CatName is required MewMewwoo!!", http.StatusBadRequest)
		return
	}
	cacheMutex.Lock()
	CatCache[len(CatCache)+1] = cat
	cacheMutex.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func getCat(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cacheMutex.RLock()
	cat, oki := CatCache[id]
	cacheMutex.RUnlock()

	if !oki {
		http.Error(w, "CatNotFound", http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "app/json")

	j, err := json.Marshal(cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// This was a simple example for Get and Post Request, and this looks un-readable on 1st glance
// How to test ter's rest.http file in same dir ./rest.http
// If using just highlight the path/url you wanna get/post and press your shortcut for go to file under the curser
// If using Vs code you can see a little blue test saying send request press that
// Else use Postman or any other software of your liking
