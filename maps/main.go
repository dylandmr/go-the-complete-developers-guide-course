package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4BF745",
	}

	var zeroValueMap map[string]string

	makeMap := make(map[string]string)

	colors["white"] = "#FFFFFF"

	makeIntMap := make(map[int]string)

	makeIntMap[10] = "#FFFFFF"

	fmt.Println("makeIntMap", makeIntMap)

	delete(makeIntMap, 10)

	fmt.Println("updated makeIntMap", makeIntMap)

	fmt.Println("colors", colors)
	fmt.Println("zeroValueMap", zeroValueMap)
	fmt.Println("makeMap", makeMap)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for", color, "is", hex)
	}
}
