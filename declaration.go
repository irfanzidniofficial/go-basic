package main

import "fmt"

func main() {

	// Type Declaration, semacam bikin tipe data alias

	type NoKTP string
	var ktpIrfan NoKTP = "1111111111"

	var contoh = "2222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpIrfan) 
	fmt.Println(contohKtp)

}