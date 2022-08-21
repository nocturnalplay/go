package main

import "fmt"

func main(){
	fmt.Println("welcome to the pointer")

	number := 40
	var ptr *int = &number

	fmt.Println("value of the pointer is:",*ptr)
	fmt.Println(number)

	*ptr = *ptr + 10

	fmt.Println("value of the pointer is:",*ptr)
	fmt.Println(number)


}