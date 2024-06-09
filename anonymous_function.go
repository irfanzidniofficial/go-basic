package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList){
	if blacklist(name){
        fmt.Println("Your are blocked", name)
    } else{
		fmt.Println("Hello", name)
	}
  
}

func main()  {
	// Anonymous Function
	blacklist := func(name string)bool{
		return name == "Anjing"
    }
	registerUser("Irfan", blacklist)
	registerUser("Anjing", blacklist)

	fmt.Println("---------------")

	// Bisa juga seperti ini
	registerUser("Anjing", func(name string) bool{
		return name == "Anjing"

	})


	
}