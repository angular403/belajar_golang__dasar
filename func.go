package main

import "fmt"

func helloWorld() {
	fmt.Println("Hello World function")
}

func name(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string{
		if name == ""{
			return "Hello Bro"
		}else{
			return "Hello " + name
		}
	} 

func getFullName()(string, string){
	return "Andrew", "Wiliam"
}

func main() {

	for i := 0; i <= 10; i++ {
		helloWorld()
	}

	name("Andrew","Wiliam")

	fmt.Println(getHello("Andrew"))


	fmt.Println("=====================================")

	fisrtNamed, lastNamed := getFullName()
	fmt.Println(fisrtNamed, lastNamed)
}
