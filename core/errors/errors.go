package errors

import (
	"net/http"
	"sequence-api/core/config"

	"github.com/labstack/echo/v4"
)

// HTTPErrorHandler http error handler
func HTTPErrorHandler(err error, c echo.Context) {
	var code int
	var message interface{}

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = map[string]interface{}{"message": he.Message}
	} else if he, ok := err.(config.Error); ok {
		code = http.StatusBadRequest
		message = he
	} else {
		code = http.StatusBadRequest
		message = map[string]string{"message": err.Error()}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, message)
		}

		if err != nil {
			c.Logger().Error(err)
		}
	}
}
