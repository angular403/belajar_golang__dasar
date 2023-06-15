package main

import "fmt"

func mapb() {
	person :=
		map[string]string{
			"name":    "Andrew",
			"address": "Jakarta",
		}

	person["title"] = "Progammer coding"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["Author"] = "Andrew"
	book["wrong"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(len(book))
	fmt.Println(book)

}
