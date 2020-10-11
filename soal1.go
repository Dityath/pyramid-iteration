package main

import (
	"fmt"
	"strings"
)

func main() {

	var loop int

	fmt.Print("Masukin angka yang ingin di masukin: ")
	fmt.Scanln(&loop)

	for i := 1; i <= loop; i++ {
		fmt.Printf("%s\n", strings.Repeat("*", i))
	}

	loop--

	for i := loop; i >= 1; i-- {
		fmt.Printf("%s\n", strings.Repeat("*", i))
	}
}
