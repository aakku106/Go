package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Cat struct {
	CatName string `json:"catName"`
	CatAge  uint8  `json:"catAge"`
}

var CatCache = make(map[int]Cat)

func main() {

	mux := http.NewServeMux()
	server := http.Server{}
	mux.HandleFunc("POST /cat", createCat)
	mux.HandleFunc("GET /cat/{id}", getCat)
	server.Handler = mux
	fmt.Println("Server Listening At localhost")
	server.ListenAndServe()

}

func createCat(w http.ResponseWriter, r *http.Request) {
	var cat Cat
	json.NewDecoder(r.Body).Decode(&cat)
	CatCache[len(CatCache)+1] = cat
	w.WriteHeader(http.StatusNoContent)
}

func getCat(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	cat, _ := CatCache[id]
	w.Header().Set("content-type", "app/json")
	j, _ := json.Marshal(cat)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// This still feels entidimading
// Lest abort this POST think, and focus back on serving HTML or other kind of document fro our server insted of one line of text which only renders insider pre tag look in 0.0015/try4/main.go
