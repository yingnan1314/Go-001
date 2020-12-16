package user

type User struct {
	Id int
	Nickname string
}

func (u User) GetUserInfo() User {
	return User{1,"nick"}
}