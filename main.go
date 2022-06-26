package main

import (
	"fmt"
	"math/rand"
	"mymodule/pkg/calc"
)

func main(){
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(calc.Add(2,8))
}
