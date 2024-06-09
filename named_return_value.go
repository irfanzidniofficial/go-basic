package main

import "fmt"

// Parameter ketiganya tipe data sring, maka cukup tulis salah satu
func getCompleteName()(firstName, middleName, lastName string){
	firstName = "Muhammad"
	middleName = "Irfan"
    lastName = "Zidni"

	return firstName, middleName, lastName
}

func main()  {
	firstName, middleName, lastName := getCompleteName()
    fmt.Println(firstName, middleName, lastName)
	
}