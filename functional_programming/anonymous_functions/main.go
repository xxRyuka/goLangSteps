package main

import "fmt"

func main() {

	topla := func(a, b int) int {
		return a + b
	}

	carp := func(a int, b int) int {
		return a * b
	}

	// I I F E - Immediately Invoked Function Expression

	func() {
		fmt.Println("Merhaba from IIFE")
		fmt.Println("Sonuc:", topla(3, 4))
		fmt.Println(carp(3, 4))
	}() // IIFE sonu () koyulmalıdır
}
