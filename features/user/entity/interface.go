package entity

type UserDataInterface interface{
	Register(data UserCore) (row int, err error)
	Login(email, password string) (UserCore, string, error)
	ReadSpecificUser(id string) (user UserCore, err error)
	DeleteUser(id string) (err error)
}

type UserUseCaseInterface interface{
	Register(data UserCore) (row int, err error)
	Login(email, password string) (UserCore, string, error)
	ReadSpecificUser(id string) (user UserCore, err error)
	DeleteUser(id string) (err error)
}