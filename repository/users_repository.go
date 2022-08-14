package repository


type IUsersRepository interface{
	Find()
	Create()
	Update()
	Delete()
}

type usersRepository struct{
}


func NewUsersRepository() IUsersRepository {
	return &usersRepository
}


func Find(){

}

func Create(){

}

func Update(){

}

func Delete(){

}
