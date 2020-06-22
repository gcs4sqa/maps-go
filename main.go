package main

import "fmt"

func main() {

	// var colors map[string]string

	//colors := make(map[string]string)

	colors :=
		map[string]string{
			"red":   "#ff0000",
			"green": "#45678",
			"white": "#ffffff",
		}

	// colors["white"] = "#ffffff"
	// colors["red"] = "#ff0000"
	// colors["green"] = "#45678"

	// delete(colors, "white")

	printMap(colors)
	// fmt.Println(colors)

}

func printMap(c map[string]string) {

	for key, value := range c {
		fmt.Println("the Hex for", key, "is", value)
	}

}
