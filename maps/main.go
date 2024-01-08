package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#FF0000",
		"white": "#FFFFFF",
	}

	updateMap(colors)

	for _, r := range colors {
		fmt.Println(r)
	}

}

func updateMap(c map[string]string) {
	c["red"] = "#00000"
}
