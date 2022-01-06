package main

import (
	"fmt"
)

func main() {
	x := "Moneypenny" // if change name x , out  put is neither

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}
}
