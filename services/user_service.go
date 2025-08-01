package services

import (
	"crud-test-go/models"
	"errors"
	"strings"
)

var users []models.User

func GetAllUsers() []models.User {
	return users
}

func AddUser(u models.User) (models.User, error) {
	for _, user := range users {
		if user.Name == u.Name {
			return models.User{}, errors.New("username already taken")
		}
	}
	u.ID = len(users) + 1
	users = append(users, u)
	return u, nil
}

func DeleteTestUsers() int {
	original := users
	users = nil
	count := 0

	for _, u := range original {
		if strings.HasPrefix(u.Name, "You Can't See ME ") || strings.HasPrefix(u.Email, "my") {
			count++
			continue // Skip adding test users to the new list
		}
		users = append(users, u)
	}
	return count
}