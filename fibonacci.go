package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	pp := 0 //前項
	p := 1 //項
	fipo := 0 //フィボナッチ
	
	return func() int{
		fipo = p + pp
	    pp = p
		p = fipo
		
		return fipo
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
