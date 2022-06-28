package main

import "fmt"

type Vertex struct {
  X int
	Y int
}

func main(){
	v := Vertex{2, 4}

  //ポインタ pはvのポイントアドレスを持ってる
  p := &v

  p.X = 1e9

  fmt.Println(p.X)
}
