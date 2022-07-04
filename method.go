package main

import "fmt"

type Square struct{
	X, Y int
}

//Squareのレシーバーでメソッドが呼べる　※レシーバー.関数名
func (s Square) Area() int {
	return s.X * s.Y
}

//ポインタ型にしないと元の値で計算される,レシーバの値を変更する場合一般的
func Scale(s *Square, f int) {
	s.X = s.X * f
	s.Y = s.Y * f
}

func main(){
	s := Square{2, 4}
	Scale(&s, 10)

	fmt.Println(s.Area())
}
