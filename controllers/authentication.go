package controllers

import (
	"app/services"
	"app/views/pages"
	"net/http"

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

func (ac *AuthenticationController) HandleLogout(c echo.Context) error {
	auth_service := services.AuthService{
		CTX: c,
	}

	auth_service.SignOut()

	return c.Redirect(http.StatusPermanentRedirect, "/login")
}
