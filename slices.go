package main

import "fmt"

func main(){
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:5]
	fmt.Println(s)

	fmt.Printf("%T\n", s) //　[]int 可変長

	//これもスライス。こっちの方が実用的
	composite := []int{4, 6, 8, 10, 12, 14}
	fmt.Println(composite)

	//組み込みmake関数スライス作れる
	dy := 5
	ms := make([]int, dy)

	fmt.Println(ms)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}

	fmt.Printf("構造体の配列%T\n", st)

	var sn[]int
	 if sn == nil {
		fmt.Println("nilでっせ")
	 }

	 var ss[]string

	 if ss == nil {
		fmt.Println("nilでっせ")
	 }

	 for i := 0; i < len(s); i++{
		fmt.Println(s[i])
	 }

	 //rangeのほうがつかいそう
	 rs = string[]{"aa","bb","cc","dd","ee"}

	 for _, v = range rs{
		fmt.Println(v)
	 }

}
