package main

import "fmt"

func Constant() {
	const (
		fistName = "Andrew"
		lastName = "Wiliam"
		age      = 25
	)

	fmt.Println(fistName)
	fmt.Println(lastName)
	fmt.Println(age)
}