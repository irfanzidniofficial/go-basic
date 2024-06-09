package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 :=Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Bandung"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah menjadi Bandung


	// Pointer -> Reference 
	address1 :=Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer

	address2.City = "Bandung"
	fmt.Println(address1) // data berubah m
	fmt.Println(address2) // data berubah


}