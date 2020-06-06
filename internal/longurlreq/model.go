package longurlreq

import (
	"github.com/labstack/echo/v4"
)

type LongURL struct {
	URL string `json:"name" validate:"required"`
}

func newModel(c echo.Context) (*LongURL, error) {
	l := new(LongURL)
	if err := c.Bind(l); err != nil {
		return nil, err
	}
	return l, nil
}
