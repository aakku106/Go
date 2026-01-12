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
	if request.Method != "GET" {
		// "GET" is same do http.MethodGet(As i already explained in main.go file)
		http.Error(writer, "Only Get here", http.StatusMethodNotAllowed)
	}
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

/*
Here whenever we do

❯ curl -X POST http://localhost:8080/contact \                     ─╯
     -H "Content-Type: application/json" \
     -d '{"email":"aakku@gmail.com","msg":"hello"}'

we get (for user)

Only Get here
weeeeeeeeeeeeeeeeeeeeaaaaaaaaaaaaaaaaaa%

for server

Starting The server in port 8080
Body: &{0x1400000e210 <nil> <nil> false true {{} {0 0}} false false false 0x100992160}
Body: &{0x1400000e210 <nil> <nil> false true {{} {0 0}} false false false 0x100992160}
Body content: {"email":"aakku@gmail.com","msg":"hello"}

Althow we had if condition to filter out only GET, but since we dont have return in that if block
whatever below will strill run and response (Could be dengerious in producton, this is just to show/test)
*/

//NEXT: ./main3.go we will look into URL into deep
