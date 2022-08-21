package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	// fmt.Println("welome to the struct operation")
	// hitesh := User{Name: "bhadri",Email: "bhadri2002@gmail.com",Status: true,Age: 19}
	// fmt.Println(hitesh)
	// fmt.Printf("%+v",hitesh)
	// fmt.Println()

	days := []string{"sunday","monday","tuesday","wednesday","thursday","friday","saturday"}

	for _,v := range days{
		fmt.Println(v)
	}
}
