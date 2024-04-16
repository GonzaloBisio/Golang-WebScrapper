package user_service

import (
	m "Golang/models"
	user_repository "Golang/repositories/user.repository"
)

func Create(user m.User) error {

	err := user_repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Users, error) {

	users, err := user_repository.Read()
	if err != nil {
		return users, err
	}
	return users, nil
}

func Update(user m.User, userID string) error {

	err := user_repository.Update(user, userID)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userID string) error {

	err := user_repository.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}
