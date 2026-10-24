package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Initilizating server...")
	mux := http.NewServeMux()
	// NewServeMux allocates and returns a new [ServeMux].
	/*
		type ServeMux struct {
			mu     sync.RWMutex
			tree   routingNode
			index  routingIndex
			mux121 serveMux121 // used only when GODEBUG=httpmuxgo121=1
		}
	*/
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/cat", handleCat)
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Back Master CCN"))
}
func handleCat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Only Cat are Welcomed here !!!"))
}

/*
In this simple example we made use of Mux
which simpally means multiplexer
// ServeMux is an HTTP request multiplexer.
// It matches the URL of each incoming request against a list of registered
// patterns and calls the handler for the pattern that
// most closely matches the URL.
//
*/
