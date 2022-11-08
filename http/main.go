package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bodyBytes := make([]byte, 99999)
	// resp.Body.Read(bodyBytes)
	// fmt.Println(string(bodyBytes))

	rw := customWriter{} //replacement for os.Stdout

	io.Copy(rw, resp.Body)
}

func (customWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("Bytes written:", len(bs))

	return len(bs), nil
}
