package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	"strconv"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (v ProductValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8081"
	}
	e := echo.New()
	v := validator.New()
	products := []map[int]string{{1: "product1"}, {2: "product2"}}

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

	e.GET("/products", func(context echo.Context) error {
		return context.JSON(http.StatusOK, products)
	})

	e.POST("/product", func(context echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}
		var requestBody body
		e.Validator = &ProductValidator{v}

		if err := context.Bind(&requestBody); err != nil {
			return err
		}
		if err := context.Validate(requestBody); err != nil {
			return err
		}
		/*		if err:=v.Struct(requestBody);err != nil{
					return err
				}
		*/
		product := map[int]string{
			len(products) + 1: requestBody.Name,
		}
		products = append(products, product)
		return context.JSON(http.StatusCreated, product)
	})
	e.PUT("/product/:id", func(context echo.Context) error {
		var product map[int]string
		pID, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			return err
		}
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}
		var requestBody body
		e.Validator = &ProductValidator{v}

		if err := context.Bind(&requestBody); err != nil {
			return err
		}

		if err := context.Validate(requestBody); err != nil {
			return err
		}

		for _, p := range products {
			if _, f := p[pID]; f == true {
				p[pID] = requestBody.Name
				product = p
			}
		}

		if product == nil {
			return context.JSON(http.StatusNotFound, "product not found")
		}

		return context.JSON(http.StatusCreated, product)
	})

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
