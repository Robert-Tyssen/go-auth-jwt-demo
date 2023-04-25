package domain

import "context"

type User struct {
	Id       string
	Email    string
	Password string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
}
