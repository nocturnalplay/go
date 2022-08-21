package main

import (
	"fmt"
)

func main(){
	fmt.Println("welome to the map server")

	lang := make(map[string]string)

	lang["js"] = "javascript"
	lang["py"] = "python"

	for key,value := range lang{
		fmt.Println(key,":--:",value)
	}
}