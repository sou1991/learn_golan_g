package repository


type IUsersRepository interface{
	Find()
	Create()
	Update()
	Delete()
}

struct usersRepository struct{
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
