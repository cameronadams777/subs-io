package controllers

import (
	"app/services"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

type OAuthController struct{}

func (oac OAuthController) HandleOAuthCallback(c echo.Context) error {
	ctx := context.WithValue(
    c.Request().Context(),
    gothic.ProviderParamKey,
    c.Param("provider"),
  )

  user, err := gothic.CompleteUserAuth(c.Response(), c.Request().WithContext(ctx))

	if err != nil {
		return err
	}

  // TODO: Check if user exists based on email. If not, create one. If so,
  // and the provider is not currently connected, connect and continue.

	session, _ := gothic.Store.Get(c.Request(), services.SESSION_NAME)

	session.Values["user"] = user

	session_sav_err := session.Save(c.Request(), c.Response().Writer)
	if session_sav_err != nil {
    fmt.Println("Session Err:", session_sav_err)
		return c.String(http.StatusInternalServerError, "An error occurred creating session")
	}

	return c.Redirect(http.StatusPermanentRedirect, "/")
}

func (oac OAuthController) HandleOAuthIndex(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))

	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(c.Response(), c.Request().WithContext(ctx)); err == nil {
		return c.Redirect(http.StatusPermanentRedirect, "/")
	}

	gothic.BeginAuthHandler(c.Response(), c.Request().WithContext(ctx))

	return nil
}

func (oac OAuthController) HandleOAuthLogout(c echo.Context) error {
  fmt.Println("Provider:", c.Param("provider"))
	ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))

	gothic.Logout(c.Response(), c.Request().WithContext(ctx))

	session, _ := gothic.Store.Get(c.Request(), services.SESSION_NAME)
	session.Values["user"] = nil

	session_sav_err := session.Save(c.Request(), c.Response().Writer)
	if session_sav_err != nil {
    fmt.Println("Session Err:", session_sav_err)
		return c.String(http.StatusInternalServerError, "An error occurred destroying session")
	}

	return c.Redirect(http.StatusPermanentRedirect, "/")
}
