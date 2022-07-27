package main

import (
	"fmt"
	"github.com/sou1991/learn_golan_g/pkg/calc"
)

func main(){
	//外部パッケージをインポートしたExported(公開された)メソッド名は大文字である必要がある。
	fmt.Println(calc.Add(2,8))
}
