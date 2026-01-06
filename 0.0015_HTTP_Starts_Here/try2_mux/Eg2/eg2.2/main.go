package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/cat", handleCat)
	mux.HandleFunc("/dog", handleDog)

	server := http.Server{}
	server.Handler = mux
	server.ListenAndServe()

}
func handleCat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Only cat can come here, but Welcome Master CCN"))
}

func handleDog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Only Dog can come here, but Welcome Master CCN"))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is root Welcome Master CCN"))
}

/*
Now this server is listening in port 80 or simpally localhost it self
*/

// Till now w have only looking into GET request, lest look into PUT in try3_POST/main.go
