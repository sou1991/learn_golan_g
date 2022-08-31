package main

import "fmt"

func main(){
	n := []int {7, 2, 8, -9, 4, 0}

	pv := len(n)/2

	//チャネル作成
	c := make(chan int)

	go Sum(n[pv:], c)
	go Sum(n[:pv], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func Sum(n []int, c chan int){
	sum := 0
	for _, v := range n{
		sum = sum + v
	}
	//チャネルに結果を代入
	c <- sum
}
