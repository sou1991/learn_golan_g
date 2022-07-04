package main

import "fmt"

type Iuser interface{
	Create()
	//・・・メソッドシグニチャ郡を作る
}

type User struct {
	Id, Age int
	Name, Address string
}

func (u User) Create() {
	fmt.Println(u.Id)
}

func main(){
	var u Iuser = User{//レシーバー作成
		Id : 1,
		Name : "tarou",
		Age : 20,
		Address : "Kobe City",
	}
	
	fmt.Printf("uこれ何型??%T\n", u)
	u.Create()
}
