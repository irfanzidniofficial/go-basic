package main

import "fmt"

func endApp(){
	fmt.Println("End App")
	// recover yang benar (ditaruh di bagian yang di defer)
	message:=recover()
	fmt.Println("Terjadi error", message)
}

func runApp(error bool){
	defer endApp()

	if error{
        panic("Ups Error")
    }
	// contoh program recover yang salah 
	

	// message := recover()
	// fmt.Println("Terjadi error", message)
}

func main() {
	runApp(true)
	fmt.Println("Muhammad Irfan Zidni")



}