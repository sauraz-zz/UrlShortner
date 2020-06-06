package main

import (
	"net/http"

	"urlShortner/internal/longurlreq"
	//"urlShortner/pkg/mongo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	/*collection, err := mongo.Init()
	if err != nil {
		e.Logger.Fatal("Database Falied.")
	}*/

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	//e.GET("", hello)
	e.GET("/:shortURL", hello)
	e.POST("/", longurlreq.CreateShortURL)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getLongURL(c echo.Context) error {
	longURL := c.QueryParam("longURL")
	return c.String(http.StatusOK, longURL)
}

func getShortURL(c echo.Context) error {
	longURL := c.Param("shortURL")
	return c.String(http.StatusOK, longURL)
}
