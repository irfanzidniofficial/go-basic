package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
    return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
    return n.Message
}

func SaveData(id string, data any)error{
	if id == ""{
        return &validationError{"validation error"}
    }
	if id =="irfan"{
		return &notFoundError{"data not found"}
	} 
	// OK
	return nil 

}

func main() {
	err:=SaveData("irfan", nil)
	if err != nil{
		// terjadi error
		// if validationErr, ok:= err.(*validationError); ok {
		// 	fmt.Println("validation error", validationErr.Error())
		// } else if notFoundErr, ok:= err.(*notFoundError); ok{
		// 	fmt.Println("not found error", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown error", err.Error())
		// }

		switch finalError:=err.(type) {

		case *validationError:
			fmt.Println("validation error", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error", finalError.Error())
		default:
			fmt.Println("Unknown error", finalError.Error())
			
		}
	
       
    } else {
        fmt.Println("OK")
    }
	
}