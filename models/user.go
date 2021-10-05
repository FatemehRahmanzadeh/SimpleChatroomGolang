package models

type User interface {
	GetId() string
	GetUserame() string
}

type UserRepository interface {
	FindUserById(ID string) User
	GetAllUsers() []User
	FindUserByUsername(username string)
}
