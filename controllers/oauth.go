package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

type OAuthController struct{}

func (oac OAuthController) HandleOAuthCallback(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))

	_, err := gothic.CompleteUserAuth(c.Response(), c.Request().WithContext(ctx))
	if err != nil {
		return err
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
	ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))

	gothic.Logout(c.Response(), c.Request().WithContext(ctx))
	return c.Redirect(http.StatusPermanentRedirect, "/")
}
