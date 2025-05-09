package models

type UserID string

type User struct {
	UserId       UserID
	LastActivity string
	IsActive     bool
	FirstName    string
	LastName     string
}
