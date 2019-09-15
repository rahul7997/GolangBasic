package main

import "fmt"

func main() {

	i := 1

	for i <= 5 {
		fmt.Println("GO", i)
		i++

	}
	fmt.Println()

	for j := 1; j <= 100; j++ {
		fmt.Print(j, " ")
	}

}
