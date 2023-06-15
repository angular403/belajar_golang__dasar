package main

import "fmt"

func array() {
	var names [3]string
	names[0] = "Andrew"
	names[1] = "wiliam"
	names[2] = "napitupulu"
	
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])


	var values = [3] int{
		90,80,85,
	}

	fmt.Println(values)
}