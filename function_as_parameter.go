package main

import "fmt"

func sayHelloWithFiler(name string, filter func(string) string){
	filteredName:= filter(name)
	fmt.Println("Hello " +filteredName)

}



// Function type declaration

type Smile func(string) string

func sayHelloWithSmile(name string, smile Smile){
	smileName := smile(name)
	fmt.Println("Hello "+smileName)
}

func spamFilter(name string)string{
	if name == "Anjing" {
		return "..."

	} else {
		return name
	}
}

func main() {
	sayHelloWithFiler("Irfan", spamFilter)

	
	
    sayHelloWithFiler("Anjing", spamFilter)
	filter := spamFilter

	sayHelloWithFiler("Anjing", filter)

	fmt.Println("----------------------")

	sayHelloWithSmile("Anjing", spamFilter)
	
	

}