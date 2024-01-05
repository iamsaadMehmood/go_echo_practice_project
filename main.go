package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	products := []map[int]string{{1: "mobile"}, {2: "Laptops"}, {3: "Desktops"}}
	e.GET("/start", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})
	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string
		for _, p := range products {
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusOK, product)

	})
	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, product)
		return c.JSON(http.StatusOK, product)
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port))) //e.Logger.Fatal use for logging error
}
