package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("welcome to the server")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')
	num_rating,err := strconv.ParseFloat(strings.TrimSpace(input),32)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("rating added:",num_rating + 1)
	}
}