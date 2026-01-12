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

/*
What we observed here
1st lest ignore everything other, adn focus on Header,
upon using curl, you might see something clean like
Header:  map[Accept:[* / *] User-Agent:[curl/8.7.1]]
its showing user-agent(client) is curl version 8.XX.XX
its actually showing how go actually store header in memory, in go Header is :
type Header map[string][]string
here * / * means it accepts any media types in HTTP it means
it acepcts HTML,JSON,XML,or even plain text

But if we make request from browseres we wee see lots of stuffs
Accept":[]string{"text/html,application/xhtml+xml,application/xml;q=0.9,* / *;q=0.8
also with this we wee see accepted language, encoding,Cookie,Dnt,priority, useragent,
and many more things
And its browser browser espisific, some bowsers like safari focuses more on speed,batterylife and ecosystem
and other browser like firefox just dumps everythig, because it has to run/work every-where
So you may see consize header on curl, clean and platform specific from safari and verbose from firefox and chorimum
*/

func handlePortfolio(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Body:", request.Body)
	fmt.Println("Method: ", request.Method)
	fmt.Println("URL: ", request.URL)
	fmt.Println("URL.path: ", request.URL.Path)
	fmt.Printf("%#v\n", request.Header)
	// shows: http.Header{"Accept":[]string{"*/*"}, "User-Agent":[]string{"curl/8.7.1"}}
	// thats just mormating not new magic
}
func handleContact(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	if request.Method != http.MethodGet {
		log.Println("Method Not allodwed here,", http.StatusMethodNotAllowed)
	}
	fmt.Println("Body:", request.Body)
	fmt.Println("Method: ", request.Method)
	fmt.Println("URL: ", request.URL)
	fmt.Println("URL.path: ", request.URL.Path)
	fmt.Println("Header: ", request.Header)
}
