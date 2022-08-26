package controller

import (
	"github.com/sou1991/learn_golan_g/repository"
)

//フィールドみたいな感じ？？interfaceで持っとくとInmemoryのテストクラスで置き換えれる
//基本的にstructは振る舞い(メソッド)を持たないので、レシーバになる
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

func (uc usresController) Find(w http.ResponseWriter, r *http.Request){
	user, err = uc.ur.Find()

	if err != nil{
		w.WriteHeader(500)
		return
	}
}

func Create(){

}

func Update(){

}

func Delete(){

}
