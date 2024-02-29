package utils

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	log.Errorf("server error %v", err)

	errorMessage := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		errorMessage = fmt.Sprintf("%v", he.Message)
	}

	// Send the error response to the client
	if !c.Response().Committed {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errorMessage,
		})
	}
}
