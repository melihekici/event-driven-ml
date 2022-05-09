package service

import (
	"authapp/internal/dto"
	"authapp/internal/repository"
	"errors"
	"log"
)

type UserService interface {
	GetUser(userID int64) (*dto.User, error)
	AddUser(user dto.User) (*int64, error)
	DeleteUser(userID int64) error
	UpdateUser(user dto.User) (*dto.User, error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) AddUser(user dto.User) (*int64, error) {
	// validate email
	id, err := u.dao.NewUserQuery().CreateUser(user)
	if err != nil {
		log.Printf("Error creating user. %v", err)
		return nil, err
	}

	return id, nil
}

func (u *userService) GetUser(userID int64) (*dto.User, error) {
	var user *dto.User

	user, err := u.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		log.Printf("Requested user does not exist. %v", err)
		return nil, errors.New("Requested user does not exist" + err.Error())
	}

	return &dto.User{ID: user.ID, Username: user.Username, Password: "", Email: user.Email}, nil
}

func (u *userService) DeleteUser(id int64) error {
	err := u.dao.NewUserQuery().DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) UpdateUser(user dto.User) (*dto.User, error) {
	// email checking
	updatedUser, err := u.dao.NewUserQuery().UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
