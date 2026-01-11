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

func handleRoot(writer http.ResponseWriter, request *http.Request)      {}
func handlePortfolio(writer http.ResponseWriter, request *http.Request) {}
func handleCon
