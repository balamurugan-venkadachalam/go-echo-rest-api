package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8081"
	}
	e := echo.New()
	products := []map[int]string{{1: "product1"}}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "homepage, hello world !!!")
	})
	e.GET("/product/:id", func(context echo.Context) error {
		var product map[int]string
		for _, p := range products {
			for key := range p {
				pID, err := strconv.Atoi(context.Param("id"))
				if err != nil {
					return err
				}
				if pID == key {
					product = p
				}
			}
		}
		if product == nil {
			return context.JSON(http.StatusNotFound, "product not found")
		}

		return context.JSON(http.StatusOK, product)
	})
	e.Logger.Print("Listening on port 8081")
	e.Logger.Fatal(e.Start(":8081"))
}
