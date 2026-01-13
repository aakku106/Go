package try5

import (
	"fmt"
	"log"
	"net/http"
)

func InitilizeServer4() {
	fmt.Println("Starting The server 4 in port 8080")

	mux := http.NewServeMux()
	server := http.Server{}

	mux.HandleFunc("/", handleRoot4)
	//mux.HandleFunc("/contact", handleContact4)

	server.Addr = ":8080"
	server.Handler = mux

	if err := server.ListenAndServeTLS(); err != nil {
		log.Println("Error: ", err)
	}
}

func handleRoot4(writer http.ResponseWriter, request *http.Request) {
	if request.TLS == nil {
		fmt.Fprintf(writer, "This was HTTP request")
	}
	if request.TLS != nil {
		fmt.Fprintf(writer, "This was HTTPS request")
		fmt.Fprintf(writer, "TLS version: %v", request.TLS.Version)
		fmt.Fprintf(writer, "TLS version: %v", request.TLS.CipherSuite)
		fmt.Fprintf(writer, "TLS version: %v", request.TLS.PeerCertificates)
		fmt.Fprintf(writer, "TLS version: %v", request.TLS.HandshakeComplete)
		fmt.Fprintf(writer, "TLS version: %v", request.TLS.ServerName)
	}
}

// Ours request was not https, so it always prints its was HTTP request
/*
NOTE: TLS is not application layer thing, its transport layer security warper around TCP

NOTE: TLS is not he part of url.

So we are only checking if the request was tls or not
request.TLS == nil    → HTTP
request.TLS != nil    → HTTPS


Before any HTTP exists:
	1.	TCP connection opens
	2.	TLS handshake happens
	3.	Encryption keys are negotiated
	4.	Certificates are validated
	5.	Only then does HTTP begin

By the time your handler runs, TLS is already done.

request.TLS is the post-mortem report of that handshake.
*/
