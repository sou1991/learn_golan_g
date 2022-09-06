package tests

import (
	"github.com/sou1991/learn_golan_g/entity"
)

type UserRepositoryMock struct{
}

func (urm *UserRepositoryMock) Find() (user []entity.User err error){
	user := []entity.User{}
	return
}
