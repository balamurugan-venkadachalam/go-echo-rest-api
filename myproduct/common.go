package myproduct

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

//MiddlewareFunc func(HandlerFunc) HandlerFunc
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		fmt.Println("inside server header middleware")
		mid := next(context)
		fmt.Println("inside server header after method call middleware")
		return mid
	}
}

func Start() {
	e.Use(ServerHeader)
	e.Pre(middleware.RemoveTrailingSlash(), middleware.BodyLimit("100K"))
	e.GET("/", homePage)
	e.GET("/products", getProducts)
	e.POST("/product", postProduct)
	e.PUT("/product/:id", putProduct)
	e.DELETE("/product/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))

}
