package main

import "fmt"

func main(){

	var name = "farhan iqbal"
	var e = name[1]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)
	
	fmt.Println("Farhan")
	fmt.Println("Maulana")
	fmt.Println("Iqbal")

	fmt.Println(len("Farhan"))
	fmt.Println("Maulana"[1])
	fmt.Println("Iqbal"[3])
}
