package main

import "fmt"

func main()  {
	name:="john"
	if(name == "irfan") {
		fmt.Println("Hello Irfan")
	} else if name=="joko"{
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, boleh kenalan gak?")
	}
	
	length := len(name)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama terlalu pendek")
	}
}