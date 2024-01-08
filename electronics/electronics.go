package electronics

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

func Start() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	e.POST("/products", createProduct)
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
