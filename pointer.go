package main

import "fmt"

type Store struct{
	Id int
	Name string
}

func main(){
	n := "HanmaBaki"

	//＆でポイント型生成できる.変数nのメモリアドレスを保有してる
	np := &n

	fmt.Printf("変数nのメモリアドレス :%v\n", np)

	//*でポインタ型npが指しているメモリアドレスの値を読み取れる
	fmt.Printf("変数の値:%v\n", *np)


	//基本のキ
	s := Store{1,"hoge"}
	fmt.Println("関数の値:%v\n", Do(&s))
}

func Do(s *Store) string {
	return s.Name
}
