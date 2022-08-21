package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("welcome to the user input")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the rating:")

	//comma ok operator

	input,err := reader.ReadString('\n')

	if err == nil{
		fmt.Println("thank's for the feed back:",input)
	}
}