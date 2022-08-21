package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("welcome to the time study")

	presenttime := time.Now()

	fmt.Println(presenttime.Local().Format("01-02-2006 15:04:05 Monday"))
}