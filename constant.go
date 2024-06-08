package main

import "fmt"

func main() {
	const firstName string = "Irfan"
	const lastName string = "Zidni"

	// Multiple constant

	const (
        nickName = "Irfan"
        middleName = "Zidni"
    )
	fmt.Println(nickName)
	fmt.Printf(middleName)


}