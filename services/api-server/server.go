package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api-server/handler"
	"api-server/shop"
	shopv1 "api-server/shop/gen/shop/v1"
)

type CustomValidator struct { 
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Fprintf(os.Stdout, "Reqest Body %v\n", string(reqBody))
	}))
	rules := map[string]string{
		"Name": "isdefault",
	}
	validate := validator.New()
	validate.RegisterStructValidationMapRules(map[string]string{"Shops": "required"}, shopv1.Shops{})
	validate.RegisterStructValidationMapRules(rules, shopv1.Shop{})
	e.Validator = &CustomValidator{validator: validate}

	e.GET("/api/list", handler.GetFilenames)
	e.GET("/api/counts", handler.GetStatusCounts)
	e.GET("/api/chart_list/:filename", handler.GetCharts)
	e.DELETE("/api/deleteFile/:filename", handler.DeleteProducts)

	e.GET("/api/shops", shop.GetShops)
	e.POST("/api/shops", shop.CreateShop)
	e.Logger.Fatal(e.Start(":5000"))
}
