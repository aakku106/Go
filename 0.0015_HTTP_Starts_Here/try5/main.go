package try5

import (
	"fmt"
	"log"
	"net/http"
)

func InitilizeServer() {
	fmt.Println("Starting The server in port 8080")

	mux := http.NewServeMux()
	server := http.Server{}

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/portfolio", handlePortfolio)
	mux.HandleFunc("/contact", handleContact)

	server.Addr = ":8080"
	server.Handler = mux

	if err := server.ListenAndServe(); err != nil {
		log.Println("Error: ", err)
	}
}

// Insted of making a working serverthis time, lest understand some basic methods this time
func handleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Body:", request.Body)
	fmt.Println("Method: ", request.Method)
	fmt.Println("URL: ", request.URL)
	fmt.Println("URL.path: ", request.URL.Path)
	fmt.Println("Header: ", request.Header)
}

func handleRoot(writer http.ResponseWriter, request *http.Request)      {}
func handlePortfolio(writer http.ResponseWriter, request *http.Request) {}
func handleCon
