package main

import "fmt"

func helloWorld() {
	fmt.Println("Hello World function")
}

func name(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

func main() {

	for i := 0; i <= 10; i++ {
		helloWorld()
	}

	name("Andrew","Wiliam")

}
