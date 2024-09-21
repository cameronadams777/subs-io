package controllers

import (
	"app/views/pages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthenticationController struct{}

func (ac *AuthenticationController) HandleLoginIndex(c echo.Context) error {
	return render(c, pages.LoginIndex(
		pages.LoginIndexPageProps{
			Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
		},
	))
}

