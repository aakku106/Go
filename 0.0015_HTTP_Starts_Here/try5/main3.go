package try5

import (
	"fmt"
	"log"
	"net/http"
)

func InitilizeServer3() {
	fmt.Println("Starting The server 3 in port 8080")

	mux := http.NewServeMux()
	server := http.Server{}

	mux.HandleFunc("/", handleRoot3)
	mux.HandleFunc("/contact", handleContact3)

	server.Addr = ":8080"
	server.Handler = mux

	if err := server.ListenAndServe(); err != nil {
		log.Println("Error: ", err)
	}
}

func handleRoot3(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("URL QUERY map: ", request.URL.Query())
	fmt.Println("URL: ", request.URL)
	fmt.Println("URL Opaque: ", request.URL.Opaque)
	fmt.Println("URL OmitHost: ", request.URL.OmitHost)
	fmt.Println("URL User Username: ", request.URL.User.Username())
	fmt.Println("Scheme:", request.URL.Scheme)
	fmt.Println("Host:", request.URL.Host)
	fmt.Println("Host:", request.Header)
	fmt.Println("Path:", request.URL.Path)
	fmt.Println("RawQuery:", request.URL.RawQuery)
}

/*
Most of these feels extermly difficult, and mostly not used in http servers
Lest go one bu one

 1. request.URl: whole struct(Whatver client request sends
    <path included in this func its / and for below func it will be /contact>)

 2. URL.Query(): when you call Query, it internally
    split on &
    split on =
    and store as map[string][]string
    eg: curl "http://localhost:8080/?name=aakku&age=20&skill=go" upon calling this, Query
    break/split URL on & and = (URL decode)
    it store those values and keys like
    map[ age[20] name[aakku] skill[go] ]
    thats all, and one more thing it request was
    curl "http://localhost:8080/?name=aakku&age=20&skill=go&skill=js&specility=go" upon calling this, Queru make map like
    map[ age[20] name[aakku] skill[go js] specility[go] ]
    it cont loss and data by collasping(collasping = loosing data)
*/
func handleContact3(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("URL QUERY map: ", request.URL.Query())
	fmt.Println("URL: ", request.URL)
	fmt.Println("URL Opaque: ", request.URL.Opaque)
	fmt.Println("URL OmitHost: ", request.URL.OmitHost)
	fmt.Println("URL User Username: ", request.URL.User.Username())
	fmt.Println("Scheme:", request.URL.Scheme)
	fmt.Println("Host:", request.URL.Host)
	fmt.Println("Host:", request.Header)
	fmt.Println("Path:", request.URL.Path)
	fmt.Println("RawQuery:", request.URL.RawQuery)
}
