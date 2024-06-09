package main

import "fmt"

type Address struct {
	City, Province, Country string
}
// Pointer di function 
func ChangeCountryToIndonesia(address *Address) {
	address.Country="Indonesia"
   


}

func main() {
	var address Address =Address{}
	ChangeCountryToIndonesia(&address)
	fmt.Println(address)

	
}