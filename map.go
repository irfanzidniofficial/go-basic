package main

import "fmt"

func main() {
	var person map[string]string=map[string]string{}
	person["name"]="Irfan"
	person["address"]="Semarang"

	// Cara cepat
	person2:=map[string]string{
        "name":"Irfan",
        "address":"Jakarta",
    }

    fmt.Println(person["name"])
	fmt.Println(person2["address"])
	fmt.Println(person)
	fmt.Println(person2)

	fmt.Println("--------------------------------")

	book:=make(map[string]string)
	book["title"]="Buku Golang"
	book["author"]="Irfan"
	book["wrong"]="Salah"

	fmt.Println(book)

	// Delete data di map
	delete(book, "wrong")

	fmt.Println(book)
    

}