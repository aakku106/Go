package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Initilizating Server....")

	mux := http.NewServeMux()
	fmt.Println("Mux created")

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/cat", handleCat)
	mux.HandleFunc("/dog", handleDog)

	fmt.Println("3 route Handeler Created")
	fmt.Println("Listening In: localhost:8080")

	server := http.Server{}
	/*
			Server is the struct
		// A Server defines parameters for running an HTTP server.
		// The zero value for Server is a valid configuration.
		There are ~24 values inside this struct
	*/

	// So let us give our address 1st
	server.Addr = ":8080"
	// And also give our handeler to this server
	server.Handler = mux
	// cause our handeler are defined in mux above
	// And if we pass this server address to Listen and server method
	server.ListenAndServe()
	// We dont have to pass the arguments

	// http.ListenAndServe(":8080", mux) // And now we dont have to do this
	// Confused ?? read last comment
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is root Welcome Master CCN"))
}

func handleCat(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Only cat can come here, but Welcome Master CCN"))
}

func handleDog(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Only Dog can come here, but Welcome Master CCN"))
}

/*
The Lestion and server method, seems to be ambigious at 1st look but its just simple polymerphism
the real working LestionAndServe method is what we used in this file
And ListenAndServer used in proyer file ../main.go also go what we did in this file like:

func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

Before we passed the port/address and handelerFunction that method puts those 2 value in a struct named server and return
server.ListenAndServer which calls the exact method which we called in this file

// You can read /net/http/server.go for further clerification
*/

// Leaving server.addr value empty is also a valid thing to do the server will listen on :http look at ../Eg2/eg2.2/main.go
