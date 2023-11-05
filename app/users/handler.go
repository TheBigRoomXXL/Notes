package users

// This file regroup all HTTP handler related to the user namespace

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// @Summary Login
// @Description  Try to login. "Accept: text/html" will trigger a redirect. "Accept: application/json" will send back the user object.
// @Tags users
// @Accept       json, application/x-www-form-urlencoded
// @Produce      json, text/html
// @Param        name string "username"
// @Param        search  query string false "FTS query"
// @Success      200  {object}  users.User
// @Success      302  ""
// @Failure      422  {object}  users.User
// @Failure      401  string    "login failed"
// @Router       /users/login [post]
func PostLogin(c echo.Context) error {
	err := postLogin(c)
	fmt.Println(err)
	return err

}

func postLogin(c echo.Context) error {
	var validInput UserSerializer
	err := c.Bind(&validInput)
	if err != nil {
		return err
	}

	db := c.Get("db").(*sql.DB)
	user, err := SelectUser(db, validInput.Username)
	if err != nil {
		return err
	}

	check, err := user.checkPassword(validInput.Password)
	if err != nil {
		return err
	}

	if check {
		return c.String(
			http.StatusUnauthorized,
			"login failed: bad password or username",
		)
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   true,
	}
	sess.Values["isAuthentificated"] = true
	sess.Values["identity"] = user.Username
	sess.Save(c.Request(), c.Response())

	//TO DO: Return user object if accept JSON
	return c.Redirect(http.StatusFound, "/")
}
