package main

import (
	"net/http"
	"github.com/sou1991/learn_golan_g/repository"
	"github.com/sou1991/learn_golan_g/controller"
)

//コンストラクタ群
var repo IUsersRepository = repository.NewUsersRepository()
var ctrl IUsersController = controller.NewUsersController(repo)
var ro IUserRouter = controller.NewUsersRouter(ctrl)

func main(){
	server := http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/users/", ro.HandlerUsersRequest)
	server.ListenAndServe()
}
