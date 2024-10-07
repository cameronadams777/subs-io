package controllers

import (
	"app/services"
	"app/views/pages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ApplicationViewHandler struct {
	AuthService services.AuthService
}

func (av *ApplicationViewHandler) HandleHomeIndex(c echo.Context) error {
  app_context := get_app_context(c, av.AuthService)

	return render_with_context(
		c,
		pages.HomeIndex(pages.HomePageProps{
			Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
		}),
		app_context,
	)
}

func (av *ApplicationViewHandler) HandleNotFoundIndex(c echo.Context) error {
	return render(c, pages.NotFoundIndex())
}

func (av *ApplicationViewHandler) HandleErrorIndex(c echo.Context) error {
	return render(c, pages.ErrorIndex())
}
