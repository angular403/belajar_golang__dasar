package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5{
			break
		}
		fmt.Println("Perulangan Ke-", i)
	}


	for j := 2; j <= 10; j++{
		if j % 2 == 1{
			continue
		}
		fmt.Println(j)
	}
}