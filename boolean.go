package main

import "fmt"

func main(){
	fmt.Println("benar =", true)
	fmt.Println("salah =", false)

	var nilaiakhir = 90
	var absensi = 80

	var lulusnilaiakhir bool = nilaiakhir > 80
	var lulusabsensi bool = absensi > 80

	var lulus bool = lulusnilaiakhir && lulusabsensi

	fmt.Println(lulus)
}
