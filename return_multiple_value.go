package main

import "fmt"

func getFullName() (string, string){
	return "Irfan", "Zidni"
}

func main()  {
	// Butuh semuanya
	// firstName, lastName := getFullName()
    // fmt.Println(firstName, lastName)

	// Hanya butuh firstName saja
	firstName, _ := getFullName()
    fmt.Println(firstName)

   
	
}