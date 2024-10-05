package services

import (
	app "app"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/tiktok"
)

const (
  SESSION_NAME = "subs_io_session"
)

type AuthService struct{}

func NewAuthService(store sessions.Store) *AuthService {
	gothic.Store = store

	goth.UseProviders(
		google.New(
      app.GetConfig("GOOGLE_CLIENT_ID"),
      app.GetConfig("GOOGLE_CLIENT_SECRET"),
      "http://localhost:4000/auth/google/callback",
    ),
    tiktok.New(
      app.GetConfig("TIKTOK_CLIENT_ID"),
      app.GetConfig("TIKTOK_CLIENT_SECRET"),
      "http://localhost:4000/auth/tiktok/callback",
    ),
	)

	return &AuthService{}
}

func (s *AuthService) GetSessionUser(r *http.Request) (goth.User, error) {
	session, err := gothic.Store.Get(r, SESSION_NAME)
	if err != nil {
		return goth.User{}, err
	}

	u := session.Values["user"]
	if u == nil {
		return goth.User{}, fmt.Errorf("user is not authenticated! %v", u)
	}

	return u.(goth.User), nil
}

func (s *AuthService) StoreUserSession(c echo.Context, user goth.User) error {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := gothic.Store.Get(c.Request(), SESSION_NAME)

	session.Values["user"] = user

	err := session.Save(c.Request(), c.Response().Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return err
	}

	return nil
}

func (s *AuthService) RemoveUserSession(c echo.Context) {
	session, err := gothic.Store.Get(c.Request(), SESSION_NAME)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	session.Values["user"] = goth.User{}
	// delete the cookie immediately
	session.Options.MaxAge = -1

	session.Save(c.Request(), c.Response().Writer)
}

func RequireAuth(handlerFunc echo.HandlerFunc, auth *AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
	  _, err := auth.GetSessionUser(c.Request())
		if err != nil {
      log.Println(err)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return nil
		}

		return handlerFunc(c)
	}
}

