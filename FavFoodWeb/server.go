package main

import (
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {

	prepareDatabase()

	e := echo.New()

	e.Use(middleware.Logger())

	SetRender(e)

	e.Static("/", "public")
	// e.Static("/js/boot_test.js", "../../public/js/boot_test.js")
	// e.Static("/img/donut.png", "../../public/image/donut.png")
	// e.Static("/img/rice-cake.png", "../../public/image/rice-cake.png")

	setRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
