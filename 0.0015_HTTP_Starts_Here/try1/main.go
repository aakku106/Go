package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("HTTP Server Initiliaztion...")
	http.HandleFunc("/", rootPage)
	http.ListenAndServe(":5555", nil)
}
func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Master CCN"))
}

/*
## HTTP in Go:

Well at 1st look those function name can give a miniHeartAttack for newbees, btu they are supper simiple
1st the HanddleFunc,
this func takes 2 arguments: pattern string, handler func(ResponseWriter, *Request)
thats look scarry, put paramater patern just mean the path/url/location like / or /profile or /contacts, thats it for beginers
and the another parameter handaler is it self a function which takes to arguments ResponseWriter and pointer to Request, whicj we will look after wile,
But simpally handeler just mean what happens when the user comes/goes/navigates to the pattern(/ in our case)

In our case when user navigates to / or rather wheneer browser or client mare request on :5555(Or evern in 5555/cat or whatever it will still give same response for now)
we will get "Welcome Master CCN" in localhost:5555
here in handeler function
thse 1st ResponseWriter is the interface (Defined in /net/http/server.go)
// A ResponseWriter interface is used by an HTTP handler to
// construct an HTTP response.
which cointain 3 methods: Header,Write,WriteHeader
and what we Used here was Write() method
	// Write writes the data to the connection as part of an HTTP reply.
&
	// If WriteHeader is not called explicitly, the first call to Write
	// will trigger an implicit WriteHeader(http.StatusOK).
	// Thus explicit calls to WriteHeader are mainly used to
	// send error codes or 1xx informational responses.
*/

// If that explanation felt something straight from aliean work, than its not much of a problem
// For now just focus on How ?, and after buildign something and getting hit with erros then explore Why ?

/*
Upon repuest on localhost:5555
we got response
HTTP/1.1 200 OK
Date: Mon, 05 Jan 2026 18:04:01 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

NOTE: The date will vary accorting to your OS current time and your setted location

and upon curl we got
 curl http://localhost:5555/
Welcome Master CCN

To understand what really happening here, you need to have knowlage of client server articture aand know basic go
Things we ignored here: mux, pattern(why even 5555/cat or :5555/<anothing>/<anything>/... gives same response as :5555)
*/
