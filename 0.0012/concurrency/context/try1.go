package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	s := "cat"
	val, err := fetchData(ctx, s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(val)
	fmt.Println(ctx)
	fmt.Println(time.Since(start))
}

func fetchData(ctx context.Context, s string) (string, error) {
	ctx, cancel := context.WithTimeout()
	val, err := doingSomeWork()
	if err != nil {
		return "", err
	}
	return val, nil
}

func doingSomeWork() (string, error) {
	time.Sleep(time.Microsecond * 600)
	return "SOme HEavy WoRk DoNe", nil
}
