package encrypt

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RequestData struct {
	Message string `json:"message"`
}

func Get(c echo.Context) error {
	data := ResponseData{
		Status:  200,
		Message: "Hello, World!",
	}
	return c.JSON(http.StatusOK, data)
}

func Post(c echo.Context) error {
	data := new(RequestData)
	if err := c.Bind(&data); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}
