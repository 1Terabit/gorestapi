package main

import (
	"context"
	"fmt"
	"gorestapi/internal/handlers"
	"gorestapi/internal/middlewares"
	"log"

	"time"

	"cloud.google.com/go/firestore"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var jwtKey = []byte("tu_clave_secreta") // Cambia esto por una clave secreta más segura

// Claims es la estructura que contiene los datos del token
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT genera un nuevo token JWT
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // El token expira en 24 horas
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Leer el archivo de configuración
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error leyendo el archivo de configuración: %v", err)
	}
}

func main() {
	initConfig() // Inicializa la configuración

	e := echo.New()

	ctx := context.Background()
	projectID := viper.GetString("firestore.project_id")

	// Middleware auth & error
	e.Use(middlewares.AuthMiddleware)
	e.HTTPErrorHandler = middlewares.ErrorHandler

	// Middleware para JWT
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  jwtKey,
		TokenLookup: "header:Authorization",
		//AuthScheme:  "Bearer",
		// Puedes agregar más configuraciones aquí si es necesario
	}))

	// Rutas
	e.POST("/login", handlers.Login)
	e.POST("/users", handlers.CreateUser, echojwt.WithConfig(echojwt.Config{
		SigningKey: jwtKey,
	}))
	e.GET("/users", handlers.GetUsers, echojwt.WithConfig(echojwt.Config{
		SigningKey: jwtKey,
	}))

	// Firestore
	client, err := firestore.NewClient(ctx, projectID) // Cambia "tu-proyecto-id" por el ID de tu proyecto
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	fmt.Println("Conectado a Firestore!")

	// Obtiene el puerto de la configuración
	port := viper.GetInt("server.port")
	log.Printf("Iniciando el servidor en el puerto %d", port)
	e.Start(":" + fmt.Sprint(port))
}
