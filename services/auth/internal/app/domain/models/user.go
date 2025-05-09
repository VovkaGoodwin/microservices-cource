package models

type UserId string

type User struct {
	id       UserId
	Username string
}

func NewUser(id UserId) *User {
	return &User{
		id: id,
	}
}

func (u *User) ID() UserId {
	return u.id
}

func (u *User) WithUsername(username string) *User {
	u.Username = username
	return u
}
