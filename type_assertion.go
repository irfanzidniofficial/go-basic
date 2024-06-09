package main

import "fmt"

func random()any{
	return "OK"

}

func main() {
	var result any = random()
	// var resultString = result.(string)
	// fmt.Println(resultString)

	switch value:=result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
		
	}
}