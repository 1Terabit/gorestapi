package handlers

import (
	"gorestapi/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

var users []models.User
var nextID = 1

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Router /users [post]
func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user.ID = nextID
	nextID++
	users = append(users, user)
	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
