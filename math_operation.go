package main

import "fmt"

func math() {
	var result = 15 + 15
	fmt.Println(result)

	var a = 20
	var b = 10
	var c = a * b
	fmt.Println(c)

	var i = 20
	 i +=10
	fmt.Println(i)
	// true,false,false,true


	var value = 100
	var valueb = 200
	fmt.Println(value < valueb)
	fmt.Println(value > valueb)
	fmt.Println(value == valueb)
	fmt.Println(value != valueb)


	var ujian = 80
	var absensi = 80
	// var lulusUjian = ujian >= 80
	// var lulusAbsensi = absensi > 80
	// var lulusAkhir = lulusUjian && lulusAbsensi
	fmt.Println(ujian >= 80 && absensi >= 80)
}