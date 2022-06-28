package main

import (
	"fmt"
	"math"
)
//メソッドシグニチャを合わせる　関数も変数とのこと
func compute(hypot func(x, y float64) float64) float64{
	return hypot(5, 3)
}

func main() {
	hypot := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y)
	}
	
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
