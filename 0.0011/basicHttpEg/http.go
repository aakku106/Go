package main

import (
	"fmt"
	"io"
	"net/http"
)

func getSomething() ([]byte, error) {
	res, err := http.Get("cat.cat")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	return data, err
}
