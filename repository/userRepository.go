package repository

import (
	"errors"
	"fmt"
	uuid2 "github.com/google/uuid"
	"rest/models"
)

type UserRepository struct {
	Users []models.User
}

func (repo *UserRepository) Create(user models.User) {
	repo.Users = append(repo.Users, user)
}

func (repo *UserRepository) GetAll() []models.User {
	return repo.Users
}

func (repo *UserRepository) GetOne(uuid string) (models.User, error) {
	for _, user := range repo.Users {
		if user.UUID == uuid2.MustParse(uuid) {
			return user, nil
		}
	}
	return models.User{}, errors.New(fmt.Sprintf("User with uuid [%s] not found!", uuid))
}

func (repo *UserRepository) Delete(uuid string) error {
	for i, user := range repo.Users {
		if user.UUID == uuid2.MustParse(uuid) {
			repo.Users[i] = repo.Users[len(repo.Users)-1]
			repo.Users = repo.Users[:len(repo.Users)-1]
			return nil
		}
	}
	return errors.New(fmt.Sprintf("User with uuid [%s] not found!", uuid))
}
