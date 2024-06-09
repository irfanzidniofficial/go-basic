package main

import "fmt"

// Struct -> template data/ prototype data
type Customer struct{
	Name, Address string
	Age int
}

func (customer Customer)sayHello(name string){
	fmt.Printf("Hello ", name, "my name is", customer.Name)

}

func main(){
	var irfan Customer
	irfan.Name = "Irfan Zidni"
	irfan.Address = "Jakarta"
	irfan.Age=20

	fmt.Println(irfan)

	fmt.Println("--------------------------------")

	joko:=Customer{
		Name: "Joko",
		Address: "Bandung",
        Age: 21,
	}
	fmt.Println(joko)

	budi :=Customer{"Budi", "Surabaya", 22}
	fmt.Println(budi)

	fmt.Println("--------------------------------")

	budi.sayHello("Irfan")

}