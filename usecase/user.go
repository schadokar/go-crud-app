package usecase

import (
	"gocrudapp/model"
	"gocrudapp/repository"
	"log/slog"
)

type UserService struct {
	UserRepo repository.UserInterface
}

func (s *UserService) CreateUser(u *model.User) (int, error) {
	insertID, err := s.UserRepo.InsertUser(u)
	if err != nil {
		slog.Error("error while inserting user", err)
		return 0, err
	}

	slog.Info("user successfully inserted", slog.Int("id", insertID))
	return insertID, nil
}

func (s *UserService) FetchAllUser() ([]model.User, error) {
	users, err := s.UserRepo.FetchAllUser()
	if err != nil {
		slog.Error("error while fetching user", err)
		return users, err
	}

	slog.Info("users successfully fetched", slog.Int("total", len(users)))
	return users, nil
}

func (s *UserService) FetchUserByID(id int) (model.User, error) {
	user, err := s.UserRepo.FetchUserByID(id)
	if err != nil {
		slog.Error("error while fetching user", err)
		return user, err
	}

	slog.Info("user successfully fetched.", slog.Int("id", id))
	return user, nil
}

func (s *UserService) UpdateUserByID(id int, u *model.User) (int, error) {
	result, err := s.UserRepo.UpdateUserByID(id, u)
	if err != nil {
		slog.Error("error while deleting user", err)
		return result, err
	}

	slog.Info("user successfully updated.", slog.Int("id", id))
	return result, nil
}

func (s *UserService) DeleteUserByID(id int) (int, error) {
	result, err := s.UserRepo.DeleteUserByID(id)
	if err != nil {
		slog.Error("error while deleting user", err)
		return result, err
	}

	slog.Info("user successfully deleted.", slog.Int("id", id))
	return result, nil
}

func (s *UserService) DeleteAllUser() (int, error) {
	result, err := s.UserRepo.DeleteAllUser()
	if err != nil {
		slog.Error("error while deleting user", err)
		return result, err
	}

	slog.Info("users successfully deleted", slog.Int("total", result))
	return result, nil
}
