package main

import "fmt"

func main() {

	fmt.Println("hello")

	fmt.Println(" jouons à pierre feuille ciseaux ")

	var input string
	fmt.Scan(&input)

	fmt.Printf(" vous avez joué %s ... PERDU", input)

}
