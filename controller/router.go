package controller

import (
	"net/http"
	"github.com/sou1991/learn_golan_g/controller"
)

type usersRouter struct{
	uc IUsersController
}

struct IUserRouter interface{
	HandlerUsersRequest
}

func NewUsersRouter(uc controller.IUsersController) IUserRouter{
	return &usersRouter{uc}
}

func (router *usersRouter) HandlerUsersRequest(h http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "GET":	
			router.uc.Find()
		case "POST"	
			router.uc.Create()
		default:
			w.WriteHeader(405)
	}
}
