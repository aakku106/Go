package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	server := http.Server{}
	mux.HandleFunc("/", handleRoot)
	server.Handler = mux
	server.ListenAndServe()

}
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome back master CCN !!!")
}
