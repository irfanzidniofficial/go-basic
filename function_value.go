package main

import "fmt"


func sayGoodBye(name string)string{
	return "Good Bye " + name

}

func main()  {
	goodbye := sayGoodBye
	fmt.Println(goodbye("Irfan"))
}