package main

import (
	"fmt"
)

func main()  {
	names := [...]string{"Muhammad", "Irfan", "Zidni", "Budi", "Nugraha", "Joko", "John"}

	slice1 := names[4:6]

	fmt.Println(slice1)


	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	fmt.Println("--------------------------------")

	days:= [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2:=append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	// daysBaru = [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Libur Baru"}
	fmt.Println(daySlice1)
	fmt.Println("--------Day Slice 2.......")
	fmt.Println(daySlice2)
	fmt.Println("--------days........")
	fmt.Println(days)
	
	fmt.Println("--------------------")
	var newSlice []string = make([]string, 2,5)

	newSlice[0]="Irfan"
	newSlice[1]="Zidni"
	// newSlice[2]="Zidn" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println("--------------------")

	newSlice2:= append(newSlice, "Zid")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fmt.Println("--------------------")
	newSlice2[0]="Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fmt.Println("-------Copy Slice-------------")

	fromSlice:=days[:]
	toSlice:=make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// 
	iniArray:=[...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}