package main

import "fmt"

func main()  {
	name:= "irfan"

	switch name {
	case "irfan":
		fmt.Println("Hello irfan")
	case "zidni":
		fmt.Println("Hello zidni")
	default: 
		fmt.Println("Hi, boleh kenalan??s")
	}

	
	switch length:=len(name); length >5{
		case true:
            fmt.Println("Nama terlalu panjang")
        case false:
            fmt.Println("Nama sudah benart")
	}
	
}