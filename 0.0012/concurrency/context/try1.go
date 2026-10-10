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

type Return struct {
	vaue string
	err  error
}

func fetchData(ctx context.Context, s string) (string, error) {
	timeout := time.Microsecond * 200
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	responseChan := make(chan Return)
	go func() {
		val, err := doingSomeWork()
		responseChan <- Return{
			vaue: val,
			err:  err,
		}

	}()
	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("wtf happend it took to long time")
		case resp := <-responseChan:
			return resp.vaue, resp.err
		}
	}
}

func doingSomeWork() (string, error) {
	time.Sleep(time.Microsecond * 600)
	return "SOme HEavy WoRk DoNe", nil
}
