package main

import (
	"fmt"
	"math/rand"
	//go.modに記述しているpackage/ディレクトリ/パッケージ名※モジュールモードの場合
	"mymodule/pkg/calc"
	"runtime"
)
//グローバル変数
var c, python, java bool
const agree = true

func main(){
	fmt.Println("My favorite number is", rand.Intn(10))
	//外部パッケージをインポートしたExported(公開された)メソッド名は大文字である必要がある。
	fmt.Println(calc.Add(2,8))

	//変数
	var str = "local"
	num := 123

	fmt.Println(str,c,python,java,num)
	fmt.Println(agree)
}
