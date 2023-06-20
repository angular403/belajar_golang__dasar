package main

import "fmt"

func fors() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke-", counter)
	}

	slice := []string{"Andrew", "wiliam", "ara"}

	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}
	
	for i, value := range slice {
		fmt.Println("Index", i, "=" , value)
	}
}