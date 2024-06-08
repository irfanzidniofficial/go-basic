package main

import "fmt"

func main() {
	// counter :=1

	// for counter <= 10 {
    //     fmt.Println("Perulangan ke ", counter)
    //     counter++
    // }

	// fmt.Println("Selesai")

	for counter := 1; counter <=10; counter++ {
		fmt.Println("Perulangan ke ", counter)
		
	}
	fmt.Println("Selesai")

	names:=[] string{"Irfan", "Zidni", "John"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[1])
	}

	for index, name:= range names{
		fmt.Println("Index", index, "=", name)
	}
	
	fmt.Println("--------------------------------")
	for _, name:= range names{
		fmt.Println(name)
	}


}