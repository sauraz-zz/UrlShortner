package longurlreq

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func CreateShortURL(c echo.Context) error {

	longURL, err := newModel(c)

	if err != nil {
		return err
	}
	fmt.Println(longURL.URL)
	return nil
}
