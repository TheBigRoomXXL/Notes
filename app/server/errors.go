package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// specification: https://www.rfc-editor.org/rfc/rfc9457.html
type RFC9457ProblemDetails struct {
	Type    string `json:"type,omitempty"`    // url to documentation on the error
	Title   string `json:"title,omitempty"`   // A short, human-readable summary of the type.
	Status  int    `json:"status,omitempty"`  // The HTTP status code
	Details string `json:"details,omitempty"` //A human-readable explanation specific to this occurrence of the problem.
}

func (pb *RFC9457ProblemDetails) Error() string {
	return pb.Details
}

func RFC9457ErrorHandler(err error, c echo.Context) {
	status := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
	}
	c.Logger().Error(err)

	msg := RFC9457ProblemDetails{
		Status:  status,
		Details: err.Error(),
	}

	c.JSON(status, msg)
}
