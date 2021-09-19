package myproduct

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (v ProductValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func homePage(c echo.Context) error {
	fmt.Println("home page called")
	return c.String(http.StatusOK, "homepage, hello world !!!")
}

var products = []map[int]string{{1: "product1"}, {2: "product2"}}

func getProduct(context echo.Context) error {
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
}

func getProducts(context echo.Context) error {
	return context.JSON(http.StatusOK, products)
}

func postProduct(context echo.Context) error {
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
}

func putProduct(context echo.Context) error {
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
}

func deleteProduct(context echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}
	var index int
	for i, p := range products {
		if _, f := p[pID]; f == true {
			index = i
			product = p
		}
	}
	if product == nil {
		return context.JSON(http.StatusNotFound, "product not found")
	}
	splice := func(p []map[int]string, i int) []map[int]string {
		return append(p[:i], p[i+1:]...)
	}
	products = splice(products, index)

	return context.JSON(http.StatusCreated, product)
}
