package main

import "fmt"

func main()  {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Irfan"
	names[2] = "Zidni"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Membuat array langsung

	var values [3]int = [3] int {
		90,
		80,
	}
	

	fmt.Println(values)

	fmt.Println("--------------------------------")

	var nilai = [...]int{
		80,
		90,
		70,
	}

	fmt.Println(nilai)
	// Menghitung panjang array
	fmt.Println(len(nilai))


	
}