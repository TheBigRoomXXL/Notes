package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// specification: https://www.rfc-editor.org/rfc/rfc9457.html
type RFC9457ProblemDetails struct {
	Type     string `json:"type,omitempty"`    // url to documentation on the error
	Title    string `json:"title,omitempty"`   // A short, human-readable summary of the type.
	Status   int    `json:"status,omitempty"`  // The HTTP status code
	Details  string `json:"details,omitempty"` //A human-readable explanation specific to this occurrence of the problem.
	internal error  // Original error for record keeping
}

func (pb *RFC9457ProblemDetails) Error() string {
	return pb.internal.Error()
}

func RFC9457ErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err.Error())

	if pb, ok := err.(*RFC9457ProblemDetails); ok {
		c.Logger().Error(pb.Error())
		c.JSON(pb.Status, pb)
		return
	}

	c.JSON(http.StatusInternalServerError, RFC9457ProblemDetails{
		Status: http.StatusInternalServerError,
		Title:  "Internal Server Error",
	})
}
