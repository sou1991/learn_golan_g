package controller

import (
	"github.com/sou1991/learn_golan_g/repository"
)

//フィールドみたいな感じ？？interfaceで持っとくとInmemoryのテストクラスで置き換えれる
type usresController struct{
	ur repository.IUsersRepository
}

type IUsersController interface{
	Find()
	Create()
	Update()
	Delete()
}

//次のレイヤーの依存注入する コンストラクタ interfaceを返すによってstructを抽象化できる
func NewUsersController(ur repository.IUsersRepository) IUsersController{
	return &usresController{ur}
}

func Find(){

}

func Create(){

}

func Update(){

}

func Delete(){

}
