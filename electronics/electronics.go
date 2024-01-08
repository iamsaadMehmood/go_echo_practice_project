package electronics

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside middleware")
		return next(c)
	}
}
func preServerMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside pre middleware")
		return next(c)
	}
}

// starts the application
func Start() {
	e.Use(serverMessage)    // always called after pre can be use for admin validation if needed
	e.Pre(preServerMessage) //always called first can be use for validate token

	e.POST("/products", createProduct)
	// e.GET("/products", getProducts, serverMessage) route level middle ware
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", config.Port)))
}
