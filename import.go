package main

import (
	"fmt"
	"go-basic/helper"
)

func main() {
	result := helper.SayHello("Irfan")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses karena function awalannya huruf kecil di package helpernya
	

	
}