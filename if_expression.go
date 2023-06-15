package main

import "fmt"

func expresion() {
	name := "Andrew"

	if name == "Andrew" {
		fmt.Println("Hello Andrew")
	}else{
		fmt.Println("Nama Salah")
	}


	if length := len(name); length <4 {
		fmt.Println("Nama Terlalu Pendek")
	}else{
		fmt.Println("Nama Sudah Benar")
	}
}