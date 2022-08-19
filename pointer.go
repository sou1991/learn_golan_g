package main

import "fmt"

func main(){
	n := "HanmaBaki"

	//＆でポイント型生成できる.変数nのメモリアドレスを保有してる
	np := &n

	fmt.Printf("変数nのメモリアドレス :%v\n", np)

	//*でポインタ型npが指しているメモリアドレスの値を読み取れる
	fmt.Printf("変数の値:%v\n", *np)
}
