package playingwithjson

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
https://jsonplaceholder.typicode.com/users
*/

// Lest understand unmarshal
type Location struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Tole     string `json:"tole"`
}

type Address struct {
	Email     string `json:"email"`
	*Location `json:"location"`
}

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	*Address `json:"address"`
}

func um() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer response.Body.Close()
	fmt.Println(response.Body)
	/*
			this does give someoutput but :
		&{[] {0x14000204000} <nil> <nil>}
		thats looks like address of some slice
	*/
	// So we have to use proper Unmarsal to get data
}
func um2() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer response.Body.Close()

	var users []User

	fmt.Println("Before Unmarshal:")
	fmt.Println(users)

	r, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(r, &users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After Unmarshal:")
	fmt.Println(users)

	/*
									Here after marshal we will get value:
									[{1 Leanne Graham 0 0x1400000e198} {2 Ervin Howell 0 0x1400000e1b0} {3 Clementine Bauch 0 0x1400000e1c8} {4 Patricia Lebsack 0 0x1400000e1e0} {5 Chelsey Dietrich 0 0x1400000e1f8} {6 Mrs. Dennis Schulist 0 0x1400000e210} {7 Kurtis Weissnat 0 0x1400000e228} {8 Nicholas Runolfsdottir V 0 0x1400000e240} {9 Glenna Reichert 0 0x1400000e258} {10 Clementina DuBuque 0 0x1400000e270}]

							Which may look wired at 1st place, it is because we didnt handeled the structure of the respponse given by api,
							thats why we get only 2 readable value of
					id and name, because it matches our defined :
				Id       uint   `json:"id"`
				Name     string `json:"name"`

		Which means 1st we shall know how/what response will will get from fetching api to unmarshal it properly

	*/

}
