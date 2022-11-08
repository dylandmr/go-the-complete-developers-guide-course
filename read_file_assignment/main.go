package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name wasn't provided.")
		os.Exit(1)
	}

	filename := os.Args[1]

	fmt.Printf("Opening file titled \"%v\"\n", filename)

	file, error := os.Open(filename)

	if error != nil {
		fmt.Println("Error reading file:", error)
	}

	io.Copy(os.Stdout, file)
}
