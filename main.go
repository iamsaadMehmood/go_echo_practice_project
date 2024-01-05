package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
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
	e.Logger.Fatal(e.Start(":8080")) //e.Logger.Fatal use for logging error
}
