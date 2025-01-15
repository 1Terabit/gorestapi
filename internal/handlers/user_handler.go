package handlers

import (
	"gorestapi/internal/models"
	"log"
	"net/http"

	"gorestapi/internal/middlewares"

	"github.com/labstack/echo/v4"
)

var users []models.User
var nextID = 1

// Simulamos una base de datos de usuarios
var userDB = map[string]string{
	"user1": "password1", // username: password
	"user2": "password2",
}

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

// Login maneja el inicio de sesión y genera un token
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Validar el usuario y la contraseña
	if expectedPassword, ok := userDB[username]; !ok || expectedPassword != password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	// Generar el token si la validación es exitosa
	token, err := middlewares.GenerateJWT(username)
	if err != nil {
		log.Println("Error generating token:", err) // Log para errores en la generación del token
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
