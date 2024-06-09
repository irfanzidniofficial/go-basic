package main

import "fmt"

func getHello(name string) string{
	return "Hello " + name
}

func main()  {
	hello:=getHello("Irfan")
	fmt.Println(hello)

	fmt.Println(getHello("Zidni"))
	
}