package playingwithjson

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Go standard libary provides a built in package to work with JSON (	encoding/decoding ) and struct (serilize/deserilize)
Go has encoding/json package, has toold like marshal(convert Go data to JSON)
and unmarshal (Converting JSON to Go data)
Go's struct matches with JSON more naturally
*/
type normalStruct struct {
	name    string
	age     uint8
	address string
	email   string
}

/*
This is normal struct in Go which naturally looks liek a json format:
{
"name":"...",
"age":"...",
address:"...",
email:"..."
}
*/
func call1() {
	data := normalStruct{
		name:  "cat",
		email: "cat.cat@cat.cat",
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Fatal("An error occured while marsaling JSON")
	}
	fmt.Println(response) // thats show raw bytes,so
	fmt.Println(string(response))
	// And its still showing an empty Object (js)
	fmt.Println(data)
	// But this raw struct shows output as aspected, why ?
}

type StructWithTick struct {
	name  string `json:"name"`
	email string `email`
}

func call2() {
	data := StructWithTick{
		name:  "aakku",
		email: "aakku106@gmail.com",
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error while marsaling")
	}
	fmt.Println(string(response))
	// still after using struct tag we got and empty object {}
}

/*
So what is the real cluprit here, its :
field name in struct, it always need to be public
*/
type struct3 struct {
	Name    string
	Address string
	age     uint8
}

func call3() {
	data := struct3{
		Name:    "Adarasha Gaihre",
		Address: "Nepal",
		age:     20,
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error while marsaling")
	}
	fmt.Println(string(response))
	/*
				THe output always shows:
				{"Name":"Adarasha Gaihre","Address":"Nepal"}

				but ignores age (because it was private or starts with small case)
		here the output we got dosent look that much standard, lest put struct tag on them
	*/
}

type struct4 struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     uint8  `json:"age"`
}

func call4() {
	data := struct4{
		Name:    "Adarasha Gaihre",
		Address: "Nepal",
		Age:     20,
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error while marsaling")
	}
	fmt.Println(string(response))
	/*
		{"name":"Adarasha Gaihre","address":"Nepal","age":20}
		this one looks good, but we could do more humain reabable usinf indents
	*/
}
func call5() {
	data := struct4{
		Name:    "Adarasha Gaihre",
		Address: "Nepal",
		Age:     20,
	}
	response, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Fatal("Error while marsaling")
	}
	fmt.Println(string(response))
	/*

								{
								"name": "Adarasha Gaihre",
								"address": "Nepal",
								"age": 20
								}
				this was much readable than : {"name":"Adarasha Gaihre","address":"Nepal","age":20}
		but logically they both are same
		marsal indent just adds styling

	*/
	// actually we could also do
	response, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error while marsaling")
	}
	fmt.Println(string(response))
	/*
		 which returns
		{
		  "name": "Adarasha Gaihre",
		  "address": "Nepal",
		  "age": 20
		}
	*/
	// Or we could also do
	response, err = json.MarshalIndent(data, "", "--->")
	if err != nil {
		log.Fatal("Error while marsaling")
	}
	fmt.Println(string(response))

	/*
		 which returns
		{
		--->"name": "Adarasha Gaihre",
		--->"address": "Nepal",
		--->"age": 20
		}
	*/
}

/*
While the way encoding/json , looks wired and self contradicting,
It can be a good security guard rail it self,
for eg here:
type struct3 struct {
	Name    string
	Address string
	age     uint8
}

encoding/json can never read age, this prevents
Libraries to peek into your private fields,
accidental data leak,
frameworks silently serializing things you didn’t intend

* The code editors may throw warnings in
bad code like
type StructWithTick struct {
	name  string `json:"name"`
	email string `email`
}
saying: struct filed has tag json but not exported, it's just saying what i explained earlier, outer packages needs
public variable/filds to access data


*/

// NEXT: we will look into unmarsal in ./unMarsal.go
