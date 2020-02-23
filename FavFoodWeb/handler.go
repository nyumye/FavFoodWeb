package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func setRoute(e *echo.Echo) {
	e.GET("/top", func(c echo.Context) error {
		return c.Render(http.StatusOK, "top", nil)
	})

	e.GET("/products", func(c echo.Context) error {
		return c.Render(http.StatusOK, "products", nil)
	})

	e.GET("/fortune", func(c echo.Context) error {
		return c.Render(http.StatusOK, "fortune", nil)
	})

	// e.GET("/hello", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "hello", "ZENZAI TABE TAAAAI!!!!!")
	// })

	// e.GET("/hello2", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "hello", "OSHIRUKO DEMO IIYOOOOOOOOO!!!!!!")
	// })

	//set products api route
	e.GET("/products/api/:id", getProduct)

}

func getProduct(c echo.Context) error {
	id := c.Param("id")
	fmt.Printf("id is %v", id)
	foodInfo := findFoodDocument(bson.D{{Key: "_id", Value: id}})
	return c.JSON(http.StatusOK, foodInfo)
}
