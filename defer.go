package main

import "fmt"

func logging(){
	fmt.Println("Logging")
}

func runApplication(){
	// defer:  Dipanggil ketika function runApplication di selesai dijalankan, walaupun function di runApplication ada yang error, function logging ttetep dijalankan karena pakai defer
	defer logging()
	fmt.Println("Run Application")
}

func main()  {
	runApplication()
}