package services

import "gorestapi/internal/models"

var users []models.User
var nextID = 1

func CreateUser(user models.User) models.User {
	user.ID = nextID
	nextID++
	users = append(users, user)
	return user
}

func GetUsers() []models.User {
	return users
}
