package try5

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func InitilizeServer2() {
	fmt.Println("Starting The server in port 8080")

	mux := http.NewServeMux()
	server := http.Server{}

	mux.HandleFunc("/", handleRoot2)

	server.Addr = ":8080"
	server.Handler = mux

	if err := server.ListenAndServe(); err != nil {
		log.Println("Error: ", err)
	}
}

func handleRoot2(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Body:", request.Body)
	fmt.Fprintf(writer, "weeeeeeeeeeeeeeeeeeee")
	writer.Write([]byte("aaaaaaaaaaaaaaaaaa"))
	fmt.Println("Body:", request.Body)
	// This will print no thing cause body is io close ie: Body io.ReadCloser
	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println("read error:", err)
		return
	}
	fmt.Println("Body content:", string(bodyBytes))
}
