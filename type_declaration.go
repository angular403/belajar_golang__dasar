package main

import "fmt"

func Type_Declaration() {
	type NoKTP string
	type Married bool

	var NoKTPAndrew NoKTP = "4565646544565"
	var MarriedStatus Married = false
	fmt.Println(NoKTPAndrew)
	fmt.Println(MarriedStatus)
}