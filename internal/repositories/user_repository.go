package repositories

import "gorestapi/internal/models"

var users []models.User
var nextID = 1

func SaveUser(user models.User) models.User {
	user.ID = nextID
	nextID++
	users = append(users, user)
	return user
}

func GetAllUsers() []models.User {
	return users
}
