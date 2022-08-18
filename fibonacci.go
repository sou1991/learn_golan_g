package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	pp := 0 //前項
	p := 1 //項
	w := 0 
	
	return func() int{
		w = pp
		p = p + pp
		pp = p - pp

		return w
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
