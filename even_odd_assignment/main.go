package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		oddOrEvenText := "is odd"

		if i%2 == 0 {
			oddOrEvenText = "is even"
		}

		fmt.Println(i, oddOrEvenText)
	}
}
