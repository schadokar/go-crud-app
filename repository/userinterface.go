package repository

import "gocrudapp/model"

type UserInterface interface {
	InsertUser(*model.User) (int, error)
	FetchAllUser() ([]model.User, error)
	FetchUserByID(int) (model.User, error)
	UpdateUserByID(int, *model.User) (int, error)
	DeleteUserByID(int) (int, error)
	DeleteAllUser() (int, error)
}
