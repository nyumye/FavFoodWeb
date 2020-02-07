package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	prepareDatabase()

	e := echo.New()
	SetRender(e)

	e.Static("/", "public")
	// e.Static("/js/boot_test.js", "../../public/js/boot_test.js")
	// e.Static("/img/donut.png", "../../public/image/donut.png")
	// e.Static("/img/rice-cake.png", "../../public/image/rice-cake.png")

	e.GET("/top", func(c echo.Context) error {
		return c.Render(http.StatusOK, "top", nil)
	})

	e.GET("/products", func(c echo.Context) error {
		return c.Render(http.StatusOK, "products", nil)
	})

	// e.GET("/hello", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "hello", "ZENZAI TABE TAAAAI!!!!!")
	// })

	// e.GET("/hello2", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "hello", "OSHIRUKO DEMO IIYOOOOOOOOO!!!!!!")
	// })

	e.Logger.Fatal(e.Start(":8080"))
}
