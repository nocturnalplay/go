package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("welcome to the slice's server")

	var FruitList = []string{"apple","orange","lemon"}
	FruitList = append(FruitList, "mango","banana")
	fmt.Printf("%v\n",FruitList)

	FruitListt := append(FruitList[1:3])

	fmt.Println(FruitListt)

	highScore := make([]int,4)

	highScore[0]=  234
	highScore[1]=  23
	highScore[2]=  24
	highScore[3]=  34

	sort.Ints(highScore)

	fmt.Println(highScore)
	
}